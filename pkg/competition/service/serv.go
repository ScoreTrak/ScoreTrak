package service

import (
	"context"
	"errors"
	"github.com/ScoreTrak/ScoreTrak/pkg/competition"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/util"
	"github.com/jackc/pgconn"
)

type Serv interface {
	LoadCompetition(ctx context.Context, competition *competition.Competition) error
	FetchCoreCompetition(ctx context.Context) (*competition.Competition, error)
	FetchEntireCompetition(ctx context.Context) (*competition.Competition, error)
	ResetScores(ctx context.Context) error
	DeleteCompetition(ctx context.Context) error
}

type competitionServ struct {
	Store util.Store
}

func NewCompetitionServ(str util.Store) Serv {
	return &competitionServ{
		Store: str,
	}
}

func (svc *competitionServ) LoadCompetition(ctx context.Context, c *competition.Competition) error {
	var errAgr []error
	err := svc.Store.Config.Update(ctx, c.Config)
	if err != nil {
		return err
	}
	if len(c.Teams) != 0 {
		errAgr = append(errAgr, svc.Store.Team.Upsert(ctx, c.Teams))
	}
	if len(c.HostGroups) != 0 {
		errAgr = append(errAgr, svc.Store.HostGroup.Upsert(ctx, c.HostGroups))
	}

	if len(c.Hosts) != 0 {
		errAgr = append(errAgr, svc.Store.Host.Upsert(ctx, c.Hosts))
	}
	for i := range c.ServiceGroups {
		errAgr = append(errAgr, svc.Store.ServiceGroup.Upsert(ctx, c.ServiceGroups[i]))
	}
	if len(c.Services) != 0 {
		errAgr = append(errAgr, svc.Store.Service.Upsert(ctx, c.Services))
	}
	if len(c.Properties) != 0 {
		errAgr = append(errAgr, svc.Store.Property.Upsert(ctx, c.Properties))
	}
	if len(c.Rounds) != 0 {
		errAgr = append(errAgr, svc.Store.Round.Upsert(ctx, c.Rounds))
	}
	if len(c.Checks) != 0 {
		errAgr = append(errAgr, svc.Store.Check.Upsert(ctx, c.Checks))
	}
	if c.Report != nil {
		errAgr = append(errAgr, svc.Store.Report.Update(ctx, c.Report))
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

func (svc *competitionServ) FetchCoreCompetition(ctx context.Context) (*competition.Competition, error) {
	fls := false
	cnf, err := svc.Store.Config.Get(ctx)
	if err != nil {
		return nil, err
	}
	cnf.Enabled = &fls
	teams, err := svc.Store.Team.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	hostsGroup, err := svc.Store.HostGroup.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	hosts, err := svc.Store.Host.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	serviceGroups, err := svc.Store.ServiceGroup.GetAll(ctx)
	if config.GetStaticConfig().Queue.Use != "none" {
		for i := range serviceGroups {
			serviceGroups[i].Enabled = &fls
		}
	}
	if err != nil {
		return nil, err
	}
	services, err := svc.Store.Service.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	properties, err := svc.Store.Property.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return &competition.Competition{Config: cnf, Teams: teams, HostGroups: hostsGroup, Hosts: hosts, ServiceGroups: serviceGroups, Services: services, Properties: properties}, nil
}

func (svc *competitionServ) FetchEntireCompetition(ctx context.Context) (*competition.Competition, error) {
	cmp, err := svc.FetchCoreCompetition(ctx)
	if err != nil {
		return nil, err
	}
	rounds, err := svc.Store.Round.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	checks, err := svc.Store.Check.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	getReport, err := svc.Store.Report.Get(ctx)
	if err != nil {
		return nil, err
	}
	cmp.Rounds = rounds
	cmp.Checks = checks
	cmp.Report = getReport
	return cmp, nil
}

func (svc *competitionServ) ResetScores(ctx context.Context) error {
	err := svc.Store.Check.TruncateTable(ctx)
	if err != nil {
		return err
	}
	err = svc.Store.Round.TruncateTable(ctx)
	if err != nil {
		return err
	}
	err = svc.Store.Report.Update(ctx, &report.Report{Cache: "{}"})
	if err != nil {
		return err
	}
	return nil
}

func (svc *competitionServ) DeleteCompetition(ctx context.Context) error {
	err := svc.ResetScores(ctx)
	if err != nil {
		return err
	}
	err = svc.Store.Property.TruncateTable(ctx)
	if err != nil {
		return err
	}
	err = svc.Store.Service.TruncateTable(ctx)
	if err != nil {
		return err
	}
	err = svc.Store.Host.TruncateTable(ctx)
	if err != nil {
		return err
	}
	err = svc.Store.HostGroup.TruncateTable(ctx)
	if err != nil {
		return err
	}
	return nil
}
