package report

import (
	"time"

	"github.com/gofrs/uuid"
)

// Report is a structure that represents the state of all checks, services, teams, hosts, at a given round(typically last round).
// The report strips out all of the unnecessary details from the checks, and forwards the generated output to the client. The API can additionally perform
// any sorts of filtering on the report in case it needs to (Ex: Hiding sensitive details of team blue2 from team blue1's users)
type Report struct {
	ID        uint64    `json:"id,omitempty"`
	Cache     string    `json:"cache,omitempty" gorm:"not null;default:'{}'"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewReport() *Report {
	c := &Report{}
	c.ID = 1
	c.Cache = "{}"
	return c
}

type SimpleReport struct {
	Round uint64
	Teams map[uuid.UUID]*SimpleTeam
}

type SimpleTeam struct {
	Hosts       map[uuid.UUID]*SimpleHost
	Name        string
	Pause       bool
	Hide        bool
	TotalPoints uint64
}

type SimpleHostGroup struct {
	ID    uuid.UUID
	Name  string
	Pause bool
	Hide  bool
}

type SimpleHost struct {
	HostGroup *SimpleHostGroup
	Address   string
	Services  map[uuid.UUID]*SimpleService
	Pause     bool
	Hide      bool
}

type SimpleServiceGroup struct {
	ID      uuid.UUID
	Name    string
	Enabled bool
}

type SimpleCheck struct {
	Passed bool
	Log    string
	Err    string
}

type SimpleService struct {
	Check              *SimpleCheck
	Pause              bool
	Hide               bool
	Name               string
	DisplayName        string
	Weight             uint64
	Points             uint64
	PointsBoost        uint64
	Properties         map[string]*SimpleProperty
	SimpleServiceGroup *SimpleServiceGroup
}

type SimpleProperty struct {
	Value  string
	Status string
}
