package client

import (
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/api/client"
	"github.com/L1ghtman2k/ScoreTrak/pkg/config"
	"github.com/L1ghtman2k/ScoreTrak/pkg/master/server"
	"github.com/L1ghtman2k/ScoreTrak/pkg/property"
	"github.com/L1ghtman2k/ScoreTrak/pkg/storage/orm"
	. "github.com/L1ghtman2k/ScoreTrak/test"
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
		DataPreload(db)
		s := client.NewScoretrakClient(&url.URL{Host: fmt.Sprintf("localhost:%d", port), Scheme: "http"}, "", http.DefaultClient)
		cli := client.NewPropertyClient(s)
		Convey("Retrieving a property by ID", func() {
			retProperty, err := cli.GetByID(1)
			So(err, ShouldBeNil)
			So(retProperty.ID, ShouldEqual, 1)
			So(retProperty.Status, ShouldEqual, "View")
			So(retProperty.Value, ShouldEqual, "80")
		})
		Convey("Retrieving a property by wrong ID", func() {
			retProperty, err := cli.GetByID(20)
			So(err, ShouldNotBeNil)
			So(retProperty, ShouldBeNil)
			seer, ok := err.(*client.InvalidResponse)
			So(ok, ShouldBeTrue)
			So(seer.ResponseCode, ShouldHaveSameTypeAs, http.StatusNotFound)
		})

		Convey("Updating a property by ID", func() {
			t := property.Property{ID: 1, Value: "8080", Status: "Edit"}
			err := cli.Update(&t)
			So(err, ShouldBeNil)
			Convey("Retrieving a property by ID", func() {
				retProperty, err := cli.GetByID(1)
				So(err, ShouldBeNil)
				So(retProperty.ID, ShouldEqual, 1)
				So(retProperty.Value, ShouldEqual, "8080")
				So(retProperty.Status, ShouldEqual, "Edit")
			})
		})

		Convey("Getting all properties", func() {
			properties, err := cli.GetAll()
			So(err, ShouldBeNil)
			So(len(properties), ShouldEqual, 13)
			var IDs []uint64
			for _, hst := range properties {
				IDs = append(IDs, hst.ID)
			}
			So(IDs, ShouldContain, uint64(1))
		})

		Convey("Deleting a non existent property", func() {
			err := cli.Delete(14)
			So(err, ShouldBeNil)

			properties, err := cli.GetAll()
			So(err, ShouldBeNil)
			So(len(properties), ShouldEqual, 13)
		})

		Convey("Deleting an existent property", func() {
			err := cli.Delete(5)
			So(err, ShouldBeNil)

			properties, err := cli.GetAll()
			So(err, ShouldBeNil)
			So(len(properties), ShouldEqual, 12)
		})

		Convey("Storing a new property", func() {
			t := property.Property{ID: 20, ServiceID: 2, Status: "Edit", Key: "Port", Value: "3001"}
			err := cli.Store(&t)
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
	db.Close()
}
