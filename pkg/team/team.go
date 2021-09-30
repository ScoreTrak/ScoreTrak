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
	// Name is team name
	Name string `json:"name" gorm:"unique;not null"`
	//Pause is responsible for pausing the scoring
	Pause *bool `json:"pause,omitempty" gorm:"not null;default:false"`
	//Hide is responsible for hiding the Team on scoring table
	Hide *bool `json:"hide,omitempty" gorm:"not null;default:false"`
	//Index is team's index, which is a unique value used for templating (10.X.1.1, where X is an index) and sorting
	Index *uint64 `json:"index" gorm:"unique;"`
	//Users array stores all of the users that are part of a given Team
	Users []*user.User `gorm:"foreignkey:TeamID;association_foreignkey:ID;constraint:OnUpdate:RESTRICT,OnDelete:CASCADE" json:"-"`
	//Hosts array stores all Hosts that are part of a given Team
	Hosts []*host.Host `gorm:"foreignkey:TeamID;association_foreignkey:ID; constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT" json:"hosts,omitempty"`
}

var ErrNameIsRequired = errors.New("field Name is required")

// BeforeCreate ensures UUID is set.
func (t *Team) BeforeCreate(tx *gorm.DB) (err error) {
	if t.Name == "" {
		return ErrNameIsRequired
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
