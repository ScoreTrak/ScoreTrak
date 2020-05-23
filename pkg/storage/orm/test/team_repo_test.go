package test

import (
	"ScoreTrak/pkg/storage/orm"
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
			Convey("There should be no entries", func() {
				tr = orm.NewTeamRepo(db, l)
				ac, err := tr.GetAll()
				So(err, ShouldBeNil)
				So(len(ac), ShouldEqual, 0)
			})
		})
	})

	dropDB(db, c)
	db.Close()
}
