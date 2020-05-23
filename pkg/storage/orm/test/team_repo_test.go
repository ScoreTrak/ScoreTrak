package test

import (
	"ScoreTrak/pkg/storage/orm"
	"ScoreTrak/pkg/team"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSpec(t *testing.T) {

	c := setupConfig()
	c.DB.Cockroach.Database = "scoretrak_test_team"
	db := setupDB(c)
	l := setupLogger(c)

	Convey("Creating Tables", t, func() {
		createTables(db)
		tr := orm.NewTeamRepo(db, l)
		Reset(func() {
			cleanDB(db)
		})
		Convey("When the database is empty", func() {
			tr = orm.NewTeamRepo(db, l)
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
				})
			})

		})
	})

	dropDB(db, c)
	db.Close()
}
