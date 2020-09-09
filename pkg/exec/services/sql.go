package services

import (
	"errors"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/exec"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type Sql struct {
	Username        string
	Password        string
	Port            string
	DBType          string
	DBName          string
	Command         string
	MinExpectedRows string
	MaxExpectedRows string
}

func NewSql() *Sql {
	return &Sql{}
}

func (w *Sql) Validate() error {
	if w.Password == "" || w.Username == "" {
		return errors.New("sql service needs username, and password")
	}
	if strings.ToLower(w.DBType) != "mysql" && strings.ToLower(w.DBType) != "postgres" {
		return errors.New("DBType should either be mysql, or postgres")
	}
	if w.Command == "" {
		return errors.New("sql check needs a command parameter")
	}

	if w.MaxExpectedRows != "" {
		_, err := strconv.ParseUint(w.MaxExpectedRows, 10, 32)
		if err != nil {
			return err
		}
	}

	if w.MinExpectedRows != "" {
		_, err := strconv.ParseUint(w.MinExpectedRows, 10, 32)
		if err != nil {
			return err
		}
	}

	return nil
}

func (w *Sql) Execute(e exec.Exec) (passed bool, log string, err error) {
	var db *gorm.DB
	if w.DBType == "mysql" {
		var dsn string
		if w.DBName != "" {
			dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", w.Username, w.Password, e.Host, w.Port, w.DBName)
		} else {
			dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)", w.Username, w.Password, e.Host, w.Port)
		}
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return false, "Unable to initialize the database client", err
		}
	}

	if w.DBType == "postgres" {
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
			return false, "Unable to initialize the database client", err
		}
	}
	sqlDB, err := db.DB()
	if err != nil {
		return false, "unable to fetch the underlying sql driver, this is most likely a bug", err
	}
	defer sqlDB.Close()

	result := db.WithContext(e.Context).Raw(w.Command)

	if result.Error != nil {
		return false, "Unable to execute the provided command", result.Error
	}

	rows, err := result.Rows()
	if err != nil {
		return false, "Encountered a problem when retrieving Rows", err
	}

	var count int64
	for rows.Next() {
		count++
	}

	if w.MaxExpectedRows != "" {
		num, _ := strconv.ParseInt(w.MaxExpectedRows, 10, 64)
		result.Count(&count)
		if num <= count {
			return false, fmt.Sprintf("Number of rows was more than expected (%d <= %d)", num, count), nil
		}
	}

	if w.MinExpectedRows != "" {
		num, _ := strconv.ParseInt(w.MinExpectedRows, 10, 64)
		result.Count(&count)
		if num >= count {
			return false, fmt.Sprintf("Number of rows was less than expected (%d >= %d)", num, count), nil
		}
	}
	return true, "Success!", nil
}
