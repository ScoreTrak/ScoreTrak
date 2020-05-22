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

var _ = BeforeSuite(func() {
	var err error
	err = config.NewStaticConfig("../../../configs/dev-config.yml")
	if err != nil {
		panic(err)
	}
	c := config.GetConfig()
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
	db, err = storage.NewDb(c)
	if err != nil {
		panic(err)
	}
	l, err = logger.NewLogger(c)
	if err != nil {
		panic(err)
	}
})

var _ = AfterSuite(func() {
	db.Close()
	dbPrep.Exec(fmt.Sprintf("drop database %s", c.DB.Cockroach.Database))
	dbPrep.Close()
})

var _ = Describe("Orm", func() {
	BeforeEach(func() {
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
	})
	AfterEach(func() {
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
	})

	Describe("Team", func() {
		Context("Create and List", func() {
			tr := orm.NewTeamRepo(db, l)
			It("should be empty", func() {
				ac, err := tr.GetAll()
				Ω(err).ShouldNot(HaveOccurred())
				Ω(ac).Should(HaveLen(0))
			})
			It("should create one", func() {
				t := team.Team{ID: "TestTeam"}
				err := tr.Store(&t)
				Ω(err).ShouldNot(HaveOccurred())
			})
			It("should have one", func() {
				ac, err := tr.GetAll()
				Ω(err).ShouldNot(HaveOccurred())
				Ω(ac).Should(HaveLen(1))
			})
		})
	})
})
