package competitionservice

import (
	"context"
	"errors"
	"fmt"

	"github.com/ScoreTrak/ScoreTrak/pkg/competition"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue/queueing"
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
	Store  *util.Store
	Config queueing.Config
}

func NewCompetitionServ(str *util.Store, cfg queueing.Config) Serv {
	return &competitionServ{
		Store:  str,
		Config: cfg,
	}
}

var (
	ErrLoadCompetition = errors.New("failed to load the competition")
)

func (svc *competitionServ) LoadCompetition(ctx context.Context, c *competition.Competition) error {
	errAgr := make([]error, 0, 11)
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
	if c.Policy != nil {
		errAgr = append(errAgr, svc.Store.Policy.Update(ctx, c.Policy))
	}
	if c.Users != nil {
		errAgr = append(errAgr, svc.Store.Users.Upsert(ctx, c.Users))
	}
	errStr := ""
	for i := range errAgr {
		if errAgr[i] != nil {
			var serr *pgconn.PgError
			ok := errors.As(errAgr[i], &serr)
			if !ok || serr.Code != "23505" {
				errStr += errAgr[i].Error() + "\n"
			}
		}
	}
	if errStr != "" {
		return fmt.Errorf("%w: %s", ErrLoadCompetition, errStr)
	}
	return nil
}

func (svc *competitionServ) FetchCoreCompetition(ctx context.Context) (*competition.Competition, error) {
	fls := false
	cnf, err := svc.Store.Config.Get(ctx)
	if err != nil {
		return nil, wrapError(err, fetch)
	}
	pol, err := svc.Store.Policy.Get(ctx)
	if err != nil {
		return nil, wrapError(err, fetch)
	}
	cnf.Enabled = &fls
	teams, err := svc.Store.Team.GetAll(ctx)
	if err != nil {
		return nil, wrapError(err, fetch)
	}
	users, err := svc.Store.Users.GetAll(ctx)
	if err != nil {
		return nil, wrapError(err, fetch)
	}
	hostsGroup, err := svc.Store.HostGroup.GetAll(ctx)
	if err != nil {
		return nil, wrapError(err, fetch)
	}
	hosts, err := svc.Store.Host.GetAll(ctx)
	if err != nil {
		return nil, wrapError(err, fetch)
	}
	serviceGroups, err := svc.Store.ServiceGroup.GetAll(ctx)
	if svc.Config.Use != "none" {
		for i := range serviceGroups {
			serviceGroups[i].Enabled = &fls
		}
	}
	if err != nil {
		return nil, wrapError(err, fetch)
	}
	services, err := svc.Store.Service.GetAll(ctx)
	if err != nil {
		return nil, wrapError(err, fetch)
	}
	properties, err := svc.Store.Property.GetAll(ctx)
	if err != nil {
		return nil, wrapError(err, fetch)
	}
	return &competition.Competition{Config: cnf, Teams: teams, HostGroups: hostsGroup, Hosts: hosts, ServiceGroups: serviceGroups, Services: services, Properties: properties, Policy: pol, Users: users}, nil
}

func (svc *competitionServ) FetchEntireCompetition(ctx context.Context) (*competition.Competition, error) {
	cmp, err := svc.FetchCoreCompetition(ctx)
	if err != nil {
		return nil, wrapError(err, fetch)
	}
	rounds, err := svc.Store.Round.GetAll(ctx)
	if err != nil {
		return nil, wrapError(err, fetch)
	}
	checks, err := svc.Store.Check.GetAll(ctx)
	if err != nil {
		return nil, wrapError(err, fetch)
	}
	getReport, err := svc.Store.Report.Get(ctx)
	if err != nil {
		return nil, wrapError(err, fetch)
	}
	cmp.Rounds = rounds
	cmp.Checks = checks
	cmp.Report = getReport
	return cmp, nil
}

func (svc *competitionServ) ResetScores(ctx context.Context) error {
	err := svc.Store.Check.TruncateTable(ctx)
	if err != nil {
		return wrapError(err, reset)
	}
	err = svc.Store.Round.TruncateTable(ctx)
	if err != nil {
		return wrapError(err, reset)
	}
	err = svc.Store.Report.Update(ctx, &report.Report{Cache: "{}"})
	if err != nil {
		return wrapError(err, reset)
	}
	return nil
}

func wrapError(err error, action string) error {
	return fmt.Errorf("failed to %s competition: %w", action, err)
}

const (
	remove = "remove"
	reset  = "reset"
	fetch  = "fetch"
)

func (svc *competitionServ) DeleteCompetition(ctx context.Context) error {
	err := svc.ResetScores(ctx)
	if err != nil {
		return wrapError(err, remove)
	}
	err = svc.Store.Property.TruncateTable(ctx)
	if err != nil {
		return wrapError(err, remove)
	}
	err = svc.Store.Service.TruncateTable(ctx)
	if err != nil {
		return wrapError(err, remove)
	}
	err = svc.Store.Host.TruncateTable(ctx)
	if err != nil {
		return wrapError(err, remove)
	}
	err = svc.Store.HostGroup.TruncateTable(ctx)
	if err != nil {
		return wrapError(err, remove)
	}
	return nil
}
