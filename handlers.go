package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pborman/uuid"

	"github.com/pascaloseko/rest-api-service/models"
)

var persons = make(map[string]models.Person)

// handleGetPersons handles HTTP requests of the form:
//     GET /persons?pageNumber=1&pageSize=300
// 1. extractsPagination in request
// 2. .sorts persons
// 3. responds with subset of persons in page.
func handleGetPersons(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("pageNumber"))
	start, _ := strconv.Atoi(r.FormValue("pageSize"))

	var person models.Person

	if count > 10 || count < 1 {
		count = 10
	}

	if start < 0 {
		start = 0
	}

	persons, err := person.GetPersons(start, count)
	if err == nil {
		fmt.Printf("something happened: %+v\n", err)
		respondWithError(w, "Unable to encode response", http.StatusBadRequest)
		return
	}
	respondWithJSON(w, http.StatusOK, persons)
}

// handleGetPerson handles HTTP requests of the form:
//     GET /persons/{personid}
func handleGetPerson(w http.ResponseWriter, r *http.Request) {
	personID := mux.Vars(r)["uuid"]

	person, err := models.GetPerson(personID)
	if err != nil {
		log.Printf("Person with id does not exist: %+v", err)
		respondWithError(w, "Person with id does not exist", http.StatusNotFound)
		return
	}
	if err != nil {
		return
	}
	respondWithJSON(w, http.StatusOK, person)
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
	var err error
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	var newPa models.Person

	newPa.UUID = uuid.New()

	json.Unmarshal(body, &newPa)
	err = newPa.NewPerson()
	if err != nil {
		log.Println(err)
		respondWithError(w, "cannot add person", http.StatusBadRequest)
		return
	}
	respondWithJSON(w, http.StatusCreated, newPa)
}

func handleUpdatePerson(w http.ResponseWriter, r *http.Request) {
	personID := mux.Vars(r)["uuid"]

	person, err := models.GetPerson(personID)
	if err != nil {
		log.Printf("Person with id does not exist: %+v", err)
		respondWithError(w, "Person with id does not exist", http.StatusNotFound)
		return
	}
	if err != nil {
		return
	}
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	if err := json.Unmarshal(body, &person); err != nil {
		log.Printf("Invalid json in request: %v", err)
		respondWithError(w, "Invalid json in request: %v", http.StatusBadRequest)
	}

	if err, ok := person.Valid(); !ok {
		log.Println(err)
		respondWithError(w, "Wrong values", http.StatusBadRequest)
		return
	}

	err = models.UpdatePerson()

	if err != nil {
		respondWithError(w, "Cannot update person", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusAccepted, p)
}

func requestParamAsInt(r *http.Request, key string) (int, error) {
	valStr := r.FormValue(key)
	val, err := strconv.Atoi(valStr)
	if err != nil {
		return 0, err
	}
	return val, nil
}

func personMapToSlice(mp map[string]models.Person) []models.Person {
	var ps []models.Person

	for _, p := range mp {
		ps = append(ps, p)
	}
	return ps
}

func respondWithError(w http.ResponseWriter, message string, code int) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
