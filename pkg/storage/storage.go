package storage

import (
	"errors"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var ErrDBNotSupported = errors.New("not supported db")

// NewDB creates an instance of database based on config
func NewDB(c Config) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	switch c.Use {
	case "postgresql":
		db, err = newPostgreSQL(c)
	case "cockroach":
		db, err = newCockroach(c)
	case "sqlite":
		db, err = newSqlite(c)
	default:
		return nil, ErrDBNotSupported
	}
	if err != nil {
		return nil, err
	}
	return db, nil
}

var ErrIncompleteCertInformationProvided = errors.New("you provided some, but not all certificate information")

func newPostgreSQL(config Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s",
		config.Host,
		config.Port,
		config.UserName,
		config.Database)

	if config.Password != "" {
		dsn += " password=" + config.Password
	}
	switch {
	case config.ClientCA != "" && config.ClientSSLKey != "" && config.ClientSSLCert != "": // mTLS
		dsn += fmt.Sprintf(" ssl=true sslmode=verify-full sslrootcert=%s sslkey=%s sslcert=%s",
			config.ClientCA, config.ClientSSLKey, config.ClientSSLCert)
	case config.ClientCA != "" && config.ClientSSLKey == "" && config.ClientSSLCert == "": // OneWayTLS
		dsn += fmt.Sprintf(" ssl=true sslmode=verify-full sslrootcert=%s", config.ClientCA)
	case config.ClientCA != "" || config.ClientSSLKey != "" || config.ClientSSLCert != "":
		return nil, fmt.Errorf("%w, CA: %s, Key: %s, Cert: %s",
			ErrIncompleteCertInformationProvided, config.ClientCA, config.ClientSSLKey, config.ClientSSLCert)
	default:
		dsn += " sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{TablePrefix: config.Prefix}})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// newCockroach is internal method used for initializing cockroach db instance.
// It modifies few cockroachdb options like kv.range.backpressure_range_size_multiplier and gc.ttlseconds that
// allows for a single large value to be changed frequently
func newCockroach(config Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s",
		config.Host,
		config.Port,
		config.UserName,
		config.Database)

	if config.Password != "" {
		dsn += " password=" + config.Password
	}
	switch {
	case config.ClientCA != "" && config.ClientSSLKey != "" && config.ClientSSLCert != "": // mTLS
		dsn += fmt.Sprintf(" ssl=true sslmode=verify-full sslrootcert=%s sslkey=%s sslcert=%s",
			config.ClientCA, config.ClientSSLKey, config.ClientSSLCert)
	case config.ClientCA != "" && config.ClientSSLKey == "" && config.ClientSSLCert == "": // OneWayTLS
		dsn += fmt.Sprintf(" ssl=true sslmode=verify-full sslrootcert=%s", config.ClientCA)
	case config.ClientCA != "" || config.ClientSSLKey != "" || config.ClientSSLCert != "":
		return nil, fmt.Errorf("%w, CA: %s, Key: %s, Cert: %s",
			ErrIncompleteCertInformationProvided, config.ClientCA, config.ClientSSLKey, config.ClientSSLCert)
	default:
		dsn += " sslmode=disable"
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		TablePrefix: config.Prefix,
	}})
	if err != nil {
		return nil, err
	}

	// Cockroachdb customization
	if config.Cockroach.ConfigureZones {
		err := db.Exec("ALTER RANGE default CONFIGURE ZONE USING gc.ttlseconds = ?;", config.Cockroach.DefaultZoneConfig.GcTtlseconds).Error
		if err != nil {
			return nil, err
		}
		err = db.Exec("SET CLUSTER SETTING kv.range.backpressure_range_size_multiplier= ?;", config.Cockroach.DefaultZoneConfig.BackpressureRangeSizeMultiplier).Error
		if err != nil {
			return nil, err
		}
	} else {
		log.Println("You have chosen not to allow master configure database zones. Make sure you set gc.ttlseconds to something below 1200, so that report generation is not affected")
	}

	return db, nil
}

func newSqlite(c Config) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(c.Database), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		TablePrefix: c.Prefix,
	}})
	if err != nil {
		return nil, err
	}

	// SQlite customization
	// Enabled foreign key support
	if res := db.Exec("PRAGMA foreign_keys = ON", nil); res.Error != nil {
		panic(err)
	}
	return db, nil
}

type Config struct {
	Use           string `default:"cockroach"`
	Prefix        string `default:""`
	Host          string `default:"localhost"`
	Port          string `default:"26257"`
	UserName      string `default:"root"`
	Password      string `default:""`
	ClientCA      string `default:""`
	ClientSSLKey  string `default:""`
	ClientSSLCert string `default:""`
	Database      string `default:"scoretrak"`
	Migrate       bool   `default:"false"`
	Seed          bool   `default:"false"`
	Cockroach     struct {
		ConfigureZones    bool `default:"true"`
		DefaultZoneConfig struct {
			GcTtlseconds                    uint64 `default:"600"`
			BackpressureRangeSizeMultiplier uint64 `default:"0"`
		}
	}
}
