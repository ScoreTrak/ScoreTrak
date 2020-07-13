package handler

import (
	"encoding/json"
	"github.com/L1ghtman2k/ScoreTrak/pkg/check"
	"github.com/L1ghtman2k/ScoreTrak/pkg/logger"
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
	genericGetByID(c.svc, c.log, "GetAllByRoundID", "RoundID", w, r)
}

func (c *checkController) GetByRoundServiceID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	rID, err := strconv.ParseUint(params["RoundID"], 10, 64)
	if err != nil {
		c.log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sID, err := strconv.ParseUint(params["ServiceID"], 10, 64)
	if err != nil {
		c.log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sg, err := c.svc.GetByRoundServiceID(rID, sID)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
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
