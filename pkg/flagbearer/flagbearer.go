package flagbearer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/oklog/ulid/v2"
	"github.com/scoretrak/scoretrak/pkg/eventsv2/messages"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-co-op/gocron/v2"
	"github.com/scoretrak/scoretrak/internal/entities"
	"github.com/scoretrak/scoretrak/internal/entities/check"
	"github.com/scoretrak/scoretrak/internal/entities/hostservice"
	"github.com/scoretrak/scoretrak/internal/entities/mixins"
	"github.com/scoretrak/scoretrak/internal/entities/round"
	"github.com/scoretrak/scoretrak/internal/entities/service"
	"github.com/scoretrak/scoretrak/pkg/config"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/fx"
)

// FlagBearer is responsible for holding the logic on how to start rounds,
//
//	calculate expected services to be scored,
//	create the checks to be filed by workers when scored.
type FlagBearer struct {
	dbClient *entities.Client
	logger   *otelzap.SugaredLogger
	js       jetstream.JetStream
	//tracer   *trace.Tracer
}

//type FlagBearerPublisherOpts func()
//type FlagBearerPublisher interface {
//	Publish([]byte) error
//}

type FlagBearerOption func(flagbearer *FlagBearer)

func NewFlagBearer(entitiesClient *entities.Client, js jetstream.JetStream, logger *otelzap.SugaredLogger) *FlagBearer {
	return &FlagBearer{dbClient: entitiesClient, js: js, logger: logger}
}

func WithLogger(logger *otelzap.SugaredLogger) FlagBearerOption {
	return func(f *FlagBearer) {
		f.logger = logger
	}
}

const (
	CRON_TAG_GREEN_FLAG     = "green_flag"
	CRON_TAG_CHECKERED_FLAG = "checkered_flag"
)

func RegisterFlagBearerCronJob(lc fx.Lifecycle, el gocron.Elector, cfg config.Config, s gocron.Scheduler, fb *FlagBearer, logger *otelzap.SugaredLogger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Infoln("registering green flag job")
			_, err := s.NewJob(
				gocron.CronJob(cfg.Scheduler.Jobs.RoundStarter.CronSpec, true),
				gocron.NewTask(func() {
					logger.Infoln("determining leadership")
					if el.IsLeader(ctx) == nil {
						logger.Infoln("I am the leader")
						fb.StartNextRound(ctx)
					} else {
						logger.Infoln("I am not the leader")
					}
					logger.Infoln("done")
				}),
				gocron.WithTags(CRON_TAG_GREEN_FLAG),
				gocron.WithTags(CRON_TAG_CHECKERED_FLAG),
			)
			if err != nil {
				return err
			}

			_, err = s.NewJob(
				gocron.CronJob(cfg.Scheduler.Jobs.RoundFinisher.CronSpec, true),
				gocron.NewTask(func() {
					logger.Infoln("determining leadership")
					if el.IsLeader(ctx) == nil {
						logger.Infoln("I am the leader")
						fb.EndRound(ctx)
					} else {
						logger.Infoln("I am not the leader")
					}
					logger.Infoln("done")
				}),
				gocron.WithTags(CRON_TAG_CHECKERED_FLAG),
			)
			if err != nil {
				return err
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Infoln("removing jobs")
			s.RemoveByTags(CRON_TAG_GREEN_FLAG, CRON_TAG_CHECKERED_FLAG)
			return nil
		},
	})
}

func (f *FlagBearer) EndRound(ctx context.Context) error {
	// ctx, span := f.tracer.Start(context.Background(), "ending-round")
	// defer span.End()
	_, err := f.dbClient.Round.Query().Order(round.ByRoundNumber(sql.OrderDesc())).First(ctx)
	if err != nil {
		return err
	}

	//if currentRound.Status != round.StatusCalculated {
	//
	//} else {
	//
	//}

	return nil
}

// StartNextRound looks at the current state and setups data for scorer to score on.
func (f *FlagBearer) StartNextRound(ctx context.Context) error {
	// ctx, span := f.tracer.Start(context.Background(), "starting-new-round")
	// defer span.End()
	// TODO: Move logic on starting, running round to this component

	// Announce what competition we are starting a new round for
	_, err := f.dbClient.Competition.Get(ctx, mixins.FIRST_ID)
	if err != nil {
		return err
	}

	f.logger.Infow("Starting New Round", "competition", mixins.FIRST_ID)

	// Create New Round
	r, err := f.createNewRound(ctx)
	if err != nil {
		return err
	}

	// Determine which services are score-able
	scoreableHostServices, err := f.getScoreableServices(ctx, r.RoundNumber, r.ID, false)
	if err != nil {
		return err
	}

	// Create checks
	_, err = f.createChecks(ctx, scoreableHostServices, r)
	if err != nil {
		return err
	}

	// Get Scoreable Host Services again with Checks
	scoreableHostServicesWithChecks, err := f.getScoreableServices(ctx, r.RoundNumber, r.ID, true)
	if err != nil {
		return err
	}

	// send host service details to be scored
	err = f.publishQueueMessages(ctx, r, scoreableHostServicesWithChecks)
	if err != nil {
		return err
	}

	f.logger.Infow("Updating round status")
	// Update round to ongoing since host services queue messages have been sent
	err = f.updateRound(ctx, r, round.StatusOngoing)
	if err != nil {
		return err
	}

	return nil
}

func (f *FlagBearer) createNewRound(ctx context.Context) (*entities.Round, error) {
	//ctx, span := f.tracer.Start(context.Background(), "create-new-round")
	//defer span.End()
	var newRound *entities.Round

	// Check if no rounds exists as this would indicate the start of the competition
	numOfRounds, err := f.dbClient.Round.Query().Count(ctx)
	if err != nil {
		f.logger.Errorw("Unable to count number of rounds for competition")
		return nil, err
	}

	// Create first round if none exists
	if numOfRounds == 0 {
		// Round start from the number 1, not zero
		newRound, err = f.dbClient.Round.Create().SetRoundNumber(1).SetStartedAt(time.Now()).Save(ctx)
		if err != nil {
			f.logger.Panicw(err.Error())
		}
		return newRound, nil
	}

	// Find previous round
	previousRound, err := f.dbClient.Round.Query().Order(round.ByRoundNumber(sql.OrderDesc())).First(ctx)
	if err != nil {
		f.logger.Errorf("Unable to get previous round: %v\n", err)
		return nil, err
	}

	// Previous round found. Increment the round number for the new round
	// Inspect the status of the previous round if it exists
	// Criticize round if it has not finished
	f.logger.Infof("Latest Round: %d (%s) has %s", previousRound.RoundNumber, previousRound.ID, previousRound.Status)

	if previousRound.Status != round.StatusFinished {
		f.logger.Infow("The previous round is either still ongoing or only started. Need to shut it down and classify it as did not complete in time")
		f.logger.Infow("We will flag that round as bad")
		err := f.dbClient.Round.UpdateOne(previousRound).SetStatus(round.StatusIncomplete).Exec(ctx)
		if err != nil {
			f.logger.Panicf("Unable to flag the previous round as incomplete: %v", err.Error())
		}
	} else {
		f.logger.Infow("Previous round completed successfully")
	}

	// Create a new round
	newRound, err = f.dbClient.Round.Create().SetRoundNumber(previousRound.RoundNumber + 1).SetStartedAt(time.Now()).Save(ctx)
	if err != nil {
		f.logger.Panicw(err.Error())
	}

	return newRound, nil
}

func (f *FlagBearer) updateRound(ctx context.Context, r *entities.Round, status round.Status) error {
	err := f.dbClient.Round.UpdateOne(r).SetStatus(status).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (f *FlagBearer) getCompetition(ctx context.Context) *entities.Competition {
	//ctx, span := f.tracer.Start(context.Background(), "get-competition")
	//defer span.End()

	comp, err := f.dbClient.Competition.Get(ctx, mixins.FIRST_ID)
	if err != nil {
		f.logger.Panicw(err.Error())
	}
	if comp == nil {
		f.logger.Panicw("competition does no exist")
	}
	if comp.Pause == true {
		f.logger.Panicw(err.Error())
	}
	return comp
}

func (f *FlagBearer) getScoreableServices(ctx context.Context, roundNumber int, roundId string, withChecks bool) ([]*entities.HostService, error) {

	// Get all services that are not paused
	services, err := f.dbClient.Service.Query().Where(service.PauseEQ(false)).All(ctx)
	if err != nil {
		return nil, err
	}

	// Filter services by the round frequency set for them
	filteredServices := f.filterServicesByRoundFrequency(roundNumber, services)
	serviceIds := []string{}
	for _, filteredService := range filteredServices {
		serviceIds = append(serviceIds, filteredService.ID)
	}

	// Get host services for scoreable services that are not individually paused.
	scoreableHostServices, err := f.dbClient.HostService.Query().Where(hostservice.HasServiceWith(service.IDIn(serviceIds...)), hostservice.PauseEQ(false)).WithProperties().WithService().WithTeam().WithHost().WithChecks(
		func(cq *entities.CheckQuery) {
			cq.Where(check.RoundID(roundId))
		}).All(ctx)
	if err != nil {
		return nil, err
		//sugaredLoggerWithCtx.Panicf("%v\n", err)
	}

	if withChecks {
		// Ensure each hostservice has a check
		for _, hs := range scoreableHostServices {
			if len(hs.Edges.Checks) == 0 {
				f.logger.Errorw("Scoreable Host service does not have a check", "id", hs.ID)
			} else {
				f.logger.Infow("The check exist", "id", hs.ID)
			}
		}
	}

	return scoreableHostServices, nil
}

func (f *FlagBearer) createChecks(ctx context.Context, scoreableHostServices []*entities.HostService, r *entities.Round) ([]*entities.Check, error) {
	// Create initial checks
	var initialChecks []*entities.CheckCreate
	for _, hs := range scoreableHostServices {
		chck := f.dbClient.Check.Create().SetHostservice(hs).SetRound(r).SetRoundID(r.ID)
		initialChecks = append(initialChecks, chck)
	}

	f.logger.Infow("Creating x number of checks", "num", len(initialChecks))
	// Save initial checks
	checks, err := f.dbClient.Check.CreateBulk(initialChecks...).Save(ctx)
	if err != nil {
		return nil, err
	}
	f.logger.Infof("Number of checks create %d", len(checks))
	f.logger.Infof("%v", checks[0])
	numOfChecks, err := f.dbClient.Check.Query().Where(check.RoundID(r.ID)).Count(ctx)
	f.logger.Infof("Number of checks created %d", numOfChecks)

	return checks, nil
}

func (f *FlagBearer) publishQueueMessages(ctx context.Context, r *entities.Round, scoreableHostServices []*entities.HostService) error {
	for _, hs := range scoreableHostServices {
		msg := messages.ChecksCreatedMessage{
			RoundID:       r.ID,
			TeamID:        hs.Edges.Team.ID,
			RoundNumber:   r.RoundNumber,
			HostAddress:   hs.Edges.Host.Address,
			CheckID:       hs.Edges.Checks[0].ID,
			ServiceType:   hs.Edges.Service.Type,
			HostServiceID: hs.ServiceID,
			Properties:    buildPropertiesFromList(hs.Edges.Properties),
		}

		payload, err := json.Marshal(msg)
		if err != nil {
			return err
		}

		_, err = f.js.PublishMsgAsync(&nats.Msg{
			Data: payload,
			//Subject: "PRINT",
			Subject: fmt.Sprintf("CHECKS.%s.created", hs.Edges.Team.ID),
		}, jetstream.WithMsgID(ulid.Make().String()))
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *FlagBearer) publishWatermillQueueMessages(ctx context.Context, r *entities.Round, scoreableHostServices []*entities.HostService) error {
	var msgs []*message.Message
	for _, hs := range scoreableHostServices {
		msg := messages.ChecksCreatedMessage{
			RoundID:       r.ID,
			RoundNumber:   r.RoundNumber,
			HostAddress:   hs.Edges.Host.Address,
			CheckID:       hs.Edges.Checks[0].ID,
			ServiceType:   hs.Edges.Service.Type,
			HostServiceID: hs.ServiceID,
			Properties:    buildPropertiesFromList(hs.Edges.Properties),
		}
		payload, err := json.Marshal(msg)
		if err != nil {
			return err
		}

		queuemsh := message.NewMessage(watermill.NewULID(), payload)
		msgs = append(msgs, queuemsh)
	}
	//err := f.queuePublisher.Publish(entries.TOPIC_HOST_SERVICE_SCORE, msgs...)
	//if err != nil {
	//	return err
	//}

	return nil
}

func (f *FlagBearer) filterServicesByRoundFrequency(roundNumber int, services []*entities.Service) []*entities.Service {
	var filteredServices []*entities.Service
	for _, s := range services {
		if IsServiceScoreableInRound(roundNumber, s.RoundFrequency, s.RoundDelay) {
			filteredServices = append(filteredServices, s)
		}
	}

	return filteredServices
}

func buildPropertiesFromList(properties []*entities.Property) map[string]string {
	newProperties := make(map[string]string)
	for _, prop := range properties {
		newProperties[prop.Key] = prop.Value
	}
	return newProperties
}

func IsServiceScoreableInRound(roundNumber int, roundFreq int, roundDelay int) bool {
	// Check if current round has passed the round delay
	if roundNumber >= roundDelay+1 {
		// Check if round is not the first round ???
		// ex. 1st round and 0 round delay will result in passing
		// ex. 4th round with 2 round delay will result in passing
		// ex. 2nd round with 2 round delay will result in passing.
		// Delay means not running until after the requirest round delay.
		// so if 3 round delay is requested, we will not run until after the 3rd round.
		if roundNumber-roundDelay != 0 {

			// Check if round frequency has been met.
			if (roundNumber-roundDelay)%roundFreq == 0 {
				return true
			}
		}
	}
	return false
}
