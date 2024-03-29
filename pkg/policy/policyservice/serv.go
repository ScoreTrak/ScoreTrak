package policyservice

import (
	"context"
	"fmt"

	"github.com/ScoreTrak/ScoreTrak/pkg/policy"
	repo2 "github.com/ScoreTrak/ScoreTrak/pkg/policy/policyrepo"
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

func (svc *policyServ) Get(ctx context.Context) (*policy.Policy, error) {
	p, err := svc.repo.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to make a get call: %w", err)
	}
	return p, nil
}

func (svc *policyServ) Update(ctx context.Context, u *policy.Policy) error {
	if err := svc.repo.Update(ctx, u); err != nil {
		return fmt.Errorf("failed to make update call: %w", err)
	}
	return nil
}
