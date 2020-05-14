package handler

import (
	"ScoreTrak/pkg/check"
	"ScoreTrak/pkg/logger"
	"net/http"
)

type checkController struct {
	log logger.LogInfoFormat
	svc check.Serv
}

func NewCheckController(log logger.LogInfoFormat, svc check.Serv) *checkController {
	return &checkController{log, svc}
}

func (c *checkController) GetAllByTeamRoundID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (c *checkController) GetByTeamRoundServiceID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
