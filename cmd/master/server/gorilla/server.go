package gorilla

import (
	"errors"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/host"
	"github.com/ScoreTrak/ScoreTrak/pkg/host_group"
	"github.com/ScoreTrak/ScoreTrak/pkg/logger"
	"github.com/ScoreTrak/ScoreTrak/pkg/property"
	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	"github.com/ScoreTrak/ScoreTrak/pkg/round"
	"github.com/ScoreTrak/ScoreTrak/pkg/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/service_group"
	"github.com/ScoreTrak/ScoreTrak/pkg/team"
	"github.com/gorilla/mux"
	"go.uber.org/dig"
	"gorm.io/gorm"
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

	err = db.AutoMigrate(&team.Team{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&report.Report{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&config.DynamicConfig{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&host_group.HostGroup{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&service_group.ServiceGroup{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&host.Host{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&round.Round{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&service.Service{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&check.Check{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&property.Property{})
	if err != nil {
		panic(err)
	}

	return nil
}

// Start start serving the application
func (ds *dserver) Start() error {
	var cfg config.StaticConfig
	ds.cont.Invoke(func(c config.StaticConfig) { cfg = c })
	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), ds.router)
		if err != nil {
			ds.logger.Error(err)
		}
	}()
	return nil
}
