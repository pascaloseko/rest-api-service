package main

import (
	"encoding/json"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// handleGetPersons handles HTTP requests of the form:
//     GET /persons?pageNumber=1&pageSize=300
// 1. extractsPagination in request
// 2. .sorts persons
// 3. responds with subset of persons in page.
func handleGetPersons(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "get persons not implemented yet!", http.StatusNotImplemented)
}

// handleGetPerson handles HTTP requests of the form:
//     GET /persons/{personid}
func handleGetPerson(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "get person not implemented yet!", http.StatusNotImplemented)
}

// handleNewPerson handles HTTP requests of the form:
//     POST /persons
//         {"name": "John Doe", "Age": 22}
// 1. reads request body into person
// 2. assigns an ID
// 3. validates
// 4. adds person to list of persons
// 5. responds with inserted person.
func handleNewPerson(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "create new person not implemented yet!", http.StatusNotImplemented)
}

func handleUpdatePerson(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "update person not implemented yet!", http.StatusNotImplemented)
}
