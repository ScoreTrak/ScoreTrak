package handler

import (
	"github.com/L1ghtman2k/ScoreTrak/pkg/logger"
	"github.com/L1ghtman2k/ScoreTrak/pkg/property"
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
	tm := &property.Property{}
	genericStore(t.svc, tm, t.log, "Store", w, r)
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
