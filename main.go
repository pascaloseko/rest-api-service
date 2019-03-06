package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	// prints out server config details
	p("Persons", version(), "started at ", ":"+os.Getenv("PORT"))

	r.PathPrefix("/persons/{uuid}").Methods(http.MethodGet).HandlerFunc(handleGetPerson)
	r.PathPrefix("/persons").Methods(http.MethodPost).HandlerFunc(handleNewPerson)
	r.PathPrefix("/persons").Methods(http.MethodGet).HandlerFunc(indexHandler)

	r.PathPrefix("/persons/{uuid}").Methods(http.MethodPut).HandlerFunc(handleUpdatePerson)
	r.PathPrefix("/persons/{uuid}").Methods(http.MethodDelete).HandlerFunc(handleDelete)

	// server := &http.Server{
	// 	Addr:           config.Address,
	// 	Handler:        r,
	// 	ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
	// 	WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
	// 	MaxHeaderBytes: 1 << 20,
	// }

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}
