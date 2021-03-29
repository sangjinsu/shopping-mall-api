package rest

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func RunAPI(address string) error {
	return RunAPIWithHandler(address)
}

func RunAPIWithHandler(address string) error {
	router := fiber.New()
	dbUser := os.Getenv("USER")
	dbPassword := os.Getenv("PASSWORD")
	dbAddress := os.Getenv("DBADDRESS")
	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbAddress + ")/shoppingmallapi?charset=utf8mb4&parseTime=True&loc=Local"
	handler, err := NewHandler(dsn)
	if err != nil {
		return err
	}

	router.Get("/products", handler.Products)
	router.Get("/promotions", handler.Promotions)

	user := router.Group("/user")
	user.Post("/:id/signout", handler.SignOut)
	user.Get("/:id/orders", handler.Orders)

	users := router.Group("/users")
	users.Post("/signin", handler.SignIn)
	users.Post("", handler.AddUser)

	return router.ListenTLS(address, "./cert.pem", "./key.pem")
}
