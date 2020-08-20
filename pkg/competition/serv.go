package competition

import (
	"errors"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/di/repo"
	"github.com/jackc/pgconn"
)

type Serv interface {
	LoadCompetition(*Competition) error
	FetchCoreCompetition() (*Competition, error)
	FetchEntireCompetition() (*Competition, error)
}

type configServ struct {
	Store repo.Store
}

func NewCompetitionServ(str repo.Store) Serv {
	return &configServ{
		Store: str,
	}
}

func (svc *configServ) LoadCompetition(c *Competition) error {
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
	for i, _ := range c.ServiceGroups {
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
	for i, _ := range errAgr {
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

func (svc *configServ) FetchCoreCompetition() (*Competition, error) {
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

func (svc *configServ) FetchEntireCompetition() (*Competition, error) {
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
	report, err := svc.Store.Report.Get()
	if err != nil {
		return nil, err
	}
	cmp.Rounds = rounds
	cmp.Checks = checks
	cmp.Report = report
	return cmp, nil
}
