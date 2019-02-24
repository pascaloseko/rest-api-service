package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// prints out server config details
	p("Persons", version(), "started at ", config.Address)

	r.PathPrefix("/persons/{id}").Methods(http.MethodPut).HandlerFunc(handleUpdatePerson)
	r.PathPrefix("/persons/{id}").Methods(http.MethodGet).HandlerFunc(handleGetPerson)
	r.PathPrefix("/persons").Methods(http.MethodPost).HandlerFunc(handleNewPerson)
	r.PathPrefix("/persons").Methods(http.MethodGet).HandlerFunc(handleGetPersons)

	server := &http.Server{
		Addr:           config.Address,
		Handler:        r,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}