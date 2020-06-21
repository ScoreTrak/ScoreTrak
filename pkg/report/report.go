package report

import (
	"time"
)

type Report struct {
	ID        uint64    `json:"id,omitempty"`
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
	Round uint64
	Teams map[uint64]SimpleTeam
}

type SimpleTeam struct {
	Hosts map[uint64]SimpleHost
}

type SimpleHost struct {
	Services map[uint64]SimpleService
}

type SimpleService struct {
	Passed      bool
	Log         string
	Err         string
	Points      uint64
	PointsBoost uint64
	Properties  map[string]string
}
