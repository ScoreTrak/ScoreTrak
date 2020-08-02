package handler

import (
	"encoding/json"
	"errors"
	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/logger"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
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
	rID, err := strconv.ParseUint(params["RoundID"], 10, 32)
	if err != nil {
		c.log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sg, err := c.svc.GetAllByRoundID(uint(rID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			c.log.Error(err)
		}
		return
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(sg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		c.log.Error(err)
	}
}

func (c *checkController) GetByRoundServiceID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	rID, err := strconv.ParseUint(params["RoundID"], 10, 32)
	if err != nil {
		c.log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sID, err := uuidResolver("ServiceID", r)
	if err != nil {
		c.log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sg, err := c.svc.GetByRoundServiceID(uint(rID), sID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			c.log.Error(err)
		}
		return
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(sg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		c.log.Error(err)
	}
}
