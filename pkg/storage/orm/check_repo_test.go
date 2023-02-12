package orm

import (
	"context"
	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	. "github.com/ScoreTrak/ScoreTrak/pkg/config/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/host"
	"github.com/ScoreTrak/ScoreTrak/pkg/hostgroup"
	"github.com/ScoreTrak/ScoreTrak/pkg/round"
	"github.com/ScoreTrak/ScoreTrak/pkg/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/servicegroup"
	. "github.com/ScoreTrak/ScoreTrak/pkg/storage/orm/testutil"
	"github.com/ScoreTrak/ScoreTrak/pkg/team"
	"github.com/gofrs/uuid"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestCheckSpec(t *testing.T) {
	c, _ := LoadViperConfig("../../../configs/test-config.yml")
	db := SetupDB(c.DB)
	defer TruncateAllTables(db)
	Convey("Seed db and setup check repo", t, func() {
		db.Exec("INSERT INTO teams (id, name, pause) VALUES ('11111111-1111-1111-1111-111111111111', 'TeamOne', true)")
		db.Exec("INSERT INTO host_groups (id, name, pause) VALUES ('11111111-1111-1111-1111-111111111111', 'HostGroup1', true)")
		db.Exec("INSERT INTO hosts (id, address, team_id, host_group_id, pause, edit_host) VALUES ('11111111-1111-1111-1111-111111111111', '10.0.0.5', '11111111-1111-1111-1111-111111111111', '11111111-1111-1111-1111-111111111111', true, false)")
		db.Exec("INSERT INTO service_groups (id, name, enabled) VALUES ('11111111-1111-1111-1111-111111111111', 'ServiceGroup1', true)")
		cr := NewCheckRepo(db)
		Reset(func() {
			ctx := context.Background()
			TruncateTable(ctx, &team.Team{}, db)
			TruncateTable(ctx, &hostgroup.HostGroup{}, db)
			TruncateTable(ctx, &host.Host{}, db)
			TruncateTable(ctx, &servicegroup.ServiceGroup{}, db)
			TruncateTable(ctx, &check.Check{}, db)
			TruncateTable(ctx, &round.Round{}, db)
			TruncateTable(ctx, &service.Service{}, db)
		})
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
				db.Exec("INSERT INTO services (id, service_group_id, host_id, name) VALUES ('11111111-1111-1111-1111-111111111111', '11111111-1111-1111-1111-111111111111', '11111111-1111-1111-1111-111111111111', 'TestService1')")
				db.Exec("INSERT INTO services (id, service_group_id, host_id, name) VALUES ('22222222-2222-2222-2222-222222222222', '11111111-1111-1111-1111-111111111111', '11111111-1111-1111-1111-111111111111', 'TestService2')")
				db.Exec("INSERT INTO rounds (id, start) VALUES (1, ?)", time.Now())
				db.Exec("INSERT INTO rounds (id, start) VALUES (2, ?)", time.Now())
				db.Table("services").Count(&count)
				So(count, ShouldEqual, 2)
				db.Table("rounds").Count(&count)
				So(count, ShouldEqual, 2)
				Reset(func() {
					ctx := context.Background()
					TruncateTable(ctx, &round.Round{}, db)
					TruncateTable(ctx, &service.Service{}, db)
				})
				Convey("Creating a sample check and associating with check_service 5 and round 1", func() {
					c := []*check.Check{{Log: "TestLog", RoundID: 1, ServiceID: uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111")}}
					err := cr.Store(context.Background(), c)
					So(err, ShouldBeNil)
					Convey("Should be Allowed", func() {

						ac, err := cr.GetAll(context.Background())
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
						Convey("Then Querying By ID", func() {
							ss, err := cr.GetByRoundServiceID(context.Background(), 1, uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"))
							So(err, ShouldBeNil)
							So(ss.Log, ShouldEqual, "TestLog")
							So(ss.ServiceID, ShouldEqual, uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"))
							So(ss.RoundID, ShouldEqual, 1)
						})

						Convey("Storing round with same round id and check_service id should not be allowed", func() {
							c := []*check.Check{{Log: "TestLogSomethingElse", RoundID: 1, ServiceID: uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111")}}
							err := cr.Store(context.Background(), c)
							So(err, ShouldNotBeNil)
						})

						Convey("Then Querying By round ID", func() {
							ss, err := cr.GetByRoundServiceID(context.Background(), 20, uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"))
							So(err, ShouldNotBeNil)
							So(ss, ShouldBeNil)
						})

						Convey("Then Querying By check_service ID", func() {
							ss, err := cr.GetByRoundServiceID(context.Background(), 1, uuid.Nil)
							So(err, ShouldNotBeNil)
							So(ss, ShouldBeNil)
						})

						Convey("Then Deleting the check should be allowed", func() {
							err = cr.Delete(context.Background(), 1, uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"))
							So(err, ShouldBeNil)
							ac, err = cr.GetAll(context.Background())
							So(err, ShouldBeNil)
							So(len(ac), ShouldEqual, 0)
						})

						Convey("Then Deleting a wrong entry Should output one entry", func() {

							err = cr.Delete(context.Background(), 0, uuid.Nil)
							So(err, ShouldNotBeNil)

							err = cr.Delete(context.Background(), 0, uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"))
							So(err, ShouldNotBeNil)

							err = cr.Delete(context.Background(), 1, uuid.Nil)
							So(err, ShouldNotBeNil)

							err = cr.Delete(context.Background(), 1, uuid.FromStringOrNil("22222222-2222-2222-2222-222222222222"))
							So(err, ShouldNotBeNil)

							err = cr.Delete(context.Background(), 2, uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"))
							So(err, ShouldNotBeNil)

							ac, err := cr.GetAll(context.Background())
							So(err, ShouldBeNil)
							So(len(ac), ShouldEqual, 1)

						})

						Convey("Then adding more checks. One with similar round, and the other with similar check_service", func() {
							c1 := check.Check{Log: "TestLog2", ServiceID: uuid.FromStringOrNil("22222222-2222-2222-2222-222222222222"), RoundID: 1}
							c2 := check.Check{Log: "TestLog", ServiceID: uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"), RoundID: 2}
							checks := []*check.Check{&c1, &c2}
							err = cr.Store(context.Background(), checks)
							So(err, ShouldBeNil)
							Convey("Query by round ID should return 2 entries", func() {
								checks, err := cr.GetAllByRoundID(context.Background(), 1)
								So(err, ShouldBeNil)
								So(len(checks), ShouldEqual, 2)
							})
							Convey("Query by round ID and Service ID should return an entry", func() {
								chk, err := cr.GetByRoundServiceID(context.Background(), 1, uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"))
								So(err, ShouldBeNil)
								So(chk.Log, ShouldEqual, "TestLog")
							})
						})
					})
				})
				Convey("Creating a check with wrong check_service should not be allowed", func() {
					s := []*check.Check{{Log: "TestLog", ServiceID: uuid.FromStringOrNil("11111111-0000-1111-1111-111111111111"), RoundID: 1}}
					err := cr.Store(context.Background(), s)
					So(err, ShouldNotBeNil)
					ac, err := cr.GetAll(context.Background())
					So(err, ShouldBeNil)
					So(len(ac), ShouldEqual, 0)
				})

				Convey("Creating a check with wrong round should not be allowed", func() {
					s := []*check.Check{{Log: "TestLog", ServiceID: uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"), RoundID: 4}}
					err := cr.Store(context.Background(), s)
					So(err, ShouldNotBeNil)
					ac, err := cr.GetAll(context.Background())
					So(err, ShouldBeNil)
					So(len(ac), ShouldEqual, 0)
				})

			})
		})

	})
}
