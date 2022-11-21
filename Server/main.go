package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"

	"fmt"
	"log"
	_ "github.com/lib/pq"
	db "gorm.io/gorm"
	models "Server/src/models"
)

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
    db, err := db.Open(postgres.Open(connString), &db.Config{})
    if err != nil {
        log.Fatal(err)
    }

	db.AutoMigrate(models.SignUp{})

  	fmt.Println("Successfully connected!")

	initializerRouter()
}