package gorilla

import (
	"errors"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/logger"
	sutil "github.com/ScoreTrak/ScoreTrak/pkg/storage/orm/util"
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
	return nil
}

func (ds *dserver) LoadTables(db *gorm.DB) (err error) {
	err = sutil.CreateAllTables(db)
	if err != nil {
		return err
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
