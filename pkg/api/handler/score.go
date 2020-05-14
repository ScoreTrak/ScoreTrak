package handler

import (
	"ScoreTrak/pkg/logger"
	"net/http"
)

type scoreController struct {
	log logger.LogInfoFormat
}

func NewScoreController(log logger.LogInfoFormat) *scoreController {
	return &scoreController{log}
}

func (*scoreController) GetScore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (*scoreController) GetScores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
