package service_group_service

import (
	"context"
	"errors"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/platform"
	"github.com/ScoreTrak/ScoreTrak/pkg/platform/worker"
	"github.com/ScoreTrak/ScoreTrak/pkg/service_group"
	repo2 "github.com/ScoreTrak/ScoreTrak/pkg/service_group/service_group_repo"
	"github.com/gofrs/uuid"
	"strings"
)

type Serv interface {
	Delete(ctx context.Context, id uuid.UUID) error
	GetAll(ctx context.Context) ([]*service_group.ServiceGroup, error)
	GetByID(ctx context.Context, id uuid.UUID) (*service_group.ServiceGroup, error)
	Store(ctx context.Context, u *service_group.ServiceGroup) error
	Update(ctx context.Context, u *service_group.ServiceGroup) error
	Redeploy(ctx context.Context, id uuid.UUID) error
}

type serviceGroupServ struct {
	repo repo2.Repo
	p    platform.Platform
}

func (svc *serviceGroupServ) Redeploy(ctx context.Context, id uuid.UUID) error {
	if !(svc.p != nil && config.GetStaticConfig().Queue.Use != "none") {
		return errors.New("queue was not established, or platform is none, please manually redeploy the workers")
	}

	serGrp, err := svc.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if *serGrp.Enabled {
		return errors.New("check_service group must first be disabled")
	}
	wr := worker.Info{Topic: serGrp.Name, Label: serGrp.Label}
	err = svc.p.RemoveWorkers(ctx, wr)
	if err != nil && !strings.Contains(err.Error(), "not found") {
		return fmt.Errorf("scoretrak encountered an error while removing the workersvc. Please, delete the workers manually. Details: %v", err)
	}
	err = svc.p.DeployWorkers(ctx, wr)
	if err != nil {
		return fmt.Errorf("scoretrak encountered an error while deploying the workersvc. Please, investigate the issue, or create the workers manually. Details: %v", err.Error())
	}
	return nil
}

func NewServiceGroupServ(repo repo2.Repo) Serv {
	return &serviceGroupServ{
		repo: repo,
	}
}

func (svc *serviceGroupServ) Delete(ctx context.Context, id uuid.UUID) error {
	return svc.repo.Delete(ctx, id)
}

func (svc *serviceGroupServ) GetAll(ctx context.Context) ([]*service_group.ServiceGroup, error) {
	return svc.repo.GetAll(ctx)
}

func (svc *serviceGroupServ) GetByID(ctx context.Context, id uuid.UUID) (*service_group.ServiceGroup, error) {
	return svc.repo.GetByID(ctx, id)
}

func (svc *serviceGroupServ) Store(ctx context.Context, u *service_group.ServiceGroup) error {
	return svc.repo.Store(ctx, u)
}

func (svc *serviceGroupServ) Update(ctx context.Context, u *service_group.ServiceGroup) error {
	return svc.repo.Update(ctx, u)
}
