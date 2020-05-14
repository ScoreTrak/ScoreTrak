package server

import (
	"ScoreTrak/pkg/check"
	"ScoreTrak/pkg/config"
	"ScoreTrak/pkg/host"
	"ScoreTrak/pkg/host_group"
	"ScoreTrak/pkg/logger"
	"ScoreTrak/pkg/property"
	"ScoreTrak/pkg/round"
	"ScoreTrak/pkg/service"
	"ScoreTrak/pkg/service_group"
	"ScoreTrak/pkg/swarm"
	"ScoreTrak/pkg/team"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"go.uber.org/dig"
	"net/http"
)

type dserver struct {
	router *mux.Router
	cont   *dig.Container
	logger logger.LogInfoFormat
}

func NewServer(e *mux.Router, c *dig.Container, l logger.LogInfoFormat) *dserver {
	return &dserver{
		router: e,
		cont:   c,
		logger: l,
	}
}

func (ds *dserver) SetupDB() error {
	var db *gorm.DB

	if err := ds.cont.Invoke(func(d *gorm.DB) {
		db = d
	}); err != nil {
		return err
	}

	db.AutoMigrate(&team.Team{})
	db.AutoMigrate(&check.Check{})
	db.AutoMigrate(&config.Config{})
	db.AutoMigrate(&host.Host{})
	db.AutoMigrate(&host_group.HostGroup{})
	db.AutoMigrate(&property.Property{})
	db.AutoMigrate(&round.Round{})
	db.AutoMigrate(&service.Service{})
	db.AutoMigrate(&service_group.ServiceGroup{})
	db.AutoMigrate(&swarm.Swarm{})

	return nil
}

// Start start serving the application
func (ds *dserver) Start() error {
	var cfg *config.Config
	if err := ds.cont.Invoke(func(c *config.Config) { cfg = c }); err != nil {
		return err
	}
	return http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), ds.router)

}