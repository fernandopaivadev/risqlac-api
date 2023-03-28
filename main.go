package main

import (
	"log"
	"risqlac-api/application"
	"risqlac-api/environment"
	"risqlac-api/infra"
)

func main() {
	environment.Load()

	err := infra.Database.Connect()

	if err != nil {
		log.Fatalln(err.Error())
	}

	application.Server.Setup()
	application.Server.LoadUserRoutes()
	application.Server.LoadProductRoutes()

	err = application.Server.Start()

	if err != nil {
		log.Fatalln(err.Error())
	}
}
