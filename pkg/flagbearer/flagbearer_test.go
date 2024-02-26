package flagbearer_test

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/scoretrak/scoretrak/pkg/eventsv2"
	"github.com/scoretrak/scoretrak/pkg/eventsv2/consumers"
	"github.com/scoretrak/scoretrak/pkg/eventsv2/streams"
	"github.com/scoretrak/scoretrak/pkg/scorer"
	"github.com/scoretrak/scoretrak/pkg/scorer/executors"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"time"

	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/scoretrak/scoretrak/internal/entities"
	"github.com/scoretrak/scoretrak/internal/entities/check"
	"github.com/scoretrak/scoretrak/internal/entities/enttest"
	"github.com/scoretrak/scoretrak/internal/entities/round"
	"github.com/scoretrak/scoretrak/pkg/config"
	"github.com/scoretrak/scoretrak/pkg/flagbearer"
	"github.com/scoretrak/scoretrak/pkg/storage/seed"
	"github.com/scoretrak/scoretrak/pkg/telemetry"
)

var _ = Describe("Flagbearer", func() {
	var fb *flagbearer.FlagBearer
	var dbClient *entities.Client
	var logger *otelzap.SugaredLogger
	var scr *scorer.Scorer

	var js jetstream.JetStream
	var pccc jetstream.ConsumeContext
	var ccccc jetstream.ConsumeContext
	var csccc jetstream.ConsumeContext

	BeforeEach(func() {
		t := GinkgoT()
		ctx := context.Background()
		cfg, err := config.NewScoreTrakConfig()
		if err != nil {
			t.Error(err)
		}
		dbClient = enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
		logger, err = telemetry.NewLogger()
		if err != nil {
			t.Error(err)
		}

		scr = scorer.NewScorer(scorer.WithLogger(logger))
		scr.Handle(scorer.SERVICE_DNS, executors.ScoreDns)
		scr.Handle(scorer.SERVICE_HTTP, executors.ScoreHttp)
		scr.Handle(scorer.SERVICE_PING, executors.ScorePing)

		js, err = eventsv2.NewNats(cfg)
		if err != nil {
			t.Error(err)
		}

		_, err = streams.NewPrintStream(ctx, js)
		if err != nil {
			t.Error(err)
		}
		_, err = streams.NewCheckStream(ctx, js)
		if err != nil {
			t.Error(err)
		}

		pc, err := consumers.NewPrintConsumer(ctx, js)
		if err != nil {
			t.Error(err)
		}
		pccc, err = pc.Start(ctx)
		if err != nil {
			t.Error(err)
		}
		ccc, err := consumers.NewChecksCreatedConsumer(ctx, js, dbClient, scr, logger)
		if err != nil {
			t.Error(err)
		}
		ccccc, err = ccc.Start(ctx)
		if err != nil {
			t.Error(err)
		}
		csc, err := consumers.NewChecksScoredConsumer(ctx, js, dbClient, scr, logger)
		if err != nil {
			t.Error(err)
		}
		csccc, err = csc.Start(ctx)
		if err != nil {
			t.Error(err)
		}

		fb = flagbearer.NewFlagBearer(dbClient, js, logger)
	})

	Describe("Ensure round score-able calculation", func() {
		It("5th Round with 2 round delay", func() {
			running := flagbearer.IsServiceScoreableInRound(5, 1, 2)
			gomega.Expect(running).To(gomega.BeTrue())
		})
		It("1st Round with 0 round delay", func() {
			running := flagbearer.IsServiceScoreableInRound(1, 1, 0)
			gomega.Expect(running).To(gomega.BeTrue())
		})
		It("1st Round with 1 round delay", func() {
			running := flagbearer.IsServiceScoreableInRound(1, 1, 1)
			gomega.Expect(running).To(gomega.BeFalse())
		})
		It("1st Round with 2 round delay", func() {
			running := flagbearer.IsServiceScoreableInRound(1, 1, 2)
			gomega.Expect(running).To(gomega.BeFalse())
		})

		It("5th Round with 2 round delay", func() {
			running := flagbearer.IsServiceScoreableInRound(5, 1, 2)
			gomega.Expect(running).To(gomega.BeTrue())
		})
	})

	Describe("New Competition", func() {
		BeforeEach(func() {
			//t := GinkgoT()
			//hsh := handlers.NewHostServiceScoreHandler(dbClient, scr, logger)
			//csh := handlers.NewCheckSaveHandler(dbClient, logger)
			//hshe := entries.NewHostServiceScoreHandlerEntry(pub, sub, hsh)
			//cshe := entries.NewCheckSaveNoPublishHandlerEntry(sub, csh)
			//var err error
			//rtr, err = events.NewRouter([]*events.HandlerEntry{hshe}, []*events.NoPublishHandlerEntry{cshe}, watermillLogger)
			//if err != nil {
			//	t.Error(err)
			//}
			seed.DevSeed(context.Background(), dbClient)
		})

		It("Able to start round and create checks", func() {
			ctx := context.Background()
			//defer func(rtr *message.Router) {
			//	err := rtr.Close()
			//	if err != nil {
			//		panic(err)
			//	}
			//}(rtr)
			//go func() {
			//	err := rtr.Run(ctx)
			//	if err != nil {
			//		panic(err)
			//	}
			//}()
			//<-rtr.Running()

			err := fb.StartNextRound(ctx)
			gomega.Expect(err).ToNot(gomega.HaveOccurred())

			r, err := dbClient.Round.Query().Where(round.RoundNumber(1)).First(ctx)
			gomega.Expect(r.RoundNumber).To(gomega.Equal(1))
			gomega.Expect(r.Status).To(gomega.Equal(round.StatusOngoing))

			numOfChecks, err := dbClient.Check.Query().Where(check.RoundID(r.ID)).Count(ctx)
			gomega.Expect(err).ToNot(gomega.HaveOccurred())
			gomega.Expect(numOfChecks).To(gomega.Equal(7))

			time.Sleep(2 * time.Second)

			fmt.Println(dbClient.Check.Query().AllX(ctx))
			numOfFinishedChecks, err := dbClient.Check.Query().Where(check.ProgressStatusEQ(check.ProgressStatusFinished), check.RoundID(r.ID)).Count(ctx)
			gomega.Expect(err).ToNot(gomega.HaveOccurred())
			gomega.Expect(numOfFinishedChecks).To(gomega.Equal(7))

			//samer, err := dbClient.Round.Query().Where(round.RoundNumber(1)).First(ctx)
			//gomega.Expect(samer.RoundNumber).To(gomega.Equal(1))
			//gomega.Expect(samer.Status).To(gomega.Equal(round.StatusFinished))
			//
			//// Start second round
			//
			//err = fb.StartNextRound(ctx)
			//gomega.Expect(err).ToNot(gomega.HaveOccurred())
			//
			//r2, err := dbClient.Round.Query().Where(round.RoundNumber(2)).First(ctx)
			//gomega.Expect(r2.RoundNumber).To(gomega.Equal(2))
			//gomega.Expect(r2.Status).To(gomega.Equal(round.StatusOngoing))
			//
			//numOfChecks2, err := dbClient.Check.Query().Where(check.RoundID(r2.ID)).Count(ctx)
			//gomega.Expect(err).ToNot(gomega.HaveOccurred())
			//gomega.Expect(numOfChecks2).To(gomega.Equal(7))
			//
			//time.Sleep(2 * time.Second)

		})

	})

	AfterEach(func() {
		err := dbClient.Close()
		if err != nil {
			return
		}
		pccc.Stop()
		ccccc.Stop()
		csccc.Stop()
		// err = rtr.Close()
	})
})
