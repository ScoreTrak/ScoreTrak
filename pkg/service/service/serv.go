package service

import (
	"context"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/cmd/master/run"
	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue/queueing"
	"github.com/ScoreTrak/ScoreTrak/pkg/service"
	repo2 "github.com/ScoreTrak/ScoreTrak/pkg/service/repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/util"
	"github.com/gofrs/uuid"
	"time"
)

type Serv interface {
	Delete(ctx context.Context, id uuid.UUID) error
	GetAll(ctx context.Context) ([]*service.Service, error)
	GetByID(ctx context.Context, id uuid.UUID) (*service.Service, error)
	Store(ctx context.Context, u []*service.Service) error
	Update(ctx context.Context, u *service.Service) error
	TestService(ctx context.Context, id uuid.UUID) (*check.Check, error)
}

type serviceServ struct {
	repo repo2.Repo
	q    queue.Queue
	r    util.Store
}

func NewServiceServ(repo repo2.Repo) Serv {
	return &serviceServ{
		repo: repo,
	}
}

func (svc *serviceServ) Delete(ctx context.Context, id uuid.UUID) error {
	return svc.repo.Delete(ctx, id)
}

func (svc *serviceServ) GetAll(ctx context.Context) ([]*service.Service, error) {
	return svc.repo.GetAll(ctx)
}

func (svc *serviceServ) GetByID(ctx context.Context, id uuid.UUID) (*service.Service, error) {
	return svc.repo.GetByID(ctx, id)
}

func (svc *serviceServ) Store(ctx context.Context, u []*service.Service) error {
	return svc.repo.Store(ctx, u)
}

func (svc *serviceServ) Update(ctx context.Context, u *service.Service) error {
	return svc.repo.Update(ctx, u)
}

func (svc *serviceServ) TestService(ctx context.Context, id uuid.UUID) (*check.Check, error) {
	ser, err := svc.r.Service.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	p, err := svc.r.Property.GetAllByServiceID(ctx, id)
	if err != nil {
		return nil, err
	}
	h, err := svc.r.Host.GetByID(ctx, ser.HostID)
	if err != nil {
		return nil, err
	}
	serGrp, err := svc.r.ServiceGroup.GetByID(ctx, ser.ServiceGroupID)
	if err != nil {
		return nil, err
	}
	response, berr, err := svc.q.Send([]*queueing.ScoringData{
		{Service: queueing.QService{ID: id, Name: ser.Name, Group: serGrp.Name}, Host: *h.Address, Deadline: time.Now().Add(time.Second * 5), RoundID: 0, Properties: run.PropertyToMap(p)},
	})
	if response == nil || len(response) != 1 || err != nil {
		return nil, fmt.Errorf("something went wrong, either check took too long to execute, or the workers did not receive the check. err: %v, berr: %v", err, berr)
	}
	if berr != nil {
		response[0].Err += berr.Error()
	}
	return &check.Check{
		ServiceID: response[0].Service.ID,
		RoundID:   response[0].RoundID,
		Log:       response[0].Log,
		Err:       response[0].Err,
		Passed:    &response[0].Passed,
	}, nil
}
