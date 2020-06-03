package handler

import (
	"ScoreTrak/pkg/host_group"
	"ScoreTrak/pkg/logger"
	"ScoreTrak/pkg/storage/orm"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type hostGroupController struct {
	log logger.LogInfoFormat
	svc host_group.Serv
}

func NewHostGroupController(log logger.LogInfoFormat, svc host_group.Serv) *hostGroupController {
	return &hostGroupController{log, svc}
}

func (s *hostGroupController) Store(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var sg host_group.HostGroup
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

func (s *hostGroupController) Delete(w http.ResponseWriter, r *http.Request) {

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
		http.Redirect(w, r, "/host_group", http.StatusNotModified)
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		s.log.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func (s *hostGroupController) GetByID(w http.ResponseWriter, r *http.Request) {

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

func (s *hostGroupController) GetAll(w http.ResponseWriter, r *http.Request) {

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

func (s *hostGroupController) Update(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	decoder := json.NewDecoder(r.Body)
	var sg host_group.HostGroup
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
