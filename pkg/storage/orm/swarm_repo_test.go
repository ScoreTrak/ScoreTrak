package orm

import (
	"github.com/L1ghtman2k/ScoreTrak/pkg/config"
	"github.com/L1ghtman2k/ScoreTrak/pkg/service_group"
	"github.com/L1ghtman2k/ScoreTrak/pkg/swarm"
	. "github.com/L1ghtman2k/ScoreTrak/test"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestSwarmSpec(t *testing.T) {
	var c *config.StaticConfig
	autoTest := os.Getenv("AUTO_TEST")
	if autoTest == "TRUE" {
		c = NewConfigClone(SetupConfig("../../../configs/test-config.yml"))
	} else {
		c = NewConfigClone(SetupConfig("dev-config.yml"))
	}
	c.DB.Cockroach.Database = "scoretrak_test_orm_swarm"
	c.Logger.FileName = "swarm_test.log"
	db := SetupDB(c)
	l := SetupLogger(c)
	t.Parallel() //t.Parallel should be placed after SetupDB because gorm has race conditions on Hook register
	Convey("Creating Swarm and Service Group tables along with their foreign keys", t, func() {
		db.AutoMigrate(&service_group.ServiceGroup{})
		db.AutoMigrate(&swarm.Swarm{})
		db.Model(&swarm.Swarm{}).AddForeignKey("service_group_id", "service_groups(id)", "CASCADE", "RESTRICT")
		sr := NewSwarmRepo(db, l)
		Convey("When all tables are empty", func() {
			Convey("Should output no entry", func() {
				gt, err := sr.GetAll()
				So(err, ShouldBeNil)
				So(len(gt), ShouldEqual, 0)
			})
			Convey("Creating a sample swarm should not be allowed", func() {
				s := swarm.Swarm{Label: "TestSwarmLabel"}
				err := sr.Store(&s)
				So(err, ShouldNotBeNil)
				ac, err := sr.GetAll()
				So(err, ShouldBeNil)
				So(len(ac), ShouldEqual, 0)
			})
			Convey("Load Sample Service Group", func() {
				var count int
				db.Exec("INSERT INTO service_groups (id, name) VALUES (7, 'TestServiceGroup')")
				db.Exec("INSERT INTO service_groups (id, name) VALUES (2, 'TestServiceGroup2')")
				db.Table("service_groups").Count(&count)
				So(count, ShouldEqual, 2)

				Convey("Creating a sample swarm and associating with service group id 2", func() {
					s := swarm.Swarm{Label: "TestSwarmLabel", ServiceGroupID: 7}
					err := sr.Store(&s)
					Convey("Should be Allowed", func() {
						So(err, ShouldBeNil)
						ac, err := sr.GetAll()
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
						Convey("Then Querying By ID", func() {
							ss, err := sr.GetByID(s.ID)
							So(err, ShouldBeNil)
							So(ss.Label, ShouldEqual, "TestSwarmLabel")
							So(ss.ServiceGroupID, ShouldEqual, 7)
						})

						Convey("Then Querying By wrong ID", func() {
							ss, err := sr.GetByID(s.ID + 1)
							So(err, ShouldNotBeNil)
							So(ss, ShouldBeNil)
						})

						Convey("Then updating to different service_group should not change anything", func() {
							s.ServiceGroupID = 2
							err = sr.Update(&s)
							So(err, ShouldBeNil)
							ac, err = sr.GetAll()
							So(err, ShouldBeNil)
							So(len(ac), ShouldEqual, 1)
							So(ac[0].ServiceGroupID, ShouldEqual, 7)
						})
						Convey("Then adding swarm with the same ServiceGroupID should not be allowed", func() {
							s2 := swarm.Swarm{Label: "TestServiceTwo", ServiceGroupID: 7}
							err = sr.Store(&s2)
							So(err, ShouldNotBeNil)
						})

						Convey("Then Deleting the swarm should be allowed", func() {
							err = sr.Delete(s.ID)
							So(err, ShouldBeNil)
							ac, err = sr.GetAll()
							So(err, ShouldBeNil)
							So(len(ac), ShouldEqual, 0)
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

						Convey("Then adding swarm with the same Label should be allowed", func() {
							s2 := swarm.Swarm{Label: "TestSwarmLabel", ServiceGroupID: 2}
							err = sr.Store(&s2)
							So(err, ShouldBeNil)
						})
					})
				})
				Convey("Creating a swarm with wrong service_group should not be allowed", func() {
					s := swarm.Swarm{Label: "TestSwarm", ServiceGroupID: 88}
					err := sr.Store(&s)
					So(err, ShouldNotBeNil)
					ac, err := sr.GetAll()
					So(err, ShouldBeNil)
					So(len(ac), ShouldEqual, 0)
				})
			})
		})
		Reset(func() {
			db.DropTableIfExists(&swarm.Swarm{})
			db.DropTableIfExists(&service_group.ServiceGroup{})
		})
	})
	DropDB(db, c)
	db.Close()
}
