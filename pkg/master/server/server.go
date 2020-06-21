package server

import (
	"ScoreTrak/pkg/check"
	"ScoreTrak/pkg/config"
	"ScoreTrak/pkg/host"
	"ScoreTrak/pkg/host_group"
	"ScoreTrak/pkg/logger"
	"ScoreTrak/pkg/property"
	"ScoreTrak/pkg/report"
	"ScoreTrak/pkg/round"
	"ScoreTrak/pkg/service"
	"ScoreTrak/pkg/service_group"
	"ScoreTrak/pkg/swarm"
	"ScoreTrak/pkg/team"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"go.uber.org/dig"
	"math"
	"net/http"
	"time"
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
	ds.cont.Invoke(func(d *gorm.DB) {
		db = d
	})

	var tm time.Time
	res, err := db.Raw("SELECT current_timestamp;").Rows()
	if err != nil {
		panic(err)
	}
	defer res.Close()
	for res.Next() {
		res.Scan(&tm)
	}
	timeDiff := time.Since(tm)
	if float64(time.Second*2) < math.Abs(float64(timeDiff)) {
		panic(errors.New(
			fmt.Sprintf("time difference between master host, and database host are is large. Please synchronize time\n(The difference should not exceed 2 seconds)\nTime on database:%s\nTime on master:%s", tm.String(), time.Now())))
	}

	db.AutoMigrate(&team.Team{})
	db.AutoMigrate(&check.Check{})
	db.AutoMigrate(&config.DynamicConfig{})
	db.AutoMigrate(&host.Host{})
	db.AutoMigrate(&host_group.HostGroup{})
	db.AutoMigrate(&property.Property{})
	db.AutoMigrate(&round.Round{})
	db.AutoMigrate(&service.Service{})
	db.AutoMigrate(&service_group.ServiceGroup{})
	db.AutoMigrate(&swarm.Swarm{})
	db.AutoMigrate(&report.Report{})
	db.Model(&check.Check{}).AddForeignKey("service_id", "services(id)", "CASCADE", "RESTRICT")
	db.Model(&check.Check{}).AddForeignKey("round_id", "rounds(id)", "CASCADE", "RESTRICT")
	db.Model(&host.Host{}).AddForeignKey("host_group_id", "host_groups(id)", "RESTRICT", "RESTRICT")
	db.Model(&host.Host{}).AddForeignKey("team_name", "teams(name)", "RESTRICT", "RESTRICT")
	db.Model(&property.Property{}).AddForeignKey("service_id", "services(id)", "CASCADE", "RESTRICT")
	db.Model(&service.Service{}).AddForeignKey("service_group_id", "service_groups(id)", "RESTRICT", "RESTRICT")
	db.Model(&service.Service{}).AddForeignKey("host_id", "hosts(id)", "RESTRICT", "RESTRICT")
	db.Model(&swarm.Swarm{}).AddForeignKey("service_group_id", "service_groups(id)", "CASCADE", "RESTRICT")

	return nil
}

// Start start serving the application
func (ds *dserver) Start() error {
	var cfg *config.StaticConfig
	ds.cont.Invoke(func(c *config.StaticConfig) { cfg = c })
	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), ds.router)
		if err != nil {
			ds.logger.Error(err)
		}
	}()
	return nil
}
