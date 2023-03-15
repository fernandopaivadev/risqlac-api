package server

import "risqlac-api/environment"

func Start() {
	serverPort := ":" + environment.Variables.ServerPort

	err := Instance.Listen(serverPort)

	if err != nil {
		panic("Error starting server: " + err.Error())
	}
}
