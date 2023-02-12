package orm

import (
	"context"
	. "github.com/ScoreTrak/ScoreTrak/pkg/config/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/host"
	"github.com/ScoreTrak/ScoreTrak/pkg/hostgroup"
	. "github.com/ScoreTrak/ScoreTrak/pkg/storage/orm/testutil"
	"github.com/ScoreTrak/ScoreTrak/pkg/team"
	"github.com/gofrs/uuid"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHostGroupSpec(t *testing.T) {
	c, _ := LoadViperConfig("../../../configs/test-config.yml")
	db := SetupDB(c.DB)
	ctx := context.Background()
	defer TruncateAllTables(db)
	Convey("Creating Host Group Table", t, func() {
		hg := NewHostGroupRepo(db)
		Convey("When the Teams table is empty", func() {
			Convey("There should be no entries", func() {
				ac, err := hg.GetAll(ctx)
				So(err, ShouldBeNil)
				So(len(ac), ShouldEqual, 0)
			})

			Convey("Adding an valid entry", func() {
				var err error
				h := []*hostgroup.HostGroup{{ID: uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"), Name: "host group"}}
				err = hg.Store(ctx, h)
				So(err, ShouldBeNil)
				Convey("Then making sure the entry exists", func() {
					ac, err := hg.GetAll(ctx)
					So(err, ShouldBeNil)
					So(len(ac), ShouldEqual, 1)
					So(ac[0].ID, ShouldEqual, uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"))
					So(ac[0].Name, ShouldEqual, "host group")
				})

				Convey("Then getting entry by id", func() {
					ac, err := hg.GetByID(ctx, uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"))
					So(err, ShouldBeNil)
					So(ac.ID, ShouldEqual, uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"))
					So(ac.Name, ShouldEqual, "host group")
				})

				Convey("Then Querying By wrong ID", func() {
					ss, err := hg.GetByID(ctx, uuid.FromStringOrNil("43333333-3333-3333-3333-333333333333"))
					So(err, ShouldNotBeNil)
					So(ss, ShouldBeNil)
				})

				Convey("Then Deleting a wrong entry", func() {
					err = hg.Delete(ctx, uuid.FromStringOrNil("23333333-3333-3333-3333-333333333333"))
					So(err, ShouldNotBeNil)
					Convey("Should output one entry", func() {
						ac, err := hg.GetAll(ctx)
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
					})
				})

				Convey("Then Deleting the added entry", func() {
					err = hg.Delete(ctx, uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"))
					So(err, ShouldBeNil)
					Convey("Should output no entries", func() {
						ac, err := hg.GetAll(ctx)
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 0)
					})
				})

				Convey("Then Updating Pause to true", func() {
					tru := true
					newHostGroup := &hostgroup.HostGroup{Pause: &tru}
					Convey("For the wrong entry should not update anything", func() {
						newHostGroup.ID = uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111")
						err = hg.Update(ctx, newHostGroup)
						So(err, ShouldBeNil)
						ac, err := hg.GetAll(ctx)
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
						So(*(ac[0].Pause), ShouldBeFalse)
					})
					Convey("For the correct entry should update", func() {
						newHostGroup.ID = uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111")
						err = hg.Update(ctx, newHostGroup)
						So(err, ShouldBeNil)
						ac, err := hg.GetAll(ctx)
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
						So(*(ac[0].Pause), ShouldBeTrue)
					})
				})

				Convey("Creating Hosts Table", func() {
					var count int64
					Convey("Associating a single Host with a host group", func() {
						db.Exec("INSERT INTO teams (id, name, pause) VALUES ('44444444-4444-4444-4444-444444444444', 'TeamFour', false)")
						db.Exec("INSERT INTO hosts (id, address, team_id, host_group_id) VALUES ('44444444-4444-4444-4444-444444444444', '192.168.1.1', '44444444-4444-4444-4444-444444444444', '11111111-1111-1111-1111-111111111111')")
						db.Table("hosts").Count(&count)
						So(count, ShouldEqual, 1)
						// Ignoring this check as hostgroup model is not properly setup to fail when it has hosts
						Convey("Delete a host group without deleting a host", func() {
							err = hg.Delete(ctx, uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"))
							So(err, ShouldNotBeNil)
							ac, err := hg.GetAll(ctx)
							So(err, ShouldBeNil)
							So(len(ac), ShouldEqual, 1)
						})
						Convey("Deleting a host then deleting a host group", func() {
							db.Exec("DELETE FROM hosts WHERE id='44444444-4444-4444-4444-444444444444'")
							err = hg.Delete(ctx, uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"))
							So(err, ShouldBeNil)
							ac, err := hg.GetAll(ctx)
							So(err, ShouldBeNil)
							So(len(ac), ShouldEqual, 0)
						})

						Convey("Updating a team pause without deleting a host should not yield error", func() {
							tru := true
							h[0].Pause = &tru
							err = hg.Update(ctx, h[0])
							So(err, ShouldBeNil)
						})

						Reset(func() {
							TruncateTable(ctx, &host.Host{}, db)
							TruncateTable(ctx, &team.Team{}, db)
						})
					})
				})
			})
		})

		Reset(func() {
			TruncateTable(ctx, &host.Host{}, db)
		})
	})

}
