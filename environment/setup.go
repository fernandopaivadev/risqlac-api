package environment

import "github.com/joho/godotenv"

func Setup() {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading environment variables")
	}
}
