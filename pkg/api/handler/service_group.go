package handler

import (
	"encoding/json"
	"errors"
	"github.com/L1ghtman2k/ScoreTrak/pkg/config"
	"github.com/L1ghtman2k/ScoreTrak/pkg/logger"
	"github.com/L1ghtman2k/ScoreTrak/pkg/platform"
	"github.com/L1ghtman2k/ScoreTrak/pkg/platform/worker"
	"github.com/L1ghtman2k/ScoreTrak/pkg/queue"
	"github.com/L1ghtman2k/ScoreTrak/pkg/service_group"
	"github.com/L1ghtman2k/ScoreTrak/pkg/storage/orm"
	"github.com/qor/validations"
	"net/http"
)

type serviceGroupController struct {
	log logger.LogInfoFormat
	svc service_group.Serv
	p   platform.Platform
	q   queue.Queue
}

func NewServiceGroupController(log logger.LogInfoFormat, svc service_group.Serv, p platform.Platform, q queue.Queue) *serviceGroupController {
	return &serviceGroupController{log, svc, p, q}
}

func (s *serviceGroupController) Store(w http.ResponseWriter, r *http.Request) {
	tm := &service_group.ServiceGroup{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(tm)
	if err != nil {
		s.log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if s.p != nil && !tm.SkipPlatform && config.GetStaticConfig().Queue.Use != "none" {
		if tm.Enabled != nil && *tm.Enabled == true {
			http.Error(w, "if you are letting scoretrak manage the workers, Enabled can be set to true, only after workers are deployed.", http.StatusPreconditionFailed)
			s.log.Error(err)
			return
		}
		wr := worker.Info{Topic: tm.Name, Label: tm.Label}
		err := s.p.DeployWorkers(wr) //Todo: Make sure that worker container is not allocated multiple times (Currently workers are duplicated)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			s.log.Error(err)
			return
		}
	}
	err = s.svc.Store(tm)
	if err != nil {
		_, ok := err.(*validations.Error)
		if ok {
			http.Error(w, err.Error(), http.StatusPreconditionFailed)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		s.log.Error(err)
		return
	}
}

func (s *serviceGroupController) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := idResolver(s.svc, "id", r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.log.Error(err)
		return
	}
	idUint, ok := id.(uint64)
	if !ok {
		http.Error(w, "failed to retrieve the id", http.StatusInternalServerError)
		s.log.Error(err)
		return
	}
	serviceGrp, err := s.svc.GetByID(idUint)
	if err != nil {
		err = errors.New("failed to retrieve the object")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.log.Error(err)
		return
	}
	err = s.svc.Delete(idUint)
	if err != nil {
		_, ok := err.(*orm.NoRowsAffected)
		if ok {
			http.Redirect(w, r, "/team", http.StatusNotModified)
		} else {
			http.Error(w, err.Error(), http.StatusConflict)
			s.log.Error(err)
		}
		return
	}
	if s.p != nil && config.GetStaticConfig().Queue.Use != "none" {
		wr := worker.Info{Topic: serviceGrp.Name}
		err := s.p.RemoveWorkers(wr)
		if err != nil {
			http.Error(w, err.Error()+"\nNote: Element was removed from database", http.StatusInternalServerError)
			s.log.Error(err)
		}
	}
}

func (s *serviceGroupController) GetByID(w http.ResponseWriter, r *http.Request) {
	genericGetByID(s.svc, s.log, "GetByID", "id", w, r)
}

func (s *serviceGroupController) GetAll(w http.ResponseWriter, r *http.Request) {
	genericGet(s.svc, s.log, "GetAll", w, r)
}

func (s *serviceGroupController) Redeploy(w http.ResponseWriter, r *http.Request) {
	if !(s.p != nil && config.GetStaticConfig().Queue.Use != "none") {
		http.Error(w, "Queue was not established, or platform is none, please manually redeploy the workers", http.StatusBadRequest)
		return
	}

	id, err := idResolver(s.svc, "id", r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.log.Error(err)
		return
	}
	idUint, ok := id.(uint64)
	if !ok {
		http.Error(w, "failed to retrieve the id", http.StatusInternalServerError)
		return
	}
	serGrp, err := s.svc.GetByID(idUint)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		s.log.Error(err)
		return
	}
	if *serGrp.Enabled == true {
		http.Error(w, "service group must first be disabled", http.StatusPreconditionFailed)
		return
	}
	wr := worker.Info{Topic: serGrp.Name, Label: serGrp.Name}
	err = s.p.RemoveWorkers(wr)
	if err != nil {
		http.Error(w, "scoretrak encountered an error while removing the workers. Please delete the workers manually. Details:\n"+err.Error(), http.StatusPreconditionFailed)
		s.log.Error(err)
		return
	}
	err = s.p.DeployWorkers(wr)
	if err != nil {
		http.Error(w, "scoretrak encountered an error while deploying the workers. Please create the workers manually. Details:\n"+err.Error(), http.StatusPreconditionFailed)
		s.log.Error(err)
		return
	}
}

func (s *serviceGroupController) Update(w http.ResponseWriter, r *http.Request) {
	tm := service_group.ServiceGroup{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&tm)
	if err != nil {
		s.log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := idResolver(s.svc, "id", r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.log.Error(err)
		return
	}
	idUint, ok := id.(uint64)
	if !ok {
		http.Error(w, "failed to retrieve the id", http.StatusInternalServerError)
		s.log.Error(err)
		return
	}
	serviceGrp, err := s.svc.GetByID(idUint)
	if err != nil {
		err = errors.New("failed to retrieve the object")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.log.Error(err)
		return
	}
	if s.p != nil && !tm.SkipPlatform && config.GetStaticConfig().Queue.Use != "none" {
		if (tm.Enabled != nil && *tm.Enabled == true) || (tm.Enabled == nil && *serviceGrp.Enabled == true) {
			if !s.q.Ping(tm) {
				err = errors.New("failed to ping the worker queue, ensure that workers are up and running")
				http.Error(w, err.Error(), http.StatusInternalServerError)
				s.log.Error(err)
				return
			}
		}
	}
	genericUpdate(s.svc, tm, s.log, "Update", "id", w, r)
}
