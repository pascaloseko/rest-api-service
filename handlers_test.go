package main

import (
	"net/http"
	"net/http/httptest"
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
