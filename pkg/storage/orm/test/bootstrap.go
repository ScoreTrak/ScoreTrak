package test

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
	"ScoreTrak/pkg/storage"
	"ScoreTrak/pkg/swarm"
	"ScoreTrak/pkg/team"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

func setupDB(c *config.StaticConfig) *gorm.DB {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s sslmode=disable",
		c.DB.Cockroach.Host,
		c.DB.Cockroach.Port,
		c.DB.Cockroach.UserName)
	dbPrep, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	dbPrep.Exec(fmt.Sprintf("drop database if exists  %s", c.DB.Cockroach.Database))
	dbPrep.Exec(fmt.Sprintf("create database if not exists  %s", c.DB.Cockroach.Database))
	dbPrep.Close()
	db, err := storage.NewDb(c)
	if err != nil {
		panic(err)
	}
	return db
}

func createTables(db *gorm.DB) {
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
	db.Model(&check.Check{}).AddForeignKey("service_id", "services(id)", "RESTRICT", "RESTRICT")
	db.Model(&check.Check{}).AddForeignKey("round_id", "rounds(id)", "RESTRICT", "RESTRICT")
	db.Model(&host.Host{}).AddForeignKey("host_group_id", "host_groups(id)", "RESTRICT", "RESTRICT")
	db.Model(&host.Host{}).AddForeignKey("team_id", "teams(id)", "RESTRICT", "RESTRICT")
	db.Model(&property.Property{}).AddForeignKey("service_id", "services(id)", "RESTRICT", "RESTRICT")
	db.Model(&service.Service{}).AddForeignKey("service_group_id", "service_groups(id)", "RESTRICT", "RESTRICT")
	db.Model(&service.Service{}).AddForeignKey("host_id", "hosts(id)", "RESTRICT", "RESTRICT")
	db.Model(&swarm.Swarm{}).AddForeignKey("service_group_id", "service_groups(id)", "RESTRICT", "RESTRICT")
}

func setupConfig() *config.StaticConfig {
	var err error
	err = config.NewStaticConfig("dev-config.yml")
	if err != nil {
		panic(err)
	}
	return config.GetConfig()
}

func newConfigClone(c *config.StaticConfig) *config.StaticConfig {
	cnf := config.StaticConfig{}
	err := copier.Copy(&cnf, &c)
	if err != nil {
		panic(err)
	}
	return &cnf
}

func cleanDB(db *gorm.DB) {
	db.DropTableIfExists(&check.Check{})
	db.DropTableIfExists(&property.Property{})
	db.DropTableIfExists(&swarm.Swarm{})
	db.DropTableIfExists(&service.Service{})
	db.DropTableIfExists(&host.Host{})
	db.DropTableIfExists(&host_group.HostGroup{})
	db.DropTableIfExists(&round.Round{})
	db.DropTableIfExists(&service_group.ServiceGroup{})
	db.DropTableIfExists(&team.Team{})
	db.DropTableIfExists(&config.DynamicConfig{})
}

func dropDB(db *gorm.DB, c *config.StaticConfig) {
	db.Exec(fmt.Sprintf("drop database %s", c.DB.Cockroach.Database))
}

func setupLogger(c *config.StaticConfig) logger.LogInfoFormat {
	l, err := logger.NewLogger(c)
	if err != nil {
		panic(err)
	}
	return l
}
