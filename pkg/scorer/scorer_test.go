package scorer_test

import (
	"context"
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/scoretrak/scoretrak/pkg/scorer"
	"github.com/scoretrak/scoretrak/pkg/scorer/executors"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

var _ = Describe("Scorer", func() {
	var scorer1 *scorer.Scorer
	dnsProperties := []executors.DNSProperties{
		{
			RecordType:     "A",
			Host:           "one.one.one.one",
			Lookup:         "one.one.one.one",
			ExpectedOutput: "1.1.1.1",
		},
	}
	httpProperties := []executors.HTTPProperties{
		{
			ExpectedOutput:       "",
			ExpectedResponseCode: "200",
			Scheme:               "https",
			Host:                 "google.com",
			Port:                 "443",
			Path:                 "/",
			Subdomain:            "www",
		},
		{
			ExpectedOutput:       "",
			ExpectedResponseCode: "200",
			Scheme:               "https",
			Host:                 "google.com",
			Port:                 "443",
			Path:                 "/",
			Subdomain:            "www4",
		},
	}
	pingProperties := []executors.PingProperties{
		{
			Host:     "1.1.1.1",
			Protocol: "ipv4",
			Attempts: "3",
		},
	}

	BeforeEach(func() {
		l, _ := zap.NewDevelopment()
		ol := otelzap.New(l).Sugar()
		scorer1 = scorer.NewScorer(
			scorer.WithLogger(ol),
		)

		//scorer1.Handle("caldav", executors.ScoreCalDav)
		scorer1.Handle(scorer.SERVICE_DNS, executors.ScoreDns) //Just use google or cloudflare dns
		//scorer1.Handle("ftp", executors.ScoreFTP)
		scorer1.Handle(scorer.SERVICE_HTTP, executors.ScoreHttp) //Just user any website
		//scorer1.Handle("imap", executors.ScoreImap) //Iguess
		//scorer1.Handle("ldap", executors.ScoreLdap) //dummy server needed as well
		scorer1.Handle(scorer.SERVICE_PING, executors.ScorePing) //Any server
		//scorer1.Handle("smb", executors.ScoreSmb)
		//scorer1.Handle("sql", executors.ScoreSQL) //need dummy server
		//scorer1.Handle("ssh", executors.ScoreSSH) // dummy ssh server needed
		//scorer1.Handle("winrm", executors.ScoreWinrm) // IDK

	})

	Describe("Testing Scorer with dns handler", func() {
		It("Testing https://www.google.com", func() {
			ctx := context.Background()
			outcome := scorer1.Score(ctx, scorer.SERVICE_DNS, dnsProperties[0])
			gomega.Expect(outcome.Error).To(gomega.BeNil())
		})
	})

	Describe("Testing Scorer with http handler", func() {
		It("Testing https://www.google.com", func() {
			ctx := context.Background()
			outcome := scorer1.Score(ctx, scorer.SERVICE_HTTP, httpProperties[0])
			gomega.Expect(outcome.Error).To(gomega.BeNil())
		})
		It("Testing https://www4.google.com", func() {
			ctx := context.Background()
			outcome := scorer1.Score(ctx, scorer.SERVICE_HTTP, httpProperties[1])
			gomega.Expect(outcome.Error).To(gomega.Not(gomega.BeNil()))
		})
	})

	Describe("Testing Scorer with ping handler", func() {
		It("Testing 1.1.1.1", func() {
			ctx := context.Background()
			outcome := scorer1.Score(ctx, scorer.SERVICE_PING, pingProperties[0])
			gomega.Expect(outcome.Error).To(gomega.BeNil())
		})
	})

})
