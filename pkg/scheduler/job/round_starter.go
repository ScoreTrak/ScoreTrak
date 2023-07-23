package job

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/hostservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/mixins"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/round"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/events/entries"
	"github.com/ScoreTrak/ScoreTrak/pkg/events/handlers"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/robfig/cron/v3"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/fx"
	"time"
)

type RoundStarterJob struct {
	cronspec       string
	entryId        cron.EntryID
	entitiesClient *entities.Client
	pub            message.Publisher
	logger         *otelzap.SugaredLogger
}

func NewRoundStarterJob(cfg *config.Config, ec *entities.Client, p message.Publisher, logger *otelzap.SugaredLogger) *RoundStarterJob {
	return &RoundStarterJob{
		cronspec:       cfg.Scheduler.Jobs.RoundStarter.CronSpec,
		entitiesClient: ec,
		pub:            p,
		logger:         logger,
	}
}

func RegisterRoundStarterJob(lc fx.Lifecycle, c *cron.Cron, roundStarterJob *RoundStarterJob, logger *otelzap.SugaredLogger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			entryId, err := c.AddJob(roundStarterJob.cronspec, roundStarterJob)
			if err != nil {
				return err
			}
			roundStarterJob.setEntryId(entryId)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			c.Remove(roundStarterJob.entryId)
			return nil
		},
	})
}

func (s *RoundStarterJob) setEntryId(entryId cron.EntryID) {
	s.entryId = entryId
}

func (s *RoundStarterJob) Run() {
	timeoutDuration := time.Second * 55
	timeoutTimer := time.NewTimer(timeoutDuration)
	roundFinalizerTicker := time.NewTicker(5 * time.Second)
	ctx, cancelFunc := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancelFunc()
	sugaredLoggerWithCtx := s.logger.Ctx(ctx)

	// Get Competition
	comp, err := s.entitiesClient.Competition.Get(ctx, mixins.FIRST_ID)
	if err != nil {
		sugaredLoggerWithCtx.Panicw(err.Error())
	}
	if comp == nil {
		sugaredLoggerWithCtx.Panicw("competition does no exist")
	}
	if comp.Pause == true {
		sugaredLoggerWithCtx.Panicw(err.Error())
	}

	// Check if previous round has finished
	// Fail safely or attempt to schedule this job in the next 5 minutes

	// Possibly mark that round as failed to be finished. and start the next round.
	// Then run a function to generate a report based on the policy set for the competition like, if rounds are not completed
	// still score them. Or just.

	var newRound *entities.Round
	previousRound, err := s.entitiesClient.Round.Query().Order(round.ByRoundNumber(sql.OrderDesc())).First(ctx)
	// Check if err or previous round does not exist
	if err != nil || previousRound == nil {
		sugaredLoggerWithCtx.Infof("Unable to get previous round: %v\n", err)
		sugaredLoggerWithCtx.Infow("Will need to assume this is the start of the first round")
		// Create a new round
		newRound, err = s.entitiesClient.Round.Create().SetRoundNumber(1).SetStartedAt(time.Now()).Save(ctx)
		if err != nil {
			sugaredLoggerWithCtx.Panicw(err.Error())
		}
	} else {
		// Inspect the status of the previous round if it exists
		// Critize round if it has not finished
		sugaredLoggerWithCtx.Infof("Latest Round: %d (%s) has %s", previousRound.RoundNumber, previousRound.ID, previousRound.Status)
		if previousRound.Status != round.StatusFinished {
			sugaredLoggerWithCtx.Infow("The previous round is either still ongoing or only started. Need to shut it down and classify it as did not complete in time")
			sugaredLoggerWithCtx.Infow("We will flag that round as bad")
			err := s.entitiesClient.Round.UpdateOne(previousRound).SetStatus(round.StatusIncomplete).Exec(ctx)
			if err != nil {
				sugaredLoggerWithCtx.Panicf("Unable to flag the previous round as incomplete: %v", err.Error())
			}
		} else {
			sugaredLoggerWithCtx.Infow("Previous round completed successfully")
		}
		// Create a new round
		newRound, err = s.entitiesClient.Round.Create().SetRoundNumber(previousRound.RoundNumber + 1).SetStartedAt(time.Now()).Save(ctx)
		if err != nil {
			sugaredLoggerWithCtx.Panicw(err.Error())
		}
	}

	sugaredLoggerWithCtx.Infof("Starting scoring for %s, %s", comp.ID, comp.Name)
	sugaredLoggerWithCtx.Infof("Started New Round %d (%s)", newRound.RoundNumber, newRound.ID)

	// Get host services to score on
	sugaredLoggerWithCtx.Infow("Getting enabled services")
	services, err := s.entitiesClient.Service.Query().Where(service.PauseEQ(false)).All(ctx)
	if err != nil {
		sugaredLoggerWithCtx.Panicf("Unable to get services: %v\n", err)
	}
	filteredServices := filterServices(newRound.RoundNumber, services)
	serviceIds := []string{}
	for _, filteredService := range filteredServices {
		serviceIds = append(serviceIds, filteredService.ID)
	}

	hostServices, err := s.entitiesClient.HostService.Query().Where(hostservice.HasServiceWith(service.IDIn(serviceIds...)), hostservice.PauseEQ(false)).WithProperties().WithService().WithTeam().WithHost().All(ctx)
	if err != nil {
		sugaredLoggerWithCtx.Panicf("%v\n", err)
	}
	numOfHostServices := len(hostServices)
	sugaredLoggerWithCtx.Infof("Number of Host Services: %d", numOfHostServices)

	// Publish to be scored host services to queue
	for _, hostService := range hostServices {
		payload, err := handlers.NewHostServiceScorePayload(newRound.ID, newRound.RoundNumber, hostService.Edges.Host.Address, hostService.Edges.Service.Type, hostService.ID, hostService.Edges.Properties).Bytes()
		if err != nil {
			sugaredLoggerWithCtx.Panicf("%v\n", err)
		}
		msg := message.NewMessage(watermill.NewULID(), payload)
		err = s.pub.Publish(entries.TOPIC_HOST_SERVICE_SCORE, []*message.Message{msg}...)
		if err != nil {
			sugaredLoggerWithCtx.Panicf("%v\n", err)
		}
	}

	// Update Round status
	err = s.entitiesClient.Round.UpdateOne(newRound).SetStatus(round.StatusOngoing).SetNumOfIntendedChecks(numOfHostServices).Exec(ctx)
	if err != nil {
		sugaredLoggerWithCtx.Panicf("%v", err)
	}

	// Loop every 5 seconds to check if all checks have been saved. LOL. not the best way to do this.

	for {
		select {
		case <-ctx.Done():
			// Stop timers and tickers
			timeoutTimer.Stop()
			roundFinalizerTicker.Stop()
			sugaredLoggerWithCtx.Infow("Context has finished")
			return
		case <-timeoutTimer.C:
			roundFinalizerTicker.Stop()
			sugaredLoggerWithCtx.Infow("Function has been timeout")
			cancelFunc()
		case <-roundFinalizerTicker.C:
			closed, err := s.closeRound(ctx)
			if err != nil {
				sugaredLoggerWithCtx.Panicw(err.Error())
			}
			if closed {
				timeoutTimer.Stop()
				roundFinalizerTicker.Stop()
				sugaredLoggerWithCtx.Infow("Round closed")
				cancelFunc()
			}
		}
	}

}

func (s *RoundStarterJob) closeRound(ctx context.Context) (bool, error) {
	sugaredLoggerWithCtx := s.logger.Ctx(ctx)

	comp, err := s.entitiesClient.Competition.Get(ctx, mixins.FIRST_ID)
	if err != nil {
		return false, err
	}
	if comp == nil {
		return false, err
	}
	currentRound, err := s.entitiesClient.Round.Query().Order(round.ByRoundNumber(sql.OrderDesc())).First(ctx)
	if err != nil {
		return false, err
	}
	if currentRound.Status == round.StatusStarted {
		return false, err
	} else if currentRound.Status == round.StatusFinished {
		return false, err
	} else if currentRound.Status == round.StatusOngoing {
		// Checks if all checks have been fulfilled
		currentNumberOfChecks, _ := currentRound.QueryChecks().Count(ctx)
		if currentNumberOfChecks == currentRound.NumOfIntendedChecks {

			// Update round to finished
			sugaredLoggerWithCtx.Infow("WE ARE READY TO SCORE")
			s.entitiesClient.Round.UpdateOne(currentRound).SetStatus(round.StatusFinished).ExecX(ctx)

			// Update hostservice reports
			hss, err := s.entitiesClient.HostService.Query().WithService().WithTeam().All(ctx)
			if err != nil {
				return false, err
			}
			sugaredLoggerWithCtx.Infof("# of hss %d", len(hss))
			for _, hs := range hss {
				checks, err := s.entitiesClient.Check.Query().Where(check.HasHostserviceWith(hostservice.ID(hs.ID))).All(ctx)
				if err != nil {
					return false, err
				}
				points := 0

				for _, chk := range checks {
					if chk.Passed {
						points += hs.Edges.Service.Weight + hs.Edges.Service.PointBoost
					}
				}

				latestCheck, err := s.entitiesClient.Check.Query().Where(check.HasHostserviceWith(hostservice.ID(hs.ID))).Order(check.ByCreateTime(sql.OrderDesc())).First(ctx)
				if err != nil {
					return false, err
				}

				id, err := s.entitiesClient.HostServiceReport.Create().SetTeam(hs.Edges.Team).SetService(hs.Edges.Service).SetHostservice(hs).SetPoints(points).SetPassing(latestCheck.Passed).SetLatestCheckTime(latestCheck.CreateTime).OnConflict().UpdateNewValues().ID(ctx)
				if err != nil {
					return false, err
				}
				sugaredLoggerWithCtx.Infow("created host report", "id", id, "points", points)
			}

			// Update team reports

			// Update Comp Round ID
			err = s.entitiesClient.Competition.UpdateOne(comp).SetCurrentRoundID(currentRound.ID).Exec(ctx)
			if err != nil {
				return false, err
			}

			return true, nil

		}
	}
	return false, nil
}

func filterServices(roundNumber int, services []*entities.Service) []*entities.Service {
	var filteredServices []*entities.Service
	for _, s := range services {
		if isServiceScorableInRound(roundNumber, s.RoundFrequency, s.RoundDelay) {
			filteredServices = append(filteredServices, s)
		}
	}

	return filteredServices
}

func isServiceScorableInRound(roundNumber int, roundFreq int, roundDelay int) bool {
	if roundNumber >= roundDelay+1 {
		if roundNumber-roundDelay != 0 {
			if (roundNumber-roundDelay)%roundFreq == 0 {
				return true
			}
		}
	}
	return false
}
