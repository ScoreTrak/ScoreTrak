package orm

import (
	"context"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	. "github.com/ScoreTrak/ScoreTrak/pkg/config/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/host"
	"github.com/ScoreTrak/ScoreTrak/pkg/hostgroup"
	. "github.com/ScoreTrak/ScoreTrak/pkg/storage/orm/testutil"
	"github.com/gofrs/uuid"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHostGroupSpec(t *testing.T) {
	var c config.StaticConfig
	autoTest := os.Getenv("AUTO_TEST")
	if autoTest == "TRUE" {
		c = NewTestConfigClone("../../../configs/test-config.yml")
	} else {
		c = NewConfigClone(SetupConfig("dev-config.yml"))
	}
	c.DB.Cockroach.Database = "scoretrak_test_orm_host_group"
	db := SetupCockroachDB(c.DB)
	t.Parallel() //t.Parallel should be placed after SetupCockroachDB because gorm has race conditions on Hook register
	ctx := context.Background()
	Convey("Creating Host Group Table", t, func() {
		db.AutoMigrate(&hostgroup.HostGroup{})
		hg := NewHostGroupRepo(db)
		Reset(func() {
			db.Migrator().DropTable(&hostgroup.HostGroup{})
		})
		Convey("When the Teams table is empty", func() {
			Convey("There should be no entries", func() {
				ac, err := hg.GetAll(ctx)
				So(err, ShouldBeNil)
				So(len(ac), ShouldEqual, 0)
			})

			Convey("Adding an valid entry", func() {
				var err error
				h := []*hostgroup.HostGroup{{ID: uuid.FromStringOrNil("33333333-3333-3333-3333-333333333333"), Name: "host group"}}
				err = hg.Store(ctx, h)
				So(err, ShouldBeNil)
				Convey("Then making sure the entry exists", func() {
					ac, err := hg.GetAll(ctx)
					So(err, ShouldBeNil)
					So(len(ac), ShouldEqual, 1)
					So(ac[0].ID, ShouldEqual, uuid.FromStringOrNil("33333333-3333-3333-3333-333333333333"))
					So(ac[0].Name, ShouldEqual, "host group")
				})

				Convey("Then getting entry by id", func() {
					ac, err := hg.GetByID(ctx, uuid.FromStringOrNil("33333333-3333-3333-3333-333333333333"))
					So(err, ShouldBeNil)
					So(ac.ID, ShouldEqual, uuid.FromStringOrNil("33333333-3333-3333-3333-333333333333"))
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
					err = hg.Delete(ctx, uuid.FromStringOrNil("33333333-3333-3333-3333-333333333333"))
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
						newHostGroup.ID = uuid.FromStringOrNil("33333333-3333-3333-3333-333333333333")
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
					db.AutoMigrate(&host.Host{})
					Convey("Associating a single Host with a host group", func() {
						db.Exec("INSERT INTO hosts (id, address, team_id, host_group_id) VALUES ('44444444-4444-4444-4444-444444444444', '192.168.1.1', '44444444-4444-4444-4444-444444444444', '33333333-3333-3333-3333-333333333333')")
						db.Table("hosts").Count(&count)
						So(count, ShouldEqual, 1)
						Convey("Delete a host group without deleting a host", func() {
							err = hg.Delete(ctx, uuid.FromStringOrNil("33333333-3333-3333-3333-333333333333"))
							So(err, ShouldNotBeNil)
							ac, err := hg.GetAll(ctx)
							So(err, ShouldBeNil)
							So(len(ac), ShouldEqual, 1)
						})
						Convey("Deleting a host then deleting a host group", func() {
							db.Exec("DELETE FROM hosts WHERE id='44444444-4444-4444-4444-444444444444'")
							err = hg.Delete(ctx, uuid.FromStringOrNil("33333333-3333-3333-3333-333333333333"))
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
							db.Migrator().DropTable(&host.Host{})
						})
					})
				})
			})
		})
	})
	DropDB(db, c)

}
