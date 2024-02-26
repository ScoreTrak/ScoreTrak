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
//var _ = Describe("DnsScorer", func() {
//	var dnsp1, dnsp2, dnsp3 *executors.DNSProperties
//
//	BeforeEach(func() {
//		dnsp1 = &executors.DNSProperties{
//			Host:           "dns.google.com",
//			Lookup:         "dns.google.com",
//			ExpectedOutput: "8.8.8.8",
//		}
//
//		dnsp2 = &executors.DNSProperties{
//			Host:           "one.one.one.one",
//			Lookup:         "one.one.one.one",
//			ExpectedOutput: "1.1.1.1",
//		}
//
//		dnsp3 = &executors.DNSProperties{
//			Host:           "dns.google.com",
//			Lookup:         "apple.com",
//			ExpectedOutput: "1.1.1.1",
//		}
//	})
//
//	Describe("Testing Dns Scoring", func() {
//		It("When looking up dns.google.com", func() {
//			bs, _ := json.Marshal(dnsp1)
//			outcome := executors.ScoreDns(context.Background(), bs)
//			gomega.Expect(outcome.Error).ToNot(gomega.HaveOccurred())
//			gomega.Expect(outcome.Passed).To(gomega.BeTrue())
//		})
//
//		It("When looking up one.one.one.one", func() {
//			bs, _ := json.Marshal(dnsp2)
//			outcome := executors.ScoreDns(context.Background(), bs)
//			gomega.Expect(outcome.Error).ToNot(gomega.HaveOccurred())
//			gomega.Expect(outcome.Passed).To(gomega.BeTrue())
//		})
//
//		It("When looking up apple.com", func() {
//			bs, _ := json.Marshal(dnsp3)
//			outcome := executors.ScoreDns(context.Background(), bs)
//			gomega.Expect(outcome.Error).To(gomega.HaveOccurred())
//			gomega.Expect(outcome.Passed).To(gomega.BeFalse())
//		})
//	})
//
//})
