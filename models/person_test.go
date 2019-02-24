package models

import (
	"log"
	"testing"

	"github.com/jinzhu/gorm"

	. "github.com/selvatico/go-mocket"
)

var DB *gorm.DB

func SetUpTests() *gorm.DB {
	Catcher.Register()
	Catcher.Logging = true
	db, err := gorm.Open(DriverName, "connection_string")
	if err != nil {
		log.Println("Something went wrong here")
	}
	DB = db
	return db
}

func TestResponses(t *testing.T) {
	SetUpTests()
	t.Run("Catch by arguments", func(t *testing.T) {
		// Important: Use database files here (snake_case) and not struct variables (CamelCase)
		// eg: first_name, last_name, date_of_birth NOT FirstName, LastName or DateOfBirth
		commonReply := []map[string]interface{}{{"name": "FirstLast", "age": "30"}}
		Catcher.Reset().NewMock().WithArgs(int64(27)).WithReply(commonReply)
		result := GetPersons()
		if len(result) != 1 {
			t.Fatalf("Returned sets is not equal to 1. Received %d", len(result))
		}
		// all other checks from reply
	})
}
