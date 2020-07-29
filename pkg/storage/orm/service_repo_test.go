package orm

import (
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/check"
	"github.com/L1ghtman2k/ScoreTrak/pkg/config"
	"github.com/L1ghtman2k/ScoreTrak/pkg/host"
	"github.com/L1ghtman2k/ScoreTrak/pkg/property"
	"github.com/L1ghtman2k/ScoreTrak/pkg/service"
	"github.com/L1ghtman2k/ScoreTrak/pkg/service_group"
	. "github.com/L1ghtman2k/ScoreTrak/test"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func FTPSpec(t *testing.T) {
	var c config.StaticConfig
	autoTest := os.Getenv("AUTO_TEST")
	if autoTest == "TRUE" {
		c = NewConfigClone(SetupConfig("../../../configs/test-config.yml"))
	} else {
		c = NewConfigClone(SetupConfig("dev-config.yml"))
	}
	c.DB.Cockroach.Database = "scoretrak_test_orm_service"
	c.Logger.FileName = "service_test.log"
	db := SetupDB(c.DB)
	l := SetupLogger(c.Logger)
	t.Parallel() //t.Parallel should be placed after SetupDB because gorm has race conditions on Hook register
	Convey("Creating Service, Service Group, Host Tables along with their foreign keys", t, func() {
		db.AutoMigrate(&service.Service{})
		db.AutoMigrate(&service_group.ServiceGroup{})
		db.AutoMigrate(&host.Host{})
		sr := NewServiceRepo(db, l)
		Convey("When all tables are empty", func() {
			Convey("Should output no entry", func() {
				gt, err := sr.GetAll()
				So(err, ShouldBeNil)
				So(len(gt), ShouldEqual, 0)
			})
			Convey("Creating a sample service should not be allowed", func() {
				s := service.Service{Name: "FTP"}
				err := sr.Store(&s)
				So(err, ShouldNotBeNil)
				ac, err := sr.GetAll()
				So(err, ShouldBeNil)
				So(len(ac), ShouldEqual, 0)
			})
			Convey("Load Sample Service Group, And Host Data", func() {
				var count int64
				db.Exec("INSERT INTO hosts (id, address) VALUES (5, '192.168.1.2')")
				db.Exec("INSERT INTO hosts (id, address) VALUES (4, '192.168.1.1')")
				db.Exec("INSERT INTO service_groups (id, name) VALUES (7, 'FTPGroup')")
				db.Exec("INSERT INTO service_groups (id, name) VALUES (2, 'FTPGroup2')")
				db.Table("hosts").Count(&count)
				So(count, ShouldEqual, 2)
				db.Table("service_groups").Count(&count)
				So(count, ShouldEqual, 2)

				Convey("Creating a sample service and associating with host id 5, and service group id 2", func() {
					s := service.Service{Name: "FTP", ServiceGroupID: 2, HostID: 5}
					err := sr.Store(&s)
					Convey("Should be Allowed", func() {
						So(err, ShouldBeNil)
						ac, err := sr.GetAll()
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)

						Convey("Then Querying By ID", func() {
							ss, err := sr.GetByID(s.ID)
							So(err, ShouldBeNil)
							So(ss.Name, ShouldEqual, "FTP")
							So(ss.ServiceGroupID, ShouldEqual, 2)
							So(ss.HostID, ShouldEqual, 5)

						})

						Convey("Then Querying By wrong ID", func() {
							ss, err := sr.GetByID(s.ID + 1)
							So(err, ShouldNotBeNil)
							So(ss, ShouldBeNil)
						})

						Convey("Then Deleting a wrong entry", func() {
							err = sr.Delete(s.ID + 1)
							So(err, ShouldNotBeNil)
							Convey("Should output one entry", func() {
								ac, err := sr.GetAll()
								So(err, ShouldBeNil)
								So(len(ac), ShouldEqual, 1)
							})
						})

						Convey("Then updating to non existent host should not be allowed", func() {
							s.HostID = 20
							err = sr.Update(&s)
							So(err, ShouldNotBeNil)
						})
						Convey("Then updating to different existent host/service_group should be allowed", func() {
							s.HostID = 4
							err = sr.Update(&s)
							So(err, ShouldBeNil)
						})
						Convey("Then adding service with the same name should be allowed", func() {
							s2 := service.Service{Name: "FTP", ServiceGroupID: 7, HostID: 5}
							err = sr.Store(&s2)
							So(err, ShouldBeNil)
						})
						Convey("Then updating regular fields should be allowed", func() {
							tru := true
							s.Enabled = &tru
							s.Name = "DifferentTestName"
							s.RoundUnits = 3
							rd := uint32(2)
							s.RoundDelay = &rd
							s.Points = 5
							err = sr.Update(&s)
							So(err, ShouldBeNil)
							ac, err = sr.GetAll()
							So(err, ShouldBeNil)
							So(len(ac), ShouldEqual, 1)
							So(ac[0].Name, ShouldEqual, "DifferentTestName")
							So(*(ac[0].RoundDelay), ShouldEqual, 2)
						})
						Convey("Then updating Round Delay to something larger than Round Units should not be allowed", func() {
							rd := uint32(5)
							s.RoundDelay = &rd
							Convey("Round Units set", func() {
								s.RoundUnits = 3
								err = sr.Update(&s)
								So(err, ShouldNotBeNil)
								ac, err = sr.GetAll()
								So(err, ShouldBeNil)
								So(len(ac), ShouldEqual, 1)
								So(*(ac[0].RoundDelay), ShouldEqual, 0)
							})
							Convey("Round Units not set", func() {
								err = sr.Update(&s)
								So(err, ShouldNotBeNil)
								ac, err = sr.GetAll()
								So(err, ShouldBeNil)
								So(len(ac), ShouldEqual, 1)
								So(*(ac[0].RoundDelay), ShouldEqual, 0)
							})

						})
					})
				})
				Convey("Create Multiple services", func() {
					s1 := service.Service{Name: "FTP", ServiceGroupID: 2, HostID: 4}
					s2 := service.Service{Name: "FTP", ServiceGroupID: 7, HostID: 5}
					err := sr.Store(&s1)
					So(err, ShouldBeNil)
					err = sr.Store(&s2)
					So(err, ShouldBeNil)
					Convey("Loading Checks And Properties Tables with sample data", func() {
						db.AutoMigrate(&property.Property{})
						db.AutoMigrate(&check.Check{})
						db.Exec(fmt.Sprintf("INSERT INTO checks (service_id, round_id, passed) VALUES (%d, 999, false)", s1.ID))
						db.Exec(fmt.Sprintf("INSERT INTO checks (service_id, round_id, passed) VALUES (%d, 333, false)", s2.ID))
						db.Exec(fmt.Sprintf("INSERT INTO properties (id, service_id, key, value) VALUES (5, %d, 'sample_key', 'sample_value')", s1.ID))
						db.Exec(fmt.Sprintf("INSERT INTO properties (id, service_id, key, value) VALUES (33, %d, 'sample_key', 'sample_value')", s2.ID))
						db.Table("checks").Count(&count)
						So(count, ShouldEqual, 2)
						db.Table("properties").Count(&count)
						So(count, ShouldEqual, 2)
						Convey("Deleting the service", func() {
							err = sr.Delete(s1.ID)
							So(err, ShouldBeNil)
							Convey("Should also delete checks and properties associated with the deleted service", func() {
								var count int64
								db.Table("properties").Count(&count)
								So(count, ShouldEqual, 1)
								db.Table("checks").Count(&count)
								So(count, ShouldEqual, 1)
							})
						})
						Reset(func() {
							db.Migrator().DropTable(&property.Property{})
							db.Migrator().DropTable(&check.Check{})
						})
					})
				})

				Convey("Creating a sample service with wrong service_group should not be allowed", func() {
					s := service.Service{Name: "FTP", ServiceGroupID: 88, HostID: 5}
					err := sr.Store(&s)
					So(err, ShouldNotBeNil)
					ac, err := sr.GetAll()
					So(err, ShouldBeNil)
					So(len(ac), ShouldEqual, 0)
				})

			})
		})
		Reset(func() {
			db.Migrator().DropTable(&service.Service{})
			db.Migrator().DropTable(&host.Host{})
			db.Migrator().DropTable(&service_group.ServiceGroup{})
		})
	})
	DropDB(db, c)

}
