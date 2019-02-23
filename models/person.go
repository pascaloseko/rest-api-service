package models

import (
	"errors"
	"strings"
	"github.com/jinzhu/gorm"
)

// Person details of a person
type Person struct {
	gorm.Model
	UUID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

//Valid to check validity of a person's details
func (p Person) Valid() error {
	msg := ""
	if len(strings.TrimSpace(p.UUID)) == 0 {
		msg += "id cannot be empty"
	}

	if len(strings.TrimSpace(p.Name)) == 0 {
		msg += "name cannot be empty, "
	}

	if p.Age < 0 {
		msg += "age cannot be less than zero"
	}
	strings.TrimSuffix(msg, ", ")

	if len(msg) > 0 {
		return errors.New(msg)
	}
	return nil
}
