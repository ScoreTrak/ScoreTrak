package orm

import (
	"errors"
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/host_group"
	"github.com/L1ghtman2k/ScoreTrak/pkg/logger"
	"gorm.io/gorm"
)

type hostGroupRepo struct {
	db  *gorm.DB
	log logger.LogInfoFormat
}

func NewHostGroupRepo(db *gorm.DB, log logger.LogInfoFormat) host_group.Repo {
	return &hostGroupRepo{db, log}
}

func (h *hostGroupRepo) Delete(id uint64) error {
	h.log.Debugf("deleting the hostGroup with id : %h", id)
	result := h.db.Delete(&host_group.HostGroup{}, "id = ?", id)
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

func (h *hostGroupRepo) GetAll() ([]*host_group.HostGroup, error) {
	h.log.Debug("get all the hostGroups")

	hostGroups := make([]*host_group.HostGroup, 0)
	err := h.db.Find(&hostGroups).Error
	if err != nil {
		h.log.Debug("not a single hostGroup found")
		return nil, err
	}
	return hostGroups, nil
}

func (h *hostGroupRepo) GetByID(id uint64) (*host_group.HostGroup, error) {
	h.log.Debugf("get hostGroup details by id : %h", id)

	hstgrp := &host_group.HostGroup{}
	err := h.db.Where("id = ?", id).First(hstgrp).Error
	if err != nil {
		h.log.Errorf("hostGroup not found with id : %h, reason : %v", id, err)
		return nil, err
	}
	return hstgrp, nil
}

func (h *hostGroupRepo) Store(hstgrp *host_group.HostGroup) error {
	h.log.Debugf("creating the hostGroup with id : %v", hstgrp.ID)

	err := h.db.Create(hstgrp).Error
	if err != nil {
		h.log.Errorf("error while creating the hostGroup, reason : %v", err)
		return err
	}

	return nil
}

func (h *hostGroupRepo) Update(hstgrp *host_group.HostGroup) error {
	h.log.Debugf("updating the hostGroup, id : %v", hstgrp.ID)
	err := h.db.Model(hstgrp).Updates(host_group.HostGroup{Name: hstgrp.Name, Enabled: hstgrp.Enabled}).Error
	if err != nil {
		h.log.Errorf("error while updating the hostGroup, reason : %v", err)
		return err
	}
	return nil
}
