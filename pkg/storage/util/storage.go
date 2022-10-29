package util

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage"
	"log"
	"math"
	"time"

	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/check/checkrepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/config/configrepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/host"
	"github.com/ScoreTrak/ScoreTrak/pkg/host/hostrepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/hostgroup"
	"github.com/ScoreTrak/ScoreTrak/pkg/hostgroup/hostgrouprepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy/policyrepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/property"
	"github.com/ScoreTrak/ScoreTrak/pkg/property/propertyrepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	"github.com/ScoreTrak/ScoreTrak/pkg/report/reportrepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/round"
	"github.com/ScoreTrak/ScoreTrak/pkg/round/roundrepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/service/servicerepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/servicegroup"
	"github.com/ScoreTrak/ScoreTrak/pkg/servicegroup/servicegrouprepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/team"
	"github.com/ScoreTrak/ScoreTrak/pkg/team/teamrepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/user"
	"github.com/ScoreTrak/ScoreTrak/pkg/user/userrepo"
	"github.com/jackc/pgconn"
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
func AutoMigrate(db *gorm.DB, storageConfig storage.Config) error {
	if storageConfig.AutoMigrate {
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
		err = db.AutoMigrate(&hostgroup.HostGroup{})
		if err != nil {
			return err
		}
		err = db.AutoMigrate(&servicegroup.ServiceGroup{})
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
	db.Table("report").Count(&count)
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
