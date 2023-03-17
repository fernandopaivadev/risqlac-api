package main

import (
	"risqlac-api/database"
	"risqlac-api/environment"
)

func main() {
	environment.Load()
	database.Connect()

	server.Setup()
	server.LoadUserRoutes()
	server.LoadProductRoutes()
	server.Start()
}
