package handler

import (
	"encoding/json"
	"github.com/L1ghtman2k/ScoreTrak/pkg/logger"
	"github.com/L1ghtman2k/ScoreTrak/pkg/master/run"
	"github.com/L1ghtman2k/ScoreTrak/pkg/queue"
	"github.com/L1ghtman2k/ScoreTrak/pkg/queue/queueing"
	"github.com/L1ghtman2k/ScoreTrak/pkg/service"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

type serviceController struct {
	log logger.LogInfoFormat
	svc service.Serv
	q   queue.Queue
	r   run.RepoStore
}

func NewServiceController(log logger.LogInfoFormat, svc service.Serv, q queue.Queue, r run.RepoStore) *serviceController {
	return &serviceController{log, svc, q, r}
}

func (s *serviceController) Store(w http.ResponseWriter, r *http.Request) {
	tm := &service.Service{}
	genericStore(s.svc, tm, s.log, "Store", w, r)
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
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 64)
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
