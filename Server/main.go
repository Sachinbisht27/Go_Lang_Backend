package main

import (
	routes "Server/src/routes"
	_ "github.com/lib/pq"
	db "Server/src/database"
)


func main() {
	db.DataMigration()
	routes.HandlerRouting()
}