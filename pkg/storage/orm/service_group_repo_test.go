package orm

import (
	"context"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	. "github.com/ScoreTrak/ScoreTrak/pkg/config/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/service_group"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage"
	. "github.com/ScoreTrak/ScoreTrak/pkg/storage/orm/util"
	"github.com/gofrs/uuid"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestServiceGroupSpec(t *testing.T) {
	var c config.StaticConfig
	autoTest := os.Getenv("AUTO_TEST")
	if autoTest == "TRUE" {
		c = NewConfigClone(SetupConfig("../../../configs/test-config.yml"))
	} else {
		c = NewConfigClone(SetupConfig("dev-config.yml"))
	}
	c.DB.Cockroach.Database = "scoretrak_test_orm_service_group"
	db := storage.SetupDB(c.DB)
	ctx := context.Background()
	t.Parallel() //t.Parallel should be placed after SetupDB because gorm has race conditions on Hook register
	Convey("Creating Service Group Tables", t, func() {
		db.AutoMigrate(&service_group.ServiceGroup{})
		sgr := NewServiceGroupRepo(db)

		Convey("When the Service Group table is empty", func() {
			Convey("There should be no entries", func() {
				ac, err := sgr.GetAll(ctx)
				So(err, ShouldBeNil)
				So(len(ac), ShouldEqual, 0)
			})

			Convey("Adding a valid entry", func() {
				var err error
				s := service_group.ServiceGroup{Name: "TestServiceGroup"}
				err = sgr.Store(ctx, &s)
				So(err, ShouldBeNil)

				Convey("Should create an entry in the database", func() {
					ac, err := sgr.GetAll(ctx)
					So(err, ShouldBeNil)
					So(len(ac), ShouldEqual, 1)
				})

				Convey("And then creating an entry with same name", func() {
					tru := true
					t2 := service_group.ServiceGroup{Name: "TestServiceGroup", Enabled: &tru}
					err = sgr.Store(ctx, &t2)
					So(err, ShouldNotBeNil)
					Convey("Should not create a new entry", func() {
						ac, err := sgr.GetAll(ctx)
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
					})
				})

				Convey("Then Deleting a wrong entry", func() {
					err = sgr.Delete(ctx, uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"))
					So(err, ShouldNotBeNil)
					Convey("Should output one entry", func() {
						ac, err := sgr.GetAll(ctx)
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
					})
				})

				Convey("Then Deleting the added entry", func() {
					err = sgr.Delete(ctx, s.ID)
					So(err, ShouldBeNil)
					Convey("Should output no entries", func() {
						ac, err := sgr.GetAll(ctx)
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 0)
					})
				})

				Convey("Then Retrieving entry by ID", func() {
					sg, err := sgr.GetByID(ctx, s.ID)
					So(err, ShouldBeNil)
					Convey("Should output the inserted entry", func() {
						So(sg.Name, ShouldEqual, "TestServiceGroup")
						So(*(sg.Enabled), ShouldBeFalse)
					})
				})

				Convey("Then Querying By wrong ID", func() {
					ss, err := sgr.GetByID(ctx, uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"))
					So(err, ShouldNotBeNil)
					So(ss, ShouldBeNil)
				})

				Convey("Then Updating Enabled to true", func() {
					tru := true
					newSgr := &service_group.ServiceGroup{Enabled: &tru}
					Convey("For the wrong entry should not update anything", func() {
						newSgr.ID = uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111")
						err = sgr.Update(ctx, newSgr)
						So(err, ShouldBeNil)
						ac, err := sgr.GetAll(ctx)
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
						So(*(ac[0].Enabled), ShouldBeFalse)

					})
					Convey("For the correct entry should update", func() {
						newSgr.ID = s.ID
						err = sgr.Update(ctx, newSgr)
						So(err, ShouldBeNil)
						ac, err := sgr.GetAll(ctx)
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
						So(*(ac[0].Enabled), ShouldBeTrue)
					})
				})
				Convey("Creating Service Table", func() {
					var count int64
					db.AutoMigrate(&service.Service{})
					Convey("Then associating one check_service with the check_service group", func() {
						db.Exec(fmt.Sprintf("INSERT INTO services (id, service_group_id, host_id, name) VALUES ('55555555-5555-5555-5555-555555555555', '%s', '55555555-5555-5555-5555-555555555555', 'TestService')", s.ID.String()))
						db.Table("services").Count(&count)
						So(count, ShouldEqual, 1)
						Convey("Then Deleting the check_service group should be restricted", func() {
							err = sgr.Delete(ctx, s.ID)
							So(err, ShouldNotBeNil)
							ac, err := sgr.GetAll(ctx)
							So(err, ShouldBeNil)
							So(len(ac), ShouldEqual, 1)
							db.Table("services").Count(&count)
							So(count, ShouldEqual, 1)
						})
					})
					Reset(func() {
						db.Migrator().DropTable(&service.Service{})
					})
				})
			})
		})
		Reset(func() {
			db.Migrator().DropTable(&service_group.ServiceGroup{})
		})
	})
	DropDB(db, c)

}
