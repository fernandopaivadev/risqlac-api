package main

import (
	"risqlac-api/database"
	"risqlac-api/environment"
	"risqlac-api/routes"
)

func main() {
	environment.Load()
	database.Connect()

	routes.Setup()
	routes.User()
	routes.Product()

	err := routes.App.Listen(":3000")

	if err != nil {
		panic("Error starting server: " + err.Error())
	}
}
