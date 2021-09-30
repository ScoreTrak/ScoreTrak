package util

import (
	"errors"

	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/check/check_repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/config/config_repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/host"
	"github.com/ScoreTrak/ScoreTrak/pkg/host/host_repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/host_group"
	"github.com/ScoreTrak/ScoreTrak/pkg/host_group/host_group_repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy/policy_repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/property"
	"github.com/ScoreTrak/ScoreTrak/pkg/property/property_repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	"github.com/ScoreTrak/ScoreTrak/pkg/report/report_repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/round"
	"github.com/ScoreTrak/ScoreTrak/pkg/round/round_repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/service/service_repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/service_group"
	"github.com/ScoreTrak/ScoreTrak/pkg/service_group/service_group_repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/team"
	"github.com/ScoreTrak/ScoreTrak/pkg/team/team_repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/user"
	"github.com/ScoreTrak/ScoreTrak/pkg/user/user_repo"
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

//Store is a single collection of all Repositories
type Store struct {
	Round        round_repo.Repo
	Host         host_repo.Repo
	HostGroup    host_group_repo.Repo
	Service      service_repo.Repo
	ServiceGroup service_group_repo.Repo
	Team         team_repo.Repo
	Check        check_repo.Repo
	Property     property_repo.Repo
	Config       config_repo.Repo
	Report       report_repo.Repo
	Policy       policy_repo.Repo
	Users        user_repo.Repo
}

//CreateAllTables migrates all tables
func CreateAllTables(db *gorm.DB) error {
	err := db.AutoMigrate(&team.Team{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&user.User{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&policy.Policy{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&report.Report{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&config.DynamicConfig{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&host_group.HostGroup{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&service_group.ServiceGroup{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&host.Host{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&round.Round{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&service.Service{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&check.Check{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&property.Property{})
	if err != nil {
		return err
	}
	return nil
}

func LoadConfig(db *gorm.DB, cnf *config.DynamicConfig) error {
	var count int64
	db.Table("config").Count(&count)
	if count != 1 {
		err := db.Create(cnf).Error
		if err != nil {
			var serr *pgconn.PgError
			ok := errors.As(err, &serr)
			if ok && serr.Code == "23505" {
				dcc := &config.DynamicConfig{}
				db.Take(dcc)
				*cnf = *dcc
			} else {
				return err
			}
		}
	}
	return nil
}

func LoadReport(db *gorm.DB) error {
	var count int64
	if count != 1 {
		err := db.Create(report.NewReport()).Error
		if err != nil {
			var serr *pgconn.PgError
			ok := errors.As(err, &serr)
			if !ok || serr.Code != "23505" {
				return err
			}
		}
	}
	return nil
}
