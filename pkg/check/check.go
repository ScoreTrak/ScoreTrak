package check

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// Check model contains an instance of a single check performed on a single host at a given round for a given check_service. Check's ID is defined by combined key of ServiceID, and RoundID
type Check struct {
	// RoundID is a parent service of a given check
	ServiceID uuid.UUID `json:"service_id,omitempty" gorm:"primary_key;type:uuid;auto_increment:false"`

	// RoundID is a parent round of a given check
	RoundID uint64 `json:"round_id,omitempty" gorm:"primary_key;auto_increment:false"`

	// Log represents an comment/log of a check. This will be helpful for debugging purposes during the competition
	Log string `json:"log,omitempty"`

	// Err represents the details of an error if check was failed
	Err string `json:"err,omitempty"`

	// Passed tells weather a given check passed/failed
	Passed *bool `json:"passed,omitempty" gorm:"not null;default:false"`

	// Idempotent token that ensure scheduler does not update a check
	// IdempotentToken uuid.UUID `json:"idempotent_token,omitempty" gorm:"not null;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
