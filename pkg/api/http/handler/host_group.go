package handler

import (
	"encoding/json"
	"github.com/ScoreTrak/ScoreTrak/pkg/host_group"
	"github.com/ScoreTrak/ScoreTrak/pkg/logger"
	"github.com/qor/validations"
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
	var tm []*host_group.HostGroup
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&tm)
	if err != nil {
		t.log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = t.svc.Store(tm)
	if err != nil {
		_, ok := err.(*validations.Error)
		if ok {
			http.Error(w, err.Error(), http.StatusPreconditionFailed)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		t.log.Error(err)
		return
	}
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
