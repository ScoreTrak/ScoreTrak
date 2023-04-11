package storage

import (
	"context"

	"github.com/ScoreTrak/ScoreTrak/internal/entities"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

// NewDB creates an instance of database based on config
func NewDB(c *config.Config) (*entities.Client, error) {
	client, err := entities.Open(c.DB.Use, c.DB.DSN)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func AutoMigrate(dbClient *entities.Client) error {
	if err := dbClient.Schema.Create(context.Background()); err != nil {
		return err
	}
	return nil
}
