package orm_test

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
	"ScoreTrak/pkg/storage/orm"
	"ScoreTrak/pkg/swarm"
	"ScoreTrak/pkg/team"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var dbPrep *gorm.DB
var db *gorm.DB
var l logger.LogInfoFormat
var c *config.StaticConfig

var _ = Describe("Orm", func() {
	It("Sets up database", func() {
		setupConfig()
		setupLogger()
		setupDB()
	})
	It("Creates tables", func() {
		createTables()
	})
	Describe("Team", func() {
		var tr team.Repo
		Context("initially", func() {
			It("should be empty", func() {
				tr = orm.NewTeamRepo(db, l)
				ac, err := tr.GetAll()
				Ω(err).ShouldNot(HaveOccurred())
				Ω(ac).Should(HaveLen(0))
			})
		})
		Context("when new item is added", func() {
			It("it should not be empty", func() {
				var err error
				t := team.Team{ID: "TestTeam"}
				err = tr.Store(&t)
				Ω(err).ShouldNot(HaveOccurred())
				ac, err := tr.GetAll()
				Ω(err).ShouldNot(HaveOccurred())
				Ω(ac).Should(HaveLen(1))
			})
		})
	})

	It("Drops database", func() {
		dropDB()
	})
})

func setupDB() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s sslmode=disable",
		c.DB.Cockroach.Host,
		c.DB.Cockroach.Port,
		c.DB.Cockroach.UserName)
	dbPrep, err = gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	dbPrep.Exec(fmt.Sprintf("drop database if exists  %s", c.DB.Cockroach.Database))
	dbPrep.Exec(fmt.Sprintf("create database if not exists  %s", c.DB.Cockroach.Database))
	db, err = storage.NewDb(c)
	if err != nil {
		panic(err)
	}
}

func createTables() {
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

func setupConfig() {
	var err error
	err = config.NewStaticConfig("../../../configs/dev-config.yml")
	if err != nil {
		panic(err)
	}
	c = config.GetConfig()
}

func cleanDB() {
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

func dropDB() {
	db.Close()
	dbPrep.Exec(fmt.Sprintf("drop database %s", c.DB.Cockroach.Database))
	dbPrep.Close()
}

func setupLogger() {
	var err error
	l, err = logger.NewLogger(c)
	if err != nil {
		panic(err)
	}
}
