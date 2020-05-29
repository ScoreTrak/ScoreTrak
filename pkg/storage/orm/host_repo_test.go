package orm

import (
	"ScoreTrak/pkg/host"
	. "ScoreTrak/test"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHostSpec(t *testing.T) {
	c := NewConfigClone(SetupConfig("dev-config.yml"))
	c.DB.Cockroach.Database = "scoretrak_test_orm_host"
	c.Logger.FileName = "host_test_repo.log"
	db := SetupDB(c)
	l := SetupLogger(c)
	t.Parallel() //t.Parallel should be placed after SetupDB because gorm has race conditions on Hook register
	Convey("Creating Host Table", t, func() {
		db.AutoMigrate(&host.Host{})
		hr := NewHostRepo(db, l)
		Reset(func() {
			db.DropTableIfExists(&host.Host{})
		})
		Convey("When the Teams table is empty", func() {
			Convey("There should be no entries", func() {
				ac, err := hr.GetAll()
				So(err, ShouldBeNil)
				So(len(ac), ShouldEqual, 0)
			})

			Convey("Adding an valid entry", func() {
				var err error
				b := false
				s := "127.0.0.1"
				h := host.Host{ID: 3, Address: &s, Enabled: &b}
				err = hr.Store(&h)
				So(err, ShouldBeNil)
				Convey("Then making sure the entry exists", func() {
					ac, err := hr.GetAll()
					So(err, ShouldBeNil)
					So(len(ac), ShouldEqual, 1)
					So(ac[0].ID, ShouldEqual, 3)
					So(*(ac[0].Address), ShouldEqual, "127.0.0.1")
					So(*(ac[0].Enabled), ShouldBeFalse)
				})

				Convey("Then getting entry by id", func() {
					ac, err := hr.GetByID(3)
					So(err, ShouldBeNil)
					So(ac.ID, ShouldEqual, 3)
					So(*(ac.Address), ShouldEqual, "127.0.0.1")
					So(*(ac.Enabled), ShouldBeFalse)
				})

				Convey("Then Deleting a wrong entry", func() {
					err = hr.Delete(2)
					So(err, ShouldBeNil)
					Convey("Should output one entry", func() {
						ac, err := hr.GetAll()
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
					})
				})

				Convey("Then Deleting the added entry", func() {
					err = hr.Delete(3)
					So(err, ShouldBeNil)
					Convey("Should output no entries", func() {
						ac, err := hr.GetAll()
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 0)
					})
				})

				Convey("Then Updating Enabled to true", func() {
					b := true
					newHostGroup := host.Host{Enabled: &b}
					Convey("For the wrong entry should not update anything", func() {
						newHostGroup.ID = 1
						err = hr.Update(&newHostGroup)
						So(err, ShouldBeNil)
						ac, err := hr.GetAll()
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
						So(*(ac[0].Enabled), ShouldBeFalse)
					})
					Convey("For the correct entry should update", func() {
						newHostGroup.ID = 3
						err = hr.Update(&newHostGroup)
						So(err, ShouldBeNil)
						ac, err := hr.GetAll()
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
						So(*(ac[0].Enabled), ShouldBeTrue)
					})
				})
			})
		})
	})
	DropDB(db, c)
	db.Close()
}
