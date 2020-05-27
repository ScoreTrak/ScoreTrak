package orm

import (
	"ScoreTrak/pkg/config"
	. "ScoreTrak/test"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestConfigSpec(t *testing.T) {
	c := NewConfigClone(SetupConfig("dev-config.yml"))
	c.DB.Cockroach.Database = "scoretrak_test_config"
	c.Logger.FileName = "config_test_repo.log"
	db := SetupDB(c)
	l := SetupLogger(c)
	t.Parallel() //t.Parallel should be placed after SetupDB because gorm has race conditions on Hook register
	Convey("Creating Config Table and Insert sample config", t, func() {
		db.AutoMigrate(&config.DynamicConfig{})
		db.Exec("INSERT INTO config (id, round_duration, enabled) VALUES (1, 60, true)")
		var count int
		db.Table("config").Count(&count)
		So(count, ShouldEqual, 1)

		cr := NewConfigRepo(db, l)

		Convey("Retrieving all config properties", func() {
			dn, err := cr.Get()
			So(err, ShouldBeNil)
			So(*(dn.Enabled), ShouldBeTrue)
			So(dn.RoundDuration, ShouldEqual, 60)
		})

		Convey("Updating the config properties should not return errors", func() {
			fls := false
			dn := config.DynamicConfig{RoundDuration: 5, Enabled: &fls}
			err := cr.Update(&dn)

			dnr, err := cr.Get()
			So(err, ShouldBeNil)
			So(*(dnr.Enabled), ShouldBeFalse)
			So(dnr.RoundDuration, ShouldEqual, 5)
		})

		Reset(func() {
			db.DropTableIfExists(&config.DynamicConfig{})
		})
	})
	DropDB(db, c)
	db.Close()
}
