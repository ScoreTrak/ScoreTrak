package orm

import (
	"context"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	. "github.com/ScoreTrak/ScoreTrak/pkg/config/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/host"
	"github.com/ScoreTrak/ScoreTrak/pkg/host_group"
	"github.com/ScoreTrak/ScoreTrak/pkg/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage"
	. "github.com/ScoreTrak/ScoreTrak/pkg/storage/orm/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/team"
	"github.com/gofrs/uuid"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestHostSpec(t *testing.T) {
	var c config.StaticConfig
	autoTest := os.Getenv("AUTO_TEST")
	if autoTest == "TRUE" {
		c = NewConfigClone(SetupConfig("../../../configs/test-config.yml"))
	} else {
		c = NewConfigClone(SetupConfig("dev-config.yml"))
	}
	c.DB.Cockroach.Database = "scoretrak_test_orm_host"
	db := storage.SetupDB(c.DB)
	t.Parallel() //t.Parallel should be placed after SetupDB because gorm has race conditions on Hook register
	ctx := context.Background()
	Convey("Creating Host Table", t, func() {
		db.AutoMigrate(&host.Host{})
		db.AutoMigrate(&team.Team{})
		db.Exec("INSERT INTO teams (id, name, enabled) VALUES ('11111111-1111-1111-1111-111111111111', 'HostGroup1', true)")
		db.Exec("INSERT INTO teams (id, name, enabled) VALUES ('22222222-2222-2222-2222-222222222222', 'HostGroup2', false)")
		hr := NewHostRepo(db)
		Reset(func() {
			db.Migrator().DropTable(&host.Host{})
		})
		Convey("When the Host table is empty", func() {
			Convey("There should be no entries", func() {
				ac, err := hr.GetAll(ctx)
				So(err, ShouldBeNil)
				So(len(ac), ShouldEqual, 0)
			})

			Convey("Adding an valid entry", func() {
				var err error
				b := false
				tr := true
				s := "127.0.0.1"
				saddresses := "192.168.0.202/20,127.0.0.1,google.com,test.ubnetdef.org"
				h := []*host.Host{{ID: uuid.FromStringOrNil("33333333-3333-3333-3333-333333333333"), Address: s, AddressListRange: &saddresses, Enabled: &b, EditHost: &tr, TeamID: uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111")}}
				err = hr.Store(ctx, h)
				So(err, ShouldBeNil)
				Convey("Then making sure the entry exists", func() {
					ac, err := hr.GetAll(ctx)
					So(err, ShouldBeNil)
					So(len(ac), ShouldEqual, 1)
					So(ac[0].ID, ShouldEqual, uuid.FromStringOrNil("33333333-3333-3333-3333-333333333333"))
					So(ac[0].Address, ShouldEqual, "127.0.0.1")
					So(*(ac[0].Enabled), ShouldBeFalse)
				})

				Convey("Then getting entry by id", func() {
					ac, err := hr.GetByID(ctx, uuid.FromStringOrNil("33333333-3333-3333-3333-333333333333"))
					So(err, ShouldBeNil)
					So(ac.ID, ShouldEqual, uuid.FromStringOrNil("33333333-3333-3333-3333-333333333333"))
					So(ac.Address, ShouldEqual, "127.0.0.1")
					So(*(ac.Enabled), ShouldBeFalse)
				})

				Convey("Then Querying By wrong ID", func() {
					ss, err := hr.GetByID(ctx, uuid.FromStringOrNil("23333333-3333-3333-3333-333333333333"))
					So(err, ShouldNotBeNil)
					So(ss, ShouldBeNil)
				})

				Convey("Then Deleting a wrong entry", func() {
					err = hr.Delete(ctx, uuid.FromStringOrNil("22222222-2222-2222-2222-222222222222"))
					So(err, ShouldNotBeNil)
					Convey("Should output one entry", func() {
						ac, err := hr.GetAll(ctx)
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
					})
				})

				Convey("Then Deleting the added entry", func() {
					err = hr.Delete(ctx, uuid.FromStringOrNil("33333333-3333-3333-3333-333333333333"))
					So(err, ShouldBeNil)
					Convey("Should output no entries", func() {
						ac, err := hr.GetAll(ctx)
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 0)
					})
				})

				Convey("Then Updating address to correct value", func() {
					newHost := host.Host{Address: "google.com"}
					newHost.ID = uuid.FromStringOrNil("33333333-3333-3333-3333-333333333333")
					err = hr.Update(ctx, &newHost)
					So(err, ShouldBeNil)
				})

				Convey("Then Updating address to incorrect value", func() {
					newHost := host.Host{Address: "googleZZZ.com"}
					newHost.ID = uuid.FromStringOrNil("33333333-3333-3333-3333-333333333333")
					err = hr.Update(ctx, &newHost)
					So(err, ShouldNotBeNil)
				})

				Convey("Then Changing to incorrect AllowedAddressRange", func() {
					b := "8.8.8.8,google.com"
					newHost := host.Host{AddressListRange: &b}
					newHost.ID = uuid.FromStringOrNil("33333333-3333-3333-3333-333333333333")
					err = hr.Update(ctx, &newHost)
					So(err, ShouldNotBeNil)
				})

				Convey("Then Changing to empty AllowedAddressRange", func() {
					b := ""
					newHost := host.Host{AddressListRange: &b}
					newHost.ID = uuid.FromStringOrNil("33333333-3333-3333-3333-333333333333")
					err = hr.Update(ctx, &newHost)
					So(err, ShouldBeNil)
				})

				Convey("Then Changing to correct AllowedAddressRange", func() {
					b := "127.0.0.1/29"
					newHost := host.Host{AddressListRange: &b}
					newHost.ID = uuid.FromStringOrNil("33333333-3333-3333-3333-333333333333")
					err = hr.Update(ctx, &newHost)
					So(err, ShouldBeNil)
				})

				Convey("Then Updating Enabled to true", func() {
					b := true
					newHost := host.Host{Enabled: &b}
					Convey("For the wrong entry should not update anything", func() {
						newHost.ID = uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111")
						err = hr.Update(ctx, &newHost)
						So(err, ShouldBeNil)
						ac, err := hr.GetAll(ctx)
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
						So(*(ac[0].Enabled), ShouldBeFalse)
					})
					Convey("For the correct entry should update", func() {
						newHost.ID = uuid.FromStringOrNil("33333333-3333-3333-3333-333333333333")
						err = hr.Update(ctx, &newHost)
						So(err, ShouldBeNil)
						ac, err := hr.GetAll(ctx)
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
						So(*(ac[0].Enabled), ShouldBeTrue)
					})
				})

				Convey("Then add a host group", func() {
					db.AutoMigrate(&host_group.HostGroup{})
					Reset(func() {
						db.Migrator().DropTable(&host.Host{})
						db.Migrator().DropTable(&host_group.HostGroup{})
					})
					uuid.FromStringOrNil("")
					db.Exec("INSERT INTO host_groups (id, name, enabled) VALUES ('11111111-1111-1111-1111-111111111111', 'HostGroup1', true)")
					db.Exec("INSERT INTO host_groups (id, name, enabled) VALUES ('22222222-2222-2222-2222-222222222222', 'HostGroup2', false)")
					var count int64
					db.Table("host_groups").Count(&count)
					So(count, ShouldEqual, 2)
					Convey("Adding a new host with host group foreign key", func() {
						address := "127.0.0.1"
						tru := true
						hstg2 := uuid.FromStringOrNil("33333333-3333-3333-3333-333333333333")
						newHost := []*host.Host{{ID: uuid.FromStringOrNil("44444444-4444-4444-4444-444444444444"), HostGroupID: &hstg2, Address: address, EditHost: &tru, TeamID: uuid.FromStringOrNil("22222222-2222-2222-2222-222222222222")}}
						err := hr.Store(ctx, newHost)
						So(err, ShouldBeNil)
					})
					Convey("Updating a host with host group foreign key", func() {
						hstg1 := uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111")
						h[0].HostGroupID = &hstg1
						err := hr.Update(ctx, h[0])
						So(err, ShouldBeNil)
					})
					Convey("Updating a host with an invalid host group foreign key(Skipping this for now since foreign keys dont behave in a same way in prod, and in testing)", func() {
						hstg10 := uuid.FromStringOrNil("444333333-3333-3333-3333-333333333333")
						h[0].HostGroupID = &hstg10
						err := hr.Update(ctx, h[0])
						So(err, ShouldNotBeNil)
					})
				})
				Convey("Then add a team", func() {
					db.AutoMigrate(&team.Team{})
					Reset(func() {
						db.Migrator().DropTable(&host.Host{})
						db.Migrator().DropTable(&team.Team{})
					})
				})
				Convey("Then add a check_service", func() {
					db.AutoMigrate(&service.Service{})
					Reset(func() {
						db.Migrator().DropTable(&service.Service{})
					})
				})
			})
		})
	})
	DropDB(db, c)

}
