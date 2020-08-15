package handler

import (
	"encoding/json"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/logger"
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = c.svc.Update(sg)
	if err != nil {
		_, ok := err.(*validations.Error)
		if ok {
			http.Error(w, err.Error(), http.StatusPreconditionFailed)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		c.log.Error(err)
		return
	}
}

func (c *configController) Get(w http.ResponseWriter, r *http.Request) {
	genericGet(c.svc, c.log, "Get", w, r)
}

func (c configController) ResetScores(w http.ResponseWriter, r *http.Request) {
	cnf, err := c.svc.Get()
	if err != nil {
		c.log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if cnf.Enabled != nil && *cnf.Enabled == false {
		err := c.svc.ResetScores()
		if err != nil {
			c.log.Error(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		http.Error(w, "Competition must be disabled first", http.StatusPreconditionFailed)
		return
	}

}

func (c configController) DeleteCompetition(w http.ResponseWriter, r *http.Request) {
	cnf, err := c.svc.Get()
	if err != nil {
		c.log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if cnf.Enabled != nil && *cnf.Enabled == false {
		err := c.svc.DeleteCompetition()
		if err != nil {
			c.log.Error(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		http.Error(w, "Competition must be disabled first", http.StatusPreconditionFailed)
		return
	}

}

type staticConfigController struct {
	log logger.LogInfoFormat
	svc config.StaticServ
}

func NewStaticConfigController(log logger.LogInfoFormat, svc config.StaticServ) *staticConfigController {
	return &staticConfigController{log, svc}
}

func (c staticConfigController) Get(w http.ResponseWriter, r *http.Request) {
	genericGet(c.svc, c.log, "Get", w, r)
}
