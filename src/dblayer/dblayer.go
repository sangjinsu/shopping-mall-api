package dblayer

import "github.com/shopping-mall-api/src/models"

type DBLayer interface {
	Products() ([]models.Product, error)
	Promotions() ([]models.Product, error)
	CustomerByName(string, string) (models.Customer, error)
	CustomerByID(int) (models.Customer, error)
	Product(uint) (models.Product, error)
	AddUser(models.Customer) (models.Customer, error)
	SignIn(models.Customer) (models.Customer, error)
	SignOut(int) error
	CustomerOrdersById(int) ([]models.Order, error)
}
