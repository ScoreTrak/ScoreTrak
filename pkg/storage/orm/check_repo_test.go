package orm

import (
	"context"
	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	. "github.com/ScoreTrak/ScoreTrak/pkg/config/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/round"
	"github.com/ScoreTrak/ScoreTrak/pkg/service"
	. "github.com/ScoreTrak/ScoreTrak/pkg/storage/orm/util"
	"github.com/gofrs/uuid"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
	"time"
)

func TestCheckSpec(t *testing.T) {
	var c config.StaticConfig
	autoTest := os.Getenv("AUTO_TEST")
	if autoTest == "TRUE" {
		c = NewConfigClone(SetupConfig("../../../configs/test-config.yml"))
	} else {
		c = NewConfigClone(SetupConfig("dev-config.yml"))
	}
	c.DB.Cockroach.Database = "scoretrak_test_orm_check"
	db := SetupCockroachDB(c.DB)
	t.Parallel() //t.Parallel should be placed after SetupCockroachDB because gorm has race conditions on Hook register
	Convey("Creating Round, Service and Check tables along with their foreign keys", t, func() {
		db.AutoMigrate(&service.Service{})
		db.AutoMigrate(&round.Round{})
		db.AutoMigrate(&check.Check{})
		cr := NewCheckRepo(db)
		Convey("When all tables are empty", func() {
			Convey("Should output no entry", func() {
				gt, err := cr.GetAll(context.Background())
				So(err, ShouldBeNil)
				So(len(gt), ShouldEqual, 0)
			})
			Convey("Creating a sample check should not be allowed", func() {
				c := []*check.Check{{}}
				err := cr.Store(context.Background(), c)
				So(err, ShouldNotBeNil)
				ac, err := cr.GetAll(context.Background())
				So(err, ShouldBeNil)
				So(len(ac), ShouldEqual, 0)
			})
			Convey("Load sample services and rounds", func() {
				var count int64
				db.Exec("INSERT INTO services (id, service_group_id, host_id, name) VALUES ('55555555-5555-5555-5555-555555555555', '55555555-5555-5555-5555-555555555555', '55555555-5555-5555-5555-555555555555', 'TestService')")
				db.Exec("INSERT INTO services (id, service_group_id, host_id, name) VALUES ('66666666-6666-6666-6666-666666666666', '66666666-6666-6666-6666-666666666666', '66666666-6666-6666-6666-666666666666', 'TestService')")
				db.Exec("INSERT INTO rounds (id, start) VALUES (1, ?)", time.Now())
				db.Exec("INSERT INTO rounds (id, start) VALUES (2, ?)", time.Now())
				db.Table("services").Count(&count)
				So(count, ShouldEqual, 2)
				db.Table("rounds").Count(&count)
				So(count, ShouldEqual, 2)
				Convey("Creating a sample check and associating with check_service 5 and round 1", func() {
					c := []*check.Check{{Log: "TestLog", RoundID: 1, ServiceID: uuid.FromStringOrNil("55555555-5555-5555-5555-555555555555")}}
					err := cr.Store(context.Background(), c)
					So(err, ShouldBeNil)
					Convey("Should be Allowed", func() {

						ac, err := cr.GetAll(context.Background())
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
						Convey("Then Querying By ID", func() {
							ss, err := cr.GetByRoundServiceID(context.Background(), 1, uuid.FromStringOrNil("55555555-5555-5555-5555-555555555555"))
							So(err, ShouldBeNil)
							So(ss.Log, ShouldEqual, "TestLog")
							So(ss.ServiceID, ShouldEqual, uuid.FromStringOrNil("55555555-5555-5555-5555-555555555555"))
							So(ss.RoundID, ShouldEqual, 1)
						})

						Convey("Storing round with same round id and check_service id should not be allowed", func() {
							c := []*check.Check{{Log: "TestLogSomethingElse", RoundID: 1, ServiceID: uuid.FromStringOrNil("55555555-5555-5555-5555-555555555555")}}
							err := cr.Store(context.Background(), c)
							So(err, ShouldNotBeNil)
						})

						Convey("Then Querying By round ID", func() {
							ss, err := cr.GetByRoundServiceID(context.Background(), 20, uuid.FromStringOrNil("55555555-5555-5555-5555-555555555555"))
							So(err, ShouldNotBeNil)
							So(ss, ShouldBeNil)
						})

						Convey("Then Querying By check_service ID", func() {
							ss, err := cr.GetByRoundServiceID(context.Background(), 1, uuid.Nil)
							So(err, ShouldNotBeNil)
							So(ss, ShouldBeNil)
						})

						Convey("Then Deleting the check should be allowed", func() {
							err = cr.Delete(context.Background(), 1, uuid.FromStringOrNil("55555555-5555-5555-5555-555555555555"))
							So(err, ShouldBeNil)
							ac, err = cr.GetAll(context.Background())
							So(err, ShouldBeNil)
							So(len(ac), ShouldEqual, 0)
						})

						Convey("Then Deleting a wrong entry Should output one entry", func() {

							err = cr.Delete(context.Background(), 0, uuid.Nil)
							So(err, ShouldNotBeNil)

							err = cr.Delete(context.Background(), 0, uuid.FromStringOrNil("55555555-5555-5555-5555-555555555555"))
							So(err, ShouldNotBeNil)

							err = cr.Delete(context.Background(), 1, uuid.Nil)
							So(err, ShouldNotBeNil)

							err = cr.Delete(context.Background(), 1, uuid.FromStringOrNil("66666666-6666-6666-6666-666666666666"))
							So(err, ShouldNotBeNil)

							err = cr.Delete(context.Background(), 2, uuid.FromStringOrNil("55555555-5555-5555-5555-555555555555"))
							So(err, ShouldNotBeNil)

							ac, err := cr.GetAll(context.Background())
							So(err, ShouldBeNil)
							So(len(ac), ShouldEqual, 1)

						})

						Convey("Then adding more checks. One with similar round, and the other with similar check_service", func() {
							c1 := check.Check{Log: "TestLog2", ServiceID: uuid.FromStringOrNil("66666666-6666-6666-6666-666666666666"), RoundID: 1}
							c2 := check.Check{Log: "TestLog", ServiceID: uuid.FromStringOrNil("55555555-5555-5555-5555-555555555555"), RoundID: 2}
							checks := []*check.Check{&c1, &c2}
							err = cr.Store(context.Background(), checks)
							So(err, ShouldBeNil)
							Convey("Query by round ID should return 2 entries", func() {
								checks, err := cr.GetAllByRoundID(context.Background(), 1)
								So(err, ShouldBeNil)
								So(len(checks), ShouldEqual, 2)
							})
							Convey("Query by round ID and Service ID should return an entry", func() {
								chk, err := cr.GetByRoundServiceID(context.Background(), 1, uuid.FromStringOrNil("55555555-5555-5555-5555-555555555555"))
								So(err, ShouldBeNil)
								So(chk.Log, ShouldEqual, "TestLog")
							})
						})
					})
				})
				Convey("Creating a check with wrong check_service should not be allowed", func() {
					s := []*check.Check{{Log: "TestLog", ServiceID: uuid.FromStringOrNil("95995555-5555-5555-5555-555555555555"), RoundID: 1}}
					err := cr.Store(context.Background(), s)
					So(err, ShouldNotBeNil)
					ac, err := cr.GetAll(context.Background())
					So(err, ShouldBeNil)
					So(len(ac), ShouldEqual, 0)
				})

				Convey("Creating a check with wrong round should not be allowed", func() {
					s := []*check.Check{{Log: "TestLog", ServiceID: uuid.FromStringOrNil("55555555-5555-5555-5555-555555555555"), RoundID: 4}}
					err := cr.Store(context.Background(), s)
					So(err, ShouldNotBeNil)
					ac, err := cr.GetAll(context.Background())
					So(err, ShouldBeNil)
					So(len(ac), ShouldEqual, 0)
				})

			})
		})
		Reset(func() {
			db.Migrator().DropTable(&check.Check{})
			db.Migrator().DropTable(&round.Round{})
			db.Migrator().DropTable(&service.Service{})
		})
	})
	DropDB(db, c)

}
