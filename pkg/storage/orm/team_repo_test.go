package orm

import (
	"github.com/L1ghtman2k/ScoreTrak/pkg/config"
	"github.com/L1ghtman2k/ScoreTrak/pkg/host"
	"github.com/L1ghtman2k/ScoreTrak/pkg/team"
	. "github.com/L1ghtman2k/ScoreTrak/test"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestTeamSpec(t *testing.T) {
	var c config.StaticConfig
	autoTest := os.Getenv("AUTO_TEST")
	if autoTest == "TRUE" {
		c = NewConfigClone(SetupConfig("../../../configs/test-config.yml"))
	} else {
		c = NewConfigClone(SetupConfig("dev-config.yml"))
	}
	c.DB.Cockroach.Database = "scoretrak_test_orm_team"
	c.Logger.FileName = "team_test.log"
	db := SetupDB(c.DB)
	l := SetupLogger(c.Logger)
	t.Parallel() //t.Parallel should be placed after SetupDB because gorm has race conditions on Hook register
	Convey("Creating Team Tables", t, func() {
		db.AutoMigrate(&team.Team{})
		tr := NewTeamRepo(db, l)
		Reset(func() {
			db.Migrator().DropTable(&team.Team{})
		})
		Convey("When the Teams table is empty", func() {
			Convey("There should be no entries", func() {
				ac, err := tr.GetAll()
				So(err, ShouldBeNil)
				So(len(ac), ShouldEqual, 0)
			})

			Convey("Adding an entry with empty name", func() {
				var err error
				tru := true
				t := team.Team{Name: "", Enabled: &tru}
				err = tr.Store(&t)
				So(err, ShouldNotBeNil)

				Convey("Should output no entry", func() {
					ac, err := tr.GetAll()
					So(err, ShouldBeNil)
					So(len(ac), ShouldEqual, 0)
				})
			})

			Convey("Adding a valid entry", func() {
				var err error
				t := team.Team{Name: "TestTeam"}
				err = tr.Store(&t)
				So(err, ShouldBeNil)
				Convey("Should output one entry", func() {
					ac, err := tr.GetAll()
					So(err, ShouldBeNil)
					So(len(ac), ShouldEqual, 1)
					So(ac[0].Name, ShouldEqual, "TestTeam")
				})

				Convey("Then Deleting a wrong entry", func() {
					err = tr.DeleteByName("TestTeamWRONG")
					So(err, ShouldNotBeNil)
					Convey("Should output one entry", func() {
						ac, err := tr.GetAll()
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
					})
				})
				Convey("Then Deleting the added entry", func() {
					err = tr.DeleteByName("TestTeam")
					So(err, ShouldBeNil)
					Convey("Should output no entries", func() {
						ac, err := tr.GetAll()
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 0)
					})
				})

				Convey("Then Retrieving entry by Name", func() {
					tm, err := tr.GetByName("TestTeam")
					So(err, ShouldBeNil)
					Convey("Should output the inserted entry", func() {
						So(tm.Name, ShouldEqual, "TestTeam")
						So(*(tm.Enabled), ShouldBeTrue)
					})
				})

				Convey("Then Querying By wrong Name", func() {
					ss, err := tr.GetByName("WrongTeamName")
					So(err, ShouldNotBeNil)
					So(ss, ShouldBeNil)
				})

				Convey("Then Updating Enabled to true", func() {
					fls := false
					newTeam := &team.Team{Enabled: &fls}
					Convey("For the wrong entry should not update anything", func() {
						newTeam.ID = t.ID + 1
						err = tr.Update(newTeam)
						So(err, ShouldBeNil)
						ac, err := tr.GetAll()
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
						So(*(ac[0].Enabled), ShouldBeTrue)

					})
					Convey("For the correct entry should update", func() {
						newTeam.Name = "TestTeam"
						newTeam.ID = t.ID
						err = tr.Update(newTeam)
						So(err, ShouldBeNil)
						ac, err := tr.GetAll()
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
						So(*(ac[0].Enabled), ShouldBeFalse)

					})
				})

				Convey("Creating Hosts Table", func() {
					var count int64
					db.AutoMigrate(&host.Host{})
					Convey("Associating a single Host with a team", func() {
						db.Exec("INSERT INTO hosts (id, address, team_name) VALUES (4, '192.168.1.1', 'TestTeam')")
						db.Table("hosts").Count(&count)
						So(count, ShouldEqual, 1)
						Convey("DeleteByName a team without deleting a host", func() {
							err = tr.DeleteByName("TestTeam")
							So(err, ShouldNotBeNil)
							ac, err := tr.GetAll()
							So(err, ShouldBeNil)
							So(len(ac), ShouldEqual, 1)
						})
						Convey("Deleting a host then deleting a team", func() {
							db.Exec("DELETE FROM hosts WHERE id=4")
							err = tr.DeleteByName("TestTeam")
							So(err, ShouldBeNil)
							ac, err := tr.GetAll()
							So(err, ShouldBeNil)
							So(len(ac), ShouldEqual, 0)
						})

						Convey("Updating a team enabled without deleting a host should not yield error", func() {
							tru := true
							t.Enabled = &tru
							err = tr.Update(&t)
							So(err, ShouldBeNil)
						})

						Reset(func() {
							db.Migrator().DropTable(&host.Host{})
						})
					})
				})
			})
		})
	})
	DropDB(db, c)

}
