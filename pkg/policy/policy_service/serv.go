package policy_service

import (
	"context"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy"
	repo2 "github.com/ScoreTrak/ScoreTrak/pkg/policy/policy_repo"
)

type Serv interface {
	Get(ctx context.Context) (*policy.Policy, error)
	Update(ctx context.Context, u *policy.Policy) error
}

type policyServ struct {
	repo repo2.Repo
}

func NewPolicyServ(repo repo2.Repo) Serv {
	return &policyServ{
		repo: repo,
	}
}

func (svc *policyServ) Get(ctx context.Context) (*policy.Policy, error) { return svc.repo.Get(ctx) }

func (svc *policyServ) Update(ctx context.Context, u *policy.Policy) error {
	return svc.repo.Update(ctx, u)
}
