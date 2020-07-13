package storage

import (
	"errors"
	"fmt"

	"github.com/L1ghtman2k/ScoreTrak/pkg/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/qor/validations"
)

var db *gorm.DB

func GetGlobalDB() *gorm.DB {
	return db
}

func LoadDB(c config.DB) (*gorm.DB, error) {
	var err error
	if db == nil {
		db, err = NewDB(c)
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}

func NewDB(c config.DB) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	if c.Use == "cockroach" {
		db, err = newCockroach(c)
	} else {
		return nil, errors.New("not supported db")
	}
	if err != nil {
		return nil, err
	}
	validations.RegisterCallbacks(db)
	db.BlockGlobalUpdate(true)
	return db, nil
}

func newCockroach(c config.DB) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable",
		c.Cockroach.Host,
		c.Cockroach.Port,
		c.Cockroach.UserName,
		c.Cockroach.Database)

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	if c.Cockroach.ConfigureZones {
		err := db.Exec("ALTER RANGE default CONFIGURE ZONE USING gc.ttlseconds = $1;", c.Cockroach.DefaultZoneConfig.GcTtlseconds).Error
		if err != nil {
			return nil, err
		}
		err = db.Exec("SET CLUSTER SETTING kv.range.backpressure_range_size_multiplier= $1;", c.Cockroach.DefaultZoneConfig.BackpressureRangeSizeMultiplier).Error
		if err != nil {
			return nil, err
		}
	} else {
		fmt.Println("You have chosen not to allow master configure database zones. Make sure you set gc.ttlseconds to something below 1200, so that report generation is not affected")
	}
	return db, nil
}
