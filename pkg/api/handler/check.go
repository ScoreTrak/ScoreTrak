package handler

import (
	"ScoreTrak/pkg/check"
	"ScoreTrak/pkg/logger"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

type checkController struct {
	log logger.LogInfoFormat
	svc check.Serv
}

func NewCheckController(log logger.LogInfoFormat, svc check.Serv) *checkController {
	return &checkController{log, svc}
}

func (c *checkController) GetAllByRoundID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	rID, err := strconv.ParseUint(params["RoundID"], 10, 64)
	if err != nil {
		c.log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sg, err := c.svc.GetAllByRoundID(rID)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			c.log.Error(err)
		}
		return
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(sg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		c.log.Error(err)
	}
}

func (c *checkController) GetByRoundServiceID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	rID, err := strconv.ParseUint(params["RoundID"], 10, 64)
	if err != nil {
		c.log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sID, err := strconv.ParseUint(params["ServiceID"], 10, 64)
	if err != nil {
		c.log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sg, err := c.svc.GetByRoundServiceID(rID, sID)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			c.log.Error(err)
		}
		return
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(sg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		c.log.Error(err)
	}
}
