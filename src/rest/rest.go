package rest

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func RunAPI(app *fiber.App, address string) error {
	return RunAPIWithHandler(app, address)
}

func RunAPIWithHandler(app *fiber.App, address string) error {
	dbUser := os.Getenv("USER")
	dbPassword := os.Getenv("PASSWORD")
	dbAddress := os.Getenv("DBADDRESS")
	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbAddress + ")/shoppingmallapi?charset=utf8mb4&parseTime=True&loc=Local"
	handler, err := NewHandler(dsn)
	if err != nil {
		return err
	}

	app.Get("/products", handler.Products)
	app.Get("/promotions", handler.Promotions)

	user := app.Group("/user")
	user.Post("/:id/signout", handler.SignOut)
	user.Get("/:id/orders", handler.Orders)

	users := app.Group("/users")
	users.Post("/signin", handler.SignIn)
	users.Post("", handler.AddUser)

	return nil
}
