package handler

import (
	"ScoreTrak/pkg/config"
	"ScoreTrak/pkg/storage/orm"
	"ScoreTrak/pkg/team"
	. "ScoreTrak/test"
	"encoding/json"
	"github.com/gorilla/mux"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestCheckSpec(t *testing.T) {
	var c *config.StaticConfig
	autoTest := os.Getenv("AUTO_TEST")
	if autoTest == "TRUE" {
		c = NewConfigClone(SetupConfig("../../../configs/test-config.yml"))
	} else {
		c = NewConfigClone(SetupConfig("dev-config.yml"))
	}
	c.DB.Cockroach.Database = "scoretrak_test_api_team"
	c.Logger.FileName = "team_test_repo.log"
	db := SetupDB(c)
	l := SetupLogger(c)
	t.Parallel() //t.Parallel should be placed after SetupDB because gorm has race conditions on Hook register
	Convey("Initializing check repo and controller", t, func() {
		DataPreload(db)
		cr := orm.NewTeamRepo(db, l)
		ctrl := NewTeamController(l, cr)
		Convey("Retrieving a team by ID", func() {
			//https://stackoverflow.com/questions/34435185/unit-testing-for-functions-that-use-gorilla-mux-url-parameters
			r, _ := http.NewRequest("GET", "/team/TeamOne", nil)
			w := httptest.NewRecorder()
			vars := map[string]string{
				"id": "TeamOne",
			}
			r = mux.SetURLVars(r, vars)
			ctrl.GetByID(w, r)
			var t team.Team
			err := json.Unmarshal([]byte(w.Body.String()), &t)
			Convey("Should return the team, and correct parrameters for that team", func() {
				So(err, ShouldBeNil)
				So(w.Code, ShouldEqual, http.StatusOK)
				So(t.ID, ShouldEqual, "TeamOne")
				So(*(t.Enabled), ShouldBeTrue)
			})
		})
		Reset(func() {
			CleanAllTables(db)
		})

	})
	DropDB(db, c)
	db.Close()
}
