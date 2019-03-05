package models

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// Person details of a person
type Person struct {
	ID        int64     `json:"id,omitempty"`
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

//GetAllPersons ...
func (p *Person) GetAllPersons() (persons []Person, err error) {
	statement := "SELECT id,uuid,name,age,created_at FROM person ORDER BY age DESC"
	rows, err := GetDB().Query(statement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	persons = []Person{}

	for rows.Next() {
		var ps Person
		if err := rows.Scan(&ps.ID, &ps.UUID, &ps.Name, &ps.Age, &ps.Timestamp); err != nil {
			return nil, err
		}

		persons = append(persons, ps)
	}

	return persons, nil
}

//NewPerson Create new person
func (p *Person) NewPerson() (err error) {
	statement := "INSERT INTO person (uuid, name, age, created_at) VALUES($1, $2, $3, $4) RETURNING *;"
	stmt, err := GetDB().Prepare(statement)
	if err != nil {
		log.Printf("something happened: %v", err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(p.UUID, p.Name, p.Age, time.Now()).Scan(&p.ID, &p.UUID, &p.Name, &p.Age, &p.Timestamp)
	return
}

// GetPersons get all persons in the db
func (p *Person) GetPersons(start, count int) (persons []Person, err error) {
	statement := fmt.Sprintf("SELECT id, uuid, name, age, created_at FROM person LIMIT %d OFFSET %d", count, start)
	rows, err := GetDB().Query(statement)
	if err != nil {
		return nil, err
	}

	rows.Close()

	for rows.Next() {
		ps := Person{}
		if err = rows.Scan(&ps.ID, &ps.UUID, &ps.Name, &ps.Age, &ps.Timestamp); err != nil {
			return
		}

		persons = append(persons, ps)
	}
	return persons, nil
}

// GetPerson retrieves one person
func GetPerson(uuid string) (p Person, err error) {
	p = Person{}

	err = GetDB().QueryRow("SELECT id, uuid, name, age, created_at FROM person WHERE uuid = $1", uuid).
		Scan(&p.ID, &p.UUID, &p.Name, &p.Age, &p.Timestamp)
	if err != nil {
		log.Printf("cannot retrieve person with the id: %v", err)
	}
	return
}

// UpdatePerson updates person based on uuid
func (p *Person) UpdatePerson() (err error) {
	statement := "UPDATE person SET name = $1, age = $2, created_at = $3 WHERE uuid = $4 returning id,uuid,name,age,created_at"
	stmt, err := GetDB().Prepare(statement)
	if err != nil {
		return
	}

	defer stmt.Close()
	err = stmt.QueryRow(p.Name, p.Age, p.Timestamp, p.UUID).Scan(&p.ID, &p.UUID, &p.Name, &p.Age, &p.Timestamp)
	return
}

// Delete ..
func (p *Person) Delete() (err error) {
	_, err = GetDB().Exec("DELETE FROM person WHERE uuid = $1", p.UUID)
	return
}
