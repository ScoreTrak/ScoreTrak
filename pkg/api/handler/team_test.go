package handler

import (
	"ScoreTrak/pkg/storage/orm"
	. "ScoreTrak/test"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCheckSpec(t *testing.T) {
	c := NewConfigClone(SetupConfig("dev-config.yml"))
	c.DB.Cockroach.Database = "scoretrak_test_api_team"
	c.Logger.FileName = "team_test_repo.log"
	db := SetupDB(c)
	l := SetupLogger(c)
	DataPreload(db)
	t.Parallel() //t.Parallel should be placed after SetupDB because gorm has race conditions on Hook register
	Convey("Initializing check repo and controller", t, func() {
		cr := orm.NewTeamRepo(db, l)
		_ = NewTeamController(l, cr)
		Convey("", func() {

		})
	})
	DropDB(db, c)
	db.Close()
}
