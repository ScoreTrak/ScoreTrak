package util

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/host"
	"github.com/ScoreTrak/ScoreTrak/pkg/hostgroup"
	"github.com/ScoreTrak/ScoreTrak/pkg/property"
	"github.com/ScoreTrak/ScoreTrak/pkg/round"
	"github.com/ScoreTrak/ScoreTrak/pkg/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/servicegroup"
	"github.com/ScoreTrak/ScoreTrak/pkg/team"
	"github.com/ScoreTrak/ScoreTrak/pkg/user"
	"log"
	"math"
	"time"

	"github.com/ScoreTrak/ScoreTrak/pkg/check/checkrepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/config/configrepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/host/hostrepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/hostgroup/hostgrouprepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy/policyrepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/property/propertyrepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	"github.com/ScoreTrak/ScoreTrak/pkg/report/reportrepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/round/roundrepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/service/servicerepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/servicegroup/servicegrouprepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/team/teamrepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/user/userrepo"
	"gorm.io/gorm"
)

// Store is a single collection of all Repositories
type Store struct {
	Round        roundrepo.Repo
	Host         hostrepo.Repo
	HostGroup    hostgrouprepo.Repo
	Service      servicerepo.Repo
	ServiceGroup servicegrouprepo.Repo
	Team         teamrepo.Repo
	Check        checkrepo.Repo
	Property     propertyrepo.Repo
	Config       configrepo.Repo
	Report       reportrepo.Repo
	Policy       policyrepo.Repo
	Users        userrepo.Repo
}

// CheckDBTimeSync gets the current time reported by the database and return an error if out of sync
func CheckDBTimeSync(db *gorm.DB, staticConfig config.StaticConfig) error {
	if staticConfig.DB.Use != "sqlite" {
		var tm time.Time
		res, err := db.Raw("SELECT current_timestamp;").Rows()
		if err != nil || res.Err() != nil {
			return err
		}

		defer func(res *sql.Rows) {
			err := res.Close()
			if err != nil {
				log.Fatalln(fmt.Errorf("unable to close the database connection properly: %w", err))
			}
		}(res)

		for res.Next() {
			err := res.Scan(&tm)
			if err != nil {
				return err
			}
		}

		err = DatabaseOutOfSync(tm, staticConfig)
		return err
	}
	return nil
}

var ErrTimeDifferenceTooLarge = errors.New("time difference between master host, and database host are is large. The difference should not exceed 2 seconds")

// DatabaseOutOfSync ensures that drift between database is not larger than DatabaseMaxTimeDriftSeconds
func DatabaseOutOfSync(dbTime time.Time, config config.StaticConfig) error {
	timeDiff := time.Since(dbTime)
	if float64(time.Second*time.Duration(config.DatabaseMaxTimeDriftSeconds)) < math.Abs(float64(timeDiff)) {
		return fmt.Errorf("%w: Time on database:%s, Time on master:%s", ErrTimeDifferenceTooLarge, dbTime.String(), time.Now())
	}
	return nil
}

// AutoMigrate migrates all tables
func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&team.Team{}, &user.User{}, &policy.Policy{}, &report.Report{}, &config.DynamicConfig{}, &hostgroup.HostGroup{}, &servicegroup.ServiceGroup{}, &host.Host{}, &round.Round{}, &service.Service{}, &check.Check{}, &property.Property{})
	if err != nil {
		return err
	}
	return nil
}

func NewRepoStore(roundrepo roundrepo.Repo, hostrepo hostrepo.Repo, hostgrouprepo hostgrouprepo.Repo, servicerepo servicerepo.Repo, servicegrouprepo servicegrouprepo.Repo, teamrepo teamrepo.Repo, checkrepo checkrepo.Repo, propertyrepo propertyrepo.Repo, configrepo configrepo.Repo, reportrepo reportrepo.Repo, policyrepo policyrepo.Repo, userrepo userrepo.Repo) *Store {
	return &Store{
		Round:        roundrepo,
		Host:         hostrepo,
		HostGroup:    hostgrouprepo,
		Service:      servicerepo,
		ServiceGroup: servicegrouprepo,
		Team:         teamrepo,
		Check:        checkrepo,
		Property:     propertyrepo,
		Config:       configrepo,
		Report:       reportrepo,
		Policy:       policyrepo,
		Users:        userrepo,
	}
}
