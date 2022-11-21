package database

import(
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"log"
	models "Server/src/models"
)

const (
    HOST = "localhost"
    PORT = 5432
    USER = "postgres"
    PASSWORD = "Lavish@123"
    DBNAME = "react_app"
)

var db *gorm.DB

func DataMigration(){
	connString := fmt.Sprintf(
        "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        HOST, PORT, USER, PASSWORD, DBNAME,
    )
    db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }

	db.AutoMigrate(models.SignUp{})
    fmt.Println("Successfully connected!")
}