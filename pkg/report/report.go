package report

import (
	"ScoreTrak/pkg/round"
	"ScoreTrak/pkg/team"
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

func RecalculateReport(team []*team.Team, round round.Round) (simpleTeams map[uint64]SimpleTeam, err error) {
	simpleTeams = make(map[uint64]SimpleTeam)
	for _, t := range team {
		st := SimpleTeam{}
		st.Hosts = make(map[uint64]SimpleHost)
		for _, h := range t.Hosts {
			sh := SimpleHost{}
			sh.Services = make(map[uint64]SimpleService)
			for _, s := range h.Services {
				var points uint64
				for _, c := range s.Checks {
					if *c.Passed {
						points += s.Points
					}
				}
				params := map[string]string{}
				for _, p := range s.Properties {
					if p.Status != "Hide" {
						params[p.Key] = p.Value
					}
				}
				if len(s.Checks) != 0 {
					lastCheck := s.Checks[len(s.Checks)-1]
					if lastCheck.RoundID == round.ID {
						sh.Services[s.ID] = SimpleService{Passed: *lastCheck.Passed, Log: lastCheck.Log, Err: lastCheck.Err, Points: points, Properties: params, PointsBoost: s.PointsBoost}
					} else {
						sh.Services[s.ID] = SimpleService{Passed: false, Log: "Service was not checked because it was disabled", Err: "", Points: points, Properties: params, PointsBoost: s.PointsBoost}
					}
				}
			}
			st.Hosts[h.ID] = sh
		}
		simpleTeams[t.ID] = st
	}
	return simpleTeams, nil
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
