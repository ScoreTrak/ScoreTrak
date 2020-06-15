package handler

import (
	"ScoreTrak/pkg/logger"
	"ScoreTrak/pkg/service"
	"net/http"
)

type serviceController struct {
	log logger.LogInfoFormat
	svc service.Serv
}

func NewServiceController(log logger.LogInfoFormat, svc service.Serv) *serviceController {
	return &serviceController{log, svc}
}

func (s *serviceController) Store(w http.ResponseWriter, r *http.Request) {
	tm := &service.Service{}
	genericStore(s.svc, tm, s.log, "Store", w, r)
}

func (s *serviceController) Delete(w http.ResponseWriter, r *http.Request) {
	genericDelete(s.svc, s.log, "Delete", "id", w, r)
}

func (s *serviceController) GetByID(w http.ResponseWriter, r *http.Request) {
	genericGetByID(s.svc, s.log, "GetByID", "id", w, r)
}

func (s *serviceController) GetAll(w http.ResponseWriter, r *http.Request) {
	genericGet(s.svc, s.log, "GetAll", w, r)
}

func (s *serviceController) Update(w http.ResponseWriter, r *http.Request) {
	tm := &service.Service{}
	genericUpdate(s.svc, tm, s.log, "Update", "id", w, r)
}
