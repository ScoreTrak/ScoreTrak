package orm

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	. "github.com/ScoreTrak/ScoreTrak/test"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
	"time"
)

func TestConfigSpec(t *testing.T) {
	var c config.StaticConfig
	autoTest := os.Getenv("AUTO_TEST")
	if autoTest == "TRUE" {
		c = NewConfigClone(SetupConfig("../../../configs/test-config.yml"))
	} else {
		c = NewConfigClone(SetupConfig("dev-config.yml"))
	}
	c.DB.Cockroach.Database = "scoretrak_test_orm_config"
	c.Logger.FileName = "config_test.log"
	db := SetupDB(c.DB)
	l := SetupLogger(c.Logger)
	t.Parallel() //t.Parallel should be placed after SetupDB because gorm has race conditions on Hook register
	Convey("Creating Config Table and Insert sample config", t, func() {
		db.AutoMigrate(&config.DynamicConfig{})
		db.AutoMigrate(&report.Report{})
		db.Exec("INSERT INTO config (id, round_duration, enabled) VALUES (1, 60, true)")
		db.Exec("INSERT INTO report (id, cache, updated_at) VALUES (1, '{}', ?)", time.Now())
		var count int64
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
			dn := config.DynamicConfig{RoundDuration: 25, Enabled: &fls}
			err := cr.Update(&dn)
			So(err, ShouldBeNil)
			dnr, err := cr.Get()
			So(err, ShouldBeNil)
			So(*(dnr.Enabled), ShouldBeFalse)
			So(dnr.RoundDuration, ShouldEqual, 25)
		})

		Reset(func() {
			db.Migrator().DropTable(&config.DynamicConfig{})
		})
	})
	DropDB(db, c)

}
