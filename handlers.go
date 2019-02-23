package main

import (
	"net/http"
)

// handleGetPersons handles HTTP requests of the form:
//     GET /persons?pageNumber=1&pageSize=300
// 1. extractsPagination in request
// 2. .sorts persons
// 3. responds with subset of persons in page.
func handleGetPersons(w http.ResponseWriter, r *http.Request){}

// handleGetPerson handles HTTP requests of the form:
//     GET /persons/{personid}
func handleGetPerson(w http.ResponseWriter, r *http.Request){}

// handleNewPerson handles HTTP requests of the form:
//     POST /persons
//         {"name": "John Doe", "Age": 22}
// 1. reads request body into person
// 2. assigns an ID
// 3. validates
// 4. adds person to list of persons
// 5. responds with inserted person.
func handleNewPerson(w http.ResponseWriter, r *http.Request){}