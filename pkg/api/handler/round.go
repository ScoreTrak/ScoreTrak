package handler

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/logger"
	"github.com/ScoreTrak/ScoreTrak/pkg/round"
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

func (c *roundController) GetLastRound(w http.ResponseWriter, r *http.Request) {
	genericGet(c.svc, c.log, "GetLastRound", w, r)
}

func (c *roundController) GetAll(w http.ResponseWriter, r *http.Request) {
	genericGet(c.svc, c.log, "GetAll", w, r)
}

func (c *roundController) GetByID(w http.ResponseWriter, r *http.Request) {
	genericGetByID(c.svc, c.log, "GetByID", "id", w, r)
}
