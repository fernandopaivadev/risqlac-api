package main

import (
	"log"
	"risqlac-api/infra"
)

func main() {
	infra.Environment.Load()

	err := infra.Database.Connect()

	if err != nil {
		log.Fatalln(err.Error())
	}

	Server.Setup()
	Server.LoadUserRoutes()
	Server.LoadProductRoutes()

	err = Server.Start()

	if err != nil {
		log.Fatalln(err.Error())
	}
}
