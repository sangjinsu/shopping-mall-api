package rest

import (
	"github.com/gofiber/fiber"
)

func RunAPI(address string) error {
	return RunAPIWithHandler(address)
}

func RunAPIWithHandler(address string) error {
	router := fiber.New()

	handler, err := NewHandler()
	if err != nil {
		return err
	}

	router.Get("/products", handler.Products)
	router.Get("/promotions", handler.Promotions)

	user := router.Group("/user")
	user.Post("/:id/signout", handler.SignOut)
	user.Get("/:id/orders", handler.Orders)

	users := router.Group("/users")
	users.Post("/charge", handler.Charge)
	users.Post("/signin", handler.SignIn)
	users.Post("", handler.AddUser)

	return router.Listen(address)

}
