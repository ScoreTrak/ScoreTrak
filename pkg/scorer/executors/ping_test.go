package executors_test

//
//import (
//	"context"
//	"encoding/json"
//	"fmt"
//	"github.com/scoretrak/scoretrak/pkg/scorer/executors"
//	. "github.com/onsi/ginkgo/v2"
//	"github.com/onsi/gomega"
//)
//
//var _ = Describe("PingScorer", func() {
//	var pingp1, pingp2, pingp3 *executors.PingProperties
//
//	BeforeEach(func() {
//		pingp1 = &executors.PingProperties{
//			Host:     "1.1.1.1",
//			Protocol: "ipv4",
//			Attempts: "3",
//		}
//
//		pingp2 = &executors.PingProperties{
//			Host:     "8.8.8.8",
//			Protocol: "ipv4",
//			Attempts: "3",
//		}
//
//		pingp3 = &executors.PingProperties{
//			Host:     "169.254.0.1",
//			Protocol: "ipv4",
//			Attempts: "1",
//		}
//	})
//
//	Describe("Testing PING Scoring", func() {
//		It("When pinging example.com succeeds", func() {
//			bs, _ := json.Marshal(pingp1)
//			outcome := executors.ScoreHttp(context.Background(), bs)
//			fmt.Println(outcome.Error.Error())
//			gomega.Expect(outcome.Passed).To(gomega.BeTrue())
//		})
//
//		It("When pinging google.com succeeds", func() {
//			bs, _ := json.Marshal(pingp2)
//			outcome := executors.ScoreHttp(context.Background(), bs)
//			fmt.Println(outcome.Error.Error())
//			gomega.Expect(outcome.Passed).To(gomega.BeTrue())
//		})
//
//		It("When pinging app1e.com fails", func() {
//			bs, _ := json.Marshal(pingp3)
//			outcome := executors.ScoreHttp(context.Background(), bs)
//			gomega.Expect(outcome.Passed).To(gomega.BeFalse())
//		})
//	})
//
//})
