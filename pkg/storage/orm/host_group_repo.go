package orm

import (
	"context"
	"errors"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/host_group"
	"github.com/ScoreTrak/ScoreTrak/pkg/host_group/host_group_repo"

	"github.com/ScoreTrak/ScoreTrak/pkg/storage/orm/testutil"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type hostGroupRepo struct {
	db *gorm.DB
}

func NewHostGroupRepo(db *gorm.DB) host_group_repo.Repo {
	return &hostGroupRepo{db}
}

func (h *hostGroupRepo) Delete(ctx context.Context, id uuid.UUID) error {
	result := h.db.WithContext(ctx).Delete(&host_group.HostGroup{}, "id = ?", id)
	if result.Error != nil {
		errMsg := fmt.Sprintf("error while deleting the host with id : %d", id)
		return errors.New(errMsg)
	}
	if result.RowsAffected == 0 {
		return &NoRowsAffected{"no model found"}
	}
	return nil
}

func (h *hostGroupRepo) GetAll(ctx context.Context) ([]*host_group.HostGroup, error) {
	hostGroups := make([]*host_group.HostGroup, 0)
	err := h.db.WithContext(ctx).Find(&hostGroups).Error
	if err != nil {
		return nil, err
	}
	return hostGroups, nil
}

func (h *hostGroupRepo) GetByID(ctx context.Context, id uuid.UUID) (*host_group.HostGroup, error) {
	hstgrp := &host_group.HostGroup{}
	err := h.db.WithContext(ctx).Where("id = ?", id).First(hstgrp).Error
	if err != nil {
		return nil, err
	}
	return hstgrp, nil
}

func (h *hostGroupRepo) Store(ctx context.Context, hstgrp []*host_group.HostGroup) error {

	err := h.db.WithContext(ctx).Create(hstgrp).Error
	if err != nil {
		return err
	}

	return nil
}

func (h *hostGroupRepo) Upsert(ctx context.Context, hstgrp []*host_group.HostGroup) error {
	err := h.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(hstgrp).Error
	if err != nil {
		return err
	}
	return nil
}

func (h *hostGroupRepo) Update(ctx context.Context, hstgrp *host_group.HostGroup) error {
	err := h.db.WithContext(ctx).Model(hstgrp).Updates(host_group.HostGroup{Name: hstgrp.Name, Pause: hstgrp.Pause, Hide: hstgrp.Hide}).Error
	if err != nil {
		return err
	}
	return nil
}

func (h *hostGroupRepo) TruncateTable(ctx context.Context) (err error) {
	err = testutil.TruncateTable(ctx, &host_group.HostGroup{}, h.db)
	if err != nil {
		return err
	}
	return nil
}
