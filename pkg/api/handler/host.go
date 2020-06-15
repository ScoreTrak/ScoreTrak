package handler

import (
	"ScoreTrak/pkg/host"
	"ScoreTrak/pkg/logger"
	"net/http"
)

type hostController struct {
	log logger.LogInfoFormat
	svc host.Serv
}

func NewHostController(log logger.LogInfoFormat, svc host.Serv) *hostController {
	return &hostController{log, svc}
}

func (t *hostController) Store(w http.ResponseWriter, r *http.Request) {
	tm := &host.Host{}
	genericStore(t.svc, tm, t.log, "Store", w, r)
}

func (t *hostController) Delete(w http.ResponseWriter, r *http.Request) {
	genericDelete(t.svc, t.log, "Delete", "id", w, r)
}

func (t *hostController) GetByID(w http.ResponseWriter, r *http.Request) {
	genericGetByID(t.svc, t.log, "GetByID", "id", w, r)
}

func (t *hostController) GetAll(w http.ResponseWriter, r *http.Request) {
	genericGet(t.svc, t.log, "GetAll", w, r)
}

func (t *hostController) Update(w http.ResponseWriter, r *http.Request) {
	tm := &host.Host{}
	genericUpdate(t.svc, tm, t.log, "Update", "id", w, r)
}
