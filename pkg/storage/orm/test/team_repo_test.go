package test

import (
	"ScoreTrak/pkg/host"
	"ScoreTrak/pkg/storage/orm"
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
		tr := orm.NewTeamRepo(db, l)
		Reset(func() {
			db.DropTableIfExists(&team.Team{})
		})
		Convey("When the Teams table is empty", func() {
			Convey("There should be no entries", func() {
				ac, err := tr.GetAll()
				So(err, ShouldBeNil)
				So(len(ac), ShouldEqual, 0)
			})
			Convey("Adding an entry", func() {
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

						Convey("Deleting a host then deleting a table", func() {
							db.Exec("DELETE FROM hosts WHERE id=4")
							err = tr.Delete("TestTeam")
							So(err, ShouldBeNil)
							ac, err := tr.GetAll()
							So(err, ShouldBeNil)
							So(len(ac), ShouldEqual, 0)
						})

					})
					Reset(func() {
						db.DropTableIfExists(&host.Host{})
					})
				})
			})
		})
	})

	DropDB(db, c)
	db.Close()
}
