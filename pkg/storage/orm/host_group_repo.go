package orm

import (
	"context"
	"errors"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/host_group"
	"github.com/ScoreTrak/ScoreTrak/pkg/logger"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/orm/util"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type hostGroupRepo struct {
	db  *gorm.DB
	log logger.LogInfoFormat
}

func NewHostGroupRepo(db *gorm.DB, log logger.LogInfoFormat) host_group.Repo {
	return &hostGroupRepo{db, log}
}

func (h *hostGroupRepo) Delete(ctx context.Context, id uuid.UUID) error {
	h.log.Debugf("deleting the hostGroup with id : %h", id)
	result := h.db.WithContext(ctx).Delete(&host_group.HostGroup{}, "id = ?", id)
	if result.Error != nil {
		errMsg := fmt.Sprintf("error while deleting the host with id : %d", id)
		h.log.Errorf(errMsg)
		return errors.New(errMsg)
	}

	if result.RowsAffected == 0 {
		return &NoRowsAffected{"no model found"}
	}

	return nil
}

func (h *hostGroupRepo) GetAll(ctx context.Context) ([]*host_group.HostGroup, error) {
	h.log.Debug("get all the hostGroups")

	hostGroups := make([]*host_group.HostGroup, 0)
	err := h.db.WithContext(ctx).Find(&hostGroups).Error
	if err != nil {
		h.log.Debug("not a single hostGroup found")
		return nil, err
	}
	return hostGroups, nil
}

func (h *hostGroupRepo) GetByID(ctx context.Context, id uuid.UUID) (*host_group.HostGroup, error) {
	h.log.Debugf("get hostGroup details by id : %h", id)

	hstgrp := &host_group.HostGroup{}
	err := h.db.WithContext(ctx).Where("id = ?", id).First(hstgrp).Error
	if err != nil {
		h.log.Errorf("hostGroup not found with id : %h, reason : %v", id, err)
		return nil, err
	}
	return hstgrp, nil
}

func (h *hostGroupRepo) Store(ctx context.Context, hstgrp []*host_group.HostGroup) error {

	err := h.db.WithContext(ctx).Create(hstgrp).Error
	if err != nil {
		h.log.Errorf("error while creating the hostGroup, reason : %v", err)
		return err
	}

	return nil
}

func (h *hostGroupRepo) Upsert(ctx context.Context, hstgrp []*host_group.HostGroup) error {
	err := h.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(hstgrp).Error
	if err != nil {
		h.log.Errorf("error while creating the user, reason : %v", err)
		return err
	}
	return nil
}

func (h *hostGroupRepo) Update(ctx context.Context, hstgrp *host_group.HostGroup) error {
	h.log.Debugf("updating the hostGroup, id : %v", hstgrp.ID)
	err := h.db.WithContext(ctx).Model(hstgrp).Updates(host_group.HostGroup{Name: hstgrp.Name, Enabled: hstgrp.Enabled}).Error
	if err != nil {
		h.log.Errorf("error while updating the hostGroup, reason : %v", err)
		return err
	}
	return nil
}

func (h *hostGroupRepo) TruncateTable(ctx context.Context) (err error) {
	err = util.TruncateTable(ctx, &host_group.HostGroup{}, h.db)
	if err != nil {
		return err
	}
	return nil
}
