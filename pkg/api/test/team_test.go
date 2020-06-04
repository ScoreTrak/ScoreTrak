package client

import (
	"ScoreTrak/pkg/api/client"
	"ScoreTrak/pkg/config"
	"ScoreTrak/pkg/master/server"
	"ScoreTrak/pkg/storage/orm"
	"ScoreTrak/pkg/team"
	. "ScoreTrak/test"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"net"
	"net/http"
	"net/url"
	"os"
	"testing"
)

func TestTeamSpec(t *testing.T) {
	var c *config.StaticConfig
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
		Convey("Retrieving a team by ID", func() {
			retTeam, err := cli.GetByID("TeamOne")
			So(err, ShouldBeNil)
			So(retTeam.ID, ShouldEqual, "TeamOne")
			So(*(retTeam.Enabled), ShouldBeTrue)
		})
		Convey("Retrieving a team by wrong ID", func() {
			retTeam, err := cli.GetByID("TeamWrong")
			So(err, ShouldNotBeNil)
			So(retTeam, ShouldBeNil)
			seer, ok := err.(*client.InvalidResponse)
			So(ok, ShouldBeTrue)
			So(seer.ResponseCode, ShouldHaveSameTypeAs, http.StatusNotFound)
		})

		Convey("Updating a team by ID", func() {
			fls := false
			t := team.Team{ID: "TeamOne", Enabled: &fls}
			err := cli.Update(&t)
			So(err, ShouldBeNil)
			Convey("Retrieving a team by ID", func() {
				retTeam, err := cli.GetByID("TeamOne")
				So(err, ShouldBeNil)
				So(retTeam.ID, ShouldEqual, "TeamOne")
				So(*(retTeam.Enabled), ShouldBeFalse)
			})
		})

		Convey("Getting all teams by ID", func() {
			teams, err := cli.GetAll()
			So(err, ShouldBeNil)
			So(len(teams), ShouldEqual, 4)
			var IDs []string
			for _, team := range teams {
				IDs = append(IDs, team.ID)
			}
			So(IDs, ShouldContain, "TeamTwo")
		})

		Convey("Deleting a team that doesnt have child hosts by ID", func() {
			err := cli.Delete("TeamOne")
			So(err, ShouldBeNil)
			Convey("Getting all teams", func() {
				teams, err := cli.GetAll()
				So(err, ShouldBeNil)
				So(len(teams), ShouldEqual, 3)
			})
		})

		Convey("Deleting a team that does have child hosts by ID", func() {
			err := cli.Delete("TeamTwo")
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
			err := cli.Delete("TeamWrong")
			So(err, ShouldBeNil)
		})

		Convey("Storing a new team", func() {
			fls := false
			t := team.Team{ID: "TeamFive", Enabled: &fls}
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
