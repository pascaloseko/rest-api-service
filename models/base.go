package models

import (
	"fmt"
	"log"
	"os"

	"database/sql"
	//
	_ "github.com/lib/pq"
)

// database
var db *sql.DB

func init() {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	dbURI := fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s", username, dbName, password)
	fmt.Println(dbURI)

	conn, err := sql.Open("postgres", dbURI)
	if err != nil {
		log.Printf("cannot connect to db: %v", err)
	}
	db = conn
}

//GetDB Returns a handle to the db object
func GetDB() (database *sql.DB) {
	db.Close()
	return db
}
