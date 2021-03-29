package rest

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/shopping-mall-api/src/dblayer"
	"github.com/shopping-mall-api/src/models"
)

type IHandler interface {
	Products(c *fiber.Ctx)
	Promotions(c *fiber.Ctx)
	AddUser(c *fiber.Ctx)
	SignIn(c *fiber.Ctx)
	SignOut(c *fiber.Ctx)
	Orders(c *fiber.Ctx)
}

type Handler struct {
	db dblayer.DBLayer
}

func NewHandler(dbname string) (*Handler, error) {
	db, err := dblayer.NewORM(dbname)
	if err != nil {
		return nil, err
	}
	return &Handler{db: db}, nil
}

func (h *Handler) Products(c *fiber.Ctx) error {
	if h.db == nil {
		c.JSON(fiber.Map{"status": http.StatusInternalServerError, "error": errors.New("server database error")})
		return nil
	}
	products, err := h.db.Products()
	if err != nil {
		c.JSON(fiber.Map{"status": http.StatusInternalServerError, "error": err.Error()})
		return err
	}
	c.JSON(fiber.Map{"status": http.StatusOK, "products": products})
	return nil
}

func (h *Handler) Promotions(c *fiber.Ctx) error {
	if h.db == nil {
		c.JSON(fiber.Map{"status": http.StatusInternalServerError, "error": errors.New("server database error")})
		return nil
	}
	promotions, err := h.db.Promotions()
	if err != nil {
		c.JSON(fiber.Map{"status": http.StatusInternalServerError, "error": err.Error()})
		return err
	}
	c.JSON(fiber.Map{"status": http.StatusOK, "promotions": promotions})
	return nil
}

func (h *Handler) SignIn(c *fiber.Ctx) error {
	if h.db == nil {
		c.JSON(fiber.Map{"status": http.StatusInternalServerError, "error": errors.New("server database error")})
		return nil
	}
	var customer models.Customer
	err := c.JSONP(&customer)
	if err != nil {
		c.JSON(fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		return err
	}
	customer, err = h.db.SignIn(customer.Email, customer.Pass)
	if err != nil {
		if err != dblayer.ErrINVALIDPASSWORD {
			c.JSON(fiber.Map{"status": http.StatusForbidden, "error": err.Error()})
			return err
		}
		c.JSON(fiber.Map{"status": http.StatusInternalServerError, "error": err.Error()})
		return err
	}
	c.JSON(fiber.Map{"status": http.StatusOK, "customer": customer})
	return nil
}

func (h *Handler) AddUser(c *fiber.Ctx) error {
	if h.db == nil {
		c.JSON(fiber.Map{"status": http.StatusInternalServerError, "error": errors.New("server database error")})
		return nil
	}
	var customer models.Customer
	err := c.JSONP(&customer)
	if err != nil {
		c.JSON(fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
	}
	customer, err = h.db.AddUser(customer)
	if err != nil {
		c.JSON(fiber.Map{"status": http.StatusInternalServerError, "error": err.Error()})
		return err
	}
	c.JSON(fiber.Map{"status": http.StatusOK, "customer": customer})
	return nil
}

func (h *Handler) SignOut(c *fiber.Ctx) error {
	if h.db == nil {
		c.JSON(fiber.Map{"status": http.StatusInternalServerError, "error": errors.New("server database error")})
		return nil
	}
	p := c.Params("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		return err
	}
	err = h.db.SignOut(id)
	if err != nil {
		c.JSON(fiber.Map{"status": http.StatusInternalServerError, "error": err.Error()})
		return err
	}
	return nil
}

func (h *Handler) Orders(c *fiber.Ctx) error {
	if h.db == nil {
		c.JSON(fiber.Map{"status": http.StatusInternalServerError, "error": errors.New("server database error")})
		return nil
	}
	p := c.Params("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		return err
	}
	orders, err := h.db.CustomerOrdersByID(id)
	if err != nil {
		c.JSON(fiber.Map{"status": http.StatusInternalServerError, "error": err.Error()})
		return err
	}
	c.JSON(fiber.Map{"status": http.StatusOK, "orders": orders})
	return nil
}
