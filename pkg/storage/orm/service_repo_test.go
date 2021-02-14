package orm

import (
	"context"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	. "github.com/ScoreTrak/ScoreTrak/pkg/config/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/host"
	"github.com/ScoreTrak/ScoreTrak/pkg/property"
	"github.com/ScoreTrak/ScoreTrak/pkg/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/service_group"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage"
	. "github.com/ScoreTrak/ScoreTrak/pkg/storage/orm/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/team"
	"github.com/gofrs/uuid"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestServiceSpec(t *testing.T) {
	var c config.StaticConfig
	autoTest := os.Getenv("AUTO_TEST")
	if autoTest == "TRUE" {
		c = NewConfigClone(SetupConfig("../../../configs/test-config.yml"))
	} else {
		c = NewConfigClone(SetupConfig("dev-config.yml"))
	}
	c.DB.Cockroach.Database = "scoretrak_test_orm_service"
	db := storage.SetupDB(c.DB)
	ctx := context.Background()
	t.Parallel() //t.Parallel should be placed after SetupDB because gorm has race conditions on Hook register
	Convey("Creating Service, Service Group, Host Tables along with their foreign keys", t, func() {
		db.AutoMigrate(&service.Service{})
		db.AutoMigrate(&service_group.ServiceGroup{})
		db.AutoMigrate(&team.Team{})
		db.AutoMigrate(&host.Host{})
		sr := NewServiceRepo(db)
		Convey("When all tables are empty", func() {
			Convey("Should output no entry", func() {
				gt, err := sr.GetAll(ctx)
				So(err, ShouldBeNil)
				So(len(gt), ShouldEqual, 0)
			})
			Convey("Creating a sample check_service should not be allowed", func() {
				s := []*service.Service{{Name: "FTP"}}
				err := sr.Store(ctx, s)
				So(err, ShouldNotBeNil)
				ac, err := sr.GetAll(ctx)
				So(err, ShouldBeNil)
				So(len(ac), ShouldEqual, 0)
			})
			Convey("Load Sample Service Group, And Host Data", func() {
				var count int64
				db.Exec("INSERT INTO teams (id, name, pause) VALUES ('11111111-1111-1111-1111-111111111111', 'TeamOne', true)")
				db.Exec("INSERT INTO teams (id, name, pause) VALUES ('22222222-2222-2222-2222-222222222222', 'TeamTwo', false)")
				db.Exec("INSERT INTO hosts (id, address, team_id) VALUES ('55555555-5555-5555-5555-555555555555', '192.168.1.2', '11111111-1111-1111-1111-111111111111')")
				db.Exec("INSERT INTO hosts (id, address, team_id) VALUES ('44444444-4444-4444-4444-444444444444', '192.168.1.1', '22222222-2222-2222-2222-222222222222')")
				db.Exec("INSERT INTO service_groups (id, name) VALUES ('77777777-7777-7777-7777-777777777777', 'FTPGroup')")
				db.Exec("INSERT INTO service_groups (id, name) VALUES ('22222222-2222-2222-2222-222222222222', 'FTPGroup2')")
				db.Table("hosts").Count(&count)
				So(count, ShouldEqual, 2)
				db.Table("service_groups").Count(&count)
				So(count, ShouldEqual, 2)

				Convey("Creating a sample check_service and associating with host id 5, and check_service group id 2", func() {
					s := []*service.Service{{Name: "FTP", ServiceGroupID: uuid.FromStringOrNil("22222222-2222-2222-2222-222222222222"), HostID: uuid.FromStringOrNil("55555555-5555-5555-5555-555555555555")}}
					err := sr.Store(ctx, s)
					Convey("Should be Allowed", func() {
						So(err, ShouldBeNil)
						ac, err := sr.GetAll(ctx)
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)

						Convey("Then Querying By ID", func() {
							ss, err := sr.GetByID(ctx, s[0].ID)
							So(err, ShouldBeNil)
							So(ss.Name, ShouldEqual, "FTP")
							So(ss.ServiceGroupID, ShouldEqual, uuid.FromStringOrNil("22222222-2222-2222-2222-222222222222"))
							So(ss.HostID, ShouldEqual, uuid.FromStringOrNil("55555555-5555-5555-5555-555555555555"))

						})

						Convey("Then Querying By wrong ID", func() {
							ss, err := sr.GetByID(ctx, uuid.FromStringOrNil("24422222-2222-2222-2222-222222222222"))
							So(err, ShouldNotBeNil)
							So(ss, ShouldBeNil)
						})

						Convey("Then Deleting a wrong entry", func() {
							err = sr.Delete(ctx, uuid.FromStringOrNil("24422222-2222-2222-2222-222222222222"))
							So(err, ShouldNotBeNil)
							Convey("Should output one entry", func() {
								ac, err := sr.GetAll(ctx)
								So(err, ShouldBeNil)
								So(len(ac), ShouldEqual, 1)
							})
						})

						Convey("Then updating to non existent host should not be allowed", func() {
							s[0].HostID = uuid.FromStringOrNil("22222233-2222-2222-2222-222222222222")
							err = sr.Update(ctx, s[0])
							So(err, ShouldNotBeNil)
						})
						Convey("Then updating to different existent host/service_group should be allowed", func() {
							s[0].HostID = uuid.FromStringOrNil("44444444-4444-4444-4444-444444444444")
							err = sr.Update(ctx, s[0])
							So(err, ShouldBeNil)
						})
						Convey("Then adding check_service with the same name should be allowed", func() {
							s2 := []*service.Service{{Name: "FTP", ServiceGroupID: uuid.FromStringOrNil("77777777-7777-7777-7777-777777777777"), HostID: uuid.FromStringOrNil("55555555-5555-5555-5555-555555555555")}}
							err = sr.Store(ctx, s2)
							So(err, ShouldBeNil)
						})
						Convey("Then updating regular fields should be allowed", func() {
							tru := true
							s[0].Pause = &tru
							s[0].Name = "SSH"
							s[0].RoundUnits = 3
							rd := uint64(2)
							s[0].RoundDelay = &rd
							x := uint64(5)
							s[0].Weight = &x
							err = sr.Update(ctx, s[0])
							So(err, ShouldBeNil)
							ac, err = sr.GetAll(ctx)
							So(err, ShouldBeNil)
							So(len(ac), ShouldEqual, 1)
							So(ac[0].Name, ShouldEqual, "SSH")
							So(*(ac[0].RoundDelay), ShouldEqual, 2)
						})
						Convey("Then updating Round Delay to something larger than Round Units should not be allowed", func() {
							rd := uint64(5)
							s[0].RoundDelay = &rd
							Convey("Round Units set", func() {
								s[0].RoundUnits = 3
								err = sr.Update(ctx, s[0])
								So(err, ShouldNotBeNil)
								ac, err = sr.GetAll(ctx)
								So(err, ShouldBeNil)
								So(len(ac), ShouldEqual, 1)
								So(*(ac[0].RoundDelay), ShouldEqual, 0)
							})
							Convey("Round Units not set", func() {
								err = sr.Update(ctx, s[0])
								So(err, ShouldNotBeNil)
								ac, err = sr.GetAll(ctx)
								So(err, ShouldBeNil)
								So(len(ac), ShouldEqual, 1)
								So(*(ac[0].RoundDelay), ShouldEqual, 0)
							})

						})
					})
				})
				Convey("Create Multiple services", func() {
					s1 := []*service.Service{{Name: "FTP", ServiceGroupID: uuid.FromStringOrNil("22222222-2222-2222-2222-222222222222"), HostID: uuid.FromStringOrNil("44444444-4444-4444-4444-444444444444")}}
					s2 := []*service.Service{{Name: "FTP", ServiceGroupID: uuid.FromStringOrNil("77777777-7777-7777-7777-777777777777"), HostID: uuid.FromStringOrNil("55555555-5555-5555-5555-555555555555")}}
					err := sr.Store(ctx, s1)
					So(err, ShouldBeNil)
					err = sr.Store(ctx, s2)
					So(err, ShouldBeNil)
					Convey("Loading Checks And Properties Tables with sample data", func() {
						db.AutoMigrate(&property.Property{})
						db.AutoMigrate(&check.Check{})
						db.Exec(fmt.Sprintf("INSERT INTO checks (service_id, round_id, passed) VALUES ('%s', 999, false)", s1[0].ID.String()))
						db.Exec(fmt.Sprintf("INSERT INTO checks (service_id, round_id, passed) VALUES ('%s', 333, false)", s2[0].ID.String()))
						db.Exec(fmt.Sprintf("INSERT INTO properties (service_id, key, value) VALUES ('%s', 'sample_key', 'sample_value')", s1[0].ID.String()))
						db.Exec(fmt.Sprintf("INSERT INTO properties (service_id, key, value) VALUES ('%s', 'sample_key', 'sample_value')", s2[0].ID.String()))
						db.Table("checks").Count(&count)
						So(count, ShouldEqual, 2)
						db.Table("properties").Count(&count)
						So(count, ShouldEqual, 2)
						Convey("Deleting the check_service", func() {
							err = sr.Delete(ctx, s1[0].ID)
							So(err, ShouldBeNil)
							Convey("Should also delete checks and properties associated with the deleted check_service", func() {
								var count int64
								db.Table("properties").Count(&count)
								So(count, ShouldEqual, 1)
								db.Table("checks").Count(&count)
								So(count, ShouldEqual, 1)
							})
						})
						Reset(func() {
							db.Migrator().DropTable(&property.Property{})
							db.Migrator().DropTable(&check.Check{})
						})
					})
				})

				Convey("Creating a sample check_service with wrong service_group should not be allowed", func() {
					s := []*service.Service{{Name: "FTP", ServiceGroupID: uuid.FromStringOrNil("77777777-7777-7777-7777-777777777557"), HostID: uuid.FromStringOrNil("55555555-5555-5555-5555-555555555555")}}
					err := sr.Store(ctx, s)
					So(err, ShouldNotBeNil)
					ac, err := sr.GetAll(ctx)
					So(err, ShouldBeNil)
					So(len(ac), ShouldEqual, 0)
				})

			})
		})
		Reset(func() {
			db.Migrator().DropTable(&service.Service{})
			db.Migrator().DropTable(&host.Host{})
			db.Migrator().DropTable(&service_group.ServiceGroup{})
		})
	})
	DropDB(db, c)

}
