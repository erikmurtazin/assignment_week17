package api

import (
	"assignment_week17/db"
	"assignment_week17/models"
	"encoding/json"
	"log"
	"net/http"
	"reflect"
)

func (s *Server) HandleMongoRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		writeJSON(w, http.StatusMethodNotAllowed, models.ApiError{Error: "method not allowed"})
		return
	}
	var req db.DbReqest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		writeJSON(w, http.StatusMethodNotAllowed, models.ApiError{Error: "not enough arguments"})
		return
	}
	defer r.Body.Close()
	responseFromDb, err := s.Db.FetchDataFromMongo(req)
	if err != nil {
		log.Println(err)
		res := models.NewMongoResponse(1, "error", *responseFromDb)
		writeJSON(w, http.StatusBadRequest, res)
		return
	}
	if reflect.ValueOf(*responseFromDb).IsZero() {
		res := models.NewMongoResponse(2, "no records", *responseFromDb)
		writeJSON(w, http.StatusOK, res)
	} else {
		res := models.NewMongoResponse(0, "success", *responseFromDb)
		writeJSON(w, http.StatusOK, res)

	}
}

func (s *Server) handleInMemoryRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var req models.InMemory
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Println(err)
			writeJSON(w, http.StatusBadRequest, models.ApiError{Error: "invalid request"})
			return
		}
		defer r.Body.Close()
		value := reflect.ValueOf(req)
		for i := range value.NumField() {
			if fieldValue := value.Field(i); fieldValue.String() == "" {
				writeJSON(w, http.StatusBadRequest, models.ApiError{Error: "fields must not be empty"})
				return
			}
		}
		if s.Memory[req.Key] == "" {
			s.Memory[req.Key] = req.Value
			writeJSON(w, http.StatusOK, req)
		} else {
			writeJSON(w, http.StatusOK, models.ApiError{Error: "key is already in use"})
		}

	case "GET":
		r.ParseForm()
		key := r.Form.Get("key")
		if key == "" {
			writeJSON(w, http.StatusBadRequest, models.ApiError{Error: "invalid request"})
			return
		}
		if s.Memory[key] != "" {
			res := models.NewInmemoryResponse(key, s.Memory[key])
			writeJSON(w, http.StatusOK, res)
		} else {
			writeJSON(w, http.StatusOK, models.ApiError{Error: "no records"})
		}
	default:
		writeJSON(w, http.StatusMethodNotAllowed, models.ApiError{Error: "method not allowed"})
		return
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
