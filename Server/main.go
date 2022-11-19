package main

import (
	"fmt"
	"log"
	"net/http"
	"Server/router"
)

func main() {
	r := router.Router()
	fmt.Println("Starting the server on Port 9000...")

	log.Fatal(http.ListenAndServe(":9000", r))
}