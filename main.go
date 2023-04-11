package main

import (
	"log"
	"risqlac-api/application"
	"risqlac-api/environment"
	"risqlac-api/infrastructure"
)

func main() {
	environment.Load()

	err := infrastructure.Database.Connect()

	if err != nil {
		log.Fatalln(err.Error())
	}

	application.Server.Setup()
	application.Server.LoadDefaultRoutes()
	application.Server.LoadSessionRoutes()
	application.Server.LoadUserRoutes()
	application.Server.LoadProductRoutes()

	err = application.Server.Start()

	if err != nil {
		log.Fatalln(err.Error())
	}
}
