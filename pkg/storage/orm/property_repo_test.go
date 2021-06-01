package orm

import (
	"context"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	. "github.com/ScoreTrak/ScoreTrak/pkg/config/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/property"
	"github.com/ScoreTrak/ScoreTrak/pkg/service"
	. "github.com/ScoreTrak/ScoreTrak/pkg/storage/orm/testutil"
	"github.com/gofrs/uuid"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestPropertySpec(t *testing.T) {
	var c config.StaticConfig
	autoTest := os.Getenv("AUTO_TEST")
	if autoTest == "TRUE" {
		c = NewConfigClone(SetupConfig("../../../configs/test-config.yml"))
	} else {
		c = NewConfigClone(SetupConfig("dev-config.yml"))
	}
	c.DB.Cockroach.Database = "scoretrak_test_orm_property"
	db := SetupCockroachDB(c.DB)
	ctx := context.Background()
	t.Parallel() //t.Parallel should be placed after SetupCockroachDB because gorm has race conditions on Hook register
	Convey("Creating Property and Property tables along with their foreign keys", t, func() {
		db.AutoMigrate(&service.Service{})
		db.AutoMigrate(&property.Property{})
		cr := NewPropertyRepo(db)
		Convey("When all tables are empty", func() {
			Convey("Should output no entry", func() {
				gt, err := cr.GetAll(ctx)
				So(err, ShouldBeNil)
				So(len(gt), ShouldEqual, 0)
			})
			Convey("Creating a sample property should not be allowed", func() {
				c := []*property.Property{{}}
				err := cr.Store(ctx, c)
				So(err, ShouldNotBeNil)
				ac, err := cr.GetAll(ctx)
				So(err, ShouldBeNil)
				So(len(ac), ShouldEqual, 0)
			})
			Convey("Load sample services and rounds", func() {
				var count int64
				db.Exec("INSERT INTO services (id, service_group_id, host_id, name) VALUES ('55555555-5555-5555-5555-555555555555', '99999999-9999-9999-9999-999999999999', '99999999-9999-9999-9999-999999999999', 'TestService')")
				db.Exec("INSERT INTO services (id, service_group_id, host_id, name) VALUES ('66666666-6666-6666-6666-666666666666', '99999999-9999-9999-9999-999999999999', '99999999-9999-9999-9999-999999999999', 'TestService')")
				db.Table("services").Count(&count)
				So(count, ShouldEqual, 2)

				Convey("Creating a sample property and associating with check_service 5 and round 1", func() {
					str := "TestValue"
					c := []*property.Property{{Key: "TestKey", ServiceID: uuid.FromStringOrNil("55555555-5555-5555-5555-555555555555"), Value: &str, Status: "Edit"}}
					err := cr.Store(ctx, c)
					Convey("Should be Allowed", func() {
						So(err, ShouldBeNil)
						ac, err := cr.GetAll(ctx)
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
						Convey("Then Querying By ID", func() {
							ss, err := cr.GetByServiceIDKey(ctx, uuid.FromStringOrNil("55555555-5555-5555-5555-555555555555"), "TestKey")
							So(err, ShouldBeNil)
							So(ss.Key, ShouldEqual, "TestKey")
							So(ss.ServiceID, ShouldEqual, uuid.FromStringOrNil("55555555-5555-5555-5555-555555555555"))
							So(*ss.Value, ShouldEqual, "TestValue")
						})

						Convey("Then Querying By wrong ID", func() {
							ss, err := cr.GetByServiceIDKey(ctx, uuid.FromStringOrNil("55555555-5555-5555-5555-555555555522"), "TestKey")
							So(err, ShouldNotBeNil)
							So(ss, ShouldBeNil)
						})

						Convey("Then Deleting a wrong entry", func() {
							err = cr.Delete(ctx, uuid.FromStringOrNil("55555555-5555-5555-5555-555555555522"), "TestKey")
							So(err, ShouldNotBeNil)
							Convey("Should output one entry", func() {
								ac, err := cr.GetAll(ctx)
								So(err, ShouldBeNil)
								So(len(ac), ShouldEqual, 1)
							})
						})

						Convey("Then Deleting the property should be allowed", func() {
							err = cr.Delete(ctx, uuid.FromStringOrNil("55555555-5555-5555-5555-555555555555"), "TestKey")
							So(err, ShouldBeNil)
							ac, err = cr.GetAll(ctx)
							So(err, ShouldBeNil)
							So(len(ac), ShouldEqual, 0)
						})

						Convey("Then Updating the property Description, Status", func() {
							c[0].Status = "Hide"
							str := ""
							c[0].Value = &str
							err = cr.Update(ctx, c[0])
							So(err, ShouldBeNil)
							ac, err = cr.GetAll(ctx)
							So(err, ShouldBeNil)
							So(len(ac), ShouldEqual, 1)
							So(ac[0].Status, ShouldEqual, "Hide")
						})

						Convey("Then Updating the property Status to an invalid value", func() {
							c[0].Status = "SomeBadStatus"
							err = cr.Update(ctx, c[0])
							So(err, ShouldNotBeNil)
							ac, err = cr.GetAll(ctx)
							So(err, ShouldBeNil)
							So(len(ac), ShouldEqual, 1)
							So(ac[0].Status, ShouldEqual, "Edit")
						})

						Convey("Then Updating the property Value", func() {
							str := "AnotherValue"
							c[0].Value = &str
							err = cr.Update(ctx, c[0])
							So(err, ShouldBeNil)
							ac, err = cr.GetAll(ctx)
							So(err, ShouldBeNil)
							So(len(ac), ShouldEqual, 1)
							So(*ac[0].Value, ShouldEqual, "AnotherValue")
							So(ac[0].Status, ShouldEqual, "Edit")
						})

					})
				})
				Convey("Creating a property with wrong check_service should not be allowed", func() {
					str := "TestValue"
					s := []*property.Property{{Key: "TestKey", ServiceID: uuid.FromStringOrNil("55521555-5555-5555-5555-555555555555"), Value: &str}}
					err := cr.Store(ctx, s)
					So(err, ShouldNotBeNil)
					ac, err := cr.GetAll(ctx)
					So(err, ShouldBeNil)
					So(len(ac), ShouldEqual, 0)
				})
			})
		})
		Reset(func() {
			db.Migrator().DropTable(&property.Property{})
			db.Migrator().DropTable(&service.Service{})
		})
	})
	DropDB(db, c)

}
