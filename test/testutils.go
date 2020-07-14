package test

import (
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/check"
	"github.com/L1ghtman2k/ScoreTrak/pkg/config"
	"github.com/L1ghtman2k/ScoreTrak/pkg/host"
	"github.com/L1ghtman2k/ScoreTrak/pkg/host_group"
	"github.com/L1ghtman2k/ScoreTrak/pkg/logger"
	"github.com/L1ghtman2k/ScoreTrak/pkg/property"
	"github.com/L1ghtman2k/ScoreTrak/pkg/round"
	"github.com/L1ghtman2k/ScoreTrak/pkg/service"
	"github.com/L1ghtman2k/ScoreTrak/pkg/service_group"
	"github.com/L1ghtman2k/ScoreTrak/pkg/storage"
	"github.com/L1ghtman2k/ScoreTrak/pkg/team"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"net/http/httptest"
	"time"
)

func SetupDB(c storage.Config) *gorm.DB {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s sslmode=disable",
		c.Cockroach.Host,
		c.Cockroach.Port,
		c.Cockroach.UserName)
	dbPrep, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	dbPrep.Exec(fmt.Sprintf("drop database if exists  %s", c.Cockroach.Database))
	dbPrep.Exec(fmt.Sprintf("create database if not exists  %s", c.Cockroach.Database))
	dbPrep.Close()
	db, err := storage.NewDB(c)
	if err != nil {
		panic(err)
	}
	return db
}

func CreateAllTables(db *gorm.DB) {
	db.AutoMigrate(&team.Team{})
	db.AutoMigrate(&check.Check{})
	db.AutoMigrate(&config.DynamicConfig{})
	db.AutoMigrate(&host.Host{})
	db.AutoMigrate(&host_group.HostGroup{})
	db.AutoMigrate(&property.Property{})
	db.AutoMigrate(&round.Round{})
	db.AutoMigrate(&service.Service{})
	db.AutoMigrate(&service_group.ServiceGroup{})
	db.Model(&check.Check{}).AddForeignKey("service_id", "services(id)", "CASCADE", "RESTRICT")
	db.Model(&check.Check{}).AddForeignKey("round_id", "rounds(id)", "CASCADE", "RESTRICT")
	db.Model(&host.Host{}).AddForeignKey("host_group_id", "host_groups(id)", "RESTRICT", "RESTRICT")
	db.Model(&host.Host{}).AddForeignKey("team_name", "teams(name)", "RESTRICT", "RESTRICT")
	db.Model(&property.Property{}).AddForeignKey("service_id", "services(id)", "CASCADE", "RESTRICT")
	db.Model(&service.Service{}).AddForeignKey("service_group_id", "service_groups(id)", "RESTRICT", "RESTRICT")
	db.Model(&service.Service{}).AddForeignKey("host_id", "hosts(id)", "RESTRICT", "RESTRICT")
}

func SetupConfig(f string) config.StaticConfig {
	var err error
	err = config.NewStaticConfig(f)
	if err != nil {
		panic(err)
	}
	return config.GetStaticConfig()
}

func NewConfigClone(c config.StaticConfig) config.StaticConfig {
	cnf := config.StaticConfig{}
	err := copier.Copy(&cnf, &c)
	if err != nil {
		panic(err)
	}
	return cnf
}

func CleanAllTables(db *gorm.DB) {
	db.DropTableIfExists(&check.Check{})
	db.DropTableIfExists(&property.Property{})
	db.DropTableIfExists(&service.Service{})
	db.DropTableIfExists(&host.Host{})
	db.DropTableIfExists(&host_group.HostGroup{})
	db.DropTableIfExists(&round.Round{})
	db.DropTableIfExists(&service_group.ServiceGroup{})
	db.DropTableIfExists(&team.Team{})
	db.DropTableIfExists(&config.DynamicConfig{})
}

func DropDB(db *gorm.DB, c config.StaticConfig) {
	db.Exec(fmt.Sprintf("drop database %s", c.DB.Cockroach.Database))
}

func SetupLogger(c logger.Config) logger.LogInfoFormat {
	l, err := logger.NewLogger(c)
	if err != nil {
		panic(err)
	}
	return l
}

func DataPreload(db *gorm.DB) {
	var count int
	CreateAllTables(db)
	//Creating Config
	db.Exec("INSERT INTO config (id, round_duration, enabled) VALUES (1, 60, true)")
	db.Table("config").Count(&count)
	if count != 1 {
		panic("There should be 1 entry in config")
	}
	//Creating Teams
	db.Exec("INSERT INTO teams (id, name, enabled) VALUES (1, 'TeamOne', true)")
	db.Exec("INSERT INTO teams (id, name, enabled) VALUES (2, 'TeamTwo', false)")
	db.Exec("INSERT INTO teams (id, name, enabled) VALUES (3, 'TeamThree', true)")
	db.Exec("INSERT INTO teams (id, name, enabled) VALUES (4, 'TeamFour', false)")
	db.Table("teams").Count(&count)
	if count != 4 {
		panic("There should be 4 entry in teams")
	}
	//Creating Host Groups
	db.Exec("INSERT INTO host_groups (id, name, enabled) VALUES (1, 'HostGroup1', true)")
	db.Exec("INSERT INTO host_groups (id, name, enabled) VALUES (2, 'HostGroup2', false)")
	db.Exec("INSERT INTO host_groups (id, name, enabled) VALUES (3, 'HostGroup3', true)")
	db.Exec("INSERT INTO host_groups (id, name, enabled) VALUES (4, 'HostGroup4', false)")
	db.Table("host_groups").Count(&count)
	if count != 4 {
		panic("There should be 4 entry in host groups")
	}
	//Creating Service Groups
	db.Exec("INSERT INTO service_groups (id, name, enabled) VALUES (1, 'ServiceGroup1', true)")
	db.Exec("INSERT INTO service_groups (id, name, enabled) VALUES (2, 'ServiceGroup2', false)")
	db.Exec("INSERT INTO service_groups (id, name, enabled) VALUES (3, 'ServiceGroup3', true)")
	db.Exec("INSERT INTO service_groups (id, name, enabled) VALUES (4, 'ServiceGroup4', false)")
	db.Table("service_groups").Count(&count)
	if count != 4 {
		panic("There should be 4 entry in service groups")
	}
	//Creating Hosts
	db.Exec("INSERT INTO hosts (id, address, team_name, host_group_id, enabled, edit_host) VALUES (1, '10.0.0.1', NULL, NULL, true, true)")
	db.Exec("INSERT INTO hosts (id, address, team_name, host_group_id, enabled, edit_host) VALUES (2, '10.0.0.2', 'TeamTwo', NULL, false, true)")
	db.Exec("INSERT INTO hosts (id, address, team_name, host_group_id, enabled, edit_host) VALUES (3, '10.0.0.3', NULL, 3, true, false)")
	db.Exec("INSERT INTO hosts (id, address, team_name, host_group_id, enabled, edit_host) VALUES (4, '10.0.0.4', 'TeamFour', 4, false, false)")
	db.Table("hosts").Count(&count)
	if count != 4 {
		panic("There should be 4 entry in hosts")
	}
	//Creating Services
	db.Exec("INSERT INTO services (id, service_group_id, host_id, name, display_name, points, round_units, round_delay, enabled) VALUES (1, 1, 1, 'WINRM', 'host1-service1', 0, 1, 0, true)")
	db.Exec("INSERT INTO services (id, service_group_id, host_id, name, display_name, points, round_units, round_delay, enabled) VALUES (2, 2, 2, 'FTP', 'host2-service2', 40, 23, 2, true)")
	db.Exec("INSERT INTO services (id, service_group_id, host_id, name, display_name, points, round_units, round_delay, enabled) VALUES (3, 4, 2, 'SSH', 'host3-service3', 50, 3, 0, true)")
	db.Exec("INSERT INTO services (id, service_group_id, host_id, name, display_name, points, round_units, round_delay, enabled) VALUES (4, 4, 4, 'HTTP', 'host4-service4', 200, 4, 3, true)")
	db.Exec("INSERT INTO services (id, service_group_id, host_id, name, display_name, points, round_units, round_delay, enabled) VALUES (5, 1, 2, 'SSH', 'host1-service2', 30, 5, 4, false)")
	db.Exec("INSERT INTO services (id, service_group_id, host_id, name, display_name, points, round_units, round_delay, enabled) VALUES (6, 2, 1, 'SMB', 'host2-service1', 2, 5, 2, true)")
	db.Exec("INSERT INTO services (id, service_group_id, host_id, name, display_name, points, round_units, round_delay, enabled) VALUES (7, 2, 4, 'FTP', 'host3-service4', 55, 6, 3, false)")
	db.Exec("INSERT INTO services (id, service_group_id, host_id, name, display_name, points, round_units, round_delay, enabled) VALUES (8, 4, 4, 'IMAP', 'host4-service3', 44, 23, 22, false)")
	db.Table("services").Count(&count)
	if count != 8 {
		panic("There should be 8 entry in services")
	}
	//Creating Properties
	db.Exec("INSERT INTO properties (id, service_id, key, value, description, status) VALUES (1, 1, 'Port','80','Port to which connect','View')")
	db.Exec("INSERT INTO properties (id, service_id, key, value, description, status) VALUES (2, 1, 'Username','root','Username of the service','Edit')")
	db.Exec("INSERT INTO properties (id, service_id, key, value, description, status) VALUES (3, 2, 'Port','53','Port to which connect','View')")
	db.Exec("INSERT INTO properties (id, service_id, key, value, description, status) VALUES (4, 2, 'Password','changeme','Password for the account','Hide')")
	db.Exec("INSERT INTO properties (id, service_id, key, value, description, status) VALUES (5, 2, 'Username','root','Username of the service','Hide')")
	db.Exec("INSERT INTO properties (id, service_id, key, value, description, status) VALUES (6, 3, 'Password','changeme','Password for the account','Hide')")
	db.Exec("INSERT INTO properties (id, service_id, key, value, description, status) VALUES (7, 4, 'Port','443','Port to which connect','Edit')")
	db.Exec("INSERT INTO properties (id, service_id, key, value, description, status) VALUES (8, 5, 'Username','admin','Username of the service','Edit')")
	db.Exec("INSERT INTO properties (id, service_id, key, value, description, status) VALUES (9, 6, 'Password','Change.me!','Password for the account','View')")
	db.Exec("INSERT INTO properties (id, service_id, key, value, status) VALUES (10, 6, 'Content','Sample Content!', 'Hide')")
	db.Exec("INSERT INTO properties (id, service_id, key, value, description, status) VALUES (11, 6, 'EncodedContent', 'RW5jb2RlZENvbnRlbnQ=', 'Encoded content!', 'Edit')")
	db.Exec("INSERT INTO properties (id, service_id, key, value, status) VALUES (12, 6, 'domain', 'sample.com', 'View')")
	db.Exec("INSERT INTO properties (id, service_id, key, value, status) VALUES (13, 6, 'domain', 'ad.sample.com', 'View')")
	db.Table("properties").Count(&count)
	if count != 13 {
		panic("There should be 13 entry in properties")
	}
	//Creating Rounds
	db.Exec("INSERT INTO rounds (id, start, finish) VALUES (1, $1, $2)", time.Now().Add(-time.Second*60*3), time.Now().Add(-time.Second*60*3+time.Second*30))
	db.Exec("INSERT INTO rounds (id, start, finish) VALUES (2, $1, $2)", time.Now().Add(-time.Second*60*2), time.Now().Add(-time.Second*60*2+time.Second*20))
	db.Exec("INSERT INTO rounds (id, start, finish) VALUES (3, $1, $2)", time.Now().Add(-time.Second*60), time.Now().Add(-time.Second*60+time.Second*40))
	db.Exec("INSERT INTO rounds (id, start) VALUES (4, $1)", time.Now())
	db.Table("rounds").Count(&count)
	if count != 4 {
		panic("There should be 4 entry in rounds")
	}
	//Creating Checks
	db.Exec("INSERT INTO checks (round_id, service_id, log, passed) VALUES (1, 1, '',true)")
	db.Exec("INSERT INTO checks (round_id, service_id, log, passed) VALUES (2, 1, '',true)")
	db.Exec("INSERT INTO checks (round_id, service_id, log, passed) VALUES (3, 1, 'Failed because of incorrect password',false)")
	db.Exec("INSERT INTO checks (round_id, service_id, log, passed) VALUES (3, 3, '',true)")
	db.Table("checks").Count(&count)
	if count != 4 {
		panic("There should be 4 entry in checks")
	}

}

func NewJsonRecorder() *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	return w
}
