package client

import (
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/api/client"
	"github.com/L1ghtman2k/ScoreTrak/pkg/config"
	"github.com/L1ghtman2k/ScoreTrak/pkg/master/run"
	"github.com/L1ghtman2k/ScoreTrak/pkg/master/server/gorilla"
	"github.com/L1ghtman2k/ScoreTrak/pkg/service"
	"github.com/L1ghtman2k/ScoreTrak/pkg/storage/orm"
	. "github.com/L1ghtman2k/ScoreTrak/test"
	"net"
	"net/http"
	"net/url"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestServiceSpec(t *testing.T) {
	var c config.StaticConfig
	autoTest := os.Getenv("AUTO_TEST")
	if autoTest == "TRUE" {
		c = NewConfigClone(SetupConfig("../../../configs/test-config.yml"))
	} else {
		c = NewConfigClone(SetupConfig("dev-config.yml"))
	}
	c.DB.Cockroach.Database = "scoretrak_test_api_service"
	c.Logger.FileName = "service_test.log"
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
	cr := orm.NewServiceRepo(db, l)
	serviceSvc := service.NewServiceServ(cr)
	routes = append(routes, gorilla.ServiceRoutes(l, serviceSvc, nil, run.RepoStore{})...)
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
	Convey("Initializing service repo and controller", t, func() {
		DataPreload(db)
		s := client.NewScoretrakClient(&url.URL{Host: fmt.Sprintf("localhost:%d", port), Scheme: "http"}, "", http.DefaultClient)
		cli := client.NewServiceClient(s)
		Convey("Retrieving a service by ID", func() {
			retService, err := cli.GetByID(1)
			So(err, ShouldBeNil)
			So(retService.ID, ShouldEqual, 1)
			So(retService.Name, ShouldEqual, "WINRM")
			So(retService.DisplayName, ShouldEqual, "host1-service1")
		})

		Convey("Retrieving a service by wrong ID", func() {
			retService, err := cli.GetByID(20)
			So(err, ShouldNotBeNil)
			So(retService, ShouldBeNil)
			seer, ok := err.(*client.InvalidResponse)
			So(ok, ShouldBeTrue)
			So(seer.ResponseCode, ShouldHaveSameTypeAs, http.StatusNotFound)
		})

		Convey("Updating a service by ID", func() {
			t := service.Service{ID: 1, Name: "SSH", DisplayName: "name-change-test", Points: 80}
			err := cli.Update(&t)
			So(err, ShouldBeNil)
			Convey("Retrieving a service by ID", func() {
				retService, err := cli.GetByID(1)
				So(err, ShouldBeNil)
				So(retService.ID, ShouldEqual, 1)
				So(retService.Name, ShouldEqual, "SSH")
				So(retService.DisplayName, ShouldEqual, "name-change-test")
				So(retService.Points, ShouldEqual, 80)
			})
		})

		Convey("Getting all services", func() {
			services, err := cli.GetAll()
			So(err, ShouldBeNil)
			So(len(services), ShouldEqual, 8)
			var IDs []uint64
			for _, hst := range services {
				IDs = append(IDs, hst.ID)
			}
			So(IDs, ShouldContain, uint64(1))
		})

		Convey("Deleting a non existent service", func() {
			err := cli.Delete(14)
			So(err, ShouldBeNil)

			services, err := cli.GetAll()
			So(err, ShouldBeNil)
			So(len(services), ShouldEqual, 8)
		})

		Convey("Deleting an existent service", func() {
			err := cli.Delete(5)
			So(err, ShouldBeNil)

			services, err := cli.GetAll()
			So(err, ShouldBeNil)
			So(len(services), ShouldEqual, 7)
		})

		Convey("Storing a new service", func() {
			t := service.Service{ID: 20, ServiceGroupID: 1, Name: "IMAP", DisplayName: "test-display-name", HostID: 3}
			err := cli.Store(&t)
			So(err, ShouldBeNil)
			Convey("Getting all services", func() {
				properties, err := cli.GetAll()
				So(err, ShouldBeNil)
				So(len(properties), ShouldEqual, 9)
			})
		})

		Reset(func() {
			CleanAllTables(db)
		})
	})
	DropDB(db, c)
	db.Close()
}
