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
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": errors.New("server database error")})
	}
	products, err := h.db.Products()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"products": products})
}

func (h *Handler) Promotions(c *fiber.Ctx) error {
	if h.db == nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": errors.New("server database error")})
	}
	promotions, err := h.db.Promotions()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": errors.New("server database error")})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"promotions": promotions})
}

func (h *Handler) SignIn(c *fiber.Ctx) error {
	if h.db == nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": errors.New("server database error")})
	}
	var customer models.Customer
	err := c.JSONP(&customer)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	customer, err = h.db.SignIn(customer.Email, customer.Pass)
	if err != nil {
		if err != dblayer.ErrINVALIDPASSWORD {
			return c.Status(http.StatusForbidden).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"customer": customer})
}

func (h *Handler) AddUser(c *fiber.Ctx) error {
	if h.db == nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": errors.New("server database error")})
	}
	var customer models.Customer
	err := c.JSONP(&customer)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	customer, err = h.db.AddUser(customer)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"customer": customer})
}

func (h *Handler) SignOut(c *fiber.Ctx) error {
	if h.db == nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": errors.New("server database error")})
	}
	p := c.Params("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	err = h.db.SignOut(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(http.StatusOK)
}

func (h *Handler) Orders(c *fiber.Ctx) error {
	if h.db == nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": errors.New("server database error")})
	}
	p := c.Params("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	orders, err := h.db.CustomerOrdersByID(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"orders": orders})
}
