package client

import (
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/cmd/master/server/gorilla"
	"github.com/ScoreTrak/ScoreTrak/pkg/api/client"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	. "github.com/ScoreTrak/ScoreTrak/pkg/config/util"
	. "github.com/ScoreTrak/ScoreTrak/pkg/logger/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/orm"
	. "github.com/ScoreTrak/ScoreTrak/pkg/storage/orm/util"

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
	rtr := gorilla.NewRouter()
	routes := gorilla.Routes{
		gorilla.Route{
			Name:        "Index",
			Method:      "GET",
			Pattern:     "/",
			HandlerFunc: gorilla.Index,
		},
	}
	cr := orm.NewConfigRepo(db, l)
	configSvc := config.NewConfigServ(cr)
	staticConfigSvc := config.NewStaticConfigServ()
	routes = append(routes, gorilla.ConfigRoutes(l, configSvc)...)
	routes = append(routes, gorilla.StaticConfigRoutes(l, staticConfigSvc)...)
	for _, route := range routes {
		var hdler http.Handler
		hdler = route.HandlerFunc
		hdler = gorilla.Logger(hdler, route.Name) //Default Logger
		rtr.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(hdler)
	}
	rtr.Use(gorilla.JsonHeader)
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	port := listener.Addr().(*net.TCPAddr).Port
	go http.Serve(listener, rtr)
	t.Parallel() //t.Parallel should be placed after SetupDB because gorm has race conditions on Hook register
	Convey("Initializing config repo and controller", t, func() {
		CreateAllTables(db)
		DataPreload(db)
		s := client.NewScoretrakClient(&url.URL{Host: fmt.Sprintf("localhost:%d", port), Scheme: "http"}, "", http.DefaultClient)
		cli := client.NewConfigClient(s)
		cliStatic := client.NewStaticConfigClient(s)
		Convey("Retrieving a config", func() {
			retConfig, err := cli.Get()
			So(err, ShouldBeNil)
			So(retConfig.ID, ShouldEqual, 1)
			So(retConfig.RoundDuration, ShouldEqual, uint(60))
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
				So(retConfig.RoundDuration, ShouldEqual, uint(50))
				So(*(retConfig.Enabled), ShouldBeFalse)
			})
		})

		Reset(func() {
			CleanAllTables(db)
		})
	})
	DropDB(db, c)

}
