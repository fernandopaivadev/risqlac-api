package app

import "risqlac-api/environment"

func Start() {
	serverPort := ":" + environment.Get().SERVER_PORT

	err := Instance.Listen(serverPort)

	if err != nil {
		panic("Error starting server: " + err.Error())
	}
}
