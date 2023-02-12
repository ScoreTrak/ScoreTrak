package orm

import (
	"context"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	. "github.com/ScoreTrak/ScoreTrak/pkg/config/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/round"
	. "github.com/ScoreTrak/ScoreTrak/pkg/storage/orm/testutil"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestRoundSpec(t *testing.T) {
	c, _ := LoadViperConfig("../../../configs/test-config.yml")
	db := SetupDB(c.DB)
	ctx := context.Background()
	defer TruncateAllTables(db)
	Convey("Creating Round Tables", t, func() {
		rr := NewRoundRepo(db)
		Reset(func() {
			TruncateTable(ctx, &round.Round{}, db)
		})
		Convey("When the Rounds table is empty", func() {
			Convey("There should be no entries", func() {
				ac, err := rr.GetAll(ctx)
				So(err, ShouldBeNil)
				So(len(ac), ShouldEqual, 0)
			})

			Convey("Adding an entry with no ID", func() {
				var err error
				r := round.Round{}
				err = rr.Store(ctx, &r)
				So(err, ShouldNotBeNil)

				Convey("Should output no entry", func() {
					ac, err := rr.GetAll(ctx)
					So(err, ShouldBeNil)
					So(len(ac), ShouldEqual, 0)
				})
			})

			Convey("Last GetLastNonElapsingRound should output last round that has END set", func() {
				rnd, err := rr.GetLastNonElapsingRound(ctx)
				So(err, ShouldNotBeNil)
				So(rnd, ShouldBeNil)
			})

			Convey("Adding a valid entry", func() {
				var err error
				now := time.Now()
				r := round.Round{ID: 1, Finish: &now}
				err = rr.Store(ctx, &r)
				So(err, ShouldBeNil)
				Convey("Should output one entry", func() {
					ac, err := rr.GetAll(ctx)
					So(err, ShouldBeNil)
					So(len(ac), ShouldEqual, 1)
					So(ac[0].Start.UnixNano(), ShouldBeBetween, time.Now().Add(time.Second*-2).UnixNano(), time.Now().Add(time.Second*2).UnixNano())
				})

				Convey("Adding an entry with the same ID", func() {
					var err error
					r := round.Round{ID: 1}
					err = rr.Store(ctx, &r)
					So(err, ShouldNotBeNil)
					// Ignored as using sqlite conn
					//var serr *pgconn.PgError
					//ok := errors.As(err, &serr)
					//So(ok, ShouldBeTrue)
					//So(serr.Code, ShouldEqual, "23505")
				})

				Convey("Then Deleting a wrong entry", func() {
					err = rr.Delete(ctx, 3)
					So(err, ShouldNotBeNil)
					Convey("Should output one entry", func() {
						ac, err := rr.GetAll(ctx)
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
					})
				})

				Convey("Then Deleting the added entry", func() {
					err = rr.Delete(ctx, 1)
					So(err, ShouldBeNil)
					Convey("Should output no entries", func() {
						ac, err := rr.GetAll(ctx)
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 0)
					})
				})

				Convey("Then Retrieving entry by ID", func() {
					rnd, err := rr.GetByID(ctx, 1)
					So(err, ShouldBeNil)
					Convey("Should output the inserted entry", func() {
						So(rnd.ID, ShouldEqual, 1)
						So(rnd.Start.UnixNano(), ShouldBeBetween, time.Now().Add(time.Second*-2).UnixNano(), time.Now().Add(time.Second*2).UnixNano())
					})
				})

				Convey("Then Querying By wrong ID", func() {
					ss, err := rr.GetByID(ctx, r.ID+1)
					So(err, ShouldNotBeNil)
					So(ss, ShouldBeNil)
				})

				Convey("Then Updating Finish to time.Now()", func() {
					now := time.Now()
					newRound := &round.Round{Finish: &now}
					Convey("For the wrong entry should not update anything", func() {
						newRound.ID = 5
						err = rr.Update(ctx, newRound)
						So(err, ShouldBeNil)
						ac, err := rr.GetAll(ctx)
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)

					})
					Convey("For the correct entry should update", func() {
						newRound.ID = 1
						err = rr.Update(ctx, newRound)
						So(err, ShouldBeNil)
						ac, err := rr.GetAll(ctx)
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
						So((*(ac[0].Finish)).UnixNano(), ShouldBeBetween, time.Now().Add(time.Second*-1).UnixNano(), time.Now().UnixNano())
					})
				})

				Convey("Last GetLastElapsingRound should output last round", func() {
					rnd, err := rr.GetLastElapsingRound(ctx)
					So(err, ShouldNotBeNil)
					So(rnd, ShouldBeNil)
				})

				Convey("Creating more rounds", func() {
					r2 := round.Round{ID: 2, Finish: &now}
					r3 := round.Round{ID: 3, Finish: &now}
					r4 := round.Round{ID: 4}
					err = rr.Update(ctx, &r)
					So(err, ShouldBeNil)
					err = rr.Store(ctx, &r2)
					So(err, ShouldBeNil)
					err = rr.Store(ctx, &r3)
					So(err, ShouldBeNil)
					err = rr.Store(ctx, &r4)
					So(err, ShouldBeNil)
					var count int64
					db.Table("rounds").Count(&count)
					So(count, ShouldEqual, 4)
					Convey("Last GetLastNonElapsingRound should output last round that has END set", func() {
						rnd, err := rr.GetLastNonElapsingRound(ctx)
						So(err, ShouldBeNil)
						So(rnd.ID, ShouldEqual, 3)
					})
					Convey("Last GetLastElapsingRound should output last round", func() {
						rnd, err := rr.GetLastElapsingRound(ctx)
						So(err, ShouldBeNil)
						So(rnd.ID, ShouldEqual, 4)
					})
				})
				Convey("Creating Checks Table", func() {
					var count int64
					Convey("Associating a single Check with a Round", func() {
						db.Exec("INSERT INTO teams (id, name, pause) VALUES ('11111111-1111-1111-1111-111111111111', 'TeamOne', true)")
						db.Exec("INSERT INTO host_groups (id, name, pause) VALUES ('11111111-1111-1111-1111-111111111111', 'HostGroup1', true)")
						db.Exec("INSERT INTO hosts (id, address, team_id, host_group_id, pause, edit_host) VALUES ('55555555-5555-5555-5555-555555555555', '10.0.0.1', '11111111-1111-1111-1111-111111111111', '11111111-1111-1111-1111-111111111111', true, true)")
						db.Exec("INSERT INTO service_groups (id, name, enabled) VALUES ('11111111-1111-1111-1111-111111111111', 'ServiceGroup1', true)")
						db.Exec(fmt.Sprintf("INSERT INTO services (id, service_group_id, host_id, name) VALUES ('11111111-1111-1111-1111-111111111111', '11111111-1111-1111-1111-111111111111', '55555555-5555-5555-5555-555555555555', 'TestService')"))
						db.Exec(fmt.Sprintf("INSERT INTO checks (service_id, round_id, log) VALUES ('11111111-1111-1111-1111-111111111111', %d, 'TestLog')", r.ID))
						db.Table("checks").Count(&count)
						So(count, ShouldEqual, 1)
						Convey("Delete a round without deleting a check should cascade all checks", func() {
							err = rr.Delete(ctx, 1)
							So(err, ShouldBeNil)
							ac, err := rr.GetAll(ctx)
							So(err, ShouldBeNil)
							So(len(ac), ShouldEqual, 0)
							db.Table("checks").Count(&count)
							So(count, ShouldEqual, 0)
						})

						Reset(func() {
							TruncateTable(ctx, &check.Check{}, db)
						})
					})
				})
			})
		})
	})

}
