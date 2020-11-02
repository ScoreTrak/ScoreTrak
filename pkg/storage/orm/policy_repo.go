package orm

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/policy"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy/repo"
	"gorm.io/gorm"
)

type policyRepo struct {
	db *gorm.DB
}

func NewPolicyRepo(db *gorm.DB) repo.Repo {
	return &policyRepo{db}
}

func (h *policyRepo) Get() (*policy.Policy, error) {
	p := &policy.Policy{}
	p.ID = 1
	h.db.Take(p)
	return p, nil
}

func (h *policyRepo) Update(tm *policy.Policy) error {
	tm.ID = 1
	err := h.db.Model(tm).Updates(policy.Policy{AllowUnauthenticatedUsers: tm.AllowUnauthenticatedUsers, ShowPoints: tm.ShowPoints,
		AllowChangingUsernamesAndPasswords: tm.AllowChangingUsernamesAndPasswords, ShowAddresses: tm.ShowAddresses}).Error
	if err != nil {
		return err
	}
	return nil
}
