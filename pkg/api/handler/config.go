package handler

import (
	"ScoreTrak/pkg/config"
	"ScoreTrak/pkg/logger"
	"net/http"
)

type configController struct {
	log logger.LogInfoFormat
	svc config.Serv
}

func NewConfigController(log logger.LogInfoFormat, svc config.Serv) *configController {
	return &configController{log, svc}
}

func (c *configController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (c *configController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
