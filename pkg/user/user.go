package user

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	Black     = "black"
	Blue      = "blue"
	Red       = "red"
	Anonymous = ""
)

type User struct {
	ID           uuid.UUID `json:"id,omitempty" gorm:"type:uuid;primary_key;"`
	Username     string    `json:"username" gorm:"unique;not null;default:null"`
	PasswordHash string    `json:"password_hash" gorm:"not null;default:null"`
	TeamID       uuid.UUID `json:"team_id,omitempty" gorm:"type:uuid"`
	Role         string    `json:"role" gorm:"default:'blue'"`
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if u.Username != "" && !govalidator.IsAlpha(u.Username) {
		return errors.New("field Name must be alphanumeric")
	}
	if u.Role != "" {
		var validStatus bool
		for _, item := range []string{Black, Blue, Red} {
			if item == u.Role {
				validStatus = true
			}
		}
		if !validStatus {
			return errors.New("you must specify a correct role")
		}
		return nil

	}
	return nil
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		uid, err := uuid.NewV4()
		if err != nil {
			return err
		}
		u.ID = uid
	}
	return nil
}

func (u *User) IsCorrectPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}
