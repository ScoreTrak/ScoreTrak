package client

import (
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/cmd/master/server/gorilla"
	"github.com/L1ghtman2k/ScoreTrak/pkg/api/client"
	"github.com/L1ghtman2k/ScoreTrak/pkg/config"
	"github.com/L1ghtman2k/ScoreTrak/pkg/host"
	"github.com/L1ghtman2k/ScoreTrak/pkg/storage/orm"
	. "github.com/L1ghtman2k/ScoreTrak/test"
	"github.com/gofrs/uuid"
	. "github.com/smartystreets/goconvey/convey"
	"net"
	"net/http"
	"net/url"
	"os"
	"testing"
)

func TestHostSpec(t *testing.T) {
	var c config.StaticConfig
	autoTest := os.Getenv("AUTO_TEST")
	if autoTest == "TRUE" {
		c = NewConfigClone(SetupConfig("../../../configs/test-config.yml"))
	} else {
		c = NewConfigClone(SetupConfig("dev-config.yml"))
	}
	c.DB.Cockroach.Database = "scoretrak_test_api_host"
	c.Logger.FileName = "host_test.log"
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
	cr := orm.NewHostRepo(db, l)
	hostSvc := host.NewHostServ(cr)
	routes = append(routes, gorilla.HostRoutes(l, hostSvc)...)
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
	Convey("Initializing host repo and controller", t, func() {
		DataPreload(db)
		s := client.NewScoretrakClient(&url.URL{Host: fmt.Sprintf("localhost:%d", port), Scheme: "http"}, "", http.DefaultClient)
		cli := client.NewHostClient(s)
		Convey("Retrieving a host by ID", func() {
			retHost, err := cli.GetByID(uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"))
			So(err, ShouldBeNil)
			So(retHost.ID, ShouldEqual, uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"))
			So(*(retHost.Enabled), ShouldBeTrue)
			So(*(retHost.Address), ShouldEqual, "10.0.0.1")
		})
		Convey("Retrieving a host by wrong ID", func() {
			retHost, err := cli.GetByID(uuid.FromStringOrNil("20202020-2020-2020-2020-202020202020"))
			So(err, ShouldNotBeNil)
			So(retHost, ShouldBeNil)
			seer, ok := err.(*client.InvalidResponse)
			So(ok, ShouldBeTrue)
			So(seer.ResponseCode, ShouldHaveSameTypeAs, http.StatusNotFound)
		})

		Convey("Updating a host by ID", func() {
			fls := false
			addr := "192.168.20.21"
			t := host.Host{ID: uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"), Enabled: &fls, Address: &addr, EditHost: &fls}
			err := cli.Update(&t)
			So(err, ShouldBeNil)
			Convey("Retrieving a host by ID", func() {
				retHost, err := cli.GetByID(uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"))
				So(err, ShouldBeNil)
				So(retHost.ID, ShouldEqual, uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"))
				So(*(retHost.Enabled), ShouldBeFalse)
				So(*(retHost.Address), ShouldEqual, "192.168.20.21")
				So(*(retHost.EditHost), ShouldEqual, false)
			})
		})

		SkipConvey("Updating a host by ID with a wrong hostname", func() { //TODO: Change this to Convey once govalidations are enabled
			addr := "Wrong Hostname"
			t := host.Host{ID: uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"), Address: &addr}
			err := cli.Update(&t)
			So(err, ShouldNotBeNil)
			seer, ok := err.(*client.InvalidResponse)
			So(ok, ShouldBeTrue)
			So(seer.ResponseCode, ShouldHaveSameTypeAs, http.StatusPreconditionFailed)
			Convey("Retrieving a host by ID", func() {
				retHost, err := cli.GetByID(uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"))
				So(err, ShouldBeNil)
				So(*(retHost.Address), ShouldEqual, "10.0.0.1")
			})
		})

		Convey("Getting all hosts", func() {
			hosts, err := cli.GetAll()
			So(err, ShouldBeNil)
			So(len(hosts), ShouldEqual, 4)
			var IDs []uuid.UUID
			for _, hst := range hosts {
				IDs = append(IDs, hst.ID)
			}
			So(IDs, ShouldContain, uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"))
		})

		Convey("Deleting a host that doesnt have child hosts by ID", func() {
			err := cli.Delete(uuid.FromStringOrNil("33333333-3333-3333-3333-333333333333"))
			So(err, ShouldBeNil)
			Convey("Getting all hosts", func() {
				hosts, err := cli.GetAll()
				So(err, ShouldBeNil)
				So(len(hosts), ShouldEqual, 3)
			})
		})

		Convey("Deleting a host that does have child service by ID", func() {
			err := cli.Delete(uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"))
			So(err, ShouldNotBeNil)
			seer, ok := err.(*client.InvalidResponse)
			So(ok, ShouldBeTrue)
			So(seer.ResponseCode, ShouldHaveSameTypeAs, http.StatusConflict)
			Convey("Getting all hosts", func() {
				hosts, err := cli.GetAll()
				So(err, ShouldBeNil)
				So(len(hosts), ShouldEqual, 4)
			})
		})

		Convey("Deleting a non existent host", func() {
			err := cli.Delete(uuid.FromStringOrNil("55555555-5555-5555-5555-555555555555"))
			So(err, ShouldBeNil)
		})

		Convey("Storing a new host", func() {
			addr := "test.com"
			t := []*host.Host{{Address: &addr, TeamID: uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111")}}
			err := cli.Store(t)
			So(err, ShouldBeNil)
			Convey("Getting all hosts", func() {
				hosts, err := cli.GetAll()
				So(err, ShouldBeNil)
				So(len(hosts), ShouldEqual, 5)
			})
		})

		SkipConvey("Storing a new host without specifying address", func() { //TODO: Change this to Convey once govalidations are enabled
			t := []*host.Host{{TeamID: uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111")}}
			err := cli.Store(t)
			So(err, ShouldNotBeNil)
			Convey("Getting all hosts", func() {
				hosts, err := cli.GetAll()
				So(err, ShouldBeNil)
				So(len(hosts), ShouldEqual, 4)
			})
		})

		SkipConvey("Storing a new host with specifying wrong address", func() { //TODO: Change this to Convey once govalidations are enabled
			baddr := "test com"
			t := []*host.Host{{Address: &baddr, TeamID: uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111")}}
			err := cli.Store(t)
			So(err, ShouldNotBeNil)
			seer, ok := err.(*client.InvalidResponse)
			So(ok, ShouldBeTrue)
			So(seer.ResponseCode, ShouldHaveSameTypeAs, http.StatusPreconditionFailed)
			Convey("Getting all hosts", func() {
				hosts, err := cli.GetAll()
				So(err, ShouldBeNil)
				So(len(hosts), ShouldEqual, 4)
			})
		})

		Reset(func() {
			CleanAllTables(db)
		})
	})
	DropDB(db, c)

}
