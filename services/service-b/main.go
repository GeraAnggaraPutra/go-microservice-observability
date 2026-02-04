package main

import (
	"service-b/app"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("APP_ENV") == "" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("ERROR load env file : %s", err.Error())
		}
	}

	app.InitLogger()
	app.StartServer()
}
