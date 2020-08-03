package client

import (
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/cmd/master/server/gorilla"
	"github.com/ScoreTrak/ScoreTrak/pkg/api/client"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	. "github.com/ScoreTrak/ScoreTrak/pkg/config/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/host_group"
	. "github.com/ScoreTrak/ScoreTrak/pkg/logger/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/orm"
	. "github.com/ScoreTrak/ScoreTrak/pkg/storage/orm/util"

	"github.com/gofrs/uuid"
	. "github.com/smartystreets/goconvey/convey"
	"net"
	"net/http"
	"net/url"
	"os"
	"testing"
)

func TestHostGroupSpec(t *testing.T) {
	var c config.StaticConfig
	autoTest := os.Getenv("AUTO_TEST")
	if autoTest == "TRUE" {
		c = NewConfigClone(SetupConfig("../../../configs/test-config.yml"))
	} else {
		c = NewConfigClone(SetupConfig("dev-config.yml"))
	}
	c.DB.Cockroach.Database = "scoretrak_test_api_host_group"
	c.Logger.FileName = "host_group_test.log"
	db := storage.SetupDB(c.DB)
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
	cr := orm.NewHostGroupRepo(db, l)
	hostGroupSvc := host_group.NewHostGroupServ(cr)
	routes = append(routes, gorilla.HostGroupRoutes(l, hostGroupSvc)...)
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
	Convey("Initializing Host Group repo and controller", t, func() {
		CreateAllTables(db)
		DataPreload(db)
		s := client.NewScoretrakClient(&url.URL{Host: fmt.Sprintf("localhost:%d", port), Scheme: "http"}, "", http.DefaultClient)
		cli := client.NewHostGroupClient(s)
		Convey("Retrieving a Host Group by ID", func() {
			retHostGroup, err := cli.GetByID(uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"))
			So(err, ShouldBeNil)
			So(retHostGroup.ID, ShouldEqual, uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"))
			So(*(retHostGroup.Enabled), ShouldBeTrue)
		})
		Convey("Retrieving a Host Group by wrong ID", func() {
			retHostGroup, err := cli.GetByID(uuid.FromStringOrNil("55555555-5555-5555-5555-555555555555"))
			So(err, ShouldNotBeNil)
			So(retHostGroup, ShouldBeNil)
			seer, ok := err.(*client.InvalidResponse)
			So(ok, ShouldBeTrue)
			So(seer.ResponseCode, ShouldHaveSameTypeAs, http.StatusNotFound)
		})

		Convey("Updating a Host Group by ID", func() {
			fls := false
			t := host_group.HostGroup{ID: uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"), Enabled: &fls}
			err := cli.Update(&t)
			So(err, ShouldBeNil)
			Convey("Retrieving a Host Group by ID", func() {
				retHostGroup, err := cli.GetByID(uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"))
				So(err, ShouldBeNil)
				So(retHostGroup.Name, ShouldEqual, "HostGroup1")
				So(*(retHostGroup.Enabled), ShouldBeFalse)
			})
		})

		Convey("Getting all Host Group", func() {
			hostGroups, err := cli.GetAll()
			So(err, ShouldBeNil)
			So(len(hostGroups), ShouldEqual, 4)
			var IDs []uuid.UUID
			for _, tm := range hostGroups {
				IDs = append(IDs, tm.ID)
			}
			So(IDs, ShouldContain, uuid.FromStringOrNil("33333333-3333-3333-3333-333333333333"))
		})

		Convey("Deleting a Host Group that doesnt have child hosts by ID", func() {
			err := cli.Delete(uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"))
			So(err, ShouldBeNil)
			Convey("Getting all hostGroups", func() {
				hostGroups, err := cli.GetAll()
				So(err, ShouldBeNil)
				So(len(hostGroups), ShouldEqual, 3)
			})
		})

		Convey("Deleting a Host Group that does have child hosts by ID", func() {
			err := cli.Delete(uuid.FromStringOrNil("33333333-3333-3333-3333-333333333333"))
			So(err, ShouldNotBeNil)
			seer, ok := err.(*client.InvalidResponse)
			So(ok, ShouldBeTrue)
			So(seer.ResponseCode, ShouldHaveSameTypeAs, http.StatusConflict)
			Convey("Getting all hostGroups", func() {
				hostGroups, err := cli.GetAll()
				So(err, ShouldBeNil)
				So(len(hostGroups), ShouldEqual, 4)
			})
		})

		Convey("Deleting a non existent Host Group", func() {
			err := cli.Delete(uuid.FromStringOrNil("66666666-6666-6666-6666-666666666666"))
			So(err, ShouldBeNil)
		})

		Convey("Storing a new Host Group", func() {
			fls := false
			t := []*host_group.HostGroup{{Enabled: &fls, Name: "HostGroup5"}}
			err := cli.Store(t)
			So(err, ShouldBeNil)
			Convey("Getting all Host Group", func() {
				hostGroups, err := cli.GetAll()
				So(err, ShouldBeNil)
				So(len(hostGroups), ShouldEqual, 5)
			})
		})

		Convey("Storing a new Host Group with the same name", func() {
			fls := false
			t := []*host_group.HostGroup{{Enabled: &fls, Name: "HostGroup1"}}
			err := cli.Store(t)
			So(err, ShouldNotBeNil)
			seer, ok := err.(*client.InvalidResponse)
			So(ok, ShouldBeTrue)
			So(seer.ResponseCode, ShouldHaveSameTypeAs, http.StatusPreconditionFailed)
			Convey("Getting all Host Group", func() {
				hostGroups, err := cli.GetAll()
				So(err, ShouldBeNil)
				So(len(hostGroups), ShouldEqual, 4)
			})
		})

		Reset(func() {
			CleanAllTables(db)
		})
	})
	DropDB(db, c)

}
