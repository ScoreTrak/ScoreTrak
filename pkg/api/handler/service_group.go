package handler

import (
	"encoding/json"
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/config"
	"github.com/L1ghtman2k/ScoreTrak/pkg/logger"
	"github.com/L1ghtman2k/ScoreTrak/pkg/platform"
	"github.com/L1ghtman2k/ScoreTrak/pkg/platform/worker"
	"github.com/L1ghtman2k/ScoreTrak/pkg/service_group"
	"github.com/qor/validations"
	"net/http"
)

type serviceGroupController struct {
	log logger.LogInfoFormat
	svc service_group.Serv
	p   platform.Platform
}

func NewServiceGroupController(log logger.LogInfoFormat, svc service_group.Serv, p platform.Platform) *serviceGroupController {
	return &serviceGroupController{log, svc, p}
}

func (s *serviceGroupController) Store(w http.ResponseWriter, r *http.Request) {
	tm := &service_group.ServiceGroup{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(tm)
	if err != nil {
		s.log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(tm, config.GetStaticConfig(), s.p)
	if s.p != nil && (tm.AllowPlatform == nil || *tm.AllowPlatform == true) && config.GetStaticConfig().Queue.Use != "none" {
		wr := worker.Info{Topic: tm.Name, Label: tm.Name}
		err := s.p.DeployWorkers(wr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			s.log.Error(err)
			return
		}
	}
	err = s.svc.Store(tm)
	if err != nil {
		_, ok := err.(*validations.Error)
		if ok {
			http.Error(w, err.Error(), http.StatusPreconditionFailed)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		s.log.Error(err)
		return
	}
} //Todo: Use queue to ping the worker to ensure that everything is working

func (s *serviceGroupController) Delete(w http.ResponseWriter, r *http.Request) {
	genericDelete(s.svc, s.log, "Delete", "id", w, r)
} // Todo: Implement Removal of service Group

func (s *serviceGroupController) GetByID(w http.ResponseWriter, r *http.Request) {
	genericGetByID(s.svc, s.log, "GetByID", "id", w, r)
}

func (s *serviceGroupController) GetAll(w http.ResponseWriter, r *http.Request) {
	genericGet(s.svc, s.log, "GetAll", w, r)
}

func (s *serviceGroupController) Update(w http.ResponseWriter, r *http.Request) {
	tm := &service_group.ServiceGroup{}
	genericUpdate(s.svc, tm, s.log, "Update", "id", w, r)
}
