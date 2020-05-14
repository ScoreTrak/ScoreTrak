package handler

import (
	"ScoreTrak/pkg/host_group"
	"ScoreTrak/pkg/logger"
	"net/http"
)

type hostGroupController struct {
	log logger.LogInfoFormat
	svc host_group.Serv
}

func NewHostGroupController(log logger.LogInfoFormat, svc host_group.Serv) *hostGroupController {
	return &hostGroupController{log, svc}
}

func (h *hostGroupController) Store(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (h *hostGroupController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (h *hostGroupController) GetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (h *hostGroupController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (h *hostGroupController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
