package handler

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/logger"
	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	"net/http"
)

type reportController struct {
	log logger.LogInfoFormat
	svc report.Serv
}

func NewReportController(log logger.LogInfoFormat, svc report.Serv) *reportController {
	return &reportController{log, svc}
}

func (c *reportController) Get(w http.ResponseWriter, r *http.Request) {
	genericGet(c.svc, c.log, "Get", w, r)
}
