package handler

import (
	"ScoreTrak/pkg/logger"
	"ScoreTrak/pkg/team"
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
	tm := &team.Team{}
	genericStore(t.svc, tm, t.log, "Store", w, r)
}

func (t *teamController) DeleteByName(w http.ResponseWriter, r *http.Request) {
	genericDelete(t.svc, t.log, "DeleteByName", "name", w, r)
}

func (t *teamController) GetByName(w http.ResponseWriter, r *http.Request) {
	genericGetByID(t.svc, t.log, "GetByName", "name", w, r)
}

func (t *teamController) GetAll(w http.ResponseWriter, r *http.Request) {
	genericGet(t.svc, t.log, "GetAll", w, r)
}

func (t *teamController) Update(w http.ResponseWriter, r *http.Request) {
	tm := &team.Team{}
	genericUpdate(t.svc, tm, t.log, "Update", "name", w, r)
}
