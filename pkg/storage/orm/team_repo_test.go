package orm

import (
	"ScoreTrak/pkg/host"
	"ScoreTrak/pkg/team"
	. "ScoreTrak/test"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSpec(t *testing.T) {
	c := NewConfigClone(SetupConfig("dev-config.yml"))
	c.DB.Cockroach.Database = "scoretrak_test_team"
	c.Logger.FileName = "scoretrak_team.log"
	db := SetupDB(c)
	l := SetupLogger(c)
	Convey("Creating Team Tables", t, func() {
		db.AutoMigrate(&team.Team{})
		tr := NewTeamRepo(db, l)
		Reset(func() {
			db.DropTableIfExists(&team.Team{})
		})
		Convey("When the Teams table is empty", func() {
			Convey("There should be no entries", func() {
				ac, err := tr.GetAll()
				So(err, ShouldBeNil)
				So(len(ac), ShouldEqual, 0)
			})

			Convey("Adding a entry with empty name", func() {
				var err error
				tru := true
				t := team.Team{ID: "", Enabled: &tru}
				err = tr.Store(&t)
				So(err, ShouldNotBeNil)

				Convey("Should output no entry", func() {
					ac, err := tr.GetAll()
					So(err, ShouldBeNil)
					So(len(ac), ShouldEqual, 0)
				})
			})

			Convey("Adding an valid entry", func() {
				var err error
				t := team.Team{ID: "TestTeam"}
				err = tr.Store(&t)
				So(err, ShouldBeNil)
				Convey("Should output one entry", func() {
					ac, err := tr.GetAll()
					So(err, ShouldBeNil)
					So(len(ac), ShouldEqual, 1)
					So(ac[0].ID, ShouldEqual, "TestTeam")
				})

				Convey("Then Deleting a wrong entry", func() {
					err = tr.Delete("TestTeamWRONG")
					So(err, ShouldBeNil)
					Convey("Should output one entry", func() {
						ac, err := tr.GetAll()
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
					})
				})
				Convey("Then Deleting the added entry", func() {
					err = tr.Delete("TestTeam")
					So(err, ShouldBeNil)
					Convey("Should output no entries", func() {
						ac, err := tr.GetAll()
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 0)
					})
				})

				Convey("Then Retrieving entry by ID", func() {
					tm, err := tr.GetByID("TestTeam")
					So(err, ShouldBeNil)
					Convey("Should output the inserted entry", func() {
						So(tm.ID, ShouldEqual, "TestTeam")
						So(*(tm.Enabled), ShouldBeFalse)
					})
				})

				Convey("Then Updating Enabled to true", func() {
					tru := true
					new_team := &team.Team{Enabled: &tru}
					Convey("For the wrong Team should not update the entry", func() {
						new_team.ID = "WrongTestTeam"
						err = tr.Update(new_team)
						So(err, ShouldBeNil)
						ac, err := tr.GetAll()
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
						So(*(ac[0].Enabled), ShouldBeFalse)

					})
					Convey("For the correct Team should not yield error", func() {
						err = tr.Update(new_team)
						So(err, ShouldBeNil)
						ac, err := tr.GetAll()
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
						So(*(ac[0].Enabled), ShouldBeTrue)

					})
				})

				Convey("Creating Hosts Table", func() {
					db.AutoMigrate(&host.Host{})
					db.Model(&host.Host{}).AddForeignKey("team_id", "teams(id)", "RESTRICT", "RESTRICT")
					Convey("Associating a single Host with a team", func() {
						db.Exec("INSERT INTO hosts (id, address, team_id) VALUES (4, '192.168.1.1', 'TestTeam')")
						Convey("Delete a team without deleting a host", func() {
							err = tr.Delete("TestTeam")
							So(err, ShouldNotBeNil)
							ac, err := tr.GetAll()
							So(err, ShouldBeNil)
							So(len(ac), ShouldEqual, 1)
						})
						Convey("Deleting a host then deleting a team", func() {
							db.Exec("DELETE FROM hosts WHERE id=4")
							err = tr.Delete("TestTeam")
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
							db.DropTableIfExists(&host.Host{})
						})
					})
				})
			})
		})
	})
	DropDB(db, c)
	db.Close()
}
