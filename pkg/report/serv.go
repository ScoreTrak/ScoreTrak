package report

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/host_group"
	"github.com/ScoreTrak/ScoreTrak/pkg/round"
	"github.com/ScoreTrak/ScoreTrak/pkg/service_group"
	"github.com/ScoreTrak/ScoreTrak/pkg/team"
	"github.com/gofrs/uuid"
)

type Serv interface {
	Get() (*Report, error)
}

type reportServ struct {
	repo Repo
}

type reportCalculator struct {
	repo Repo
}

func NewReportServ(repo Repo) Serv {
	return &reportServ{
		repo: repo,
	}
}

func NewReportCalculator(repo Repo) *reportCalculator {
	return &reportCalculator{
		repo: repo,
	}
}

func (svc *reportServ) Get() (*Report, error) { return svc.repo.Get() }

func (svc *reportCalculator) RecalculateReport(team []*team.Team, hostGroup []*host_group.HostGroup, serviceGroups []*service_group.ServiceGroup, round round.Round) (simpleTeams map[uuid.UUID]*SimpleTeam, err error) {
	simpleTeams = make(map[uuid.UUID]*SimpleTeam)
	for _, t := range team {
		st := &SimpleTeam{Name: t.Name, Enabled: *t.Enabled}
		st.Hosts = make(map[uuid.UUID]*SimpleHost)
		for _, h := range t.Hosts {
			sh := SimpleHost{Address: *h.Address, Enabled: *h.Enabled}
			if h.HostGroupID != nil {
				for _, hG := range hostGroup {
					if hG.ID == *h.HostGroupID {
						sh.HostGroup = &SimpleHostGroup{Enabled: *hG.Enabled, ID: *h.HostGroupID, Name: hG.Name}
					}
				}
			}
			sh.Services = make(map[uuid.UUID]*SimpleService)
			for _, s := range h.Services {
				var points uint
				for _, c := range s.Checks {
					if *c.Passed {
						points += s.Points
					}
				}
				params := map[string]*SimpleProperty{}
				for _, p := range s.Properties {
					params[p.Key] = &SimpleProperty{Value: p.Value, Status: p.Status}
				}
				var simpSgr *SimpleServiceGroup
				for _, sG := range serviceGroups {
					if sG.ID == s.ServiceGroupID {
						simpSgr = &SimpleServiceGroup{sG.ID, sG.Name, *sG.Enabled}
					}
				}
				if len(s.Checks) != 0 {
					lastCheck := s.Checks[len(s.Checks)-1]
					sh.Services[s.ID] = &SimpleService{Name: s.Name, DisplayName: s.DisplayName, Enabled: *s.Enabled, Points: points, Properties: params, PointsBoost: s.PointsBoost, SimpleServiceGroup: simpSgr}
					if lastCheck.RoundID == round.ID {
						sh.Services[s.ID].Passed = *lastCheck.Passed
						sh.Services[s.ID].Log = lastCheck.Log
						sh.Services[s.ID].Err = lastCheck.Err
					} else {
						sh.Services[s.ID].Passed = false
						sh.Services[s.ID].Log = "Service was not checked because it was disabled"
						sh.Services[s.ID].Err = ""
					}
				}
			}
			st.Hosts[h.ID] = &sh
		}
		simpleTeams[t.ID] = st
	}
	return simpleTeams, nil
}
