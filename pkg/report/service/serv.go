package service

import (
	"context"
	"github.com/ScoreTrak/ScoreTrak/pkg/host_group"
	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	repo2 "github.com/ScoreTrak/ScoreTrak/pkg/report/repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/round"
	"github.com/ScoreTrak/ScoreTrak/pkg/service_group"
	"github.com/ScoreTrak/ScoreTrak/pkg/team"
	"github.com/gofrs/uuid"
)

type Serv interface {
	Get(ctx context.Context) (*report.Report, error)
}

type reportServ struct {
	repo repo2.Repo
}

type reportCalculator struct {
	repo repo2.Repo
}

func NewReportServ(repo repo2.Repo) Serv {
	return &reportServ{
		repo: repo,
	}
}

func NewReportCalculator(repo repo2.Repo) *reportCalculator {
	return &reportCalculator{
		repo: repo,
	}
}

func (svc *reportServ) Get(ctx context.Context) (*report.Report, error) { return svc.repo.Get(ctx) }

func (svc *reportCalculator) RecalculateReport(team []*team.Team, hostGroup []*host_group.HostGroup, serviceGroups []*service_group.ServiceGroup, round round.Round) (simpleTeams map[uuid.UUID]*report.SimpleTeam) {
	simpleTeams = make(map[uuid.UUID]*report.SimpleTeam)
	for _, t := range team {
		st := &report.SimpleTeam{Name: t.Name, Enabled: *t.Enabled}
		st.Hosts = make(map[uuid.UUID]*report.SimpleHost)
		for _, h := range t.Hosts {
			sh := report.SimpleHost{Address: *h.Address, Enabled: *h.Enabled}
			if h.HostGroupID != nil {
				for _, hG := range hostGroup {
					if hG.ID == *h.HostGroupID {
						sh.HostGroup = &report.SimpleHostGroup{Enabled: *hG.Enabled, ID: *h.HostGroupID, Name: hG.Name}
					}
				}
			}
			sh.Services = make(map[uuid.UUID]*report.SimpleService)
			for _, s := range h.Services {
				var points uint64
				for _, c := range s.Checks {
					if *c.Passed {
						points += *s.Weight
					}
				}
				params := map[string]*report.SimpleProperty{}
				for _, p := range s.Properties {
					params[p.Key] = &report.SimpleProperty{Value: *p.Value, Status: p.Status}
				}
				var simpSgr *report.SimpleServiceGroup
				for _, sG := range serviceGroups {
					if sG.ID == s.ServiceGroupID {
						simpSgr = &report.SimpleServiceGroup{ID: sG.ID, Name: sG.Name, Enabled: *sG.Enabled}
					}
				}
				if len(s.Checks) != 0 {
					lastCheck := s.Checks[len(s.Checks)-1]
					sh.Services[s.ID] = &report.SimpleService{Name: s.Name, DisplayName: s.DisplayName, Enabled: *s.Enabled, Points: points, Properties: params, PointsBoost: *s.PointsBoost, SimpleServiceGroup: simpSgr, Weight: *s.Weight}
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
	return simpleTeams
}
