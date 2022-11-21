package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"

	// "os"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

var DB *gorm.DB

type SignUp struct{
	ID        uint      `gorm:"primaryKey" json:"id"`
	FirstName string 	`gorm:"not null;" json:"firstname"`
	LastName  string 	`gorm:"not null;" json:"lastname"`
	Email     string	`gorm:"not null;" json:"email"`
	Password  uint		`gorm:"not null;" json:"password"`
	ConPassword  uint	`gorm:"not null;" json:"conpassword"`
}

const (
    HOST = "localhost"
    PORT = 5432
    USER = "postgres"
    PASSWORD = "Lavish@123"
    DBNAME = "react_app"
)

func initializerRouter() {
	r := mux.NewRouter()

	log.Fatal(http.ListenAndServe(":9000",r))


}


func main() {
	connString := fmt.Sprintf(
        "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        HOST, PORT, USER, PASSWORD, DBNAME,
    )
    DB, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }

	DB.AutoMigrate(&SignUp{})
	
	fmt.Println("Successfully connected!")
	// connect()

	initializerRouter()
}