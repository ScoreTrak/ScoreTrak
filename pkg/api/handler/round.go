package handler

import (
	"ScoreTrak/pkg/logger"
	"ScoreTrak/pkg/round"
	"net/http"
)

type roundController struct {
	log logger.LogInfoFormat
	svc round.Serv
}

func NewRoundController(log logger.LogInfoFormat, svc round.Serv) *roundController {
	return &roundController{log, svc}
}

func (c *roundController) GetLastRound(w http.ResponseWriter, r *http.Request) {
	genericGet(c.svc, c.log, "GetLastRound", w, r)
}
