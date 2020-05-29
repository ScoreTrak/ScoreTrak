package orm

import (
	"ScoreTrak/pkg/check"
	"ScoreTrak/pkg/host"
	"ScoreTrak/pkg/property"
	"ScoreTrak/pkg/service"
	"ScoreTrak/pkg/service_group"
	. "ScoreTrak/test"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestServiceSpec(t *testing.T) {
	c := NewConfigClone(SetupConfig("dev-config.yml"))
	c.DB.Cockroach.Database = "scoretrak_test_orm_service"
	c.Logger.FileName = "service_test_repo.log"
	db := SetupDB(c)
	l := SetupLogger(c)
	t.Parallel() //t.Parallel should be placed after SetupDB because gorm has race conditions on Hook register
	Convey("Creating Service, Service Group, Host Tables along with their foreign keys", t, func() {
		db.AutoMigrate(&service.Service{})
		db.AutoMigrate(&service_group.ServiceGroup{})
		db.AutoMigrate(&host.Host{})
		db.Model(&service.Service{}).AddForeignKey("service_group_id", "service_groups(id)", "RESTRICT", "RESTRICT")
		db.Model(&service.Service{}).AddForeignKey("host_id", "hosts(id)", "RESTRICT", "RESTRICT")
		sr := NewServiceRepo(db, l)
		Convey("When all tables are empty", func() {
			Convey("Should output no entry", func() {
				gt, err := sr.GetAll()
				So(err, ShouldBeNil)
				So(len(gt), ShouldEqual, 0)
			})
			Convey("Creating a sample service should not be allowed", func() {
				s := service.Service{Name: "TestService"}
				err := sr.Store(&s)
				So(err, ShouldNotBeNil)
				ac, err := sr.GetAll()
				So(err, ShouldBeNil)
				So(len(ac), ShouldEqual, 0)
			})
			Convey("Load Sample Service Group, And Host Data", func() {
				var count int
				db.Exec("INSERT INTO hosts (id, address) VALUES (5, '192.168.1.2')")
				db.Exec("INSERT INTO hosts (id, address) VALUES (4, '192.168.1.1')")
				db.Exec("INSERT INTO service_groups (id, name) VALUES (7, 'TestServiceGroup')")
				db.Exec("INSERT INTO service_groups (id, name) VALUES (2, 'TestServiceGroup2')")
				db.Table("hosts").Count(&count)
				So(count, ShouldEqual, 2)
				db.Table("service_groups").Count(&count)
				So(count, ShouldEqual, 2)

				Convey("Creating a sample service and associating with host id 5, and service group id 2", func() {
					s := service.Service{Name: "TestService", ServiceGroupID: 2, HostID: 5}
					err := sr.Store(&s)
					Convey("Should be Allowed", func() {
						So(err, ShouldBeNil)
						ac, err := sr.GetAll()
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)

						Convey("Then Querying By ID", func() {
							ss, err := sr.GetByID(s.ID)
							So(err, ShouldBeNil)
							So(ss.Name, ShouldEqual, "TestService")
							So(ss.ServiceGroupID, ShouldEqual, 2)
							So(ss.HostID, ShouldEqual, 5)

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
							s2 := service.Service{Name: "TestService", ServiceGroupID: 7, HostID: 5}
							err = sr.Store(&s2)
							So(err, ShouldBeNil)
						})
						Convey("Then updating regular fields should be allowed", func() {
							tru := true
							s.Enabled = &tru
							s.Name = "DifferentTestName"
							s.RoundUnits = 3
							rd := uint64(5)
							s.RoundDelay = &rd
							s.Points = 5
							err = sr.Update(&s)
							So(err, ShouldBeNil)
							ac, err = sr.GetAll()
							So(err, ShouldBeNil)
							So(len(ac), ShouldEqual, 1)
							So(ac[0].Name, ShouldEqual, "DifferentTestName")
						})
					})
				})
				Convey("Create Multiple services", func() {
					s1 := service.Service{Name: "TestService", ServiceGroupID: 2, HostID: 4}
					s2 := service.Service{Name: "TestService2", ServiceGroupID: 7, HostID: 5}
					err := sr.Store(&s1)
					So(err, ShouldBeNil)
					err = sr.Store(&s2)
					So(err, ShouldBeNil)
					Convey("Loading Checks And Properties Tables with sample data", func() {
						db.AutoMigrate(&property.Property{})
						db.AutoMigrate(&check.Check{})
						db.Model(&check.Check{}).AddForeignKey("service_id", "services(id)", "CASCADE", "RESTRICT")
						db.Model(&property.Property{}).AddForeignKey("service_id", "services(id)", "CASCADE", "RESTRICT")
						db.Exec(fmt.Sprintf("INSERT INTO checks (id, service_id, round_id, passed) VALUES (5, %d, 999, false)", s1.ID))
						db.Exec(fmt.Sprintf("INSERT INTO checks (id, service_id, round_id, passed) VALUES (42, %d, 333, false)", s2.ID))
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
								var count int
								db.Table("properties").Count(&count)
								So(count, ShouldEqual, 1)
								db.Table("checks").Count(&count)
								So(count, ShouldEqual, 1)
							})
						})
						Reset(func() {
							db.DropTableIfExists(&property.Property{})
							db.DropTableIfExists(&check.Check{})
						})
					})
				})

				Convey("Creating a sample service with wrong service_group should not be allowed", func() {
					s := service.Service{Name: "TestService", ServiceGroupID: 88, HostID: 5}
					err := sr.Store(&s)
					So(err, ShouldNotBeNil)
					ac, err := sr.GetAll()
					So(err, ShouldBeNil)
					So(len(ac), ShouldEqual, 0)
				})

			})
		})
		Reset(func() {
			db.DropTableIfExists(&service.Service{})
			db.DropTableIfExists(&host.Host{})
			db.DropTableIfExists(&service_group.ServiceGroup{})
		})
	})
	DropDB(db, c)
	db.Close()
}

//TODO: Test that only available services can be set. Ex: FTP, HTTP, etc
