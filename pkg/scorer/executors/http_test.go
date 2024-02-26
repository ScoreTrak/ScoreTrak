package executors_test

//
//import (
//	"context"
//	"encoding/json"
//	"github.com/scoretrak/scoretrak/pkg/scorer/executors"
//	. "github.com/onsi/ginkgo/v2"
//	"github.com/onsi/gomega"
//)
//
//var _ = Describe("HttpScorer", func() {
//	var httpp1, httpp2, httpp3 *executors.HTTPProperties
//
//	BeforeEach(func() {
//		httpp1 = &executors.HTTPProperties{
//			Host:                 "example.com",
//			Port:                 "80",
//			Scheme:               "http",
//			ExpectedOutput:       "",
//			ExpectedResponseCode: "200",
//			Path:                 "/",
//			Subdomain:            "www",
//		}
//
//		httpp2 = &executors.HTTPProperties{
//			Host:                 "app1e.com",
//			Port:                 "443",
//			Scheme:               "https",
//			ExpectedOutput:       "",
//			ExpectedResponseCode: "200",
//			Path:                 "/",
//			Subdomain:            "",
//		}
//
//		httpp3 = &executors.HTTPProperties{
//			Host:                 "google.com",
//			Port:                 "443",
//			Scheme:               "https",
//			ExpectedOutput:       "",
//			ExpectedResponseCode: "200",
//			Path:                 "/",
//			Subdomain:            "",
//		}
//	})
//
//	Describe("Testing HTTP Scoring", func() {
//		It("When expecting https://example.com to return 200", func() {
//			bs, _ := json.Marshal(httpp1)
//			outcome := executors.ScoreHttp(context.Background(), bs)
//			gomega.Expect(outcome.Passed).To(gomega.BeTrue())
//		})
//
//		It("When apple.com has wrong address", func() {
//			bs, _ := json.Marshal(httpp2)
//			outcome := executors.ScoreHttp(context.Background(), bs)
//			gomega.Expect(outcome.Passed).To(gomega.BeFalse())
//		})
//
//		It("When https://google.com returns 200", func() {
//			bs, _ := json.Marshal(httpp3)
//			outcome := executors.ScoreHttp(context.Background(), bs)
//			gomega.Expect(outcome.Passed).To(gomega.BeTrue())
//		})
//	})
//})
