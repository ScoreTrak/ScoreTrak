package storage

import (
	"errors"
	"fmt"

	"ScoreTrak/pkg/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/qor/validations"
)

func NewDb(c *config.StaticConfig) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	if c.DB.Use == "cockroach" {
		db, err = newCockroach(c)
	}

	if err != nil {
		return nil, errors.New("Not supported db")
	}
	validations.RegisterCallbacks(db)
	db.BlockGlobalUpdate(true)
	return db, nil
}

func newCockroach(c *config.StaticConfig) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable",
		c.DB.Cockroach.Host,
		c.DB.Cockroach.Port,
		c.DB.Cockroach.UserName,
		c.DB.Cockroach.Database)

	db, err := gorm.Open("postgres", psqlInfo)

	if err != nil {
		return nil, err
	}
	return db, nil
}
