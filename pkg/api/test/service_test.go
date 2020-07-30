package client

import (
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/cmd/master/server/gorilla"
	"github.com/L1ghtman2k/ScoreTrak/pkg/api/client"
	"github.com/L1ghtman2k/ScoreTrak/pkg/config"
	"github.com/L1ghtman2k/ScoreTrak/pkg/service"
	"github.com/L1ghtman2k/ScoreTrak/pkg/storage/orm"
	"github.com/L1ghtman2k/ScoreTrak/pkg/storage/util"
	. "github.com/L1ghtman2k/ScoreTrak/test"
	"github.com/gofrs/uuid"
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
	routes = append(routes, gorilla.ServiceRoutes(l, serviceSvc, nil, util.RepoStore{})...)
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
			retService, err := cli.GetByID(uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"))
			So(err, ShouldBeNil)
			So(retService.ID, ShouldEqual, uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"))
			So(retService.Name, ShouldEqual, "WINRM")
			So(retService.DisplayName, ShouldEqual, "host1-service1")
		})

		Convey("Retrieving a service by wrong ID", func() {
			retService, err := cli.GetByID(uuid.FromStringOrNil("11111111-1111-1111-1111-111111111333"))
			So(err, ShouldNotBeNil)
			So(retService, ShouldBeNil)
			seer, ok := err.(*client.InvalidResponse)
			So(ok, ShouldBeTrue)
			So(seer.ResponseCode, ShouldHaveSameTypeAs, http.StatusNotFound)
		})

		Convey("Updating a service by ID", func() {
			t := service.Service{ID: uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"), Name: "SSH", DisplayName: "name-change-test", Points: 80}
			err := cli.Update(&t)
			So(err, ShouldBeNil)
			Convey("Retrieving a service by ID", func() {
				retService, err := cli.GetByID(uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"))
				So(err, ShouldBeNil)
				So(retService.ID, ShouldEqual, uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"))
				So(retService.Name, ShouldEqual, "SSH")
				So(retService.DisplayName, ShouldEqual, "name-change-test")
				So(retService.Points, ShouldEqual, 80)
			})
		})

		Convey("Getting all services", func() {
			services, err := cli.GetAll()
			So(err, ShouldBeNil)
			So(len(services), ShouldEqual, 8)
			var IDs []uuid.UUID
			for _, hst := range services {
				IDs = append(IDs, hst.ID)
			}
			So(IDs, ShouldContain, uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"))
		})

		Convey("Deleting a non existent service", func() {
			err := cli.Delete(uuid.FromStringOrNil("11111111-1111-1111-1111-111111111333"))
			So(err, ShouldBeNil)

			services, err := cli.GetAll()
			So(err, ShouldBeNil)
			So(len(services), ShouldEqual, 8)
		})

		Convey("Deleting a service that exists", func() {
			err := cli.Delete(uuid.FromStringOrNil("55555555-5555-5555-5555-555555555555"))
			So(err, ShouldBeNil)

			services, err := cli.GetAll()
			So(err, ShouldBeNil)
			So(len(services), ShouldEqual, 7)
		})

		Convey("Storing a new service", func() {
			t := []*service.Service{{ID: uuid.FromStringOrNil("20202020-2020-2020-2020-202020202020"), ServiceGroupID: uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"), Name: "IMAP", DisplayName: "test-display-name", HostID: uuid.FromStringOrNil("33333333-3333-3333-3333-333333333333")}}
			err := cli.Store(t)
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

}
