package main

import (
	"ecommerce-app/cmd/server"
	"ecommerce-app/pkg/database"
	"log"
)

func main() {
	var DBConnection = database.NewDatabase()
	err := server.Run(DBConnection)
	if err != nil {
		log.Fatal(err)
		return
	}
	server.Injection()
}
