package models

import (
	"encoding/json"
	"fmt"
	"scoretrak/constants"
	"strings"
	"time"

	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

//User is a generated model from buffalo-auth, it serves as the base for username/password authentication.
type User struct {
	ID                   uuid.UUID `json:"id" db:"id"`
	CreatedAt            time.Time `json:"created_at" db:"created_at"`
	UpdatedAt            time.Time `json:"updated_at" db:"updated_at"`
	Username             string    `json:"username" db:"username"`
	PasswordHash         string    `json:"password_hash" db:"password_hash"`
	TeamID               uuid.UUID `json:"-" db:"team_id"`
	Team                 *Team     `json:"team" belongs_to:"team"`
	Password             string    `json:"-" db:"-"`
	PasswordConfirmation string    `json:"-" db:"-"`
}

// Create wraps up the pattern of encrypting the password and
// running validations. Useful when writing tests.
func (u *User) Create(tx *pop.Connection) (*validate.Errors, error) {
	if err := u.generateHash(); err != nil {
		return validate.NewErrors(), errors.WithStack(err)
	}
	return tx.ValidateAndCreate(u)
}

func (u *User) Update(tx *pop.Connection) (*validate.Errors, error) {
	if u.Password != "" {
		if err := u.generateHash(); err != nil {
			return validate.NewErrors(), errors.WithStack(err)
		}
	}
	return tx.ValidateAndUpdate(u)
}

func (u *User) generateHash() error {
	u.Username = strings.ToLower(u.Username)
	ph, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(ph)
	return nil
}

// String is not required by pop and may be deleted
func (u User) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Users is not required by pop and may be deleted
type Users []User

// String is not required by pop and may be deleted
func (u Users) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (u *User) Validate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	return validate.Validate(
		&validators.StringIsPresent{Field: u.Username, Name: "Username"},
		&validators.StringIsPresent{Field: u.PasswordHash, Name: "PasswordHash"},
		&validators.UUIDIsPresent{Field: u.TeamID, Name: "TeamID"},
		// check to see if the username address is already taken:
		&validators.FuncValidator{
			Field:   u.Username,
			Name:    "Username",
			Message: "%s is already taken",
			Fn: func() bool {
				var b bool
				q := tx.Where("username = ?", u.Username)
				if u.ID != uuid.Nil {
					q = q.Where("id != ?", u.ID)
				}
				b, err = q.Exists(u)
				if err != nil {
					return false
				}
				return !b
			},
		},

		&validators.FuncValidator{
			Name:    "TeamID",
			Message: "Provided TeamID does not exist",
			Fn: func() bool {
				var t Team
				err := tx.Find(&t, u.TeamID)
				if err != nil {
					return false
				}
				return true
			},
		},
	), err
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *User) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	return validate.Validate(
		&validators.StringIsPresent{Field: u.Password, Name: "Password"},
		&validators.StringsMatch{Name: "Password", Field: u.Password, Field2: u.PasswordConfirmation, Message: "Password does not match confirmation"},
	), err
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (u *User) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	return validate.Validate(
		&validators.FuncValidator{
			Name:    "TeamID",
			Message: fmt.Sprintf("You cannot reassign last user in last team with role \"%s\"", constants.Black),
			Fn: func() bool {
				teams := []Team{}
				user, err := GetUserByID(tx, u.ID)
				if user.TeamID != u.TeamID {
					//Query for all teams with role constants.Black
					if err = tx.Where("role = (?)", constants.Black).Eager().All(&teams); err != nil {
						return false
					}
					//Disallow deletion of the last constants.Black team
					if len(teams) == 1 && len(teams[0].Users) == 1 && teams[0].Users[0].ID == u.ID {
						return false
					}
				}
				return true
			},
		},
	), err
}

//GetUserByUsername retreives User object that matches the username parameter
func GetUserByUsername(tx *pop.Connection, n string) (User, error) {
	u := []User{}
	query := tx.Where("username = (?)", n)
	err := query.All(&u)
	if err != nil {
		return User{}, err
	}
	return u[0], err
}

func GetUserByID(tx *pop.Connection, id uuid.UUID) (User, error) {
	u := User{}
	err := tx.Find(&u, id)
	if err != nil {
		return User{}, err
	}
	return u, err
}
