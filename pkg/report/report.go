package report

import (
	"time"
)

type Report struct {
	ID        uint32    `json:"id,omitempty"`
	Cache     string    `json:"cache,omitempty" gorm:"not null;default:'{}'"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (d Report) TableName() string {
	return "report"
}

func NewReport() *Report {
	c := &Report{}
	c.ID = 1
	c.Cache = "{}"
	return c
}

type SimpleReport struct {
	Round uint32
	Teams map[uint32]*SimpleTeam
}

type SimpleTeam struct {
	Hosts map[uint32]*SimpleHost
}

type SimpleHost struct {
	Services map[uint32]*SimpleService
}

type SimpleService struct {
	Name        string
	DisplayName string
	Passed      bool
	Log         string
	Err         string
	Points      uint32
	PointsBoost uint32
	Properties  map[string]string
}
