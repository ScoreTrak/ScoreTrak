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

var db *gorm.DB

func GetGlobalDB() *gorm.DB {
	return db
}

// LoadDB serves as a singleton that initializes the value of db per package
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

var ErrDBNotSupported = errors.New("not supported db")

// NewDB creates an instance of database based on config
func NewDB(c Config) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	switch c.Use {
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

func newSqlite(c Config) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(c.Sqlite.Path), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		TablePrefix: c.Prefix,
	}})
	if err != nil {
		return nil, err
	}

	return db, nil
}

var ErrIncompleteCertInformationProvided = errors.New("you provided some, but not all certificate information")

// newCockroach is internal method used for initializing cockroach db instance.
// It modifies few cockroachdb options like kv.range.backpressure_range_size_multiplier and gc.ttlseconds that
// allows for a single large value to be changed frequently
func newCockroach(config Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s",
		config.Cockroach.Host,
		config.Cockroach.Port,
		config.Cockroach.UserName,
		config.Cockroach.Database)

	if config.Cockroach.Password != "" {
		psqlInfo += " password=" + config.Cockroach.Password
	}
	switch {
	case config.Cockroach.ClientCA != "" && config.Cockroach.ClientSSLKey != "" && config.Cockroach.ClientSSLCert != "": // mTLS
		psqlInfo += fmt.Sprintf(" ssl=true sslmode=verify-full sslrootcert=%s sslkey=%s sslcert=%s",
			config.Cockroach.ClientCA, config.Cockroach.ClientSSLKey, config.Cockroach.ClientSSLCert)
	case config.Cockroach.ClientCA != "" && config.Cockroach.ClientSSLKey == "" && config.Cockroach.ClientSSLCert == "": // OneWayTLS
		psqlInfo += fmt.Sprintf(" ssl=true sslmode=verify-full sslrootcert=%s", config.Cockroach.ClientCA)
	case config.Cockroach.ClientCA != "" || config.Cockroach.ClientSSLKey != "" || config.Cockroach.ClientSSLCert != "":
		return nil, fmt.Errorf("%w, CA: %s, Key: %s, Cert: %s",
			ErrIncompleteCertInformationProvided, config.Cockroach.ClientCA, config.Cockroach.ClientSSLKey, config.Cockroach.ClientSSLCert)
	default:
		psqlInfo += " sslmode=disable"
	}
	DB, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		TablePrefix: config.Prefix,
	}})
	if err != nil {
		return nil, err
	}
	if config.Cockroach.ConfigureZones {
		err := DB.Exec("ALTER RANGE default CONFIGURE ZONE USING gc.ttlseconds = ?;", config.Cockroach.DefaultZoneConfig.GcTtlseconds).Error
		if err != nil {
			return nil, err
		}
		err = DB.Exec("SET CLUSTER SETTING kv.range.backpressure_range_size_multiplier= ?;", config.Cockroach.DefaultZoneConfig.BackpressureRangeSizeMultiplier).Error
		if err != nil {
			return nil, err
		}
	} else {
		log.Println("You have chosen not to allow master configure database zones. Make sure you set gc.ttlseconds to something below 1200, so that report generation is not affected")
	}
	return DB, nil
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
	Sqlite struct {
		Path     string `default:"scoretrak.db"`
		Database string `default:"scoretrak"`
	}
}
