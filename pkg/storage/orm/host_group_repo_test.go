package orm

import (
	"ScoreTrak/pkg/config"
	"ScoreTrak/pkg/host"
	"ScoreTrak/pkg/host_group"
	. "ScoreTrak/test"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHostGroupSpec(t *testing.T) {
	var c *config.StaticConfig
	autoTest := os.Getenv("AUTO_TEST")
	if autoTest == "TRUE" {
		c = NewConfigClone(SetupConfig("../../../configs/test-config.yml"))
	} else {
		c = NewConfigClone(SetupConfig("dev-config.yml"))
	}
	c.DB.Cockroach.Database = "scoretrak_test_orm_host_group"
	c.Logger.FileName = "host_group_test_repo.log"
	db := SetupDB(c)
	l := SetupLogger(c)
	t.Parallel() //t.Parallel should be placed after SetupDB because gorm has race conditions on Hook register
	Convey("Creating Host Group Table", t, func() {
		db.AutoMigrate(&host_group.HostGroup{})
		hg := NewHostGroupRepo(db, l)
		Reset(func() {
			db.DropTableIfExists(&host_group.HostGroup{})
		})
		Convey("When the Teams table is empty", func() {
			Convey("There should be no entries", func() {
				ac, err := hg.GetAll()
				So(err, ShouldBeNil)
				So(len(ac), ShouldEqual, 0)
			})

			Convey("Adding an valid entry", func() {
				var err error
				h := host_group.HostGroup{ID: 3, Name: "host group"}
				err = hg.Store(&h)
				So(err, ShouldBeNil)
				Convey("Then making sure the entry exists", func() {
					ac, err := hg.GetAll()
					So(err, ShouldBeNil)
					So(len(ac), ShouldEqual, 1)
					So(ac[0].ID, ShouldEqual, 3)
					So(ac[0].Name, ShouldEqual, "host group")
				})

				Convey("Then getting entry by id", func() {
					ac, err := hg.GetByID(3)
					So(err, ShouldBeNil)
					So(ac.ID, ShouldEqual, 3)
					So(ac.Name, ShouldEqual, "host group")
				})

				Convey("Then Querying By wrong ID", func() {
					ss, err := hg.GetByID(h.ID + 1)
					So(err, ShouldNotBeNil)
					So(ss, ShouldBeNil)
				})

				Convey("Then Deleting a wrong entry", func() {
					err = hg.Delete(2)
					So(err, ShouldNotBeNil)
					Convey("Should output one entry", func() {
						ac, err := hg.GetAll()
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
					})
				})

				Convey("Then Deleting the added entry", func() {
					err = hg.Delete(3)
					So(err, ShouldBeNil)
					Convey("Should output no entries", func() {
						ac, err := hg.GetAll()
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 0)
					})
				})

				Convey("Then Updating Enabled to true", func() {
					tru := true
					newHostGroup := &host_group.HostGroup{Enabled: &tru}
					Convey("For the wrong entry should not update anything", func() {
						newHostGroup.ID = 1
						err = hg.Update(newHostGroup)
						So(err, ShouldBeNil)
						ac, err := hg.GetAll()
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
						So(*(ac[0].Enabled), ShouldBeFalse)
					})
					Convey("For the correct entry should update", func() {
						newHostGroup.ID = 3
						err = hg.Update(newHostGroup)
						So(err, ShouldBeNil)
						ac, err := hg.GetAll()
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
						So(*(ac[0].Enabled), ShouldBeTrue)
					})
				})

				Convey("Creating Hosts Table", func() {
					var count int
					db.AutoMigrate(&host.Host{})
					db.Model(&host.Host{}).AddForeignKey("host_group_id", "host_groups(id)", "RESTRICT", "RESTRICT")
					Convey("Associating a single Host with a host group", func() {
						db.Exec("INSERT INTO hosts (id, address, team_id, host_group_id) VALUES (4, '192.168.1.1', '', 3)")
						db.Table("hosts").Count(&count)
						So(count, ShouldEqual, 1)
						Convey("Delete a host group without deleting a host", func() {
							err = hg.Delete(3)
							So(err, ShouldNotBeNil)
							ac, err := hg.GetAll()
							So(err, ShouldBeNil)
							So(len(ac), ShouldEqual, 1)
						})
						Convey("Deleting a host then deleting a host group", func() {
							db.Exec("DELETE FROM hosts WHERE id=4")
							err = hg.Delete(3)
							So(err, ShouldBeNil)
							ac, err := hg.GetAll()
							So(err, ShouldBeNil)
							So(len(ac), ShouldEqual, 0)
						})

						Convey("Updating a team enabled without deleting a host should not yield error", func() {
							tru := true
							h.Enabled = &tru
							err = hg.Update(&h)
							So(err, ShouldBeNil)
						})

						Reset(func() {
							db.DropTableIfExists(&host.Host{})
						})
					})
				})
			})
		})
	})
	DropDB(db, c)
	db.Close()
}
