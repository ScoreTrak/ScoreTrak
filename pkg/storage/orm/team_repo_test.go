package orm

import (
	"context"
	"fmt"
	. "github.com/ScoreTrak/ScoreTrak/pkg/config/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/host"
	. "github.com/ScoreTrak/ScoreTrak/pkg/storage/orm/testutil"
	"github.com/ScoreTrak/ScoreTrak/pkg/team"
	"github.com/gofrs/uuid"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestTeamSpec(t *testing.T) {
	c, _ := LoadViperConfig("../../../configs/test-config.yml")
	db := SetupSqliteDB(c.DB)
	ctx := context.Background()
	Convey("Creating Team Tables", t, func() {
		db.AutoMigrate(&team.Team{})
		tr := NewTeamRepo(db)
		Reset(func() {
			db.Migrator().DropTable(&team.Team{})
		})
		Convey("When the Teams table is empty", func() {
			Convey("There should be no entries", func() {
				ac, err := tr.GetAll(ctx)
				So(err, ShouldBeNil)
				So(len(ac), ShouldEqual, 0)
			})

			Convey("Adding an entry with empty name", func() {
				var err error
				tru := true
				t := []*team.Team{{Name: "", Pause: &tru}}
				err = tr.Store(ctx, t)
				So(err, ShouldNotBeNil)

				Convey("Should output no entry", func() {
					ac, err := tr.GetAll(ctx)
					So(err, ShouldBeNil)
					So(len(ac), ShouldEqual, 0)
				})
			})

			Convey("Adding a valid entry", func() {
				var err error
				tru := true
				t := []*team.Team{{Name: "TestTeam", Pause: &tru}}
				err = tr.Store(ctx, t)
				So(err, ShouldBeNil)
				Convey("Should output one entry", func() {
					ac, err := tr.GetAll(ctx)
					So(err, ShouldBeNil)
					So(len(ac), ShouldEqual, 1)
					So(ac[0].Name, ShouldEqual, "TestTeam")
				})

				Convey("Then Deleting a wrong entry", func() {
					err = tr.DeleteByName(ctx, "TestTeamWRONG")
					So(err, ShouldNotBeNil)
					Convey("Should output one entry", func() {
						ac, err := tr.GetAll(ctx)
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
					})
				})
				Convey("Then Deleting the added entry", func() {
					err = tr.DeleteByName(ctx, "TestTeam")
					So(err, ShouldBeNil)
					Convey("Should output no entries", func() {
						ac, err := tr.GetAll(ctx)
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 0)
					})
				})

				Convey("Then Retrieving entry by Name", func() {
					tm, err := tr.GetByName(ctx, "TestTeam")
					So(err, ShouldBeNil)
					Convey("Should output the inserted entry", func() {
						So(tm.Name, ShouldEqual, "TestTeam")
						So(*(tm.Pause), ShouldBeTrue)
					})
				})

				Convey("Then Querying By wrong Name", func() {
					ss, err := tr.GetByName(ctx, "WrongTeamName")
					So(err, ShouldNotBeNil)
					So(ss, ShouldBeNil)
				})

				Convey("Then Updating Pause to true", func() {
					fls := false
					newTeam := &team.Team{Pause: &fls}
					Convey("For the wrong entry should not update anything", func() {
						newTeam.ID = uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111")
						err = tr.Update(ctx, newTeam)
						So(err, ShouldBeNil)
						ac, err := tr.GetAll(ctx)
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
						So(*(ac[0].Pause), ShouldBeTrue)

					})
					Convey("For the correct entry should update", func() {
						newTeam.Name = "TestTeam"
						newTeam.ID = t[0].ID
						err = tr.Update(ctx, newTeam)
						So(err, ShouldBeNil)
						ac, err := tr.GetAll(ctx)
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
						So(*(ac[0].Pause), ShouldBeFalse)

					})
				})

				Convey("Creating Hosts Table", func() {
					var count int64
					db.AutoMigrate(&host.Host{})
					Convey("Associating a single Host with a team", func() {
						db.Exec(fmt.Sprintf("INSERT INTO hosts (id, address, team_id) VALUES ('44444444-4444-4444-4444-444444444444', '192.168.1.1', '%s')", t[0].ID.String()))
						db.Table("hosts").Count(&count)
						So(count, ShouldEqual, 1)
						Convey("DeleteByName a team without deleting a host", func() {
							err = tr.DeleteByName(ctx, "TestTeam")
							So(err, ShouldNotBeNil)
							ac, err := tr.GetAll(ctx)
							So(err, ShouldBeNil)
							So(len(ac), ShouldEqual, 1)
						})
						Convey("Deleting a host then deleting a team", func() {
							db.Exec("DELETE FROM hosts WHERE id='44444444-4444-4444-4444-444444444444'")
							err = tr.DeleteByName(ctx, "TestTeam")
							So(err, ShouldBeNil)
							ac, err := tr.GetAll(ctx)
							So(err, ShouldBeNil)
							So(len(ac), ShouldEqual, 0)
						})

						Convey("Updating a team pause without deleting a host should not yield error", func() {
							tru := true
							t[0].Pause = &tru
							err = tr.Update(ctx, t[0])
							So(err, ShouldBeNil)
						})

						Reset(func() {
							db.Migrator().DropTable(&host.Host{})
						})
					})
				})
			})
		})
	})
	DropDB(db, c)

}
