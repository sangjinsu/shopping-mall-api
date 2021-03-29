package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/shopping-mall-api/src/rest"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	const address string = ":8080"
	log.Println("Server is running on " + address)
	err = rest.RunAPI(address)
	log.Fatal(err.Error())
}
