package main

import (
	"risqlac-api/app"
	"risqlac-api/database"
	"risqlac-api/environment"
	"risqlac-api/routes"
)

func main() {
	environment.Load()
	database.Connect()
	app.Setup()
	routes.SetupUserRoutes()
	routes.SetupProductRoutes()
	app.Start()
}
