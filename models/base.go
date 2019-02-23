package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

// database
var db *gorm.DB

func init() {
	// Load .env file
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbURI)

	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		log.Printf("cannot connect to db: %v", err)
	}

	db = conn
	// Database Migration
	db.Debug().AutoMigrate(&Person{}, &Page{})
}

//GetDB Returns a handle to the db object
func GetDB() (database *gorm.DB) {
	return db
}
