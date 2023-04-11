package competition

import (
	"time"

	"github.com/ScoreTrak/ScoreTrak/pkg/team"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// Competition is a struct that holds an aggregate of all models. This is used to upload/export competition as a file.
type Competition struct {
	ID   uuid.UUID `json:"id,omitempty" gorm:"type:uuid;primary_key;"`
	Name string
	// Owner         user.User
	Started       *bool
	RoundDuration uint64
	// Staff         []*team.Team
	// Competitors []*user.User
	// Config        *config.DynamicConfig
	// Report       *report.Report
	// HostGroups   []*hostgroup.HostGroup
	// Hosts []*host.Host
	// Services     []*service.Service
	// WorkerGroups []*workergroup.WorkerGroup
	// Rounds       []*round.Round
	// Properties   []*property.Property
	// Checks       []*check.Check
	// Users  []*user.User
	// Policy *policy.Policy
	Teams     []*team.Team `gorm:"many2many:competiton_teams"`
	StartedAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

func CurrentCompetition(competitionId uuid.UUID) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("competition_id = ?", competitionId)
	}
}

func OwnedCompetitions(ownerId uuid.UUID) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("owner_id = ?", ownerId)
	}
}

func RunningCompetitions(db *gorm.DB) *gorm.DB {
	return db.Where("started = ?", true)
}
