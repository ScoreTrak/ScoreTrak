package client

import (
	"ScoreTrak/pkg/api/client"
	"ScoreTrak/pkg/config"
	"ScoreTrak/pkg/host_group"
	"ScoreTrak/pkg/master/server"
	"ScoreTrak/pkg/storage/orm"
	. "ScoreTrak/test"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"net"
	"net/http"
	"net/url"
	"os"
	"testing"
)

func TestHostGroupSpec(t *testing.T) {
	var c *config.StaticConfig
	autoTest := os.Getenv("AUTO_TEST")
	if autoTest == "TRUE" {
		c = NewConfigClone(SetupConfig("../../../configs/test-config.yml"))
	} else {
		c = NewConfigClone(SetupConfig("dev-config.yml"))
	}
	c.DB.Cockroach.Database = "scoretrak_test_api_host_group"
	c.Logger.FileName = "host_group_test.log"
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
	cr := orm.NewHostGroupRepo(db, l)
	hostGroupSvc := host_group.NewHostGroupServ(cr)
	routes = append(routes, server.HostGroupRoutes(l, hostGroupSvc)...)
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
	Convey("Initializing hostGroup repo and controller", t, func() {
		DataPreload(db)
		s := client.NewScoretrakClient(&url.URL{Host: fmt.Sprintf("localhost:%d", port), Scheme: "http"}, "", http.DefaultClient)
		cli := client.NewHostGroupClient(s)
		Convey("Retrieving a hostGroup by ID", func() {
			retHostGroup, err := cli.GetByID(1)
			So(err, ShouldBeNil)
			So(retHostGroup.ID, ShouldEqual, 1)
		})
		Reset(func() {
			CleanAllTables(db)
		})
	})
	DropDB(db, c)
	db.Close()
}
