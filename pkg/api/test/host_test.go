package client

import (
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/api/client"
	"github.com/L1ghtman2k/ScoreTrak/pkg/config"
	"github.com/L1ghtman2k/ScoreTrak/pkg/host"
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
	rtr := server.NewRouter()
	routes := server.Routes{
		server.Route{
			Name:        "Index",
			Method:      "GET",
			Pattern:     "/",
			HandlerFunc: server.Index,
		},
	}
	cr := orm.NewHostRepo(db, l)
	hostSvc := host.NewHostServ(cr)
	routes = append(routes, server.HostRoutes(l, hostSvc)...)
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
	Convey("Initializing host repo and controller", t, func() {
		DataPreload(db)
		s := client.NewScoretrakClient(&url.URL{Host: fmt.Sprintf("localhost:%d", port), Scheme: "http"}, "", http.DefaultClient)
		cli := client.NewHostClient(s)
		Convey("Retrieving a host by ID", func() {
			retHost, err := cli.GetByID(1)
			So(err, ShouldBeNil)
			So(retHost.ID, ShouldEqual, 1)
			So(*(retHost.Enabled), ShouldBeTrue)
			So(*(retHost.Address), ShouldEqual, "10.0.0.1")
		})
		Convey("Retrieving a host by wrong ID", func() {
			retHost, err := cli.GetByID(20)
			So(err, ShouldNotBeNil)
			So(retHost, ShouldBeNil)
			seer, ok := err.(*client.InvalidResponse)
			So(ok, ShouldBeTrue)
			So(seer.ResponseCode, ShouldHaveSameTypeAs, http.StatusNotFound)
		})

		Convey("Updating a host by ID", func() {
			fls := false
			addr := "192.168.20.21"
			t := host.Host{ID: 1, Enabled: &fls, Address: &addr, EditHost: &fls}
			err := cli.Update(&t)
			So(err, ShouldBeNil)
			Convey("Retrieving a host by ID", func() {
				retHost, err := cli.GetByID(1)
				So(err, ShouldBeNil)
				So(retHost.ID, ShouldEqual, 1)
				So(*(retHost.Enabled), ShouldBeFalse)
				So(*(retHost.Address), ShouldEqual, "192.168.20.21")
				So(*(retHost.EditHost), ShouldEqual, false)
			})
		})

		Convey("Updating a host by ID with a wrong hostname", func() {
			addr := "Wrong Hostname"
			t := host.Host{ID: 1, Address: &addr}
			err := cli.Update(&t)
			So(err, ShouldNotBeNil)
			seer, ok := err.(*client.InvalidResponse)
			So(ok, ShouldBeTrue)
			So(seer.ResponseCode, ShouldHaveSameTypeAs, http.StatusPreconditionFailed)
			Convey("Retrieving a host by ID", func() {
				retHost, err := cli.GetByID(1)
				So(err, ShouldBeNil)
				So(*(retHost.Address), ShouldEqual, "10.0.0.1")
			})
		})

		Convey("Getting all hosts", func() {
			hosts, err := cli.GetAll()
			So(err, ShouldBeNil)
			So(len(hosts), ShouldEqual, 4)
			var IDs []uint64
			for _, hst := range hosts {
				IDs = append(IDs, hst.ID)
			}
			So(IDs, ShouldContain, uint64(1))
		})

		Convey("Deleting a host that doesnt have child hosts by ID", func() {
			err := cli.Delete(3)
			So(err, ShouldBeNil)
			Convey("Getting all hosts", func() {
				hosts, err := cli.GetAll()
				So(err, ShouldBeNil)
				So(len(hosts), ShouldEqual, 3)
			})
		})

		Convey("Deleting a host that does have child service by ID", func() {
			err := cli.Delete(1)
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
			err := cli.Delete(5)
			So(err, ShouldBeNil)
		})

		Convey("Storing a new host", func() {
			addr := "test.com"
			t := host.Host{Address: &addr}
			err := cli.Store(&t)
			So(err, ShouldBeNil)
			Convey("Getting all hosts", func() {
				hosts, err := cli.GetAll()
				So(err, ShouldBeNil)
				So(len(hosts), ShouldEqual, 5)
			})
		})

		Convey("Storing a new host without specifying address", func() {
			t := host.Host{}
			err := cli.Store(&t)
			So(err, ShouldNotBeNil)
			Convey("Getting all hosts", func() {
				hosts, err := cli.GetAll()
				So(err, ShouldBeNil)
				So(len(hosts), ShouldEqual, 4)
			})
		})

		Convey("Storing a new host with specifying wrong address", func() {
			baddr := "test com"
			t := host.Host{Address: &baddr}
			err := cli.Store(&t)
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
	db.Close()
}
