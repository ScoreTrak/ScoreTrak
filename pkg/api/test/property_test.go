package client

import (
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/cmd/master/server/gorilla"
	"github.com/ScoreTrak/ScoreTrak/pkg/api/client"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	. "github.com/ScoreTrak/ScoreTrak/pkg/config/util"
	. "github.com/ScoreTrak/ScoreTrak/pkg/logger/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/property"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/orm"
	. "github.com/ScoreTrak/ScoreTrak/pkg/storage/orm/util"

	"github.com/gofrs/uuid"
	"net"
	"net/http"
	"net/url"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPropertySpec(t *testing.T) {
	var c config.StaticConfig
	autoTest := os.Getenv("AUTO_TEST")
	if autoTest == "TRUE" {
		c = NewConfigClone(SetupConfig("../../../configs/test-config.yml"))
	} else {
		c = NewConfigClone(SetupConfig("dev-config.yml"))
	}
	c.DB.Cockroach.Database = "scoretrak_test_api_property"
	c.Logger.FileName = "property_test.log"
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
	cr := orm.NewPropertyRepo(db, l)
	propertySvc := property.NewPropertyServ(cr)
	routes = append(routes, gorilla.PropertyRoutes(l, propertySvc)...)
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
	Convey("Initializing property repo and controller", t, func() {
		CreateAllTables(db)
		DataPreload(db)
		s := client.NewScoretrakClient(&url.URL{Host: fmt.Sprintf("localhost:%d", port), Scheme: "http"}, "", http.DefaultClient)
		cli := client.NewPropertyClient(s)
		Convey("Retrieving a property by ID", func() {
			retProperty, err := cli.GetByID(uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"))
			So(err, ShouldBeNil)
			So(retProperty.ID, ShouldEqual, uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"))
			So(retProperty.Status, ShouldEqual, "View")
			So(retProperty.Value, ShouldEqual, "80")
		})
		Convey("Retrieving a property by wrong ID", func() {
			retProperty, err := cli.GetByID(uuid.FromStringOrNil("20202020-2020-2020-2020-202020202020"))
			So(err, ShouldNotBeNil)
			So(retProperty, ShouldBeNil)
			seer, ok := err.(*client.InvalidResponse)
			So(ok, ShouldBeTrue)
			So(seer.ResponseCode, ShouldHaveSameTypeAs, http.StatusNotFound)
		})

		Convey("Updating a property by ID", func() {
			t := property.Property{ID: uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"), Value: "8080", Status: "Edit"}
			err := cli.Update(&t)
			So(err, ShouldBeNil)
			Convey("Retrieving a property by ID", func() {
				retProperty, err := cli.GetByID(uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"))
				So(err, ShouldBeNil)
				So(retProperty.ID, ShouldEqual, uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"))
				So(retProperty.Value, ShouldEqual, "8080")
				So(retProperty.Status, ShouldEqual, "Edit")
			})
		})

		Convey("Getting all properties", func() {
			properties, err := cli.GetAll()
			So(err, ShouldBeNil)
			So(len(properties), ShouldEqual, 13)
			var IDs []uuid.UUID
			for _, hst := range properties {
				IDs = append(IDs, hst.ID)
			}
			So(IDs, ShouldContain, uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"))
		})

		Convey("Deleting a non existent property", func() {
			err := cli.Delete(uuid.FromStringOrNil("11211111-1111-1111-1111-111111111333"))
			So(err, ShouldBeNil)

			properties, err := cli.GetAll()
			So(err, ShouldBeNil)
			So(len(properties), ShouldEqual, 13)
		})

		Convey("Deleting an existent property", func() {
			err := cli.Delete(uuid.FromStringOrNil("55555555-5555-5555-5555-555555555555"))
			So(err, ShouldBeNil)

			properties, err := cli.GetAll()
			So(err, ShouldBeNil)
			So(len(properties), ShouldEqual, 12)
		})

		Convey("Storing a new property", func() {
			t := []*property.Property{{ID: uuid.FromStringOrNil("20202020-2020-2020-2020-202020202020"), ServiceID: uuid.FromStringOrNil("22222222-2222-2222-2222-222222222222"), Status: "Edit", Key: "Port", Value: "3001"}}
			err := cli.Store(t)
			So(err, ShouldBeNil)
			Convey("Getting all properties", func() {
				properties, err := cli.GetAll()
				So(err, ShouldBeNil)
				So(len(properties), ShouldEqual, 14)
			})
		})

		Reset(func() {
			CleanAllTables(db)
		})
	})
	DropDB(db, c)

}
