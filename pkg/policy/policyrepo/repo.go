package policyrepo

import (
	"context"

	"github.com/ScoreTrak/ScoreTrak/pkg/policy"
)

type Repo interface {
	Get(ctx context.Context) (*policy.Policy, error)
	Upsert(ctx context.Context, pol *policy.Policy) error
	Update(ctx context.Context, u *policy.Policy) error
}
