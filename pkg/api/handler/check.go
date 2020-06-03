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

func (c *checkController) GetAllByRoundID(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (c *checkController) GetByRoundServiceID(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
