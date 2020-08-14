package handler

import (
	"encoding/json"
	"github.com/ScoreTrak/ScoreTrak/pkg/competition"
	"github.com/ScoreTrak/ScoreTrak/pkg/logger"
	"github.com/qor/validations"
	"net/http"
)

type competitionController struct {
	log logger.LogInfoFormat
	svc competition.Serv
}

func NewCompetitionController(log logger.LogInfoFormat, svc competition.Serv) *competitionController {
	return &competitionController{log, svc}
}

func (t *competitionController) LoadCompetition(w http.ResponseWriter, r *http.Request) {
	var tm *competition.Competition
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&tm)
	if err != nil {
		t.log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = t.svc.LoadCompetition(tm)
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

func (t *competitionController) FetchCoreCompetition(w http.ResponseWriter, r *http.Request) {
	genericGet(t.svc, t.log, "FetchCoreCompetition", w, r)
}

func (t *competitionController) FetchEntireCompetition(w http.ResponseWriter, r *http.Request) {
	genericGet(t.svc, t.log, "FetchEntireCompetition", w, r)
}
