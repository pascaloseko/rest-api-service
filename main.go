package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.PathPrefix("/persons/{id}").Methods(http.MethodPut).HandlerFunc(handleUpdatePerson)
	r.PathPrefix("/persons/{id}").Methods(http.MethodGet).HandlerFunc(handleGetPerson)
	r.PathPrefix("/persons").Methods(http.MethodPost).HandlerFunc(handleNewPerson)
	r.PathPrefix("/persons").Methods(http.MethodGet).HandlerFunc(handleGetPersons)

	addr := ":9000"
	log.Printf("Listening on: %s", addr)
	err := http.ListenAndServe(addr, r)
	log.Fatalf("Quiting with error: %v", err)
}
