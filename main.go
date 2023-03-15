package main

import (
	"risqlac-api/database"
	"risqlac-api/environment"
	"risqlac-api/routes"
	"risqlac-api/server"
)

func main() {
	environment.Load()
	database.Connect()
	server.Setup()
	routes.LoadUserRoutes()
	routes.LoadProductRoutes()
	server.Start()
}
