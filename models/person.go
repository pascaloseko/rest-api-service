package models

import (
	"log"
	"strings"

	"github.com/jinzhu/gorm"
)

// Person details of a person
type Person struct {
	gorm.Model
	UUID string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Message function
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

//Valid to check validity of a person's details
func (p Person) Valid() (map[string]interface{}, bool) {
	msg := ""
	if len(strings.TrimSpace(p.UUID)) == 0 {
		return Message(false, "id cannot be empty"), false
	}

	if len(strings.TrimSpace(p.Name)) == 0 {
		return Message(false, "name cannot be empty"), false
	}

	if p.Age < 0 {
		return Message(false, "age cannot be less than zero"), false
	}
	strings.TrimSuffix(msg, ", ")

	if len(msg) > 0 {
		return Message(false, "must contain a message"), false
	}
	return nil, false
}

//CreatePerson Creates a person
func (p *Person) CreatePerson() map[string]interface{} {
	if resp, ok := p.Valid(); !ok {
		return resp
	}

	GetDB().Create(p)
	resp := Message(true, "success")
	resp["p"] = p
	return resp
}

//GetPersons get a list of persons
func GetPersons() []Person {
	var persons []Person

	err := GetDB().Table("persons").Error
	if err != nil {
		log.Printf("cannot get persons: %+v\n", err)
		return nil
	}
	return persons
}
