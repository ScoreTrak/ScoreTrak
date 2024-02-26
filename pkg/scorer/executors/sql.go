package executors

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/scoretrak/scoretrak/pkg/scorer/outcome"
	"github.com/scoretrak/scoretrak/pkg/scorer/outcome_writer"
	"log"
	"net"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SQLProperties struct {
	Username        string
	Password        string
	Host            string
	Port            string
	DBType          string
	DBName          string
	Command         string
	MinExpectedRows string
	MaxExpectedRows string
}

var ErrSQLRequiresCommand = errors.New("sql check needs a command parameter")
var ErrUnsupportedDBType = errors.New("DBType should either be mysql, or postgres")
var ErrSQLNeedsUsernameOrPassword = errors.New("sql check_service needs username, and password")

//func (w *SQL) Validate() error {
//	if w.Password == "" || w.Username == "" {
//		return ErrSQLNeedsUsernameOrPassword
//	}
//	if strings.ToLower(w.DBType) != "mysql" && strings.ToLower(w.DBType) != "postgres" {
//		return ErrUnsupportedDBType
//	}
//	if w.Command == "" {
//		return ErrSQLRequiresCommand
//	}
//
//	if w.MaxExpectedRows != "" {
//		_, err := strconv.ParseUint(w.MaxExpectedRows, 10, 64)
//		if err != nil {
//			return err
//		}
//	}
//
//	if w.MinExpectedRows != "" {
//		_, err := strconv.ParseUint(w.MinExpectedRows, 10, 64)
//		if err != nil {
//			return err
//		}
//	}
//
//	return nil
//}

func setupDB(ctx context.Context, properties *SQLProperties) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	if properties.DBType == "mysql" {
		if properties.Port == "" {
			properties.Port = "3306"
		}
		timeout, ok := ctx.Deadline()
		if ok {
			err = tcpPortDial(net.JoinHostPort(properties.Host, properties.Port), time.Until(timeout)/3)
			if err != nil {
				return nil, err
			}
		}
		var dsn string
		if properties.DBName != "" {
			dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", properties.Username, properties.Password, properties.Host, properties.Port, properties.DBName)
		} else {
			dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)", properties.Username, properties.Password, properties.Host, properties.Port)
		}
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("unable to initialize mysql client: %w", err)
		}
	}

	if properties.DBType == "postgres" {
		if properties.Port == "" {
			properties.Port = "5432"
		}
		timeout, ok := ctx.Deadline()
		if ok {
			err = tcpPortDial(net.JoinHostPort(properties.Host, properties.Port), time.Until(timeout)/3)
			if err != nil {
				return nil, err
			}
		}
		var dsn string
		if properties.DBName != "" {
			dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				properties.Host,
				properties.Port,
				properties.Username,
				properties.Password,
				properties.DBName)
		} else {
			dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable",
				properties.Host,
				properties.Port,
				properties.Username,
				properties.Password)
		}
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("unable to initialize postgres client: %w", err)
		}
	}
	return db, nil
}

func ScoreSQL(ctx context.Context, ow *outcome_writer.OutcomeWriter, properties []byte) {
	sqlproperties := &SQLProperties{}
	err := json.Unmarshal(properties, &sqlproperties)
	if err != nil {
		ow.SetError(fmt.Errorf("unable to unmarshall properties: %v", err))
		return
	}

	//err = validate.Struct(sqlproperties)
	//if err != nil {
	//	ow.SetError(fmt.Errorf("validation error: %w", err))
	//	return
	//}

	db, err := setupDB(ctx, sqlproperties)
	if err != nil {
		ow.SetError(err)
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		ow.SetError(fmt.Errorf("unable to fetch the underlying sql driver, this is most likely a bug: %w", err))
		return
	}
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {
			log.Println(fmt.Errorf("unable to close sql connection: %w", err))
		}
	}(sqlDB)

	result := db.WithContext(ctx).Raw(sqlproperties.Command)

	if result.Error != nil {
		ow.SetError(fmt.Errorf("unable to execute the provided command: %w", result.Error))
		return
	}

	rows, err := result.Rows()
	defer func(rows *sql.Rows) {
		if rows != nil {
			err := rows.Close()
			if err != nil {
				log.Println(fmt.Errorf("failed to close rows: %w", err))
			}
		}
	}(rows)
	if err != nil {
		ow.SetError(fmt.Errorf("encountered a problem when retrieving rows: %w", err))
		return
	}
	if rows.Err() != nil {
		ow.SetError(fmt.Errorf("encountered a problem during row iteration: %w", err))
		return
	}

	var count int64
	for rows.Next() {
		count++
	}

	if sqlproperties.MaxExpectedRows != "" {
		num, _ := strconv.ParseInt(sqlproperties.MaxExpectedRows, 10, 64)
		if num < count {
			ow.SetError(fmt.Errorf("%w (%d < %d)", ErrNumberOfRowsMoreThanExpected, num, count))
			return
		}
	}

	if sqlproperties.MinExpectedRows != "" {
		num, _ := strconv.ParseInt(sqlproperties.MinExpectedRows, 10, 64)
		if num > count {
			ow.SetError(fmt.Errorf("%w: (%d > %d)", ErrNumberOfRowsLessThanExpected, num, count))
			return
		}
	}
	ow.SetStatus(outcome.OUTCOME_STATUS_PASSED)
}

var ErrNumberOfRowsLessThanExpected = errors.New("number of rows was less than expected")
var ErrNumberOfRowsMoreThanExpected = errors.New("number of rows was more than expected")

// TODO(thisisibrahimd): Implement Content Check
