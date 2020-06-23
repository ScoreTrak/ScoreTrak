package handler

import (
	"github.com/L1ghtman2k/ScoreTrak/pkg/logger"
	"github.com/L1ghtman2k/ScoreTrak/pkg/service_group"
	"net/http"
)

type serviceGroupController struct {
	log logger.LogInfoFormat
	svc service_group.Serv
}

func NewServiceGroupController(log logger.LogInfoFormat, svc service_group.Serv) *serviceGroupController {
	return &serviceGroupController{log, svc}
}

func (s *serviceGroupController) Store(w http.ResponseWriter, r *http.Request) {
	tm := &service_group.ServiceGroup{}
	genericStore(s.svc, tm, s.log, "Store", w, r)
	//ToDO: Implement checking if queue is accessible (First via docker, then via ping)
}

func (s *serviceGroupController) Delete(w http.ResponseWriter, r *http.Request) {
	genericDelete(s.svc, s.log, "Delete", "id", w, r)
}

func (s *serviceGroupController) GetByID(w http.ResponseWriter, r *http.Request) {
	genericGetByID(s.svc, s.log, "GetByID", "id", w, r)
}

func (s *serviceGroupController) GetAll(w http.ResponseWriter, r *http.Request) {
	genericGet(s.svc, s.log, "GetAll", w, r)
}

func (s *serviceGroupController) Update(w http.ResponseWriter, r *http.Request) {
	tm := &service_group.ServiceGroup{}
	genericUpdate(s.svc, tm, s.log, "Update", "id", w, r)
	//ToDO: Implement checking if queue is accessible (First via docker, then via ping)
}
