package client

import (
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/api/client"
	"github.com/L1ghtman2k/ScoreTrak/pkg/config"
	"github.com/L1ghtman2k/ScoreTrak/pkg/master/server"
	"github.com/L1ghtman2k/ScoreTrak/pkg/round"
	"github.com/L1ghtman2k/ScoreTrak/pkg/storage/orm"
	. "github.com/L1ghtman2k/ScoreTrak/test"
	"net"
	"net/http"
	"net/url"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRoundSpec(t *testing.T) {
	var c config.StaticConfig
	autoTest := os.Getenv("AUTO_TEST")
	if autoTest == "TRUE" {
		c = NewConfigClone(SetupConfig("../../../configs/test-config.yml"))
	} else {
		c = NewConfigClone(SetupConfig("dev-config.yml"))
	}
	c.DB.Cockroach.Database = "scoretrak_test_api_round"
	c.Logger.FileName = "round_test.log"
	db := SetupDB(c.DB)
	l := SetupLogger(c.Logger)
	rtr := server.NewRouter()
	routes := server.Routes{
		server.Route{
			Name:        "Index",
			Method:      "GET",
			Pattern:     "/",
			HandlerFunc: server.Index,
		},
	}
	cr := orm.NewRoundRepo(db, l)
	roundSvc := round.NewRoundServ(cr)
	routes = append(routes, server.RoundRoutes(l, roundSvc)...)
	for _, route := range routes {
		var hdler http.Handler
		hdler = route.HandlerFunc
		hdler = server.Logger(hdler, route.Name) //Default Logger
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
	Convey("Initializing round repo and controller", t, func() {
		DataPreload(db)
		s := client.NewScoretrakClient(&url.URL{Host: fmt.Sprintf("localhost:%d", port), Scheme: "http"}, "", http.DefaultClient)
		cli := client.NewRoundClient(s)
		Convey("Retrieving a round by ID", func() {
			retRound, err := cli.GetLastNonElapsingRound()
			So(err, ShouldBeNil)
			So(retRound.ID, ShouldEqual, 3)
			So(retRound.Finish, ShouldNotBeNil)
		})
		Reset(func() {
			CleanAllTables(db)
		})
	})
	DropDB(db, c)
	db.Close()
}
