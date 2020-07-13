package handler

import (
	"github.com/L1ghtman2k/ScoreTrak/pkg/logger"
	"github.com/L1ghtman2k/ScoreTrak/pkg/round"
	"net/http"
)

type roundController struct {
	log logger.LogInfoFormat
	svc round.Serv
}

func NewRoundController(log logger.LogInfoFormat, svc round.Serv) *roundController {
	return &roundController{log, svc}
}

func (c *roundController) GetLastNonElapsingRound(w http.ResponseWriter, r *http.Request) {
	genericGet(c.svc, c.log, "GetLastNonElapsingRound", w, r)
}
