package orm

import (
	"ScoreTrak/pkg/host"
	"ScoreTrak/pkg/logger"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

type hostRepo struct {
	db  *gorm.DB
	log logger.LogInfoFormat
}

func NewHostRepo(db *gorm.DB, log logger.LogInfoFormat) host.Repo {
	return &hostRepo{db, log}
}

func (h *hostRepo) Delete(id uint64) error {
	h.log.Debugf("deleting the host with id : %h", id)

	if h.db.Delete(&host.Host{}, "host_id = ?", id).Error != nil {
		errMsg := fmt.Sprintf("error while deleting the host with id : %h", id)
		h.log.Errorf(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

func (h *hostRepo) GetAll() ([]*host.Host, error) {
	h.log.Debug("get all the hosts")

	hosts := make([]*host.Host, 0)
	err := h.db.Find(&hosts).Error
	if err != nil {
		h.log.Debug("not a single host found")
		return nil, err
	}
	return hosts, nil
}

func (h *hostRepo) GetByID(id uint64) (*host.Host, error) {
	h.log.Debugf("get host details by id : %h", id)

	hst := &host.Host{}
	err := h.db.Where("host_id = ?", id).First(&hst).Error
	if err != nil {
		h.log.Errorf("host not found with id : %h, reason : %v", id, err)
		return nil, err
	}
	return hst, nil
}

func (h *hostRepo) Store(hst *host.Host) error {
	h.log.Debugf("creating the host with id : %v", hst.ID)

	err := h.db.Create(&hst).Error
	if err != nil {
		h.log.Errorf("error while creating the host, reason : %v", err)
		return err
	}
	return nil
}

func (h *hostRepo) Update(hst *host.Host) error {
	h.log.Debugf("updating the host, host_id : %v", hst.ID)
	err := h.db.Model(&hst).Updates(host.Host{}).Error
	if err != nil {
		h.log.Errorf("error while updating the host, reason : %v", err)
		return err
	}
	return nil
}