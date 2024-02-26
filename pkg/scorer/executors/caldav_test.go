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
//var _ = Describe("CalDavScorer", func() {
//	var cdp1, cdp2, cdp3, cdp4, cdp5 *executors.CalDavProperties
//
//	BeforeEach(func() {
//		cdp1 = &executors.CalDavProperties{
//			Scheme:         "https",
//			Host:           "caldavserver.com",
//			Port:           "443",
//			Path:           "/DAV/calendars/ffd1d96c-b628-4c86-afe2-ae7ae36ff20a",
//			Subdomain:      "www",
//			ExpectedOutput: "",
//			Username:       "U70bc38aa",
//			Password:       "8112fa22087a4427982a90ca1853d145",
//		}
//		cdp2 = &executors.CalDavProperties{
//			Scheme:         "https",
//			Host:           "caldavserver.com",
//			Port:           "443",
//			Path:           "/DAV/calendars/3c0cb7e5-4d00-4536-8c1b-02493d6a7c17",
//			Subdomain:      "www",
//			ExpectedOutput: "",
//			Username:       "U70bc38aa",
//			Password:       "8112fa22087a4427982a90ca1853d145",
//		}
//		cdp3 = &executors.CalDavProperties{
//			Scheme:         "https",
//			Host:           "caldavserver.com",
//			Port:           "443",
//			Path:           "/DAV/calendars/22c17430-1c34-4124-ad20-77838b3ebda6",
//			Subdomain:      "www",
//			ExpectedOutput: "",
//			Username:       "U70bc38aa",
//			Password:       "8112fa22087a4427982a90ca1853d145",
//		}
//		cdp4 = &executors.CalDavProperties{
//			Scheme:         "https",
//			Host:           "caldavserver.com",
//			Port:           "443",
//			Path:           "/DAV/calendars/22c17430-1c34-4124-ad20-77838b3ebda6",
//			Subdomain:      "www",
//			ExpectedOutput: "",
//			Username:       "U70bc38aa",
//			Password:       "8112fa22087a4427982a90ca1853d14",
//		}
//		cdp5 = &executors.CalDavProperties{
//			Scheme:         "https",
//			Host:           "caldavserver.com",
//			Port:           "443",
//			Path:           "/DAV/calendars/22c17430-1c34-4124-ad20-77838bebda6",
//			Subdomain:      "www",
//			ExpectedOutput: "",
//			Username:       "U70bc38a",
//			Password:       "8112fa22087a4427982a90ca1853d145",
//		}
//	})
//	Describe("Testing CalDav Scoring", func() {
//		It("When properties are correct #1", func() {
//			bs, _ := json.Marshal(cdp1)
//			outcome := executors.ScoreCalDav(context.Background(), bs)
//			gomega.Expect(outcome.Error).ToNot(gomega.HaveOccurred())
//			gomega.Expect(outcome.Passed).To(gomega.BeTrue())
//		})
//		It("When properties are correct #2", func() {
//			bs, _ := json.Marshal(cdp2)
//			outcome := executors.ScoreCalDav(context.Background(), bs)
//			gomega.Expect(outcome.Error).ToNot(gomega.HaveOccurred())
//			gomega.Expect(outcome.Passed).To(gomega.BeTrue())
//		})
//		It("When properties are correct #3", func() {
//			bs, _ := json.Marshal(cdp3)
//			outcome := executors.ScoreCalDav(context.Background(), bs)
//			gomega.Expect(outcome.Error).ToNot(gomega.HaveOccurred())
//			gomega.Expect(outcome.Passed).To(gomega.BeTrue())
//		})
//		It("When password is incorrect", func() {
//			bs, _ := json.Marshal(cdp4)
//			outcome := executors.ScoreCalDav(context.Background(), bs)
//			gomega.Expect(outcome.Error).To(gomega.HaveOccurred())
//			gomega.Expect(outcome.Passed).To(gomega.BeFalse())
//		})
//		It("When username and url is incorrect", func() {
//			bs, _ := json.Marshal(cdp5)
//			outcome := executors.ScoreCalDav(context.Background(), bs)
//			gomega.Expect(outcome.Error).To(gomega.HaveOccurred())
//			gomega.Expect(outcome.Passed).To(gomega.BeFalse())
//		})
//	})
//})
