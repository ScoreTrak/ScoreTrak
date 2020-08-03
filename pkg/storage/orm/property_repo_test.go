package orm

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	. "github.com/ScoreTrak/ScoreTrak/pkg/config/util"
	. "github.com/ScoreTrak/ScoreTrak/pkg/logger/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/property"
	"github.com/ScoreTrak/ScoreTrak/pkg/service"
	. "github.com/ScoreTrak/ScoreTrak/pkg/storage/orm/util"

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
	c.Logger.FileName = "property_test.log"
	db := SetupDB(c.DB)
	l := SetupLogger(c.Logger)
	t.Parallel() //t.Parallel should be placed after SetupDB because gorm has race conditions on Hook register
	Convey("Creating Property and Property tables along with their foreign keys", t, func() {
		db.AutoMigrate(&service.Service{})
		db.AutoMigrate(&property.Property{})
		cr := NewPropertyRepo(db, l)
		Convey("When all tables are empty", func() {
			Convey("Should output no entry", func() {
				gt, err := cr.GetAll()
				So(err, ShouldBeNil)
				So(len(gt), ShouldEqual, 0)
			})
			Convey("Creating a sample property should not be allowed", func() {
				c := []*property.Property{{}}
				err := cr.Store(c)
				So(err, ShouldNotBeNil)
				ac, err := cr.GetAll()
				So(err, ShouldBeNil)
				So(len(ac), ShouldEqual, 0)
			})
			Convey("Load sample services and rounds", func() {
				var count int64
				db.Exec("INSERT INTO services (id, service_group_id, host_id, name) VALUES ('55555555-5555-5555-5555-555555555555', '99999999-9999-9999-9999-999999999999', '99999999-9999-9999-9999-999999999999', 'TestService')")
				db.Exec("INSERT INTO services (id, service_group_id, host_id, name) VALUES ('66666666-6666-6666-6666-666666666666', '99999999-9999-9999-9999-999999999999', '99999999-9999-9999-9999-999999999999', 'TestService')")
				db.Table("services").Count(&count)
				So(count, ShouldEqual, 2)

				Convey("Creating a sample property and associating with service 5 and round 1", func() {
					c := []*property.Property{{Key: "TestKey", ServiceID: uuid.FromStringOrNil("55555555-5555-5555-5555-555555555555"), Value: "TestValue", Status: "Edit"}}
					err := cr.Store(c)
					Convey("Should be Allowed", func() {
						So(err, ShouldBeNil)
						ac, err := cr.GetAll()
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
						Convey("Then Querying By ID", func() {
							ss, err := cr.GetByID(c[0].ID)
							So(err, ShouldBeNil)
							So(ss.Key, ShouldEqual, "TestKey")
							So(ss.ServiceID, ShouldEqual, uuid.FromStringOrNil("55555555-5555-5555-5555-555555555555"))
							So(ss.Value, ShouldEqual, "TestValue")
						})

						Convey("Then Querying By wrong ID", func() {
							ss, err := cr.GetByID(uuid.FromStringOrNil("55555555-5555-5555-5555-555555555522"))
							So(err, ShouldNotBeNil)
							So(ss, ShouldBeNil)
						})

						Convey("Then Deleting a wrong entry", func() {
							err = cr.Delete(uuid.FromStringOrNil("55555555-5555-5555-5555-555555555522"))
							So(err, ShouldNotBeNil)
							Convey("Should output one entry", func() {
								ac, err := cr.GetAll()
								So(err, ShouldBeNil)
								So(len(ac), ShouldEqual, 1)
							})
						})

						Convey("Then Deleting the property should be allowed", func() {
							err = cr.Delete(c[0].ID)
							So(err, ShouldBeNil)
							ac, err = cr.GetAll()
							So(err, ShouldBeNil)
							So(len(ac), ShouldEqual, 0)
						})

						Convey("Then Updating the property Description, Status", func() {
							c[0].Status = "Hide"
							c[0].Description = "Test Description"
							c[0].Value = ""
							err = cr.Update(c[0])
							So(err, ShouldBeNil)
							ac, err = cr.GetAll()
							So(err, ShouldBeNil)
							So(len(ac), ShouldEqual, 1)
							So(ac[0].Status, ShouldEqual, "Hide")
							So(ac[0].Description, ShouldEqual, "Test Description")
						})

						SkipConvey("Then Updating the property Status to an invalid value", func() { //TODO: Change this to Convey once govalidations are enabled
							c[0].Status = "SomeBadStatus"
							c[0].Description = "Test Description"
							err = cr.Update(c[0])
							So(err, ShouldNotBeNil)
							ac, err = cr.GetAll()
							So(err, ShouldBeNil)
							So(len(ac), ShouldEqual, 1)
							So(ac[0].Status, ShouldEqual, "Edit")
							So(ac[0].Description, ShouldEqual, "")
						})

						Convey("Then Updating the property Value", func() {
							c[0].Value = "AnotherValue"
							err = cr.Update(c[0])
							So(err, ShouldBeNil)
							ac, err = cr.GetAll()
							So(err, ShouldBeNil)
							So(len(ac), ShouldEqual, 1)
							So(ac[0].Value, ShouldEqual, "AnotherValue")
							So(ac[0].Status, ShouldEqual, "Edit")
							So(ac[0].Description, ShouldEqual, "")
						})

					})
				})
				Convey("Creating a property with wrong service should not be allowed", func() {
					s := []*property.Property{{Key: "TestKey", ServiceID: uuid.FromStringOrNil("55521555-5555-5555-5555-555555555555"), Value: "TestValue"}}
					err := cr.Store(s)
					So(err, ShouldNotBeNil)
					ac, err := cr.GetAll()
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
