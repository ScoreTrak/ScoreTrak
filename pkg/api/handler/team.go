package handler

import (
	"ScoreTrak/pkg/logger"
	"ScoreTrak/pkg/team"
	"encoding/json"
	"github.com/gorilla/mux"
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
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	decoder := json.NewDecoder(r.Body)

	var tm team.Team
	err := decoder.Decode(&tm)

	if err != nil {
		t.log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = t.svc.Store(&tm)
	if err != nil {
		t.log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (t *teamController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	params := mux.Vars(r)
	err := t.svc.Delete(params["TeamID"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		t.log.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (t *teamController) GetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	params := mux.Vars(r)
	tm, err := t.svc.GetByID(params["TeamID"])

	if tm == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		t.log.Error(err)
		return
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(tm)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		t.log.Error(err)
	}
}

func (t *teamController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	teams, err := t.svc.GetAll()
	if len(teams) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		t.log.Error(err)
		return
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(teams)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		t.log.Error(err)
	}
}

func (t *teamController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	params := mux.Vars(r)
	decoder := json.NewDecoder(r.Body)
	var tm team.Team
	err := decoder.Decode(&tm)
	tm.ID = params["TeamID"]
	if err != nil {
		t.log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = t.svc.Update(&tm)
	if err != nil {
		t.log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
