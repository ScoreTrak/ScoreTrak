package orm

import (
	"ScoreTrak/pkg/config"
	"ScoreTrak/pkg/host"
	"ScoreTrak/pkg/host_group"
	"ScoreTrak/pkg/service"
	"ScoreTrak/pkg/team"
	. "ScoreTrak/test"
	"os"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHostSpec(t *testing.T) {
	var c *config.StaticConfig
	autoTest := os.Getenv("AUTO_TEST")
	if autoTest == "TRUE" {
		c = NewConfigClone(SetupConfig("../../../configs/test-config.yml"))
	} else {
		c = NewConfigClone(SetupConfig("dev-config.yml"))
	}
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
		Convey("When the Host table is empty", func() {
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

				Convey("Then Querying By wrong ID", func() {
					ss, err := hr.GetByID(h.ID + 1)
					So(err, ShouldNotBeNil)
					So(ss, ShouldBeNil)
				})

				Convey("Then Deleting a wrong entry", func() {
					err = hr.Delete(2)
					So(err, ShouldNotBeNil)
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
					newHost := host.Host{Enabled: &b}
					Convey("For the wrong entry should not update anything", func() {
						newHost.ID = 1
						err = hr.Update(&newHost)
						So(err, ShouldBeNil)
						ac, err := hr.GetAll()
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
						So(*(ac[0].Enabled), ShouldBeFalse)
					})
					Convey("For the correct entry should update", func() {
						newHost.ID = 3
						err = hr.Update(&newHost)
						So(err, ShouldBeNil)
						ac, err := hr.GetAll()
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
						So(*(ac[0].Enabled), ShouldBeTrue)
					})
				})

				Convey("Then add a host group", func() {
					db.AutoMigrate(&host_group.HostGroup{})
					db.Model(&host.Host{}).AddForeignKey("host_group_id", "host_groups(id)", "RESTRICT", "RESTRICT")
					time.Sleep(500 * time.Second)
					Reset(func() {
						db.DropTableIfExists(&host.Host{})
						db.DropTableIfExists(&host_group.HostGroup{})
					})
					db.Exec("INSERT INTO host_groups (id, name, enabled) VALUES (1, 'HostGroup1', true)")
					db.Exec("INSERT INTO host_groups (id, name, enabled) VALUES (2, 'HostGroup2', false)")
					var count int
					db.Table("host_groups").Count(&count)
					So(count, ShouldEqual, 2)
					Convey("Adding a new host with host group foreign key", func() {
						address := "127.0.0.1"
						newHost := host.Host{ID: 4, HostGroupID: 2, Address: &address}
						err := hr.Store(&newHost)
						So(err, ShouldBeNil)
					})
					Convey("Updating a host with host group foreign key", func() {
						h.HostGroupID = 1
						err := hr.Update(&h)
						So(err, ShouldBeNil)
					})
					Convey("Updating a host with an invalid host group foreign key", func() {
						h.HostGroupID = 10
						err := hr.Update(&h)
						So(err, ShouldNotBeNil)
					})
				})
				Convey("Then add a team", func() {
					db.AutoMigrate(&team.Team{})
					db.Model(&host.Host{}).AddForeignKey("team_id", "teams(id)", "RESTRICT", "RESTRICT")
					Reset(func() {
						db.DropTableIfExists(&host.Host{})
						db.DropTableIfExists(&team.Team{})
					})
				})
				Convey("Then add a service", func() {
					db.AutoMigrate(&service.Service{})
					db.Model(&service.Service{}).AddForeignKey("host_id", "hosts(id)", "RESTRICT", "RESTRICT")
					Reset(func() {
						db.DropTableIfExists(&service.Service{})
					})
				})
			})
		})
	})
	DropDB(db, c)
	db.Close()
}
