package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/shopping-mall-api/src/rest"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	const address string = ":8080"
	log.Println("Server is running on " + address)
	err = rest.RunAPI(app, address)
	if err != nil {
		panic(err)
	}

	err = app.ListenTLS(address, "./cert.pem", "./key.pem")
	if err != nil {
		panic(err)
	}
}
