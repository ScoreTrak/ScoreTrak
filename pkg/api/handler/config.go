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
	sg := &config.DynamicConfig{}
	genericUpdate(c.svc, sg, c.log, "Update", w, r)
}

func (c *configController) Get(w http.ResponseWriter, r *http.Request) {
	genericGet(c.svc, c.log, "Get", w, r)

}
