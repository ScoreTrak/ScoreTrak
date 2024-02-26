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
//var _ = Describe("FtpScorer", func() {
//	var ftpp1, ftpp2, ftpp3, ftpp5 *executors.FTPProperties
//
//	BeforeEach(func() {
//		ftpp1 = &executors.FTPProperties{
//			Username: "dlpuser",
//			Password: "rNrKYTX9g7z3RgJRmxWuGHbeu",
//			Host:     "ftp.dlptest.com",
//			Port:     "21",
//			//Text:          "Bye",
//			//WriteFilename: "bye.txt",
//			//ReadFilename:   "1MB.zip",
//			//ExpectedOutput: "Hello",
//		}
//
//		ftpp2 = &executors.FTPProperties{
//			Username:       "",
//			Password:       "",
//			Host:           "speedtest.tele2.net",
//			Port:           "21",
//			Text:           "",
//			WriteFilename:  "",
//			ReadFilename:   "",
//			ExpectedOutput: "",
//		}
//
//		ftpp3 = &executors.FTPProperties{
//			Username:       "dlpuser",
//			Password:       "rNrKYTX9g7z3RgJRmxWuGHbeu",
//			Host:           "ftp.dlptest.com",
//			Port:           "21",
//			Text:           "Bye",
//			WriteFilename:  "bye.txt",
//			ReadFilename:   "",
//			ExpectedOutput: "",
//		}
//
//		ftpp5 = &executors.FTPProperties{
//			Username:       "dlpuser",
//			Password:       "rNrKYTX9g7z3RgJRmxWuGHbeu",
//			Host:           "ftp.dlptest.com",
//			Port:           "21",
//			Text:           "",
//			WriteFilename:  "",
//			ReadFilename:   "",
//			ExpectedOutput: "",
//		}
//	})
//
//	Describe("Testing FTP Scoring", func() {
//		It("When correct credentials are provided", func() {
//			bs, _ := json.Marshal(ftpp1)
//			outcome := executors.ScoreFTP(context.Background(), bs)
//			gomega.Expect(outcome.Passed).To(gomega.BeTrue())
//		})
//
//		It("When the ftp server is anon", func() {
//			bs, _ := json.Marshal(ftpp2)
//			outcome := executors.ScoreFTP(context.Background(), bs)
//			gomega.Expect(outcome.Passed).To(gomega.BeTrue())
//		})
//
//		It("When uploading a file", func() {
//			bs, _ := json.Marshal(ftpp3)
//			outcome := executors.ScoreFTP(context.Background(), bs)
//			gomega.Expect(outcome.Passed).To(gomega.BeTrue())
//		})
//
//		It("When reading and writing a file", func() {
//			bs, _ := json.Marshal(ftpp5)
//			outcome := executors.ScoreFTP(context.Background(), bs)
//			gomega.Expect(outcome.Passed).To(gomega.BeTrue())
//		})
//	})
//
//})
