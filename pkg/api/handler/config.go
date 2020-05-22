package handler

import (
	"ScoreTrak/pkg/config"
	"ScoreTrak/pkg/logger"
	"encoding/json"
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
	decoder := json.NewDecoder(r.Body)
	var sg config.DynamicConfig
	err := decoder.Decode(&sg)
	if err != nil {
		c.log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = c.svc.Update(&sg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *configController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	conf, err := c.svc.Get()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		c.log.Error(err)
		return
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(conf)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		c.log.Error(err)
	}

}
