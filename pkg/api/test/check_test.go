package client

import (
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/cmd/master/server/gorilla"
	"github.com/ScoreTrak/ScoreTrak/pkg/api/client"
	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	. "github.com/ScoreTrak/ScoreTrak/pkg/config/util"
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

func TestCheckSpec(t *testing.T) {
	var c config.StaticConfig
	autoTest := os.Getenv("AUTO_TEST")
	if autoTest == "TRUE" {
		c = NewConfigClone(SetupConfig("../../../configs/test-config.yml"))
	} else {
		c = NewConfigClone(SetupConfig("dev-config.yml"))
	}
	c.DB.Cockroach.Database = "scoretrak_test_api_check"
	c.Logger.FileName = "check_test.log"
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
	cr := orm.NewCheckRepo(db, l)
	checkSvc := check.NewCheckServ(cr)
	routes = append(routes, gorilla.CheckRoutes(l, checkSvc)...)
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
	Convey("Initializing check repo and controller", t, func() {
		CreateAllTables(db)
		DataPreload(db)
		s := client.NewScoretrakClient(&url.URL{Host: fmt.Sprintf("localhost:%d", port), Scheme: "http"}, "", http.DefaultClient)
		cli := client.NewCheckClient(s)
		Convey("Retrieving checks by Round ID", func() {
			retChecks, err := cli.GetAllByRoundID(3)
			So(err, ShouldBeNil)
			So(len(retChecks), ShouldEqual, 2)
		})

		Convey("Retrieving checks by Round ID and Service ID", func() {
			retChecks, err := cli.GetAllByRoundID(3)
			So(err, ShouldBeNil)
			So(len(retChecks), ShouldEqual, 2)

			for _, chck := range retChecks {
				if chck.ServiceID == uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111") {
					So(chck.Log, ShouldEqual, "Failed because of incorrect password")
					So(*(chck.Passed), ShouldBeFalse)
				} else {
					So(chck.Log, ShouldEqual, "")
					So(*(chck.Passed), ShouldBeTrue)
				}
			}

		})

		Reset(func() {
			CleanAllTables(db)
		})
	})
	DropDB(db, c)

}
