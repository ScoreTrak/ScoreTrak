package handler

import (
	"ScoreTrak/pkg/logger"
	"ScoreTrak/pkg/service_group"
	"net/http"
)

type serviceGroupController struct {
	log logger.LogInfoFormat
	svc service_group.Serv
}

func NewServiceGroupController(log logger.LogInfoFormat, svc service_group.Repo) *serviceGroupController {
	return &serviceGroupController{log, svc}
}

func (*serviceGroupController) Store(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (*serviceGroupController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (*serviceGroupController) GetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (*serviceGroupController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (*serviceGroupController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
