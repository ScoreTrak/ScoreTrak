package handler

import (
	"ScoreTrak/pkg/logger"
	"ScoreTrak/pkg/service"
	"ScoreTrak/pkg/storage/orm"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type serviceController struct {
	log logger.LogInfoFormat
	svc service.Serv
}

func NewServiceController(log logger.LogInfoFormat, svc service.Serv) *serviceController {
	return &serviceController{log, svc}
}

func (s *serviceController) Store(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	decoder := json.NewDecoder(r.Body)
	var sg service.Service
	err := decoder.Decode(&sg)
	if err != nil {
		s.log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = s.svc.Store(&sg)
	if err != nil {
		s.log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *serviceController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		s.log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = s.svc.Delete(id)
	_, ok := err.(*orm.NoRowsAffected)
	if ok {
		http.Redirect(w, r, "/service", http.StatusNotModified)
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		s.log.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func (s *serviceController) GetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		s.log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sg, err := s.svc.GetByID(id)
	if sg == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		s.log.Error(err)
		return
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(sg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		s.log.Error(err)
	}
}

func (s *serviceController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	sgs, err := s.svc.GetAll()
	if len(sgs) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		s.log.Error(err)
		return
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(sgs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		s.log.Error(err)
	}

}

func (s *serviceController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	params := mux.Vars(r)
	decoder := json.NewDecoder(r.Body)
	var sg service.Service
	err := decoder.Decode(&sg)
	if err != nil {
		s.log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sg.ID, err = strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		s.log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = s.svc.Update(&sg)
	if err != nil {
		s.log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
