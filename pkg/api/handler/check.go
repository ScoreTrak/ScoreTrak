package handler

import (
	"encoding/json"
	"errors"
	"github.com/L1ghtman2k/ScoreTrak/pkg/check"
	"github.com/L1ghtman2k/ScoreTrak/pkg/logger"
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
	genericGetByID(c.svc, c.log, "GetAllByRoundID", "RoundID", w, r)
}

func (c *checkController) GetByRoundServiceID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	rID, err := strconv.ParseUint(params["RoundID"], 10, 32)
	if err != nil {
		c.log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sID, err := strconv.ParseUint(params["ServiceID"], 10, 32)
	if err != nil {
		c.log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	rID32 := uint32(rID)
	sID32 := uint32(sID)
	sg, err := c.svc.GetByRoundServiceID(rID32, sID32)
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
