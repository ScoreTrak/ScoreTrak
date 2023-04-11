package team

import (
	"errors"
	"time"

	"github.com/ScoreTrak/ScoreTrak/pkg/host"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// Team model represents internal team model of the scoring engine.
type Team struct {
	ID uuid.UUID `json:"id,omitempty" gorm:"type:uuid;primary_key;"`
	// Name is team name
	Name string `json:"name" gorm:"unique;not null"`
	// Pause is responsible for pausing the scoring
	Pause *bool `json:"pause,omitempty" gorm:"not null;default:false"`
	// Hide is responsible for hiding the Team on scoring table
	Hide *bool `json:"hide,omitempty" gorm:"not null;default:false"`
	// Index is team's index, which is a unique value used for templating (10.X.1.1, where X is an index) and sorting
	Index *uint64 `json:"index" gorm:"unique;"`
	// Type of team. Staff or Competitiors
	// Users array stores all of the users that are part of a given Team
	// Users []*user.User `gorm:"foreignkey:UserID;association_foreignkey:ID;constraint:OnUpdate:RESTRICT,OnDelete:CASCADE" json:"-"`
	// Hosts array stores all Hosts that are part of a given Team
	Hosts         []*host.Host `gorm:"foreignkey:TeamID;association_foreignkey:ID; constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT" json:"hosts,omitempty"`
	CompetitionID uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
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
