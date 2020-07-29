package report

import (
	"github.com/L1ghtman2k/ScoreTrak/pkg/round"
	"github.com/L1ghtman2k/ScoreTrak/pkg/team"
)

type Serv interface {
	Get() (*Report, error)
}

type reportServ struct {
	repo Repo
}

func NewReportServ(repo Repo) Serv {
	return &reportServ{
		repo: repo,
	}
}
func (svc *reportServ) Get() (*Report, error) { return svc.repo.Get() }

func (svc *reportServ) RecalculateReport(team []*team.Team, round round.Round) (simpleTeams map[uint32]SimpleTeam, err error) {
	simpleTeams = make(map[uint32]SimpleTeam)
	for _, t := range team {
		st := SimpleTeam{}
		st.Hosts = make(map[uint32]*SimpleHost)
		for _, h := range t.Hosts {
			sh := SimpleHost{}
			sh.Services = make(map[uint32]*SimpleService)
			for _, s := range h.Services {
				var points uint32
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
						sh.Services[s.ID] = &SimpleService{Passed: *lastCheck.Passed, Log: lastCheck.Log, Err: lastCheck.Err, Points: points, Properties: params, PointsBoost: s.PointsBoost}
					} else {
						sh.Services[s.ID] = &SimpleService{Passed: false, Log: "Service was not checked because it was disabled", Err: "", Points: points, Properties: params, PointsBoost: s.PointsBoost}
					}
				}
			}
			st.Hosts[h.ID] = &sh
		}
		simpleTeams[t.ID] = st
	}
	return simpleTeams, nil
}
