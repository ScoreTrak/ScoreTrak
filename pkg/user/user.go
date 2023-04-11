package user

import (
	"errors"
	"time"

	"github.com/ScoreTrak/ScoreTrak/pkg/team"
	"github.com/asaskevich/govalidator"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID    `json:"id,omitempty" gorm:"type:uuid;primary_key;"`
	Username  string       `json:"username" gorm:"unique;not null;default:null"`
	Teams     []*team.Team `gorm:"many2many:user_teams"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

var ErrNameMustBeAlphanumeric = errors.New("name must be alphanumeric")
var ErrInvalidRoleSpecified = errors.New("incorrect role specified")

// BeforeSave ensures that user is part of either Black, Blue, or Red roles. It also ensures that username is alphanumeric
func (u *User) BeforeSave(_ *gorm.DB) (err error) {
	if u.Username != "" && !govalidator.IsAlphanumeric(u.Username) {
		return ErrNameMustBeAlphanumeric
	}
	// if u.Role != "" {
	// 	var validStatus bool
	// 	for _, item := range []string{Black, Blue, Red} {
	// 		if item == u.Role {
	// 			validStatus = true
	// 		}
	// 	}
	// 	if !validStatus {
	// 		return ErrInvalidRoleSpecified
	// 	}
	// 	return nil
	// }
	return nil
}

// BeforeCreate ensures UUID is set.
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

// IsCorrectPassword compares password to the hash
func (u *User) IsCorrectPassword(password string) bool {
	// err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	// return err == nil
	return false
}
