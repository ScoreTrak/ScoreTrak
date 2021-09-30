package orm

import (
	"context"

	"github.com/ScoreTrak/ScoreTrak/pkg/policy"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy/policyrepo"
	"gorm.io/gorm"
)

type policyRepo struct {
	db *gorm.DB
}

func NewPolicyRepo(db *gorm.DB) policyrepo.Repo {
	return &policyRepo{db}
}

func (h *policyRepo) Get(ctx context.Context) (*policy.Policy, error) {
	p := &policy.Policy{}
	p.ID = 1
	h.db.WithContext(ctx).Take(p)
	return p, nil
}

func (h *policyRepo) Update(ctx context.Context, tm *policy.Policy) error {
	tm.ID = 1
	err := h.db.WithContext(ctx).Model(tm).Updates(policy.Policy{AllowUnauthenticatedUsers: tm.AllowUnauthenticatedUsers, ShowPoints: tm.ShowPoints,
		AllowChangingUsernamesAndPasswords: tm.AllowChangingUsernamesAndPasswords, ShowAddresses: tm.ShowAddresses, AllowRedTeamLaunchingServiceTestsManually: tm.AllowRedTeamLaunchingServiceTestsManually}).Error
	if err != nil {
		return err
	}
	return nil
}
