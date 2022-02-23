package util

import (
	"errors"
	"fmt"
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

var ErrTimeDifferenceTooLarge = errors.New("time difference between master host, and database host are is large. The difference should not exceed 2 seconds")

// DatabaseOutOfSync ensures that drift between database is not larger than DatabaseMaxTimeDriftSeconds
func DatabaseOutOfSync(dbTime time.Time, config config.StaticConfig) error {
	timeDiff := time.Since(dbTime)
	if float64(time.Second*time.Duration(config.DatabaseMaxTimeDriftSeconds)) < math.Abs(float64(timeDiff)) {
		return fmt.Errorf("%w: Time on database:%s, Time on master:%s", ErrTimeDifferenceTooLarge, dbTime.String(), time.Now())
	}
	return nil
}

// CreateAllTables migrates all tables
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
