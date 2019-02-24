package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// database
var db *gorm.DB

func init() {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	dbURI := fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s", username, dbName, password)
	fmt.Println(dbURI)

	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		log.Printf("cannot connect to db: %v", err)
	}

	db = conn
	// Database Migration
	db.Debug().AutoMigrate(&Person{})
}

//GetDB Returns a handle to the db object
func GetDB() (database *gorm.DB) {
	return db
}
