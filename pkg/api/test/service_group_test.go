package client

import (
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/api/client"
	"github.com/L1ghtman2k/ScoreTrak/pkg/config"
	"github.com/L1ghtman2k/ScoreTrak/pkg/master/server"
	"github.com/L1ghtman2k/ScoreTrak/pkg/service_group"
	"github.com/L1ghtman2k/ScoreTrak/pkg/storage/orm"
	. "github.com/L1ghtman2k/ScoreTrak/test"
	. "github.com/smartystreets/goconvey/convey"
	"net"
	"net/http"
	"net/url"
	"os"
	"testing"
)

func TestServiceGroupSpec(t *testing.T) {
	var c config.StaticConfig
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
	routes = append(routes, server.ServiceGroupRoutes(l, serviceGroupSvc, nil)...)
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
		Convey("Retrieving a Service Group by ID", func() {
			retServiceGroup, err := cli.GetByID(1)
			So(err, ShouldBeNil)
			So(retServiceGroup.ID, ShouldEqual, 1)
			So(*(retServiceGroup.Enabled), ShouldBeTrue)
		})
		Convey("Retrieving a Service Group by wrong ID", func() {
			retServiceGroup, err := cli.GetByID(5)
			So(err, ShouldNotBeNil)
			So(retServiceGroup, ShouldBeNil)
			seer, ok := err.(*client.InvalidResponse)
			So(ok, ShouldBeTrue)
			So(seer.ResponseCode, ShouldHaveSameTypeAs, http.StatusNotFound)
		})

		Convey("Updating a Service Group by ID", func() {
			fls := false
			t := service_group.ServiceGroup{ID: 1, Enabled: &fls}
			err := cli.Update(&t)
			So(err, ShouldBeNil)
			Convey("Retrieving a Service Group by ID", func() {
				retServiceGroup, err := cli.GetByID(1)
				So(err, ShouldBeNil)
				So(retServiceGroup.Name, ShouldEqual, "ServiceGroup1")
				So(*(retServiceGroup.Enabled), ShouldBeFalse)
			})
		})

		Convey("Getting all Service Group", func() {
			serviceGroups, err := cli.GetAll()
			So(err, ShouldBeNil)
			So(len(serviceGroups), ShouldEqual, 4)
			var IDs []uint64
			for _, tm := range serviceGroups {
				IDs = append(IDs, tm.ID)
			}
			So(IDs, ShouldContain, uint64(3))
		})

		Convey("Deleting a Service Group that doesnt have child service by ID", func() {
			err := cli.Delete(3)
			So(err, ShouldBeNil)
			Convey("Getting all serviceGroups", func() {
				serviceGroups, err := cli.GetAll()
				So(err, ShouldBeNil)
				So(len(serviceGroups), ShouldEqual, 3)
			})
		})

		Convey("Deleting a Service Group that does have child service by ID", func() {
			err := cli.Delete(1)
			So(err, ShouldNotBeNil)
			seer, ok := err.(*client.InvalidResponse)
			So(ok, ShouldBeTrue)
			So(seer.ResponseCode, ShouldHaveSameTypeAs, http.StatusConflict)
			Convey("Getting all serviceGroups", func() {
				serviceGroups, err := cli.GetAll()
				So(err, ShouldBeNil)
				So(len(serviceGroups), ShouldEqual, 4)
			})
		})

		Convey("Deleting a non existent Service Group", func() {
			err := cli.Delete(6)
			So(err, ShouldBeNil)
		})

		Convey("Storing a new Service Group", func() {
			fls := false
			t := service_group.ServiceGroup{Enabled: &fls, Name: "ServiceGroup5"}
			err := cli.Store(&t)
			So(err, ShouldBeNil)
			Convey("Getting all Service Group", func() {
				serviceGroups, err := cli.GetAll()
				So(err, ShouldBeNil)
				So(len(serviceGroups), ShouldEqual, 5)
			})
		})

		Convey("Storing a new Service Group with the same name", func() {
			fls := false
			t := service_group.ServiceGroup{Enabled: &fls, Name: "ServiceGroup1"}
			err := cli.Store(&t)
			So(err, ShouldNotBeNil)
			seer, ok := err.(*client.InvalidResponse)
			So(ok, ShouldBeTrue)
			So(seer.ResponseCode, ShouldHaveSameTypeAs, http.StatusPreconditionFailed)
			Convey("Getting all Service Group", func() {
				serviceGroups, err := cli.GetAll()
				So(err, ShouldBeNil)
				So(len(serviceGroups), ShouldEqual, 4)
			})
		})

		Reset(func() {
			CleanAllTables(db)
		})
	})
	DropDB(db, c)
	db.Close()
}
