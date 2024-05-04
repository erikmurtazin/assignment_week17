package tests

import (
	"assignment_week17/api"
	"assignment_week17/db"
	"assignment_week17/models"
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInMemoryHandler(t *testing.T) {
	server := api.NewServer(nil)
	body := models.InMemory{
		Key:   "foo",
		Value: "bar",
	}
	bodyMarshal, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/in-memory", bytes.NewBuffer(bodyMarshal))
	if err != nil {
		t.Fatal(err)
	}
	server.HandleInMemoryRequest(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Error("Wrong status code")
	}
}

func TestMongoRequest(t *testing.T) {
	t.Setenv("LISTEN_ADDR", ":3000")
	t.Setenv("MONGODB_URI", "mongodb+srv://challengeUser:WUMglwNBaydH8Yvu@challenge-xzwqd.mongodb.net/getir-case-study?retryWrites=true")
	store, err := db.NewStorage()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := store.Client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()
	server := api.NewServer(store)
	body := db.DbRequest{
		StartDate: "2016-01-01",
		EndDate:   "2024-01-01",
		MinCount:  3000,
		MaxCount:  100000,
	}
	bodyMarshal, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/mongo", bytes.NewBuffer(bodyMarshal))
	if err != nil {
		t.Fatal(err)
	}
	server.HandleMongoRequest(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Error("Wrong status code")
	}
}
