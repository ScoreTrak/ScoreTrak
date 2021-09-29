package service_group_service

import (
	"context"
	"errors"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/platform"
	"github.com/ScoreTrak/ScoreTrak/pkg/platform/worker"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue"
	"github.com/ScoreTrak/ScoreTrak/pkg/service_group"
	repo2 "github.com/ScoreTrak/ScoreTrak/pkg/service_group/service_group_repo"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	q    queue.WorkerQueue
}

func (svc *serviceGroupServ) Redeploy(ctx context.Context, id uuid.UUID) error {
	if svc.p == nil || config.GetStaticConfig().Queue.Use == queue.None {
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

func NewServiceGroupServ(repo repo2.Repo, plat platform.Platform, q queue.WorkerQueue) Serv {
	return &serviceGroupServ{
		repo: repo, p: plat, q: q,
	}
}

func (svc *serviceGroupServ) Delete(ctx context.Context, id uuid.UUID) error {
	serviceGrp, err := svc.GetByID(ctx, id)
	if err != nil {
		return status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Service Group not found: %v", err),
		)
	}
	err = svc.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	if svc.p != nil && config.GetStaticConfig().Queue.Use != queue.None {
		wr := worker.Info{Topic: serviceGrp.Name, Label: serviceGrp.Label}
		err := svc.p.RemoveWorkers(ctx, wr)
		if err != nil {
			return fmt.Errorf("element was removed from database, however, we were unable to remove workers from platform. Error: %v", err)
		}
	}
	return nil
}

func (svc *serviceGroupServ) GetAll(ctx context.Context) ([]*service_group.ServiceGroup, error) {
	return svc.repo.GetAll(ctx)
}

func (svc *serviceGroupServ) GetByID(ctx context.Context, id uuid.UUID) (*service_group.ServiceGroup, error) {
	return svc.repo.GetByID(ctx, id)
}

func (svc *serviceGroupServ) Store(ctx context.Context, u *service_group.ServiceGroup) error {
	if svc.p != nil && !u.SkipHelper && config.GetStaticConfig().Queue.Use != queue.None {
		if u.Enabled != nil && *u.Enabled {
			return status.Errorf(
				codes.FailedPrecondition,
				"if you are letting scoretrak manage the workers, 'Enabled' can be set to true, only after workers are deployed.",
			)

		}
		wr := worker.Info{Topic: u.Name, Label: u.Label}
		err := svc.p.DeployWorkers(ctx, wr)
		if err != nil {
			return err
		}
	}
	return svc.repo.Store(ctx, u)
}

func (svc *serviceGroupServ) Update(ctx context.Context, u *service_group.ServiceGroup) error {
	serviceGrp, err := svc.GetByID(ctx, u.ID)
	if err != nil {
		return status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Service Group not found: %v", err),
		)
	}
	if !u.SkipHelper && config.GetStaticConfig().Queue.Use != queue.None {
		if u.Enabled != nil && *u.Enabled && !*serviceGrp.Enabled {
			err = svc.q.Ping(serviceGrp)
			if err != nil {
				return err
			}
		}
	}
	return svc.repo.Update(ctx, u)
}
