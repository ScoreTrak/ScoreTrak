package models

import (
	"encoding/json"
	"time"
	"fmt"
	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
	"scoretrak/constants"
)

// Team is used by pop to map your teams database table to your go code.
type Team struct {
	ID        uuid.UUID       `json:"id" db:"id"`
	Name      string          `json:"name" db:"name"`
	Image     nulls.ByteSlice `json:"image" db:"image"`
	Role      string          `json:"role" db:"role"`
	Users     []User          `json:"users,omitempty" has_many:"users"`
	CreatedAt time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt time.Time       `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (t Team) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Teams is not required by pop and may be deleted
type Teams []Team

// String is not required by pop and may be deleted
func (t Teams) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (t *Team) Validate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	return validate.Validate(
		&validators.StringIsPresent{Field: t.Name, Name: "Name"},
		&validators.StringIsPresent{Field: t.Role, Name: "Role"},
		&validators.FuncValidator{
			Field:   t.Name,
			Name:    "Name",
			Message: "%s is already taken",
			Fn: func() bool {
				var b bool
				q := tx.Where("name = ?", t.Name)
				if t.ID != uuid.Nil {
					q = q.Where("id != ?", t.ID)
				}
				b, err = q.Exists(t)
				if err != nil {
					return false
				}
				return !b
			},
		},
		&validators.FuncValidator{
			Field:   t.Role,
			Name:    "Role",
			Message: "%s is not an existing role",
			Fn: func() bool {
				items := []string{constants.Black, constants.Blue, constants.Red, constants.White}
				_, found := find(items, t.Role)
				return found
			},
		},
	), err
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (t *Team) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (t *Team) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	return validate.Validate(
		&validators.FuncValidator{
			Field:   t.Role,
			Name:    "Role",
			Message: fmt.Sprintf("You should have at least one team with role \"%s\"", constants.Black),
			Fn: func() bool {
				
				team := Team{}
				if err = tx.Find(&team, t.ID); err != nil {
					return false
				}
				if team.Role == constants.Black && t.Role != constants.Black{
					teams := []Team{}
					//Query for all teams with role constants.Black
					if err = tx.Where("role = (?)", constants.Black).All(&teams); err != nil {
						return false
					}
					//Disallow deletion of the last constants.Black team
					if len(teams) <= 1 {
						return false
					}
				}
				return true
			},
		},
	), err
}

//find finds weather a given string is in the slice of strings.
func find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

//GetTeamByName retreives Team Object that matches the name parameter
func GetTeamByName(tx *pop.Connection, n string) (Team, error){
	t := []Team{}
	query := tx.Where("name = (?)", n)
	err := query.All(&t)
	if err != nil {
		return Team{}, err
	}
	return t[0], err
}