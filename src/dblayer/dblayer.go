package dblayer

import "github.com/shopping-mall-api/src/models"

type DBLayer interface {
	Products() ([]models.Product, error)
	Promotions() ([]models.Product, error)
	CustomerByName(string, string) (models.Customer, error)
	CustomerByID(int) (models.Customer, error)
	Product(int) (models.Product, error)
	AddUser(models.Customer) (models.Customer, error)
	SignIn(email, pass string) (models.Customer, error)
	SignOut(int) error
	CustomerOrdersByID(int) ([]models.Order, error)
}
