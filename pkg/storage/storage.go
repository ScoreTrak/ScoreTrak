package storage

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var db *gorm.DB

func GetGlobalDB() *gorm.DB {
	return db
}

//LoadDB serves as a singleton that initializes the value of db per package
func LoadDB(c Config) (*gorm.DB, error) {
	var err error
	if db == nil {
		db, err = NewDB(c)
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}

//NewDB creates an instance of database based on config
func NewDB(c Config) (*gorm.DB, error) {
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
	return db, nil
}

// newCockroach is internal method used for initializing cockroach db instance.
// It modifies few cockroachdb options like kv.range.backpressure_range_size_multiplier and gc.ttlseconds that
// allows for a single large value to be changed frequently
func newCockroach(c Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s",
		c.Cockroach.Host,
		c.Cockroach.Port,
		c.Cockroach.UserName,
		c.Cockroach.Database)

	if c.Cockroach.Password != "" {
		psqlInfo += " password=" + c.Cockroach.Password
	}
	if c.Cockroach.ClientCA != "" && c.Cockroach.ClientSSLKey != "" && c.Cockroach.ClientSSLCert != "" {
		psqlInfo += fmt.Sprintf(" ssl=true sslmode=verify-full sslrootcert=%s sslkey=%s sslcert=%s", c.Cockroach.ClientCA, c.Cockroach.ClientSSLKey, c.Cockroach.ClientSSLCert)
	} else if c.Cockroach.ClientCA != "" || c.Cockroach.ClientSSLKey != "" || c.Cockroach.ClientSSLCert != "" {
		return nil, fmt.Errorf("you provided some, but not all certificate information. CA: %s, Key: %s, Cert: %s. If you wish to not use certificates for database, make sure to all fields are empty", c.Cockroach.ClientCA, c.Cockroach.ClientSSLKey, c.Cockroach.ClientSSLCert)
	} else {
		psqlInfo += " sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		TablePrefix: c.Prefix,
	}})
	if err != nil {
		return nil, err
	}
	if c.Cockroach.ConfigureZones {
		err := db.Exec("ALTER RANGE default CONFIGURE ZONE USING gc.ttlseconds = ?;", c.Cockroach.DefaultZoneConfig.GcTtlseconds).Error
		if err != nil {
			return nil, err
		}
		err = db.Exec("SET CLUSTER SETTING kv.range.backpressure_range_size_multiplier= ?;", c.Cockroach.DefaultZoneConfig.BackpressureRangeSizeMultiplier).Error
		if err != nil {
			return nil, err
		}
	} else {
		log.Println("You have chosen not to allow master configure database zones. Make sure you set gc.ttlseconds to something below 1200, so that report generation is not affected")
	}
	return db, nil
}

type Config struct {
	Use       string `default:"cockroach"`
	Prefix    string `default:""`
	Cockroach struct {
		Host              string `default:"cockroach"`
		Port              string `default:"26257"`
		UserName          string `default:"root"`
		Password          string `default:""`
		ClientCA          string `default:""`
		ClientSSLKey      string `default:""`
		ClientSSLCert     string `default:""`
		Database          string `default:"scoretrak"`
		ConfigureZones    bool   `default:"true"`
		DefaultZoneConfig struct {
			GcTtlseconds                    uint64 `default:"600"`
			BackpressureRangeSizeMultiplier uint64 `default:"0"`
		}
	}
}
