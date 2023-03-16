package main

import (
	"risqlac-api/database"
	"risqlac-api/environment"
	"risqlac-api/routes"
)

func main() {
	environment.Load()
	database.Connect()

	server.Setup()
	routes.User.Load(server.App)
	routes.Product.Load(server.App)
	server.Start()
}
