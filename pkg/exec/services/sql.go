package services

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/ScoreTrak/ScoreTrak/pkg/exec"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SQL struct {
	Username        string
	Password        string
	Port            string
	DBType          string
	DBName          string
	Command         string
	MinExpectedRows string
	MaxExpectedRows string
}

func NewSQL() *SQL {
	return &SQL{}
}

var ErrSQLRequiresCommand = errors.New("sql check needs a command parameter")
var ErrUnsupportedDBType = errors.New("DBType should either be mysql, or postgres")
var ErrSQLNeedsUsernameOrPassword = errors.New("sql check_service needs username, and password")

func (w *SQL) Validate() error {
	if w.Password == "" || w.Username == "" {
		return ErrSQLNeedsUsernameOrPassword
	}
	if strings.ToLower(w.DBType) != "mysql" && strings.ToLower(w.DBType) != "postgres" {
		return ErrUnsupportedDBType
	}
	if w.Command == "" {
		return ErrSQLRequiresCommand
	}

	if w.MaxExpectedRows != "" {
		_, err := strconv.ParseUint(w.MaxExpectedRows, 10, 64)
		if err != nil {
			return err
		}
	}

	if w.MinExpectedRows != "" {
		_, err := strconv.ParseUint(w.MinExpectedRows, 10, 64)
		if err != nil {
			return err
		}
	}

	return nil
}

func (w *SQL) Execute(e exec.Exec) (passed bool, logOutput string, err error) {
	var db *gorm.DB
	if w.DBType == "mysql" {
		if w.Port == "" {
			w.Port = "3306"
		}
		var dsn string
		if w.DBName != "" {
			dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", w.Username, w.Password, e.Host, w.Port, w.DBName)
		} else {
			dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)", w.Username, w.Password, e.Host, w.Port)
		}
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return false, "", fmt.Errorf("unable to initialize mysql client: %w", err)
		}
	}

	if w.DBType == "postgres" {
		if w.Port == "" {
			w.Port = "5432"
		}
		var dsn string
		if w.DBName != "" {
			dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				e.Host,
				w.Port,
				w.Username,
				w.Password,
				w.DBName)
		} else {
			dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable",
				e.Host,
				w.Port,
				w.Username,
				w.Password)
		}
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return false, "", fmt.Errorf("unable to initialize postgres client: %w", err)
		}
	}
	sqlDB, err := db.DB()
	if err != nil {
		return false, "", fmt.Errorf("unable to fetch the underlying sql driver, this is most likely a bug: %w", err)
	}
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {
			log.Println(fmt.Errorf("unable to close sql connection: %w", err))
		}
	}(sqlDB)

	result := db.WithContext(e.Context).Raw(w.Command)

	if result.Error != nil {
		return false, "", fmt.Errorf("unable to execute the provided command: %w", result.Error)
	}

	rows, err := result.Rows()
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println(fmt.Errorf("failed to close rows: %w", err))
		}
	}(rows)
	if err != nil {
		return false, "", fmt.Errorf("encountered a problem when retrieving rows: %w", err)
	}
	if rows.Err() != nil {
		return false, "", fmt.Errorf("encountered a problem during row iteration: %w", err)
	}

	var count int64
	for rows.Next() {
		count++
	}

	if w.MaxExpectedRows != "" {
		num, _ := strconv.ParseInt(w.MaxExpectedRows, 10, 64)
		if num < count {
			return false, "", fmt.Errorf("%w (%d < %d)", ErrNumberOfRowsMoreThanExpected, num, count)
		}
	}

	if w.MinExpectedRows != "" {
		num, _ := strconv.ParseInt(w.MinExpectedRows, 10, 64)
		if num > count {
			return false, "", fmt.Errorf("%w: (%d > %d)", ErrNumberOfRowsLessThanExpected, num, count)
		}
	}
	return true, Success, nil
}

var ErrNumberOfRowsLessThanExpected = errors.New("number of rows was less than expected")
var ErrNumberOfRowsMoreThanExpected = errors.New("number of rows was more than expected")

// Todo: Implement Content Check
