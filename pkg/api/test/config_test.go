package client

import (
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/api/client"
	"github.com/L1ghtman2k/ScoreTrak/pkg/config"
	"github.com/L1ghtman2k/ScoreTrak/pkg/master/server"
	"github.com/L1ghtman2k/ScoreTrak/pkg/storage/orm"
	. "github.com/L1ghtman2k/ScoreTrak/test"
	. "github.com/smartystreets/goconvey/convey"
	"net"
	"net/http"
	"net/url"
	"os"
	"testing"
)

func TestConfigSpec(t *testing.T) {
	var c config.StaticConfig
	autoTest := os.Getenv("AUTO_TEST")
	if autoTest == "TRUE" {
		c = NewConfigClone(SetupConfig("../../../configs/test-config.yml"))
	} else {
		c = NewConfigClone(SetupConfig("dev-config.yml"))
	}
	c.DB.Cockroach.Database = "scoretrak_test_api_config"
	c.Logger.FileName = "config_test.log"
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
	cr := orm.NewConfigRepo(db, l)
	configSvc := config.NewConfigServ(cr)
	staticConfigSvc := config.NewStaticConfigServ()
	routes = append(routes, server.ConfigRoutes(l, configSvc)...)
	routes = append(routes, server.StaticConfigRoutes(l, staticConfigSvc)...)
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
	Convey("Initializing config repo and controller", t, func() {
		DataPreload(db)
		s := client.NewScoretrakClient(&url.URL{Host: fmt.Sprintf("localhost:%d", port), Scheme: "http"}, "", http.DefaultClient)
		cli := client.NewConfigClient(s)
		cliStatic := client.NewStaticConfigClient(s)
		Convey("Retrieving a config", func() {
			retConfig, err := cli.Get()
			So(err, ShouldBeNil)
			So(retConfig.ID, ShouldEqual, 1)
			So(retConfig.RoundDuration, ShouldEqual, uint64(60))
			So(*(retConfig.Enabled), ShouldBeTrue)
		})

		Convey("Retrieving a static config", func() {
			retConfig, err := cliStatic.Get()
			So(err, ShouldBeNil)
			So(retConfig.DB.Use, ShouldEqual, "cockroach")
		})

		Convey("Update the config", func() {
			fls := false
			t := config.DynamicConfig{RoundDuration: 50, Enabled: &fls}
			err := cli.Update(&t)
			ShouldNotBeNil(err)
			Convey("Retrieving a config", func() {
				retConfig, err := cli.Get()
				So(err, ShouldBeNil)
				So(retConfig.ID, ShouldEqual, 1)
				So(retConfig.RoundDuration, ShouldEqual, uint64(50))
				So(*(retConfig.Enabled), ShouldBeFalse)
			})
		})

		Reset(func() {
			CleanAllTables(db)
		})
	})
	DropDB(db, c)
	db.Close()
}
