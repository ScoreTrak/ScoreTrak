package handler

import (
	"encoding/json"
	"errors"
	"github.com/ScoreTrak/ScoreTrak/pkg/logger"
	"github.com/ScoreTrak/ScoreTrak/pkg/property"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/orm"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
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
	params := mux.Vars(r)
	key := params["Key"]
	if key == "" {
		http.Error(w, "key should not be empty", http.StatusBadRequest)
		return
	}
	sID, err := uuidResolver("ServiceID", r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = t.svc.Delete(sID, key)
	if err != nil {
		_, ok := err.(*orm.NoRowsAffected)
		if ok {
			http.Redirect(w, r, "/", http.StatusNotModified)
			return
		} else {
			http.Error(w, err.Error(), http.StatusConflict)
			t.log.Error(err)
			return
		}
	}
}

func (t *propertyController) GetByServiceIDKey(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	key := params["Key"]
	if key == "" {
		http.Error(w, "key should not be empty", http.StatusBadRequest)
		return
	}
	sID, err := uuidResolver("ServiceID", r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sg, err := t.svc.GetByServiceIDKey(sID, key)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			t.log.Error(err)
		}
		return
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(sg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		t.log.Error(err)
	}
}

func (t *propertyController) GetAllByServiceID(w http.ResponseWriter, r *http.Request) {
	sID, err := uuidResolver("ServiceID", r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sg, err := t.svc.GetAllByServiceID(sID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			t.log.Error(err)
		}
		return
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(sg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		t.log.Error(err)
	}
}

func (t *propertyController) GetAll(w http.ResponseWriter, r *http.Request) {
	genericGet(t.svc, t.log, "GetAll", w, r)
}

func (t *propertyController) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	key := params["Key"]
	if key == "" {
		http.Error(w, "key should not be empty", http.StatusBadRequest)
		return
	}
	sID, err := uuidResolver("ServiceID", r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tm := &property.Property{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(tm)
	if err != nil {
		t.log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tm.ServiceID = sID
	tm.Key = key
	err = t.svc.Update(tm)
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
