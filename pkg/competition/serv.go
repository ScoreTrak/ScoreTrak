package competition

import (
	"errors"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/di/repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	"github.com/jackc/pgconn"
)

type Serv interface {
	LoadCompetition(*Competition) error
	FetchCoreCompetition() (*Competition, error)
	FetchEntireCompetition() (*Competition, error)
	ResetScores() error
	DeleteCompetition() error
}

type competitionServ struct {
	Store repo.Store
}

func NewCompetitionServ(str repo.Store) Serv {
	return &competitionServ{
		Store: str,
	}
}

func (svc *competitionServ) LoadCompetition(c *Competition) error {
	var errAgr []error
	err := svc.Store.Config.Update(c.Config)
	if err != nil {
		return err
	}
	if len(c.Teams) != 0 {
		errAgr = append(errAgr, svc.Store.Team.Upsert(c.Teams))
	}
	if len(c.HostGroups) != 0 {
		errAgr = append(errAgr, svc.Store.HostGroup.Upsert(c.HostGroups))
	}

	if len(c.Hosts) != 0 {
		errAgr = append(errAgr, svc.Store.Host.Upsert(c.Hosts))
	}
	for i := range c.ServiceGroups {
		errAgr = append(errAgr, svc.Store.ServiceGroup.Upsert(c.ServiceGroups[i]))
	}
	if len(c.Services) != 0 {
		errAgr = append(errAgr, svc.Store.Service.Upsert(c.Services))
	}
	if len(c.Properties) != 0 {
		errAgr = append(errAgr, svc.Store.Property.Upsert(c.Properties))
	}
	if len(c.Rounds) != 0 {
		errAgr = append(errAgr, svc.Store.Round.Upsert(c.Rounds))
	}
	if len(c.Checks) != 0 {
		errAgr = append(errAgr, svc.Store.Check.Upsert(c.Checks))
	}
	if c.Report != nil {
		errAgr = append(errAgr, svc.Store.Report.Update(c.Report))
	}
	errStr := ""
	for i := range errAgr {
		if errAgr[i] != nil {
			serr, ok := errAgr[i].(*pgconn.PgError)
			if !ok || serr.Code != "23505" {
				errStr += errAgr[i].Error() + "\n"
			}
		}
	}
	if errStr != "" {
		return errors.New(errStr)
	}
	return nil
}

func (svc *competitionServ) FetchCoreCompetition() (*Competition, error) {
	fls := false
	cnf, err := svc.Store.Config.Get()
	if err != nil {
		return nil, err
	}
	cnf.Enabled = &fls
	teams, err := svc.Store.Team.GetAll()
	if err != nil {
		return nil, err
	}
	hostsGroup, err := svc.Store.HostGroup.GetAll()
	if err != nil {
		return nil, err
	}
	hosts, err := svc.Store.Host.GetAll()
	if err != nil {
		return nil, err
	}
	serviceGroups, err := svc.Store.ServiceGroup.GetAll()
	if config.GetStaticConfig().Queue.Use != "none" {
		for i := range serviceGroups {
			serviceGroups[i].Enabled = &fls
		}
	}
	if err != nil {
		return nil, err
	}
	services, err := svc.Store.Service.GetAll()
	if err != nil {
		return nil, err
	}
	properties, err := svc.Store.Property.GetAll()
	if err != nil {
		return nil, err
	}
	return &Competition{Config: cnf, Teams: teams, HostGroups: hostsGroup, Hosts: hosts, ServiceGroups: serviceGroups, Services: services, Properties: properties}, nil
}

func (svc *competitionServ) FetchEntireCompetition() (*Competition, error) {
	cmp, err := svc.FetchCoreCompetition()
	if err != nil {
		return nil, err
	}
	rounds, err := svc.Store.Round.GetAll()
	if err != nil {
		return nil, err
	}
	checks, err := svc.Store.Check.GetAll()
	if err != nil {
		return nil, err
	}
	getReport, err := svc.Store.Report.Get()
	if err != nil {
		return nil, err
	}
	cmp.Rounds = rounds
	cmp.Checks = checks
	cmp.Report = getReport
	return cmp, nil
}

func (svc *competitionServ) ResetScores() error {
	err := svc.Store.Check.TruncateTable()
	if err != nil {
		return err
	}
	err = svc.Store.Round.TruncateTable()
	if err != nil {
		return err
	}
	err = svc.Store.Report.Update(&report.Report{Cache: "{}"})
	if err != nil {
		return err
	}
	return nil
}

func (svc *competitionServ) DeleteCompetition() error {
	err := svc.ResetScores()
	if err != nil {
		return err
	}
	err = svc.Store.Property.TruncateTable()
	if err != nil {
		return err
	}
	err = svc.Store.Service.TruncateTable()
	if err != nil {
		return err
	}
	err = svc.Store.ServiceGroup.TruncateTable()
	if err != nil {
		return err
	}
	err = svc.Store.Host.TruncateTable()
	if err != nil {
		return err
	}
	err = svc.Store.HostGroup.TruncateTable()
	if err != nil {
		return err
	}
	return nil
}
