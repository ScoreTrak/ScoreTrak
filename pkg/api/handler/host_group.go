package handler

import (
	"ScoreTrak/pkg/host_group"
	"ScoreTrak/pkg/logger"
	"net/http"
)

type hostGroupController struct {
	log logger.LogInfoFormat
	svc host_group.Serv
}

func NewHostGroupController(log logger.LogInfoFormat, svc host_group.Serv) *hostGroupController {
	return &hostGroupController{log, svc}
}

func (t *hostGroupController) Store(w http.ResponseWriter, r *http.Request) {
	tm := &host_group.HostGroup{}
	genericStore(t.svc, tm, t.log, "Store", w, r)
}

func (t *hostGroupController) Delete(w http.ResponseWriter, r *http.Request) {
	genericDelete(t.svc, t.log, "Delete", "id", w, r)
}

func (t *hostGroupController) GetByID(w http.ResponseWriter, r *http.Request) {
	genericGetByID(t.svc, t.log, "GetByID", "id", w, r)
}

func (t *hostGroupController) GetAll(w http.ResponseWriter, r *http.Request) {
	genericGet(t.svc, t.log, "GetAll", w, r)
}

func (t *hostGroupController) Update(w http.ResponseWriter, r *http.Request) {
	tm := &host_group.HostGroup{}
	genericUpdate(t.svc, tm, t.log, "Update", "id", w, r)
}
