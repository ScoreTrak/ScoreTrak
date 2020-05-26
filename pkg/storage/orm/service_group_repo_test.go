package orm

import (
	"ScoreTrak/pkg/service"
	"ScoreTrak/pkg/service_group"
	"ScoreTrak/pkg/swarm"
	. "ScoreTrak/test"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestServiceGroupSpec(t *testing.T) {
	c := NewConfigClone(SetupConfig("dev-config.yml"))
	c.DB.Cockroach.Database = "scoretrak_test_service_group"
	c.Logger.FileName = "service_group_repo.log"
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
				t := service_group.ServiceGroup{Name: "TestServiceGroup"}
				err = sgr.Store(&t)
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
					err = sgr.Delete(t.ID + 1)
					So(err, ShouldBeNil)
					Convey("Should output one entry", func() {
						ac, err := sgr.GetAll()
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
					})
				})

				Convey("Then Deleting the added entry", func() {
					err = sgr.Delete(t.ID)
					So(err, ShouldBeNil)
					Convey("Should output no entries", func() {
						ac, err := sgr.GetAll()
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 0)
					})
				})

				Convey("Then Retrieving entry by ID", func() {
					sg, err := sgr.GetByID(t.ID)
					So(err, ShouldBeNil)
					Convey("Should output the inserted entry", func() {
						So(sg.Name, ShouldEqual, "TestServiceGroup")
						So(*(sg.Enabled), ShouldBeFalse)
					})
				})

				Convey("Then Updating Enabled to true", func() {
					tru := true
					new_sgr := &service_group.ServiceGroup{Enabled: &tru}
					Convey("For the wrong entry should not update anything", func() {
						new_sgr.ID = t.ID + 1
						err = sgr.Update(new_sgr)
						So(err, ShouldBeNil)
						ac, err := sgr.GetAll()
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
						So(*(ac[0].Enabled), ShouldBeFalse)

					})
					Convey("For the correct entry should update", func() {
						err = sgr.Update(new_sgr)
						So(err, ShouldBeNil)
						ac, err := sgr.GetAll()
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
						So(*(ac[0].Enabled), ShouldBeTrue)
					})
				})

				Convey("Creating Swarm and Service Tables", func() {
					db.AutoMigrate(&service.Service{})
					db.AutoMigrate(&swarm.Swarm{})
					db.Model(&service.Service{}).AddForeignKey("service_group_id", "service_groups(id)", "RESTRICT", "RESTRICT")
					db.Model(&swarm.Swarm{}).AddForeignKey("service_group_id", "service_groups(id)", "CASCADE", "RESTRICT")

					Convey("Then associating one service with the service group", func() {
						db.Exec("INSERT INTO hosts (id, address, team_id) VALUES (4, '192.168.1.1', 'TestTeam')")
					})

					Convey("Then associating one swarm with the service group", func() {
						db.Exec("INSERT INTO hosts (id, address, team_id) VALUES (4, '192.168.1.1', 'TestTeam')")
					})

					Reset(func() {
						db.DropTableIfExists(&service.Service{})
						db.DropTableIfExists(&swarm.Swarm{})
					})

				})
			})
		})
		Reset(func() {
			db.DropTableIfExists(&service_group.ServiceGroup{})
		})
	})
	DropDB(db, c)
	db.Close()
}
