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
			r, _ := http.NewRequest("GET", "/team/TeamOne", nil)
			w := NewJsonRecorder()
			vars := map[string]string{
				"id": "TeamOne",
			}
			r = mux.SetURLVars(r, vars)
			ctrl.GetByID(w, r)
			var t team.Team
			err := json.Unmarshal([]byte(w.Body.String()), &t)
			So(err, ShouldBeNil)
			So(w.Code, ShouldEqual, http.StatusOK)
			So(t.ID, ShouldEqual, "TeamOne")
			So(*(t.Enabled), ShouldBeTrue)
		})

		Convey("Retrieving a team by invalid ID", func() {
			r, _ := http.NewRequest("GET", "/team/WrongTeam", nil)
			w := NewJsonRecorder()
			vars := map[string]string{
				"id": "WrongTeam",
			}
			r = mux.SetURLVars(r, vars)
			ctrl.GetByID(w, r)
			So(w.Code, ShouldEqual, http.StatusNotFound)
		})

		Convey("Retrieving all teams", func() {
			r, _ := http.NewRequest("GET", "/team", nil)
			w := NewJsonRecorder()
			ctrl.GetAll(w, r)
			So(w.Code, ShouldEqual, http.StatusOK)
			var t []team.Team
			err := json.Unmarshal([]byte(w.Body.String()), &t)
			So(err, ShouldBeNil)
			So(w.Code, ShouldEqual, http.StatusOK)
			So(len(t), ShouldEqual, 4)
		})

		Convey("Deleting a team by ID (without any dependant hosts)", func() {
			r, _ := http.NewRequest("DELETE", "/team/TeamOne", nil)
			w := NewJsonRecorder()
			vars := map[string]string{
				"id": "TeamOne",
			}
			r = mux.SetURLVars(r, vars)
			ctrl.Delete(w, r)
			So(w.Code, ShouldEqual, http.StatusOK)
			Convey("Retrieving all teams", func() {
				r, _ := http.NewRequest("GET", "/team", nil)
				w := NewJsonRecorder()
				ctrl.GetAll(w, r)
				So(w.Code, ShouldEqual, http.StatusOK)
				var t []team.Team
				err := json.Unmarshal([]byte(w.Body.String()), &t)
				So(err, ShouldBeNil)
				So(w.Code, ShouldEqual, http.StatusOK)
				So(len(t), ShouldEqual, 3)
			})
		})

		Convey("Deleting a team by ID without deleting all hosts", func() {
			r, _ := http.NewRequest("DELETE", "/team/TeamTwo", nil)
			w := NewJsonRecorder()
			vars := map[string]string{
				"id": "TeamTwo",
			}
			r = mux.SetURLVars(r, vars)
			ctrl.Delete(w, r)
			So(w.Code, ShouldEqual, http.StatusConflict)
			Convey("Retrieving all teams", func() {
				r, _ := http.NewRequest("GET", "/team", nil)
				w := NewJsonRecorder()
				ctrl.GetAll(w, r)
				So(w.Code, ShouldEqual, http.StatusOK)
				var t []team.Team
				err := json.Unmarshal([]byte(w.Body.String()), &t)
				So(err, ShouldBeNil)
				So(w.Code, ShouldEqual, http.StatusOK)
				So(len(t), ShouldEqual, 4)
			})
		})

		Convey("Deleting a team by invalid ID", func() {
			r, _ := http.NewRequest("DELETE", "/team/WrongTeam", nil)
			w := NewJsonRecorder()
			vars := map[string]string{
				"id": "WrongTeam",
			}
			r = mux.SetURLVars(r, vars)
			ctrl.Delete(w, r)
			So(w.Code, ShouldEqual, http.StatusNotModified)

			Convey("Retrieving all teams", func() {
				r, _ := http.NewRequest("GET", "/team", nil)
				w := NewJsonRecorder()
				ctrl.GetAll(w, r)
				So(w.Code, ShouldEqual, http.StatusOK)
				var t []team.Team
				err := json.Unmarshal([]byte(w.Body.String()), &t)
				So(err, ShouldBeNil)
				So(w.Code, ShouldEqual, http.StatusOK)
				So(len(t), ShouldEqual, 4)
			})
		})
		Reset(func() {
			CleanAllTables(db)
		})
	})
	DropDB(db, c)
	db.Close()
}
