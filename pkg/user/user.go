package user

import (
	"errors"
	"github.com/ScoreTrak/ScoreTrak/pkg/role"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID `json:"id,omitempty" gorm:"type:uuid;primary_key;"`
	Username     string    `json:"username" gorm:"unique,not null;default:null" valid:"required,alphanum"`
	PasswordHash string    `json:"password_hash" gorm:"not null;default: null"`
	TeamID       uuid.UUID `json:"team_id,omitempty" gorm:"type:uuid"`
	Role         string    `json:"role" gorm:"default:'blue'"`
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

func (u User) Validate(db *gorm.DB) {
	if u.Role != "" && u.Role != role.Black && u.Role != role.Blue {
		db.AddError(errors.New("you must specify a correct role"))
	}
}

func (u *User) IsCorrectPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}
