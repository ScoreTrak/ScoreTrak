package serviceservice

import (
	"context"
	"fmt"
	"time"

	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/host/hostrepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/property"
	"github.com/ScoreTrak/ScoreTrak/pkg/property/propertyrepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue/queueing"
	"github.com/ScoreTrak/ScoreTrak/pkg/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/service/servicerepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/servicegroup/servicegrouprepo"
	"github.com/gofrs/uuid"
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
	repo             servicerepo.Repo
	hostRepo         hostrepo.Repo
	propertyRepo     propertyrepo.Repo
	serviceGroupRepo servicegrouprepo.Repo
	q                queue.WorkerQueue
}

func NewServiceServ(repo servicerepo.Repo, q queue.WorkerQueue, pr propertyrepo.Repo, hr hostrepo.Repo, srgr servicegrouprepo.Repo) Serv {
	return &serviceServ{
		q:                q,
		repo:             repo,
		hostRepo:         hr,
		propertyRepo:     pr,
		serviceGroupRepo: srgr,
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
	currentService, err := svc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	properties, err := svc.propertyRepo.GetAllByServiceID(ctx, id)
	if err != nil {
		return nil, err
	}
	hosts, err := svc.hostRepo.GetByID(ctx, currentService.HostID)
	if err != nil {
		return nil, err
	}
	serGrp, err := svc.serviceGroupRepo.GetByID(ctx, currentService.ServiceGroupID)
	if err != nil {
		return nil, err
	}
	response, berr, err := svc.q.Send([]*queueing.ScoringData{
		{Service: queueing.QService{ID: id, Name: currentService.Name, Group: serGrp.Name}, MasterTime: time.Now(), Host: hosts.Address, Deadline: time.Now().Add(time.Second * 5), RoundID: 0, Properties: property.PropertiesToMap(properties)},
	})
	if response == nil || len(response) != 1 || response[0] == nil || err != nil {
		return nil, fmt.Errorf("something went wrong, either check took too long to execute, or the workers did not receive the check. err: %w, berr: %v", err, berr)
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
