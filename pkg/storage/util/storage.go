package util

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math"
	"os/user"
	"time"

	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/competition"
	"github.com/ScoreTrak/ScoreTrak/pkg/competition/competitionrepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/competitionsettings"
	"github.com/ScoreTrak/ScoreTrak/pkg/competitionsettings/competitionsettingsrepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/host"
	"github.com/ScoreTrak/ScoreTrak/pkg/hostgroup"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy"
	"github.com/ScoreTrak/ScoreTrak/pkg/property"
	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	"github.com/ScoreTrak/ScoreTrak/pkg/round"
	"github.com/ScoreTrak/ScoreTrak/pkg/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/team"
	"github.com/ScoreTrak/ScoreTrak/pkg/workergroup"
	"github.com/spf13/viper"

	"github.com/ScoreTrak/ScoreTrak/pkg/check/checkrepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/host/hostrepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/hostgroup/hostgrouprepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy/policyrepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/property/propertyrepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/report/reportrepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/round/roundrepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/service/servicerepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/team/teamrepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/user/userrepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/workergroup/workergrouprepo"
	"gorm.io/gorm"
)

// Store is a single collection of all Repositories
type Store struct {
	Competition        competitionrepo.Repo
	CompetitionSetting competitionsettingsrepo.Repo
	Round              roundrepo.Repo
	Host               hostrepo.Repo
	HostGroup          hostgrouprepo.Repo
	Service            servicerepo.Repo
	WorkerGroup        workergrouprepo.Repo
	Team               teamrepo.Repo
	Check              checkrepo.Repo
	Property           propertyrepo.Repo
	Report             reportrepo.Repo
	Policy             policyrepo.Repo
	Users              userrepo.Repo
}

// CheckDBTimeSync gets the current time reported by the database and return an error if out of sync
func CheckDBTimeSync(db *gorm.DB, staticConfig config.Config) error {
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
func DatabaseOutOfSync(dbTime time.Time, c config.Config) error {
	if c.DB.Use != "sqlite" {
		timeDiff := time.Since(dbTime)
		if float64(time.Second*time.Duration(c.DatabaseMaxTimeDriftSeconds)) < math.Abs(float64(timeDiff)) {
			return fmt.Errorf("%w: Time on database:%s, Time on master:%s", ErrTimeDifferenceTooLarge, dbTime.String(), time.Now())
		}
	}
	return nil
}

// AutoMigrate migrates all tables
func AutoMigrate(db *gorm.DB) error {
	log.Println("Starting Migration")
	err := db.AutoMigrate(
		&user.User{},
		&team.Team{},
		&competition.Competition{},
		&policy.Policy{},
		&report.Report{},
		&hostgroup.HostGroup{},
		&workergroup.WorkerGroup{},
		&host.Host{},
		&round.Round{},
		&service.Service{},
		&check.Check{},
		&property.Property{},
		&competitionsettings.CompetitionSettings{},
	)
	if err != nil {
		return err
	}
	// err := db.AutoMigrate(&user.User{})
	// if err != nil {
	// 	return err
	// }
	// err = db.AutoMigrate(&team.Team{})
	// if err != nil {
	// 	return err
	// }
	// err = db.AutoMigrate(&competition.Competition{})
	// if err != nil {
	// 	return err
	// }
	// err = db.AutoMigrate(&policy.Policy{})
	// if err != nil {
	// 	return err
	// }
	// err = db.AutoMigrate(&report.Report{})
	// if err != nil {
	// 	return err
	// }
	// err = db.AutoMigrate(&config.DynamicConfig{})
	// if err != nil {
	// 	return err
	// }
	// err = db.AutoMigrate(&hostgroup.HostGroup{})
	// if err != nil {
	// 	return err
	// }
	// err = db.AutoMigrate(&workergroup.WorkerGroup{})
	// if err != nil {
	// 	return err
	// }
	// err = db.AutoMigrate(&host.Host{})
	// if err != nil {
	// 	return err
	// }
	// err = db.AutoMigrate(&round.Round{})
	// if err != nil {
	// 	return err
	// }
	// err = db.AutoMigrate(&service.Service{})
	// if err != nil {
	// 	return err
	// }
	// err = db.AutoMigrate(&check.Check{})
	// if err != nil {
	// 	return err
	// }
	// err = db.AutoMigrate(&property.Property{})
	// if err != nil {
	// 	return err
	// }
	log.Println("Migration Completed")
	return nil
}

func NewRepoStore(roundrepo roundrepo.Repo, hostrepo hostrepo.Repo, hostgrouprepo hostgrouprepo.Repo, servicerepo servicerepo.Repo, workergrouprepo workergrouprepo.Repo, teamrepo teamrepo.Repo, checkrepo checkrepo.Repo, propertyrepo propertyrepo.Repo, reportrepo reportrepo.Repo, policyrepo policyrepo.Repo, userrepo userrepo.Repo, competitionsettingrepo competitionsettingsrepo.Repo) *Store {
	return &Store{
		Round:              roundrepo,
		Host:               hostrepo,
		HostGroup:          hostgrouprepo,
		Service:            servicerepo,
		WorkerGroup:        workergrouprepo,
		Team:               teamrepo,
		Check:              checkrepo,
		Property:           propertyrepo,
		CompetitionSetting: competitionsettingrepo,
		Report:             reportrepo,
		Policy:             policyrepo,
		Users:              userrepo,
	}
}

func TruncateTable(ctx context.Context, v interface{}, db *gorm.DB) error {
	dbType := viper.GetString("db.use")
	stmt := &gorm.Statement{DB: db}
	err := stmt.Parse(v)
	if err != nil {
		return err
	}
	if dbType == "mysql" {
		// NOT DESIRABLE as there will be data that will break the foreign key constraints.
		// However, this command is only called with others table are being turncated as well :)
		return db.WithContext(ctx).Exec(fmt.Sprintf("SET FOREIGN_KEY_CHECKS=0 ; TRUNCATE %s ; SET FOREIGN_KEY_CHECKS=1", stmt.Schema.Table)).Error
	} else if dbType == "postgresql" || dbType == "cockroach" {
		return db.WithContext(ctx).Exec(fmt.Sprintf("TRUNCATE TABLE %s CASCADE", stmt.Schema.Table)).Error
	} else if dbType == "sqlite3" {
		return db.WithContext(ctx).Exec(fmt.Sprintf("DELETE FROM %s", stmt.Schema.Table)).Error
	}
	// return storage.ErrDBNotSupported
	return errors.New("DB not supported")
}
