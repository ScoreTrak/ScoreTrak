package scheduler_test

import (
	"context"
	"github.com/nats-io/nats.go/jetstream"
	. "github.com/onsi/ginkgo/v2"
	"github.com/scoretrak/scoretrak/internal/entities"
	"github.com/scoretrak/scoretrak/internal/entities/enttest"
	"github.com/scoretrak/scoretrak/pkg/config"
	"github.com/scoretrak/scoretrak/pkg/eventsv2"
	"github.com/scoretrak/scoretrak/pkg/eventsv2/consumers"
	"github.com/scoretrak/scoretrak/pkg/eventsv2/streams"
	"github.com/scoretrak/scoretrak/pkg/flagbearer"
	"github.com/scoretrak/scoretrak/pkg/scheduler"
	"github.com/scoretrak/scoretrak/pkg/scheduler/job"
	"github.com/scoretrak/scoretrak/pkg/scorer"
	"github.com/scoretrak/scoretrak/pkg/scorer/executors"
	"github.com/scoretrak/scoretrak/pkg/telemetry"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
)

var _ = Describe("Scheduler", func() {
	var fb *flagbearer.FlagBearer
	var dbClient *entities.Client
	var logger *otelzap.SugaredLogger
	var scr *scorer.Scorer
	var schlr *scheduler.Scheduler

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

		//watermillLogger = telemetry.NewWatermillLogger(logger)
		//pub, sub, err = events.NewGochannelPubsub(cfg, watermillLogger)
		//if err != nil {
		//	t.Error(err)
		//}

		fb = flagbearer.NewFlagBearer(dbClient, js, logger)

		schlr, err = scheduler.NewCron(cfg)

		job.NewRoundStarterJob(fb, logger)

	})
})
