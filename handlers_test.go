package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetPersons(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/persons", handleGetPersons).Methods("GET")
	req, err := http.NewRequest("GET", "/persons", nil)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Errorf("HTTP status expected: 200, got %d", w.Code)
	}
}

func TestUpdatePerson(t *testing.T) {
	m := http.NewServeMux()
	m.HandleFunc("/person/", handleUpdatePerson)

	w := httptest.NewRecorder()
	json := strings.NewReader(`{"uuid":"jshhsdhs3736","name":"Updated person", "age": 27, "created_at":"2019-03-04T13:00:24.574112Z"}`)
	req, _ := http.NewRequest("PUT", "/person/jshhsdhs3736", json)
	m.ServeHTTP(w, req)

	if w.Code != 202 {
		t.Errorf("HTTP status expected: 202, got %v", w.Code)
	}
}
	