package orm

import (
	"context"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	. "github.com/ScoreTrak/ScoreTrak/pkg/config/util"
	. "github.com/ScoreTrak/ScoreTrak/pkg/storage/orm/testutil"

	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestConfigSpec(t *testing.T) {
	c, _ := LoadViperConfig("../../../configs/test-config.yml")
	db := SetupDB(c.DB)
	defer TruncateAllTables(db)
	ctx := context.Background()
	Convey("Seed db and create config repo", t, func() {
		db.Exec("INSERT INTO dynamic_configs (id, round_duration, enabled) VALUES (1, 60, true)")
		db.Exec("INSERT INTO report (id, cache, updated_at) VALUES (1, '{}', ?)", time.Now())
		var count int64
		db.Table("dynamic_configs").Count(&count)
		So(count, ShouldEqual, 1)

		cr := NewConfigRepo(db)

		Convey("Retrieving all config properties", func() {
			dn, err := cr.Get(context.Background())
			So(err, ShouldBeNil)
			So(*(dn.Enabled), ShouldBeTrue)
			So(dn.RoundDuration, ShouldEqual, 60)
		})

		Convey("Updating the config properties should not return errors", func() {
			fls := false
			dn := config.DynamicConfig{RoundDuration: 25, Enabled: &fls}
			err := cr.Update(context.Background(), &dn)
			So(err, ShouldBeNil)
			dnr, err := cr.Get(context.Background())
			So(err, ShouldBeNil)
			So(*(dnr.Enabled), ShouldBeFalse)
			So(dnr.RoundDuration, ShouldEqual, 25)
		})

		Reset(func() {
			TruncateTable(ctx, &config.DynamicConfig{}, db)
		})
	})
}
