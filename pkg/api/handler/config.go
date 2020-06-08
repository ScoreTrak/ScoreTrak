package handler

import (
	"ScoreTrak/pkg/config"
	"ScoreTrak/pkg/logger"
	"encoding/json"
	"github.com/qor/validations"
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
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(sg)
	if err != nil {
		c.log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = c.svc.Update(sg)
	if err != nil {
		_, ok := err.(*validations.Error)
		if ok {
			w.WriteHeader(http.StatusPreconditionFailed)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		c.log.Error(err)
		return
	}
}

func (c *configController) Get(w http.ResponseWriter, r *http.Request) {
	genericGet(c.svc, c.log, "Get", w, r)
}
