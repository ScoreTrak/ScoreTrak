package handler

import (
	"encoding/json"
	"github.com/ScoreTrak/ScoreTrak/pkg/logger"
	"github.com/ScoreTrak/ScoreTrak/pkg/team"
	"github.com/qor/validations"
	"net/http"
)

type teamController struct {
	log logger.LogInfoFormat
	svc team.Serv
}

func NewTeamController(log logger.LogInfoFormat, svc team.Serv) *teamController {
	return &teamController{log, svc}
}

func (t *teamController) Store(w http.ResponseWriter, r *http.Request) {
	var tm []*team.Team
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

func (t *teamController) Delete(w http.ResponseWriter, r *http.Request) {
	genericDelete(t.svc, t.log, "Delete", "id", w, r)
}

func (t *teamController) GetByID(w http.ResponseWriter, r *http.Request) {
	genericGetByID(t.svc, t.log, "GetByID", "id", w, r)
}

func (t *teamController) GetAll(w http.ResponseWriter, r *http.Request) {
	genericGet(t.svc, t.log, "GetAll", w, r)
}

func (t *teamController) Update(w http.ResponseWriter, r *http.Request) {
	tm := &team.Team{}
	genericUpdate(t.svc, tm, t.log, "Update", "id", w, r)
}
