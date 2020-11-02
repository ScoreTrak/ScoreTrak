package team

import (
	"errors"
	"github.com/ScoreTrak/ScoreTrak/pkg/host"
	"github.com/ScoreTrak/ScoreTrak/pkg/user"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// Team model represents internal team model of the scoring engine.
type Team struct {

	// this id refers to ID of a team in web.
	ID uuid.UUID `json:"id,omitempty" gorm:"type:uuid;primary_key;"`

	Name string `json:"name" gorm:"unique;not null" valid:"required,alphanum"`

	Enabled *bool `json:"enabled,omitempty" gorm:"not null;default:true"`

	Index *uint64 `json:"index" gorm:"unique;not null"`

	Users []*user.User `gorm:"foreignkey:TeamID;association_foreignkey:ID;constraint:OnUpdate:RESTRICT,OnDelete:CASCADE" json:"-"`

	Hosts []*host.Host `gorm:"foreignkey:TeamID;association_foreignkey:ID; constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT" json:"hosts,omitempty"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (t *Team) BeforeCreate(tx *gorm.DB) (err error) {
	if t.Name == "" {
		return errors.New("field Name is a mandatory parameter")
	}
	if t.ID == uuid.Nil {
		u, err := uuid.NewV4()
		if err != nil {
			return err
		}
		t.ID = u
	}
	return nil
}
