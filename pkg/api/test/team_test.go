package client

import (
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/api/client"
	"github.com/L1ghtman2k/ScoreTrak/pkg/config"
	"github.com/L1ghtman2k/ScoreTrak/pkg/master/server"
	"github.com/L1ghtman2k/ScoreTrak/pkg/storage/orm"
	"github.com/L1ghtman2k/ScoreTrak/pkg/team"
	. "github.com/L1ghtman2k/ScoreTrak/test"
	. "github.com/smartystreets/goconvey/convey"
	"net"
	"net/http"
	"net/url"
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
	c.DB.Cockroach.Database = "scoretrak_test_api_team"
	c.Logger.FileName = "team_test.log"
	db := SetupDB(c)
	l := SetupLogger(c)
	rtr := server.NewRouter()
	routes := server.Routes{
		server.Route{
			Name:        "Index",
			Method:      "GET",
			Pattern:     "/",
			HandlerFunc: server.Index,
		},
	}
	cr := orm.NewTeamRepo(db, l)
	teamSvc := team.NewTeamServ(cr)
	routes = append(routes, server.TeamRoutes(l, teamSvc)...)
	for _, route := range routes {
		var hdler http.Handler
		hdler = route.HandlerFunc
		rtr.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(hdler)
	}
	rtr.Use(server.JsonHeader)
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	port := listener.Addr().(*net.TCPAddr).Port
	go http.Serve(listener, rtr)
	t.Parallel() //t.Parallel should be placed after SetupDB because gorm has race conditions on Hook register
	Convey("Initializing team repo and controller", t, func() {
		DataPreload(db)
		s := client.NewScoretrakClient(&url.URL{Host: fmt.Sprintf("localhost:%d", port), Scheme: "http"}, "", http.DefaultClient)
		cli := client.NewTeamClient(s)
		Convey("Retrieving a team by Name", func() {
			retTeam, err := cli.GetByName("TeamOne")
			So(err, ShouldBeNil)
			So(retTeam.Name, ShouldEqual, "TeamOne")
			So(*(retTeam.Enabled), ShouldBeTrue)
		})
		Convey("Retrieving a team by wrong ID", func() {
			retTeam, err := cli.GetByName("TeamWrong")
			So(err, ShouldNotBeNil)
			So(retTeam, ShouldBeNil)
			seer, ok := err.(*client.InvalidResponse)
			So(ok, ShouldBeTrue)
			So(seer.ResponseCode, ShouldHaveSameTypeAs, http.StatusNotFound)
		})

		Convey("Updating a team by Name", func() {
			fls := false
			t := team.Team{Name: "TeamOne", Enabled: &fls}
			err := cli.UpdateByName(&t)
			So(err, ShouldBeNil)
			Convey("Retrieving a team by Name", func() {
				retTeam, err := cli.GetByName("TeamOne")
				So(err, ShouldBeNil)
				So(retTeam.Name, ShouldEqual, "TeamOne")
				So(*(retTeam.Enabled), ShouldBeFalse)
			})
		})

		Convey("Getting all teams", func() {
			teams, err := cli.GetAll()
			So(err, ShouldBeNil)
			So(len(teams), ShouldEqual, 4)
			var IDs []string
			for _, tm := range teams {
				IDs = append(IDs, tm.Name)
			}
			So(IDs, ShouldContain, "TeamTwo")
		})

		Convey("Deleting a team that doesnt have child hosts by Name", func() {
			err := cli.DeleteByName("TeamOne")
			So(err, ShouldBeNil)
			Convey("Getting all teams", func() {
				teams, err := cli.GetAll()
				So(err, ShouldBeNil)
				So(len(teams), ShouldEqual, 3)
			})
		})

		Convey("Deleting a team that does have child hosts by Name", func() {
			err := cli.DeleteByName("TeamTwo")
			So(err, ShouldNotBeNil)
			seer, ok := err.(*client.InvalidResponse)
			So(ok, ShouldBeTrue)
			So(seer.ResponseCode, ShouldHaveSameTypeAs, http.StatusConflict)
			Convey("Getting all teams", func() {
				teams, err := cli.GetAll()
				So(err, ShouldBeNil)
				So(len(teams), ShouldEqual, 4)
			})
		})

		Convey("Deleting a non existent team", func() {
			err := cli.DeleteByName("TeamWrong")
			So(err, ShouldBeNil)
		})

		Convey("Storing a new team", func() {
			fls := false
			t := team.Team{Name: "TeamFive", Enabled: &fls}
			err := cli.Store(&t)
			So(err, ShouldBeNil)
			Convey("Getting all teams", func() {
				teams, err := cli.GetAll()
				So(err, ShouldBeNil)
				So(len(teams), ShouldEqual, 5)
			})
		})

		Reset(func() {
			CleanAllTables(db)
		})
	})
	DropDB(db, c)
	db.Close()
}
