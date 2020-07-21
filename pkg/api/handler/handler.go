package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/logger"
	"github.com/L1ghtman2k/ScoreTrak/pkg/storage/orm"
	"github.com/L1ghtman2k/ScoreTrak/pkg/team"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/qor/validations"
	"net/http"
	"reflect"
	"strconv"
)

//Generic function passing and assignment

func genericGetByID(svc interface{}, log logger.LogInfoFormat, m string, idParam string, w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("panic occurred:", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()
	id, err := idResolver(svc, idParam, r)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sg, err := invokeRetMethod(svc, m, id)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Error(err)
		}
		return
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(sg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error(err)
	}
}

func genericGet(svc interface{}, log logger.LogInfoFormat, m string, w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("panic occurred:", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()
	sg, err := invokeRetMethod(svc, m)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Error(err)
		}
		return
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(sg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error(err)
	}
}

func genericUpdate(svc interface{}, g interface{}, log logger.LogInfoFormat, m string, idParam string, w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("panic occurred:", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(g)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := idResolver(svc, idParam, r)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	v := reflect.ValueOf(g).Elem()
	f := reflect.ValueOf(id)
	if _, ok := svc.(team.Serv); ok {
		v.FieldByName("Name").Set(f)
	} else {
		v.FieldByName("ID").Set(f)
	}
	err = invokeNoRetMethod(svc, m, g)
	if err != nil {
		_, ok := err.(*validations.Error)
		if ok {
			http.Error(w, err.Error(), http.StatusPreconditionFailed)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		log.Error(err)
		return
	}
}

func genericStore(svc interface{}, g interface{}, log logger.LogInfoFormat, m string, w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("panic occurred:", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(g)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = invokeNoRetMethod(svc, m, g)
	if err != nil {
		_, ok := err.(*validations.Error)
		if ok {
			http.Error(w, err.Error(), http.StatusPreconditionFailed)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		log.Error(err)
		return
	}
}

func genericDelete(svc interface{}, log logger.LogInfoFormat, m string, idParam string, w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("panic occurred:", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()
	id, err := idResolver(svc, idParam, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Error(err)
		return
	}
	err = invokeNoRetMethod(svc, m, id)
	if err != nil {
		_, ok := err.(*orm.NoRowsAffected)
		if ok {
			http.Redirect(w, r, "/team", http.StatusNotModified)
			return
		} else {
			http.Error(w, err.Error(), http.StatusConflict)
			log.Error(err)
			return
		}
	}
}

//Credit to:
// https://stackoverflow.com/questions/14116840/dynamically-call-method-on-interface-regardless-of-receiver-type
// https://stackoverflow.com/questions/8103617/call-a-struct-and-its-method-by-name-in-go
func invokeRetMethod(i interface{}, methodName string, args ...interface{}) (interface{}, error) {
	finalMethod := preInvoke(i, methodName)
	if finalMethod.IsValid() {
		inputs := make([]reflect.Value, len(args))
		for i, _ := range args {
			inputs[i] = reflect.ValueOf(args[i])
		}
		r := finalMethod.Call(inputs)

		if err, ok := r[1].Interface().(error); ok {
			return nil, err
		}
		return r[0].Interface(), nil
	}
	return nil, errors.New(fmt.Sprintf("The method name %s does not exist in %s", methodName, reflect.TypeOf(i).Name()))
}

func invokeNoRetMethod(i interface{}, methodName string, args ...interface{}) error {
	finalMethod := preInvoke(i, methodName)
	if finalMethod.IsValid() {
		inputs := make([]reflect.Value, len(args))
		for i, _ := range args {
			inputs[i] = reflect.ValueOf(args[i])
		}
		r := finalMethod.Call(inputs)

		if err, ok := r[0].Interface().(error); ok {
			return err
		}
		return nil
	}
	return errors.New(fmt.Sprintf("The method name %s does not exist in %s", methodName, reflect.TypeOf(i).Name()))
}

func preInvoke(i interface{}, methodName string) reflect.Value {
	var ptr reflect.Value
	var value reflect.Value
	var finalMethod reflect.Value

	value = reflect.ValueOf(i)
	if value.Type().Kind() == reflect.Ptr {
		ptr = value
		value = ptr.Elem()
	} else {
		ptr = reflect.New(reflect.TypeOf(i))
		temp := ptr.Elem()
		temp.Set(value)
	}
	method := value.MethodByName(methodName)
	if method.IsValid() {
		finalMethod = method
	}
	method = ptr.MethodByName(methodName)
	if method.IsValid() {
		finalMethod = method
	}
	return finalMethod
}

func idResolver(svc interface{}, idParam string, r *http.Request) (interface{}, error) {
	if idParam == "" {
		return 1, nil
	}
	var id interface{}
	params := mux.Vars(r)
	if _, ok := svc.(team.Serv); ok && idParam == "name" {
		id = params[idParam]
	} else {
		var err error
		id, err = strconv.ParseUint(params[idParam], 10, 64)
		if err != nil {
			return nil, err
		}
	}
	return id, nil
}
