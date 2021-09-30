package testutil

import (
	"context"
	"errors"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/host"
	"github.com/ScoreTrak/ScoreTrak/pkg/host_group"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy"
	"github.com/ScoreTrak/ScoreTrak/pkg/property"
	"github.com/ScoreTrak/ScoreTrak/pkg/round"
	"github.com/ScoreTrak/ScoreTrak/pkg/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/service_group"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage"
	"github.com/ScoreTrak/ScoreTrak/pkg/team"
	"github.com/ScoreTrak/ScoreTrak/pkg/user"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgconn"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

//CleanAllTables Drops all tables
func CleanAllTables(db *gorm.DB) error {
	err := db.Migrator().DropTable(&check.Check{})
	if err != nil {
		return err
	}
	err = db.Migrator().DropTable(&property.Property{})
	if err != nil {
		return err
	}
	err = db.Migrator().DropTable(&service.Service{})
	if err != nil {
		return err
	}
	err = db.Migrator().DropTable(&host.Host{})
	if err != nil {
		return err
	}
	err = db.Migrator().DropTable(&host_group.HostGroup{})
	if err != nil {
		return err
	}
	err = db.Migrator().DropTable(&round.Round{})
	if err != nil {
		return err
	}
	err = db.Migrator().DropTable(&service_group.ServiceGroup{})
	if err != nil {
		return err
	}
	err = db.Migrator().DropTable(&user.User{})
	if err != nil {
		return err
	}
	err = db.Migrator().DropTable(&team.Team{})
	if err != nil {
		return err
	}
	err = db.Migrator().DropTable(&config.DynamicConfig{})
	if err != nil {
		return err
	}
	return nil
}

//uuid1 is a uuid of the initial admin user, and admin team
var uuid1 = uuid.FromStringOrNil("00000000-0000-0000-0000-000000000001")

//CreateBlackTeam Ensures that black team exists
func CreateBlackTeam(db *gorm.DB) (err error) {
	err = db.Create([]*team.Team{{ID: uuid1, Name: "Black Team"}}).Error
	if err != nil {
		var serr *pgconn.PgError
		ok := errors.As(err, &serr)
		if !ok || serr.Code != "23505" {
			return err
		}
	}
	return nil
}

//CreateAdminUser ensures Admin user exists with the default password of changeme
func CreateAdminUser(db *gorm.DB) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("changeme"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	err = db.Create([]*user.User{{ID: uuid1, TeamID: uuid1, Username: "admin", Role: user.Black, PasswordHash: string(hashedPassword)}}).Error
	if err != nil {
		var serr *pgconn.PgError
		ok := errors.As(err, &serr)
		if !ok || serr.Code != "23505" {
			return err
		}
	}
	return nil
}

//CreatePolicy Ensure policy.Policy exists
func CreatePolicy(db *gorm.DB) (*policy.Policy, error) {
	p := &policy.Policy{ID: 1}
	err := db.Create(p).Error
	if err != nil {
		var serr *pgconn.PgError
		ok := errors.As(err, &serr)
		if !ok {
			if serr.Code != "23505" {
				panic(err)
			} else {
				db.Take(p)
			}
		}
	}
	return p, nil
}

//TruncateAllTables truncates all tables (A little faster than dropping)
func TruncateAllTables(db *gorm.DB) {
	db.Exec(fmt.Sprintf("TRUNCATE TABLE %s CASCADE", "checks"))
	db.Exec(fmt.Sprintf("TRUNCATE TABLE %s CASCADE", "properties"))
	db.Exec(fmt.Sprintf("TRUNCATE TABLE %s CASCADE", "services"))
	db.Exec(fmt.Sprintf("TRUNCATE TABLE %s CASCADE", "hosts"))
	db.Exec(fmt.Sprintf("TRUNCATE TABLE %s CASCADE", "host_groups"))
	db.Exec(fmt.Sprintf("TRUNCATE TABLE %s CASCADE", "rounds"))
	db.Exec(fmt.Sprintf("TRUNCATE TABLE %s CASCADE", "service_groups"))
	db.Exec(fmt.Sprintf("TRUNCATE TABLE %s CASCADE", "users"))
	db.Exec(fmt.Sprintf("TRUNCATE TABLE %s CASCADE", "teams"))
	db.Exec(fmt.Sprintf("TRUNCATE TABLE %s CASCADE", "config"))
}

func DataPreload(db *gorm.DB) {
	var count int64
	//Creating Config
	db.Exec("INSERT INTO config (id, round_duration, enabled) VALUES (1, 60, true)")
	db.Table("config").Count(&count)
	if count != 1 {
		panic("There should be 1 entry in config")
	}
	//Creating Teams
	db.Exec("INSERT INTO teams (id, name, pause) VALUES ('11111111-1111-1111-1111-111111111111', 'TeamOne', true)")
	db.Exec("INSERT INTO teams (id, name, pause) VALUES ('22222222-2222-2222-2222-222222222222', 'TeamTwo', false)")
	db.Exec("INSERT INTO teams (id, name, pause) VALUES ('33333333-3333-3333-3333-333333333333', 'TeamThree', true)")
	db.Exec("INSERT INTO teams (id, name, pause) VALUES ('44444444-4444-4444-4444-444444444444', 'TeamFour', false)")
	db.Table("teams").Count(&count)
	if count != 4 {
		panic("There should be 4 entry in teams")
	}
	//Creating Users
	db.Exec("INSERT INTO users (team_id, id, username, role, password_hash) VALUES ('11111111-1111-1111-1111-111111111111', '11111111-1111-1111-1111-111111111111', 'TeamOneUser', 'black', '$2a$10$SDWdEajPUVYx1wczissDSur3R0/TmNsvCi/KivaKUM2DQi0Hc9EDG')")
	db.Exec("INSERT INTO users (team_id, id, username, role, password_hash) VALUES ('22222222-2222-2222-2222-222222222222', '22222222-2222-2222-2222-222222222222', 'TeamTwoUser2', 'red', '$2a$10$SDWdEajPUVYx1wczissDSur3R0/TmNsvCi/KivaKUM2DQi0Hc9EDG')")
	db.Exec("INSERT INTO users (team_id, id, username, role, password_hash) VALUES ('22222222-2222-2222-2222-222222222222', '33333333-3333-3333-3333-333333333333', 'TeamTwoUser1', 'blue', '$2a$10$SDWdEajPUVYx1wczissDSur3R0/TmNsvCi/KivaKUM2DQi0Hc9EDG')")
	db.Table("users").Count(&count)
	if count != 3 {
		panic("There should be 3 entries in users")
	}
	//Creating Host Groups
	db.Exec("INSERT INTO host_groups (id, name, pause) VALUES ('11111111-1111-1111-1111-111111111111', 'HostGroup1', true)")
	db.Exec("INSERT INTO host_groups (id, name, pause) VALUES ('22222222-2222-2222-2222-222222222222', 'HostGroup2', false)")
	db.Exec("INSERT INTO host_groups (id, name, pause) VALUES ('33333333-3333-3333-3333-333333333333', 'HostGroup3', true)")
	db.Exec("INSERT INTO host_groups (id, name, pause) VALUES ('44444444-4444-4444-4444-444444444444', 'HostGroup4', false)")
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
		panic("There should be 4 entry in check_service groups")
	}
	//Creating Hosts
	db.Exec("INSERT INTO hosts (id, address, team_id, host_group_id, pause, edit_host) VALUES ('11111111-1111-1111-1111-111111111111', '10.0.0.1', '11111111-1111-1111-1111-111111111111', NULL, true, true)")
	db.Exec("INSERT INTO hosts (id, address, team_id, host_group_id, pause, edit_host) VALUES ('22222222-2222-2222-2222-222222222222', '10.0.0.2', '22222222-2222-2222-2222-222222222222', NULL, false, true)")
	db.Exec("INSERT INTO hosts (id, address, team_id, host_group_id, pause, edit_host) VALUES ('33333333-3333-3333-3333-333333333333', '10.0.0.3', '11111111-1111-1111-1111-111111111111', '33333333-3333-3333-3333-333333333333', true, false)")
	db.Exec("INSERT INTO hosts (id, address, team_id, host_group_id, pause, edit_host) VALUES ('44444444-4444-4444-4444-444444444444', '10.0.0.4', '44444444-4444-4444-4444-444444444444', '44444444-4444-4444-4444-444444444444', false, false)")
	db.Table("hosts").Count(&count)
	if count != 4 {
		panic("There should be 4 entry in hosts")
	}
	//Creating Services
	db.Exec("INSERT INTO services (id, service_group_id, host_id, name, display_name, weight, round_units, round_delay, pause) VALUES ('11111111-1111-1111-1111-111111111111', '11111111-1111-1111-1111-111111111111', '11111111-1111-1111-1111-111111111111', 'WINRM', 'host1-service1', 0, 1, 0, true)")
	db.Exec("INSERT INTO services (id, service_group_id, host_id, name, display_name, weight, round_units, round_delay, pause) VALUES ('22222222-2222-2222-2222-222222222222', '22222222-2222-2222-2222-222222222222', '22222222-2222-2222-2222-222222222222', 'FTP', 'host2-service2', 40, 23, 2, true)")
	db.Exec("INSERT INTO services (id, service_group_id, host_id, name, display_name, weight, round_units, round_delay, pause) VALUES ('33333333-3333-3333-3333-333333333333', '44444444-4444-4444-4444-444444444444', '22222222-2222-2222-2222-222222222222', 'SSH', 'host3-service3', 50, 3, 0, true)")
	db.Exec("INSERT INTO services (id, service_group_id, host_id, name, display_name, weight, round_units, round_delay, pause) VALUES ('44444444-4444-4444-4444-444444444444', '44444444-4444-4444-4444-444444444444', '44444444-4444-4444-4444-444444444444', 'HTTP', 'host4-service4', 200, 4, 3, true)")
	db.Exec("INSERT INTO services (id, service_group_id, host_id, name, display_name, weight, round_units, round_delay, pause) VALUES ('55555555-5555-5555-5555-555555555555', '11111111-1111-1111-1111-111111111111', '22222222-2222-2222-2222-222222222222', 'SSH', 'host1-service2', 30, 5, 4, false)")
	db.Exec("INSERT INTO services (id, service_group_id, host_id, name, display_name, weight, round_units, round_delay, pause) VALUES ('66666666-6666-6666-6666-666666666666', '22222222-2222-2222-2222-222222222222', '11111111-1111-1111-1111-111111111111', 'SMB', 'host2-service1', 2, 5, 2, true)")
	db.Exec("INSERT INTO services (id, service_group_id, host_id, name, display_name, weight, round_units, round_delay, pause) VALUES ('77777777-7777-7777-7777-777777777777', '22222222-2222-2222-2222-222222222222', '44444444-4444-4444-4444-444444444444', 'FTP', 'host3-service4', 55, 6, 3, false)")
	db.Exec("INSERT INTO services (id, service_group_id, host_id, name, display_name, weight, round_units, round_delay, pause) VALUES ('88888888-8888-8888-8888-888888888888', '44444444-4444-4444-4444-444444444444', '44444444-4444-4444-4444-444444444444', 'IMAP', 'host4-service3', 44, 23, 22, false)")
	db.Table("services").Count(&count)
	if count != 8 {
		panic("There should be 8 entry in services")
	}
	//Creating Properties
	db.Exec("INSERT INTO properties (service_id, key, value, status) VALUES ( '11111111-1111-1111-1111-111111111111', 'Port','80','View')")
	db.Exec("INSERT INTO properties (service_id, key, value, status) VALUES ( '11111111-1111-1111-1111-111111111111', 'Username','root','Edit')")
	db.Exec("INSERT INTO properties (service_id, key, value, status) VALUES ( '22222222-2222-2222-2222-222222222222', 'Port','53','View')")
	db.Exec("INSERT INTO properties (service_id, key, value, status) VALUES ( '22222222-2222-2222-2222-222222222222', 'Password','changeme','Hide')")
	db.Exec("INSERT INTO properties (service_id, key, value, status) VALUES ( '22222222-2222-2222-2222-222222222222', 'Username','root','Hide')")
	db.Exec("INSERT INTO properties (service_id, key, value, status) VALUES ( '33333333-3333-3333-3333-333333333333', 'Password','changeme','Hide')")
	db.Exec("INSERT INTO properties (service_id, key, value, status) VALUES ( '44444444-4444-4444-4444-444444444444', 'Port','443','Edit')")
	db.Exec("INSERT INTO properties (service_id, key, value, status) VALUES ( '55555555-5555-5555-5555-555555555555', 'Username','admin','Edit')")
	db.Exec("INSERT INTO properties (service_id, key, value, status) VALUES ( '66666666-6666-6666-6666-666666666666', 'Password','Change.me!','View')")
	db.Exec("INSERT INTO properties (service_id, key, value, status) VALUES ( '66666666-6666-6666-6666-666666666666', 'Content','Sample Content!', 'Hide')")
	db.Exec("INSERT INTO properties (service_id, key, value, status) VALUES ( '66666666-6666-6666-6666-666666666666', 'EncodedContent', 'RW5jb2RlZENvbnRlbnQ=', 'Edit')")
	db.Exec("INSERT INTO properties (service_id, key, value, status) VALUES ( '66666666-6666-6666-6666-666666666666', 'domain', 'sample.com', 'View')")
	db.Table("properties").Count(&count)
	if count != 12 {
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
	db.Exec("INSERT INTO checks (round_id, service_id, log, passed) VALUES (1, '11111111-1111-1111-1111-111111111111', 'Successful Check One!',true)")
	db.Exec("INSERT INTO checks (round_id, service_id, log, passed) VALUES (2, '22222222-2222-2222-2222-222222222222', 'Successful Check Two!',true)")
	db.Exec("INSERT INTO checks (round_id, service_id, log, passed) VALUES (3, '11111111-1111-1111-1111-111111111111', 'Failed because of incorrect password',false)")
	db.Exec("INSERT INTO checks (round_id, service_id, log, passed) VALUES (3, '33333333-3333-3333-3333-333333333333', 'Successful Check Four!',true)")
	db.Table("checks").Count(&count)
	if count != 4 {
		panic("There should be 4 entry in checks")
	}
}

func DropDB(db *gorm.DB, c config.StaticConfig) {
	db.Exec(fmt.Sprintf("drop database %s", c.DB.Cockroach.Database))
}

func TruncateTable(ctx context.Context, v interface{}, db *gorm.DB) error {
	stmt := &gorm.Statement{DB: db}
	err := stmt.Parse(v)
	if err != nil {
		return err
	}
	return db.WithContext(ctx).Exec(fmt.Sprintf("TRUNCATE TABLE %s CASCADE", stmt.Schema.Table)).Error //POSTGRES SPECIFIC. FOR MYSQL, CHANGE THIS TO  SET FOREIGN_KEY_CHECKS=0 ; <TRUNCATE> ; SET FOREIGN_KEY_CHECKS=1
}

//SetupCockroachDB creates a new database instance using
func SetupCockroachDB(c storage.Config) *gorm.DB {
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
