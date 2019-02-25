package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"sort"
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
	page := extractPagination(r)

	pas := personMapToSlice(persons)

	sort.SliceStable(pas, func(i, j int) bool {
		return pas[i].Name < pas[j].Name
	})

	startIndex := page.StartIndex()
	if len(pas) <= startIndex {
		log.Printf("No person record found in selected page: %+v", page)
		http.Error(w, fmt.Sprintf("No person record found in selected pageNumber (%d)", page.Number), http.StatusNotFound)
		return
	}

	endPos := int(math.Min(float64(len(pas)), float64(page.EndPosition())))

	respPas := pas[startIndex:endPos]

	rp := models.GetDB().Find(&respPas)
	if err := json.NewEncoder(w).Encode(&rp); err != nil {
		respondWithError(w, "Unable to encode response: %+v", http.StatusBadRequest)
		return
	}
	respPas = pas
	respondWithJSON(w, http.StatusOK, respPas)
}

// handleGetPerson handles HTTP requests of the form:
//     GET /persons/{personid}
func handleGetPerson(w http.ResponseWriter, r *http.Request) {
	personID := mux.Vars(r)["id"]

	person, exist := persons[personID]
	if !exist {
		log.Printf("Person with id %s does not exist", personID)
		respondWithError(w, "Person with id %s does not exist", http.StatusNotFound)
		return
	}

	err := json.NewEncoder(w).Encode(person)
	if err != nil {
		log.Printf("Error encoding results: %v", err)
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
	newPa := models.Person{}
	slice := make(map[string]string)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Unable to read request: %v", err)
		respondWithError(w, "Somethin really bad happened here: %+v", http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	if err := json.Unmarshal(body, &slice); err != nil {
		log.Printf("Invalid json in request: %v", err)
		respondWithError(w, "Invalid json in request: %v", http.StatusBadRequest)
		return
	}

	newPa.UUID = uuid.New()
	newPa.Name = slice["name"]
	newPa.Age = slice["age"]

	fmt.Println(newPa)

	// if err, ok := newPa.Valid(); !ok {
	// 	log.Println(err)
	// 	respondWithError(w, "Wrong values", http.StatusBadRequest)
	// 	return
	// }

	if err := models.GetDB().Create(&newPa); err != nil {
		fmt.Println(err)
		respondWithError(w, "Cannot create person", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusCreated, newPa)
}

func handleUpdatePerson(w http.ResponseWriter, r *http.Request) {
	personID := mux.Vars(r)["id"]

	oldPerson, exist := persons[personID]
	if !exist {
		log.Printf("Person with id %s does not exist", personID)
		respondWithError(w, "Person with id %s does not exist", http.StatusNotFound)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Unable to read request: %v", err)
		respondWithError(w, "Something went wrong", http.StatusInternalServerError)
	}

	newPa := models.Person{}
	if err := json.Unmarshal(body, &newPa); err != nil {
		log.Printf("Invalid json in request: %v", err)
		respondWithError(w, "Invalid json in request: %v", http.StatusBadRequest)
	}

	newPa.ID = oldPerson.ID
	if err, ok := newPa.Valid(); !ok {
		log.Println(err)
		respondWithError(w, "Wrong values", http.StatusBadRequest)
		return
	}

	p := models.GetDB().Save(&newPa)

	if err := p; err != nil {
		respondWithError(w, "Cannot update person", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusAccepted, p)
}

func extractPagination(r *http.Request) models.Page {
	page := models.Page{}
	var err error

	page.Number, err = requestParamAsInt(r, "pageNumber")
	if err != nil {
		page.Number = 1
	}

	page.Size, err = requestParamAsInt(r, "pageSize")
	if err != nil {
		page.Size = 100
	}

	return page
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
