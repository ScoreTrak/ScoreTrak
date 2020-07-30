package handler

import (
	"encoding/json"
	"github.com/L1ghtman2k/ScoreTrak/pkg/host"
	"github.com/L1ghtman2k/ScoreTrak/pkg/logger"
	"github.com/qor/validations"
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
	var tm []*host.Host
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
