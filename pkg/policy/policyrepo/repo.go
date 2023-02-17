package policyrepo

import (
	"context"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy"
)

type Repo interface {
	Create(context.Context, *policy.Policy) error
	Get(ctx context.Context) (*policy.Policy, error)
	Upsert(ctx context.Context, pol *policy.Policy) error
	Update(ctx context.Context, u *policy.Policy) error
}
