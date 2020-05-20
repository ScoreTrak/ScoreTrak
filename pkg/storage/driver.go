package storage

import (
	"errors"
	"fmt"

	"ScoreTrak/pkg/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewDb(c *config.StaticConfig) (*gorm.DB, error) {
	if c.DB.Use == "cockroach" {
		return newCockroach(c)
	}
	return nil, errors.New("Not supported db")
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
