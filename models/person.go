package models

import (
	"strings"

	"github.com/jinzhu/gorm"
)

// Person details of a person
type Person struct {
	gorm.Model
	UUID   string `json:"id"`
	Name string `json:"name"`
	Age  string `json:"age"`
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

	if len(strings.TrimSpace(p.Age)) == 0 {
		return Message(false, "age cannot be empty"), false
	}
	strings.TrimSuffix(msg, ", ")

	if len(msg) > 0 {
		return Message(false, "must contain a message"), false
	}
	return nil, false
}