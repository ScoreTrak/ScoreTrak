package storage

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
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
	if c.Use == "cockroach" {
		db, err = newCockroach(c)
	} else {
		return nil, ErrDBNotSupported
	}
	if err != nil {
		return nil, err
	}
	return db, nil
}

var ErrIncompleteCertInformationProvided = errors.New("you provided some, but not all certificate information")

// newCockroach is internal method used for initializing cockroach db instance.
func newCockroach(c Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s",
		c.Cockroach.Host,
		c.Cockroach.Port,
		c.Cockroach.UserName,
		c.Cockroach.Database)

	if c.Cockroach.Password != "" {
		psqlInfo += " password=" + c.Cockroach.Password
	}
	switch {
	case c.Cockroach.ClientCA != "" && c.Cockroach.ClientSSLKey != "" && c.Cockroach.ClientSSLCert != "": // mTLS
		psqlInfo += fmt.Sprintf(" ssl=true sslmode=verify-full sslrootcert=%s sslkey=%s sslcert=%s",
			c.Cockroach.ClientCA, c.Cockroach.ClientSSLKey, c.Cockroach.ClientSSLCert)
	case c.Cockroach.ClientCA != "" && c.Cockroach.ClientSSLKey == "" && c.Cockroach.ClientSSLCert == "": // OneWayTLS
		psqlInfo += fmt.Sprintf(" ssl=true sslmode=verify-full sslrootcert=%s", c.Cockroach.ClientCA)
	case c.Cockroach.ClientCA != "" || c.Cockroach.ClientSSLKey != "" || c.Cockroach.ClientSSLCert != "":
		return nil, fmt.Errorf("%w, CA: %s, Key: %s, Cert: %s",
			ErrIncompleteCertInformationProvided, c.Cockroach.ClientCA, c.Cockroach.ClientSSLKey, c.Cockroach.ClientSSLCert)
	default:
		psqlInfo += " sslmode=disable"
	}
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		TablePrefix: c.Prefix,
	}})
	if err != nil {
		return nil, err
	}
	return db, nil
}

type Config struct {
	Use       string `default:"cockroach"`
	Prefix    string `default:""`
	Cockroach struct {
		Host          string `default:"cockroach"`
		Port          string `default:"26257"`
		UserName      string `default:"root"`
		Password      string `default:""`
		ClientCA      string `default:""`
		ClientSSLKey  string `default:""`
		ClientSSLCert string `default:""`
		Database      string `default:"scoretrak"`
	}
}
