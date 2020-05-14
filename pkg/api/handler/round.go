package handler

import (
	"ScoreTrak/pkg/logger"
	"ScoreTrak/pkg/round"
	"encoding/json"
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
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rnd, err := c.svc.GetLastRound()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		c.log.Error(err)
		return
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(rnd)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		c.log.Error(err)
		return
	}
}
