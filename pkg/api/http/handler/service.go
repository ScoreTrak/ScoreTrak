package handler

import (
	"encoding/json"
	"github.com/ScoreTrak/ScoreTrak/cmd/master/run"
	"github.com/ScoreTrak/ScoreTrak/pkg/di/repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/logger"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue/queueing"
	"github.com/ScoreTrak/ScoreTrak/pkg/service"
	"github.com/qor/validations"
	"net/http"
	"time"
)

type serviceController struct {
	log logger.LogInfoFormat
	svc service.Serv
	q   queue.Queue
	r   repo.Store
}

func NewServiceController(log logger.LogInfoFormat, svc service.Serv, q queue.Queue, r repo.Store) *serviceController {
	return &serviceController{log, svc, q, r}
}

func (s *serviceController) Store(w http.ResponseWriter, r *http.Request) {
	var tm []*service.Service
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&tm)
	if err != nil {
		s.log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
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

func (s *serviceController) Delete(w http.ResponseWriter, r *http.Request) {
	genericDelete(s.svc, s.log, "Delete", "id", w, r)
}

func (s *serviceController) GetByID(w http.ResponseWriter, r *http.Request) {
	genericGetByID(s.svc, s.log, "GetByID", "id", w, r)
}

func (s *serviceController) GetAll(w http.ResponseWriter, r *http.Request) {
	genericGet(s.svc, s.log, "GetAll", w, r)
}

func (s *serviceController) Update(w http.ResponseWriter, r *http.Request) {
	tm := &service.Service{}
	genericUpdate(s.svc, tm, s.log, "Update", "id", w, r)
}

func (s *serviceController) TestService(w http.ResponseWriter, r *http.Request) {
	id, err := uuidResolver("id", r)
	if err != nil {
		s.log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)

	}
	ser, err := s.r.Service.GetByID(id)
	if err != nil {
		s.log.Error(err)
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	p, _ := s.r.Property.GetAllByServiceID(id)
	h, _ := s.r.Host.GetByID(ser.HostID)
	serGrp, _ := s.r.ServiceGroup.GetByID(ser.ServiceGroupID)

	response, berr, err := s.q.Send([]*queueing.ScoringData{
		{Service: queueing.QService{ID: id, Name: ser.Name, Group: serGrp.Name}, Host: *h.Address, Deadline: time.Now().Add(time.Second * 5), RoundID: 0, Properties: run.PropertyToMap(p)},
	})
	if berr != nil {
		response[0].Err += berr.Error()
	}
	if response == nil || err != nil {
		http.Error(w, "something went wrong, either check took too long to execute, or the workers did not receive the check", http.StatusInternalServerError)
		s.log.Error(err)
		return
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(response[0])
	if err != nil {
		s.log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
