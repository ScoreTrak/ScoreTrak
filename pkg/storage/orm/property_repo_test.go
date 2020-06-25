package orm

import (
	"github.com/L1ghtman2k/ScoreTrak/pkg/config"
	"github.com/L1ghtman2k/ScoreTrak/pkg/property"
	"github.com/L1ghtman2k/ScoreTrak/pkg/service"
	. "github.com/L1ghtman2k/ScoreTrak/test"
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
	db := SetupDB(c)
	l := SetupLogger(c)
	t.Parallel() //t.Parallel should be placed after SetupDB because gorm has race conditions on Hook register
	Convey("Creating Property and Property tables along with their foreign keys", t, func() {
		db.AutoMigrate(&service.Service{})
		db.AutoMigrate(&property.Property{})
		db.Model(&property.Property{}).AddForeignKey("service_id", "services(id)", "CASCADE", "RESTRICT")
		cr := NewPropertyRepo(db, l)
		Convey("When all tables are empty", func() {
			Convey("Should output no entry", func() {
				gt, err := cr.GetAll()
				So(err, ShouldBeNil)
				So(len(gt), ShouldEqual, 0)
			})
			Convey("Creating a sample property should not be allowed", func() {
				c := property.Property{}
				err := cr.Store(&c)
				So(err, ShouldNotBeNil)
				ac, err := cr.GetAll()
				So(err, ShouldBeNil)
				So(len(ac), ShouldEqual, 0)
			})
			Convey("Load sample services and rounds", func() {
				var count int
				db.Exec("INSERT INTO services (id, service_group_id, host_id, name) VALUES (5, 999, 999, 'TestService')")
				db.Exec("INSERT INTO services (id, service_group_id, host_id, name) VALUES (6, 999, 999, 'TestService')")
				db.Table("services").Count(&count)
				So(count, ShouldEqual, 2)

				Convey("Creating a sample property and associating with service 5 and round 1", func() {
					c := property.Property{Key: "TestKey", ServiceID: 5, Value: "TestValue", Status: "Edit"}
					err := cr.Store(&c)
					Convey("Should be Allowed", func() {
						So(err, ShouldBeNil)
						ac, err := cr.GetAll()
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
						Convey("Then Querying By ID", func() {
							ss, err := cr.GetByID(c.ID)
							So(err, ShouldBeNil)
							So(ss.Key, ShouldEqual, "TestKey")
							So(ss.ServiceID, ShouldEqual, 5)
							So(ss.Value, ShouldEqual, "TestValue")
						})

						Convey("Then Querying By wrong ID", func() {
							ss, err := cr.GetByID(c.ID + 1)
							So(err, ShouldNotBeNil)
							So(ss, ShouldBeNil)
						})

						Convey("Then Deleting a wrong entry", func() {
							err = cr.Delete(c.ID + 1)
							So(err, ShouldNotBeNil)
							Convey("Should output one entry", func() {
								ac, err := cr.GetAll()
								So(err, ShouldBeNil)
								So(len(ac), ShouldEqual, 1)
							})
						})

						Convey("Then Deleting the property should be allowed", func() {
							err = cr.Delete(c.ID)
							So(err, ShouldBeNil)
							ac, err = cr.GetAll()
							So(err, ShouldBeNil)
							So(len(ac), ShouldEqual, 0)
						})

						Convey("Then Updating the property Description, Status", func() {
							c.Status = "Hide"
							c.Description = "Test Description"
							c.Value = ""
							err = cr.Update(&c)
							So(err, ShouldBeNil)
							ac, err = cr.GetAll()
							So(err, ShouldBeNil)
							So(len(ac), ShouldEqual, 1)
							So(ac[0].Status, ShouldEqual, "Hide")
							So(ac[0].Description, ShouldEqual, "Test Description")
						})

						Convey("Then Updating the property Status to an invalid value", func() {
							c.Status = "SomeBadStatus"
							c.Description = "Test Description"
							err = cr.Update(&c)
							So(err, ShouldNotBeNil)
							ac, err = cr.GetAll()
							So(err, ShouldBeNil)
							So(len(ac), ShouldEqual, 1)
							So(ac[0].Status, ShouldEqual, "Edit")
							So(ac[0].Description, ShouldEqual, "")
						})

						Convey("Then Updating the property Value", func() {
							c.Value = "AnotherValue"
							err = cr.Update(&c)
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
					s := property.Property{Key: "TestKey", ServiceID: 999, Value: "TestValue"}
					err := cr.Store(&s)
					So(err, ShouldNotBeNil)
					ac, err := cr.GetAll()
					So(err, ShouldBeNil)
					So(len(ac), ShouldEqual, 0)
				})
			})
		})
		Reset(func() {
			db.DropTableIfExists(&property.Property{})
			db.DropTableIfExists(&service.Service{})
		})
	})
	DropDB(db, c)
	db.Close()
}
