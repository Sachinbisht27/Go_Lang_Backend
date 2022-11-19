package main

import (
	"fmt"
	"log"
	"net/http"
	"Server/router"

	// "gorm.io/driver/postgres"
  	// "gorm.io/gorm"
)

func main() {
	r := router.Router()
	fmt.Println("Starting the server on Port 9000...")

	log.Fatal(http.ListenAndServe(":9000", r))


	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	// _, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
}