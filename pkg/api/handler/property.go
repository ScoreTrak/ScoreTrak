package handler

import (
	"encoding/json"
	"github.com/ScoreTrak/ScoreTrak/pkg/logger"
	"github.com/ScoreTrak/ScoreTrak/pkg/property"
	"github.com/qor/validations"
	"net/http"
)

type propertyController struct {
	log logger.LogInfoFormat
	svc property.Serv
}

func NewPropertyController(log logger.LogInfoFormat, svc property.Serv) *propertyController {
	return &propertyController{log, svc}
}

func (t *propertyController) Store(w http.ResponseWriter, r *http.Request) {
	var tm []*property.Property
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&tm)
	if err != nil {
		t.log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = t.svc.Store(tm)
	if err != nil {
		_, ok := err.(*validations.Error)
		if ok {
			http.Error(w, err.Error(), http.StatusPreconditionFailed)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		t.log.Error(err)
		return
	}
}

func (t *propertyController) Delete(w http.ResponseWriter, r *http.Request) {
	genericDelete(t.svc, t.log, "Delete", "id", w, r)
}

func (t *propertyController) GetByID(w http.ResponseWriter, r *http.Request) {
	genericGetByID(t.svc, t.log, "GetByID", "id", w, r)
}

func (t *propertyController) GetAll(w http.ResponseWriter, r *http.Request) {
	genericGet(t.svc, t.log, "GetAll", w, r)
}

func (t *propertyController) Update(w http.ResponseWriter, r *http.Request) {
	tm := &property.Property{}
	genericUpdate(t.svc, tm, t.log, "Update", "id", w, r)
}
