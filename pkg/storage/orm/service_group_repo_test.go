package orm

import (
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/config"
	"github.com/L1ghtman2k/ScoreTrak/pkg/service"
	"github.com/L1ghtman2k/ScoreTrak/pkg/service_group"
	. "github.com/L1ghtman2k/ScoreTrak/test"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestServiceGroupSpec(t *testing.T) {
	var c config.StaticConfig
	autoTest := os.Getenv("AUTO_TEST")
	if autoTest == "TRUE" {
		c = NewConfigClone(SetupConfig("../../../configs/test-config.yml"))
	} else {
		c = NewConfigClone(SetupConfig("dev-config.yml"))
	}
	c.DB.Cockroach.Database = "scoretrak_test_orm_service_group"
	c.Logger.FileName = "service_group_test.log"
	db := SetupDB(c)
	l := SetupLogger(c)
	t.Parallel() //t.Parallel should be placed after SetupDB because gorm has race conditions on Hook register
	Convey("Creating Service Group Tables", t, func() {
		db.AutoMigrate(&service_group.ServiceGroup{})
		sgr := NewServiceGroupRepo(db, l)

		Convey("When the Service Group table is empty", func() {
			Convey("There should be no entries", func() {
				ac, err := sgr.GetAll()
				So(err, ShouldBeNil)
				So(len(ac), ShouldEqual, 0)
			})

			Convey("Adding a valid entry", func() {
				var err error
				s := service_group.ServiceGroup{Name: "TestServiceGroup"}
				err = sgr.Store(&s)
				So(err, ShouldBeNil)

				Convey("Should create an entry in the database", func() {
					ac, err := sgr.GetAll()
					So(err, ShouldBeNil)
					So(len(ac), ShouldEqual, 1)
				})

				Convey("And then creating an entry with same name", func() {
					tru := true
					t2 := service_group.ServiceGroup{Name: "TestServiceGroup", Enabled: &tru}
					err = sgr.Store(&t2)
					So(err, ShouldNotBeNil)
					Convey("Should not create a new entry", func() {
						ac, err := sgr.GetAll()
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
					})
				})

				Convey("Then Deleting a wrong entry", func() {
					err = sgr.Delete(s.ID + 1)
					So(err, ShouldNotBeNil)
					Convey("Should output one entry", func() {
						ac, err := sgr.GetAll()
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
					})
				})

				Convey("Then Deleting the added entry", func() {
					err = sgr.Delete(s.ID)
					So(err, ShouldBeNil)
					Convey("Should output no entries", func() {
						ac, err := sgr.GetAll()
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 0)
					})
				})

				Convey("Then Retrieving entry by ID", func() {
					sg, err := sgr.GetByID(s.ID)
					So(err, ShouldBeNil)
					Convey("Should output the inserted entry", func() {
						So(sg.Name, ShouldEqual, "TestServiceGroup")
						So(*(sg.Enabled), ShouldBeFalse)
					})
				})

				Convey("Then Querying By wrong ID", func() {
					ss, err := sgr.GetByID(s.ID + 1)
					So(err, ShouldNotBeNil)
					So(ss, ShouldBeNil)
				})

				Convey("Then Updating Enabled to true", func() {
					tru := true
					newSgr := &service_group.ServiceGroup{Enabled: &tru}
					Convey("For the wrong entry should not update anything", func() {
						newSgr.ID = s.ID + 1
						err = sgr.Update(newSgr)
						So(err, ShouldBeNil)
						ac, err := sgr.GetAll()
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
						So(*(ac[0].Enabled), ShouldBeFalse)

					})
					Convey("For the correct entry should update", func() {
						newSgr.ID = s.ID
						err = sgr.Update(newSgr)
						So(err, ShouldBeNil)
						ac, err := sgr.GetAll()
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
						So(*(ac[0].Enabled), ShouldBeTrue)
					})
				})
				Convey("Creating Service Table", func() {
					var count int
					db.AutoMigrate(&service.Service{})
					db.Model(&service.Service{}).AddForeignKey("service_group_id", "service_groups(id)", "RESTRICT", "RESTRICT")
					Convey("Then associating one service with the service group", func() {
						db.Exec(fmt.Sprintf("INSERT INTO services (id, service_group_id, host_id, name) VALUES (5, %d, 999, 'TestService')", s.ID))
						db.Table("services").Count(&count)
						So(count, ShouldEqual, 1)
						Convey("Then Deleting the service group should be restricted", func() {
							err = sgr.Delete(s.ID)
							So(err, ShouldNotBeNil)
							ac, err := sgr.GetAll()
							So(err, ShouldBeNil)
							So(len(ac), ShouldEqual, 1)
							db.Table("services").Count(&count)
							So(count, ShouldEqual, 1)
						})
					})
					Reset(func() {
						db.DropTableIfExists(&service.Service{})
					})
				})

				//Convey("Creating Swarm Table", func() {
				//	var count int
				//	db.AutoMigrate(&swarm.Swarm{})
				//	db.Model(&swarm.Swarm{}).AddForeignKey("service_group_id", "service_groups(id)", "CASCADE", "RESTRICT")
				//	Convey("Then associating one swarm with the service group", func() {
				//		db.Exec(fmt.Sprintf("INSERT INTO swarms (id, service_group_id, label) VALUES (4, %d, 'TestLabel')", s.ID))
				//		db.Table("swarms").Count(&count)
				//		So(count, ShouldEqual, 1)
				//		Convey("Then Deleting the service group should also delete the swarm label associated", func() {
				//			err = sgr.Delete(s.ID)
				//			So(err, ShouldBeNil)
				//			ac, err := sgr.GetAll()
				//			So(err, ShouldBeNil)
				//			So(len(ac), ShouldEqual, 0)
				//			db.Table("swarms").Count(&count)
				//			So(count, ShouldEqual, 0)
				//		})
				//
				//	})
				//	Reset(func() {
				//		db.DropTableIfExists(&swarm.Swarm{})
				//	})
				//})
			})
		})
		Reset(func() {
			db.DropTableIfExists(&service_group.ServiceGroup{})
		})
	})
	DropDB(db, c)
	db.Close()
}
