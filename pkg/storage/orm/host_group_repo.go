package orm

import (
	"ScoreTrak/pkg/host_group"
	"ScoreTrak/pkg/logger"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
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

	if h.db.Delete(&host_group.HostGroup{}, "hostGroup_id = ?", id).Error != nil {
		errMsg := fmt.Sprintf("error while deleting the hostGroup with id : %h", id)
		h.log.Errorf(errMsg)
		return errors.New(errMsg)
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
	err := h.db.Where("hostGroup_id = ?", id).First(&hstgrp).Error
	if err != nil {
		h.log.Errorf("hostGroup not found with id : %h, reason : %v", id, err)
		return nil, err
	}
	return hstgrp, nil
}

func (h *hostGroupRepo) Store(hstgrp *host_group.HostGroup) error {
	h.log.Debugf("creating the hostGroup with id : %v", hstgrp.ID)

	err := h.db.Create(&hstgrp).Error
	if err != nil {
		h.log.Errorf("error while creating the hostGroup, reason : %v", err)
		return err
	}

	return nil
}

func (h *hostGroupRepo) Update(hstgrp *host_group.HostGroup) error {
	h.log.Debugf("updating the hostGroup, hostGroup_id : %v", hstgrp.ID)
	err := h.db.Model(&hstgrp).Updates(host_group.HostGroup{}).Error
	if err != nil {
		h.log.Errorf("error while updating the hostGroup, reason : %v", err)
		return err
	}
	return nil
}
