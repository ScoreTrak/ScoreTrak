package client

import (
	"ScoreTrak/pkg/api/client"
	"ScoreTrak/pkg/config"
	"ScoreTrak/pkg/master/server"
	"ScoreTrak/pkg/service_group"
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

func TestServiceGroupSpec(t *testing.T) {
	var c *config.StaticConfig
	autoTest := os.Getenv("AUTO_TEST")
	if autoTest == "TRUE" {
		c = NewConfigClone(SetupConfig("../../../configs/test-config.yml"))
	} else {
		c = NewConfigClone(SetupConfig("dev-config.yml"))
	}
	c.DB.Cockroach.Database = "scoretrak_test_api_service_group"
	c.Logger.FileName = "service_group_test.log"
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
	cr := orm.NewServiceGroupRepo(db, l)
	serviceGroupSvc := service_group.NewServiceGroupServ(cr)
	routes = append(routes, server.ServiceGroupRoutes(l, serviceGroupSvc)...)
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
	Convey("Initializing serviceGroup repo and controller", t, func() {
		DataPreload(db)
		s := client.NewScoretrakClient(&url.URL{Host: fmt.Sprintf("localhost:%d", port), Scheme: "http"}, "", http.DefaultClient)
		cli := client.NewServiceGroupClient(s)
		Convey("Retrieving a serviceGroup by ID", func() {
			retServiceGroup, err := cli.GetByID(1)
			So(err, ShouldBeNil)
			So(retServiceGroup.ID, ShouldEqual, 1)
		})
		Reset(func() {
			CleanAllTables(db)
		})
	})
	DropDB(db, c)
	db.Close()
}
