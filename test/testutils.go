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
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http/httptest"
	"time"
)

func SetupDB(c storage.Config) *gorm.DB {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s sslmode=disable",
		c.Cockroach.Host,
		c.Cockroach.Port,
		c.Cockroach.UserName)
	dbPrep, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	dbPrep.Exec(fmt.Sprintf("drop database if exists  %s", c.Cockroach.Database))
	dbPrep.Exec(fmt.Sprintf("create database if not exists  %s", c.Cockroach.Database))
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
	db.Migrator().DropTable(&check.Check{})
	db.Migrator().DropTable(&check.Check{})
	db.Migrator().DropTable(&property.Property{})
	db.Migrator().DropTable(&service.Service{})
	db.Migrator().DropTable(&host.Host{})
	db.Migrator().DropTable(&host_group.HostGroup{})
	db.Migrator().DropTable(&round.Round{})
	db.Migrator().DropTable(&service_group.ServiceGroup{})
	db.Migrator().DropTable(&team.Team{})
	db.Migrator().DropTable(&config.DynamicConfig{})
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
	var count int64
	CreateAllTables(db)
	//Creating Config
	db.Exec("INSERT INTO config (id, round_duration, enabled) VALUES (1, 60, true)")
	db.Table("config").Count(&count)
	if count != 1 {
		panic("There should be 1 entry in config")
	}
	//Creating Teams
	db.Exec("INSERT INTO teams (id, name, enabled) VALUES ('11111111-1111-1111-1111-111111111111', 'TeamOne', true)")
	db.Exec("INSERT INTO teams (id, name, enabled) VALUES ('22222222-2222-2222-2222-222222222222', 'TeamTwo', false)")
	db.Exec("INSERT INTO teams (id, name, enabled) VALUES ('33333333-3333-3333-3333-333333333333', 'TeamThree', true)")
	db.Exec("INSERT INTO teams (id, name, enabled) VALUES ('44444444-4444-4444-4444-444444444444', 'TeamFour', false)")
	db.Table("teams").Count(&count)
	if count != 4 {
		panic("There should be 4 entry in teams")
	}
	//Creating Host Groups
	db.Exec("INSERT INTO host_groups (id, name, enabled) VALUES ('11111111-1111-1111-1111-111111111111', 'HostGroup1', true)")
	db.Exec("INSERT INTO host_groups (id, name, enabled) VALUES ('22222222-2222-2222-2222-222222222222', 'HostGroup2', false)")
	db.Exec("INSERT INTO host_groups (id, name, enabled) VALUES ('33333333-3333-3333-3333-333333333333', 'HostGroup3', true)")
	db.Exec("INSERT INTO host_groups (id, name, enabled) VALUES ('44444444-4444-4444-4444-444444444444', 'HostGroup4', false)")
	db.Table("host_groups").Count(&count)
	if count != 4 {
		panic("There should be 4 entry in host groups")
	}
	//Creating Service Groups
	db.Exec("INSERT INTO service_groups (id, name, enabled) VALUES ('11111111-1111-1111-1111-111111111111', 'ServiceGroup1', true)")
	db.Exec("INSERT INTO service_groups (id, name, enabled) VALUES ('22222222-2222-2222-2222-222222222222', 'ServiceGroup2', false)")
	db.Exec("INSERT INTO service_groups (id, name, enabled) VALUES ('33333333-3333-3333-3333-333333333333', 'ServiceGroup3', true)")
	db.Exec("INSERT INTO service_groups (id, name, enabled) VALUES ('44444444-4444-4444-4444-444444444444', 'ServiceGroup4', false)")
	db.Table("service_groups").Count(&count)
	if count != 4 {
		panic("There should be 4 entry in service groups")
	}
	//Creating Hosts
	db.Exec("INSERT INTO hosts (id, address, team_id, host_group_id, enabled, edit_host) VALUES ('11111111-1111-1111-1111-111111111111', '10.0.0.1', '11111111-1111-1111-1111-111111111111', NULL, true, true)")
	db.Exec("INSERT INTO hosts (id, address, team_id, host_group_id, enabled, edit_host) VALUES ('22222222-2222-2222-2222-222222222222', '10.0.0.2', '22222222-2222-2222-2222-222222222222', NULL, false, true)")
	db.Exec("INSERT INTO hosts (id, address, team_id, host_group_id, enabled, edit_host) VALUES ('33333333-3333-3333-3333-333333333333', '10.0.0.3', '11111111-1111-1111-1111-111111111111', '33333333-3333-3333-3333-333333333333', true, false)")
	db.Exec("INSERT INTO hosts (id, address, team_id, host_group_id, enabled, edit_host) VALUES ('44444444-4444-4444-4444-444444444444', '10.0.0.4', '44444444-4444-4444-4444-444444444444', '44444444-4444-4444-4444-444444444444', false, false)")
	db.Table("hosts").Count(&count)
	if count != 4 {
		panic("There should be 4 entry in hosts")
	}
	//Creating Services
	db.Exec("INSERT INTO services (id, service_group_id, host_id, name, display_name, points, round_units, round_delay, enabled) VALUES ('11111111-1111-1111-1111-111111111111', '11111111-1111-1111-1111-111111111111', '11111111-1111-1111-1111-111111111111', 'WINRM', 'host1-service1', 0, 1, 0, true)")
	db.Exec("INSERT INTO services (id, service_group_id, host_id, name, display_name, points, round_units, round_delay, enabled) VALUES ('22222222-2222-2222-2222-222222222222', '22222222-2222-2222-2222-222222222222', '22222222-2222-2222-2222-222222222222', 'FTP', 'host2-service2', 40, 23, 2, true)")
	db.Exec("INSERT INTO services (id, service_group_id, host_id, name, display_name, points, round_units, round_delay, enabled) VALUES ('33333333-3333-3333-3333-333333333333', '44444444-4444-4444-4444-444444444444', '22222222-2222-2222-2222-222222222222', 'SSH', 'host3-service3', 50, 3, 0, true)")
	db.Exec("INSERT INTO services (id, service_group_id, host_id, name, display_name, points, round_units, round_delay, enabled) VALUES ('44444444-4444-4444-4444-444444444444', '44444444-4444-4444-4444-444444444444', '44444444-4444-4444-4444-444444444444', 'HTTP', 'host4-service4', 200, 4, 3, true)")
	db.Exec("INSERT INTO services (id, service_group_id, host_id, name, display_name, points, round_units, round_delay, enabled) VALUES ('55555555-5555-5555-5555-555555555555', '11111111-1111-1111-1111-111111111111', '22222222-2222-2222-2222-222222222222', 'SSH', 'host1-service2', 30, 5, 4, false)")
	db.Exec("INSERT INTO services (id, service_group_id, host_id, name, display_name, points, round_units, round_delay, enabled) VALUES ('66666666-6666-6666-6666-666666666666', '22222222-2222-2222-2222-222222222222', '11111111-1111-1111-1111-111111111111', 'SMB', 'host2-service1', 2, 5, 2, true)")
	db.Exec("INSERT INTO services (id, service_group_id, host_id, name, display_name, points, round_units, round_delay, enabled) VALUES ('77777777-7777-7777-7777-777777777777', '22222222-2222-2222-2222-222222222222', '44444444-4444-4444-4444-444444444444', 'FTP', 'host3-service4', 55, 6, 3, false)")
	db.Exec("INSERT INTO services (id, service_group_id, host_id, name, display_name, points, round_units, round_delay, enabled) VALUES ('88888888-8888-8888-8888-888888888888', '44444444-4444-4444-4444-444444444444', '44444444-4444-4444-4444-444444444444', 'IMAP', 'host4-service3', 44, 23, 22, false)")
	db.Table("services").Count(&count)
	if count != 8 {
		panic("There should be 8 entry in services")
	}
	//Creating Properties
	db.Exec("INSERT INTO properties (id, service_id, key, value, description, status) VALUES (	'11111111-1111-1111-1111-111111111111', '11111111-1111-1111-1111-111111111111', 'Port','80','Port to which connect','View')")
	db.Exec("INSERT INTO properties (id, service_id, key, value, description, status) VALUES (	'22222222-2222-2222-2222-222222222222', '11111111-1111-1111-1111-111111111111', 'Username','root','Username of the service','Edit')")
	db.Exec("INSERT INTO properties (id, service_id, key, value, description, status) VALUES (	'33333333-3333-3333-3333-333333333333', '22222222-2222-2222-2222-222222222222', 'Port','53','Port to which connect','View')")
	db.Exec("INSERT INTO properties (id, service_id, key, value, description, status) VALUES (	'44444444-4444-4444-4444-444444444444', '22222222-2222-2222-2222-222222222222', 'Password','changeme','Password for the account','Hide')")
	db.Exec("INSERT INTO properties (id, service_id, key, value, description, status) VALUES (	'55555555-5555-5555-5555-555555555555', '22222222-2222-2222-2222-222222222222', 'Username','root','Username of the service','Hide')")
	db.Exec("INSERT INTO properties (id, service_id, key, value, description, status) VALUES (	'66666666-6666-6666-6666-666666666666', '33333333-3333-3333-3333-333333333333', 'Password','changeme','Password for the account','Hide')")
	db.Exec("INSERT INTO properties (id, service_id, key, value, description, status) VALUES (	'77777777-7777-7777-7777-777777777777', '44444444-4444-4444-4444-444444444444', 'Port','443','Port to which connect','Edit')")
	db.Exec("INSERT INTO properties (id, service_id, key, value, description, status) VALUES (	'88888888-8888-8888-8888-888888888888', '55555555-5555-5555-5555-555555555555', 'Username','admin','Username of the service','Edit')")
	db.Exec("INSERT INTO properties (id, service_id, key, value, description, status) VALUES (	'99999999-9999-9999-9999-999999999999', '66666666-6666-6666-6666-666666666666', 'Password','Change.me!','Password for the account','View')")
	db.Exec("INSERT INTO properties (id, service_id, key, value, status) VALUES (				'11111111-1111-1111-1111-111111111110', '66666666-6666-6666-6666-666666666666', 'Content','Sample Content!', 'Hide')")
	db.Exec("INSERT INTO properties (id, service_id, key, value, description, status) VALUES (	'11111111-1111-1111-1111-111111111100', '66666666-6666-6666-6666-666666666666', 'EncodedContent', 'RW5jb2RlZENvbnRlbnQ=', 'Encoded content!', 'Edit')")
	db.Exec("INSERT INTO properties (id, service_id, key, value, status) VALUES (				'11111111-1111-1111-1111-111111111000', '66666666-6666-6666-6666-666666666666', 'domain', 'sample.com', 'View')")
	db.Exec("INSERT INTO properties (id, service_id, key, value, status) VALUES (				'11111111-1111-1111-1111-111111110000', '66666666-6666-6666-6666-666666666666', 'domain', 'ad.sample.com', 'View')")
	db.Table("properties").Count(&count)
	if count != 13 {
		panic("There should be 13 entry in properties")
	}
	//Creating Rounds
	db.Exec("INSERT INTO rounds (id, start, finish) VALUES (1, ?, ?)", time.Now().Add(-time.Second*60*3), time.Now().Add(-time.Second*60*3+time.Second*30))
	db.Exec("INSERT INTO rounds (id, start, finish) VALUES (2, ?, ?)", time.Now().Add(-time.Second*60*2), time.Now().Add(-time.Second*60*2+time.Second*20))
	db.Exec("INSERT INTO rounds (id, start, finish) VALUES (3, ?, ?)", time.Now().Add(-time.Second*60), time.Now().Add(-time.Second*60+time.Second*40))
	db.Exec("INSERT INTO rounds (id, start) VALUES (4, ?)", time.Now())
	db.Table("rounds").Count(&count)
	if count != 4 {
		panic("There should be 4 entry in rounds")
	}
	//Creating Checks
	db.Exec("INSERT INTO checks (round_id, service_id, log, passed) VALUES (1, '11111111-1111-1111-1111-111111111111', '',true)")
	db.Exec("INSERT INTO checks (round_id, service_id, log, passed) VALUES (2, '22222222-2222-2222-2222-222222222222', '',true)")
	db.Exec("INSERT INTO checks (round_id, service_id, log, passed) VALUES (3, '11111111-1111-1111-1111-111111111111', 'Failed because of incorrect password',false)")
	db.Exec("INSERT INTO checks (round_id, service_id, log, passed) VALUES (3, '33333333-3333-3333-3333-333333333333', '',true)")
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
