package orm

import (
	"ScoreTrak/pkg/check"
	"ScoreTrak/pkg/round"
	. "ScoreTrak/test"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestRoundSpec(t *testing.T) {
	c := NewConfigClone(SetupConfig("dev-config.yml"))
	c.DB.Cockroach.Database = "scoretrak_test_orm_round"
	c.Logger.FileName = "round_test_repo.log"
	db := SetupDB(c)
	l := SetupLogger(c)
	t.Parallel() //t.Parallel should be placed after SetupDB because gorm has race conditions on Hook register
	Convey("Creating Round Tables", t, func() {
		db.AutoMigrate(&round.Round{})
		rr := NewRoundRepo(db, l)
		Reset(func() {
			db.DropTableIfExists(&round.Round{})
		})
		Convey("When the Rounds table is empty", func() {
			Convey("There should be no entries", func() {
				ac, err := rr.GetAll()
				So(err, ShouldBeNil)
				So(len(ac), ShouldEqual, 0)
			})

			Convey("Adding an entry with no ID", func() {
				var err error
				r := round.Round{}
				err = rr.Store(&r)
				So(err, ShouldNotBeNil)

				Convey("Should output no entry", func() {
					ac, err := rr.GetAll()
					So(err, ShouldBeNil)
					So(len(ac), ShouldEqual, 0)
				})
			})

			Convey("Adding a valid entry", func() {
				var err error
				r := round.Round{ID: 1}
				err = rr.Store(&r)
				So(err, ShouldBeNil)
				Convey("Should output one entry", func() {
					ac, err := rr.GetAll()
					So(err, ShouldBeNil)
					So(len(ac), ShouldEqual, 1)
					So(ac[0].Start.UnixNano(), ShouldBeBetween, time.Now().Add(time.Second*-1).UnixNano(), time.Now().UnixNano())
				})

				Convey("Then Deleting a wrong entry", func() {
					err = rr.Delete(3)
					So(err, ShouldBeNil)
					Convey("Should output one entry", func() {
						ac, err := rr.GetAll()
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
					})
				})
				Convey("Then Deleting the added entry", func() {
					err = rr.Delete(1)
					So(err, ShouldBeNil)
					Convey("Should output no entries", func() {
						ac, err := rr.GetAll()
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 0)
					})
				})

				Convey("Then Retrieving entry by ID", func() {
					rnd, err := rr.GetByID(1)
					So(err, ShouldBeNil)
					Convey("Should output the inserted entry", func() {
						So(rnd.ID, ShouldEqual, 1)
						So(rnd.Start.UnixNano(), ShouldBeBetween, time.Now().Add(time.Second*-1).UnixNano(), time.Now().UnixNano())
					})
				})

				Convey("Then Updating End to time.Now()", func() {
					now := time.Now()
					newRound := &round.Round{End: &now}
					Convey("For the wrong entry should not update anything", func() {
						newRound.ID = 5
						err = rr.Update(newRound)
						So(err, ShouldBeNil)
						ac, err := rr.GetAll()
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
						So(ac[0].End, ShouldBeNil)

					})
					Convey("For the correct entry should update", func() {
						newRound.ID = 1
						err = rr.Update(newRound)
						So(err, ShouldBeNil)
						ac, err := rr.GetAll()
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
						So((*(ac[0].End)).UnixNano(), ShouldBeBetween, time.Now().Add(time.Second*-1).UnixNano(), time.Now().UnixNano())
					})
				})

				Convey("Creating more rounds", func() {
					now := time.Now()
					r.End = &now
					r2 := round.Round{ID: 2, End: &now}
					r3 := round.Round{ID: 3, End: &now}
					r4 := round.Round{ID: 4}
					err = rr.Update(&r)
					So(err, ShouldBeNil)
					err = rr.Store(&r2)
					So(err, ShouldBeNil)
					err = rr.Store(&r3)
					So(err, ShouldBeNil)
					err = rr.Store(&r4)
					So(err, ShouldBeNil)
					var count int
					db.Table("rounds").Count(&count)
					So(count, ShouldEqual, 4)
					Convey("Last GetLastRound should output last round that has END set", func() {
						rnd, err := rr.GetLastRound()
						So(err, ShouldBeNil)
						So(rnd.ID, ShouldEqual, 3)
					})
				})
				Convey("Creating Checks Table", func() {
					var count int
					db.AutoMigrate(&check.Check{})
					db.Model(&check.Check{}).AddForeignKey("round_id", "rounds(id)", "CASCADE", "RESTRICT")
					Convey("Associating a single Check with a Round", func() {
						db.Exec(fmt.Sprintf("INSERT INTO checks (id, service_id, round_id, log) VALUES (23, 5, %d, 'TestLog')", r.ID))
						db.Table("checks").Count(&count)
						So(count, ShouldEqual, 1)
						Convey("Delete a round without deleting a check", func() {
							err = rr.Delete(1)
							So(err, ShouldBeNil)
							ac, err := rr.GetAll()
							So(err, ShouldBeNil)
							So(len(ac), ShouldEqual, 0)
							db.Table("checks").Count(&count)
							So(count, ShouldEqual, 0)
						})

						Reset(func() {
							db.DropTableIfExists(&check.Check{})
						})
					})
				})
			})
		})
	})
	DropDB(db, c)
	db.Close()
}
