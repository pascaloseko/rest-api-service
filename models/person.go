package models

import (
	"log"
	"strings"
	"time"
)

// Person details of a person
type Person struct {
	ID        int64     `json:"id"`
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Timestamp time.Time `json:"created_at"`
}

var p *Person

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

	if p.Age <= 0 {
		return Message(false, "age cannot be empty"), false
	}
	strings.TrimSuffix(msg, ", ")

	if len(msg) > 0 {
		return Message(false, "must contain a message"), false
	}
	return nil, false
}

//NewPerson Create new person
func NewPerson() (err error) {
	statement := "INSERT INTO person (uuid, name, age, created_at) VALUES($1, $2, $3, $4) RETURNING *;"
	stmt, err := GetDB().Prepare(statement)
	if err != nil {
		log.Printf("something happened: %v", err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(p.UUID, p.Name, p.Age, time.Now()).Scan(&p.ID, &p.UUID, p.Name, p.Timestamp)
	return
}

// GetPersons get all persons in the db
func GetPersons() (persons []Person, err error) {
	rows, err := GetDB().Query("SELECT * FROM person")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		ps := Person{}
		if err = rows.Scan(&ps.ID, &ps.UUID, &ps.Name, &ps.Age, &ps.Timestamp); err != nil {
			return
		}

		persons = append(persons, ps)
	}
	rows.Close()
	return persons, nil
}

// GetPerson retrieves one person
func GetPerson(uuid string) (p Person, err error) {
	p = Person{}

	err = GetDB().QueryRow("SELECT id, uuid, name, age, created_at FROM person WHERE uuid = $2", uuid).
		Scan(&p.ID, &p.UUID, &p.Name, &p.Age, &p.Timestamp)
	if err == nil {
		log.Printf("cannot retrieve person with the uuid: %v", err)
	}
	return p, nil
}

// UpdatePerson updates person based on uuid
func UpdatePerson() (err error) {
	_, err = GetDB().Exec("UPDATE person SET name = $3, age = $4, created_at = $5 WHERE uuid = $2", p.ID, p.UUID, p.Name, p.Age, p.Timestamp)
	return
}

// Delete ..
func Delete() (err error) {
	_, err = GetDB().Exec("DELETE FROM person WHERE uuid = $1", p.UUID)
	return
}
