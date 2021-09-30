package orm

import (
	"context"
	"errors"
	"fmt"

	"github.com/ScoreTrak/ScoreTrak/pkg/hostgroup"
	"github.com/ScoreTrak/ScoreTrak/pkg/hostgroup/hostgrouprepo"

	"github.com/ScoreTrak/ScoreTrak/pkg/storage/orm/testutil"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type hostGroupRepo struct {
	db *gorm.DB
}

func NewHostGroupRepo(db *gorm.DB) hostgrouprepo.Repo {
	return &hostGroupRepo{db}
}

var ErrDeletingHostGroup = errors.New("error while deleting the host group with id")

func (h *hostGroupRepo) Delete(ctx context.Context, id uuid.UUID) error {
	result := h.db.WithContext(ctx).Delete(&hostgroup.HostGroup{}, "id = ?", id)
	if result.Error != nil {
		return fmt.Errorf("%w: %d", ErrDeletingHostGroup, id)
	}
	if result.RowsAffected == 0 {
		return &NoRowsAffected{"no model found"}
	}
	return nil
}

func (h *hostGroupRepo) GetAll(ctx context.Context) ([]*hostgroup.HostGroup, error) {
	hostGroups := make([]*hostgroup.HostGroup, 0)
	err := h.db.WithContext(ctx).Find(&hostGroups).Error
	if err != nil {
		return nil, err
	}
	return hostGroups, nil
}

func (h *hostGroupRepo) GetByID(ctx context.Context, id uuid.UUID) (*hostgroup.HostGroup, error) {
	hstgrp := &hostgroup.HostGroup{}
	err := h.db.WithContext(ctx).Where("id = ?", id).First(hstgrp).Error
	if err != nil {
		return nil, err
	}
	return hstgrp, nil
}

func (h *hostGroupRepo) Store(ctx context.Context, hstgrp []*hostgroup.HostGroup) error {
	err := h.db.WithContext(ctx).Create(hstgrp).Error
	if err != nil {
		return err
	}

	return nil
}

func (h *hostGroupRepo) Upsert(ctx context.Context, hstgrp []*hostgroup.HostGroup) error {
	err := h.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(hstgrp).Error
	if err != nil {
		return err
	}
	return nil
}

func (h *hostGroupRepo) Update(ctx context.Context, hstgrp *hostgroup.HostGroup) error {
	err := h.db.WithContext(ctx).Model(hstgrp).Updates(hostgroup.HostGroup{Name: hstgrp.Name, Pause: hstgrp.Pause, Hide: hstgrp.Hide}).Error
	if err != nil {
		return err
	}
	return nil
}

func (h *hostGroupRepo) TruncateTable(ctx context.Context) (err error) {
	err = testutil.TruncateTable(ctx, &hostgroup.HostGroup{}, h.db)
	if err != nil {
		return err
	}
	return nil
}
