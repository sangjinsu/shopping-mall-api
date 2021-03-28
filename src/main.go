package main

import (
	"log"

	"github.com/shopping-mall-api/src/rest"
)

func main() {
	const address string = ":8080"
	log.Println("Server is running on " + address)
	err := rest.RunAPI(address)
	log.Fatal(err)
}
