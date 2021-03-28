package rest

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber"
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
	Charge(c *fiber.Ctx)
}

type Handler struct {
	db dblayer.DBLayer
}

func NewHandler() (*Handler, error) {
	return new(Handler), nil
}

func (h *Handler) Products(c *fiber.Ctx) {
	if h.db == nil {
		return
	}
	products, err := h.db.Products()
	if err != nil {
		c.JSON(fiber.Map{"status": http.StatusInternalServerError, "error": err.Error()})
		return
	}
	c.JSON(fiber.Map{"status": http.StatusOK, "products": products})
}

func (h *Handler) Promotions(c *fiber.Ctx) {
	if h.db == nil {
		return
	}
	promotions, err := h.db.Promotions()
	if err != nil {
		c.JSON(fiber.Map{"status": http.StatusInternalServerError, "error": err.Error()})
		return
	}
	c.JSON(fiber.Map{"status": http.StatusOK, "promotions": promotions})
}

func (h *Handler) SignIn(c *fiber.Ctx) {
	if h.db == nil {
		return
	}
	var customer models.Customer
	err := c.JSONP(&customer)
	if err != nil {
		c.JSON(fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
	}
	customer, err = h.db.SignIn(customer)
	if err != nil {
		c.JSON(fiber.Map{"status": http.StatusInternalServerError, "error": err.Error()})
		return
	}
	c.JSON(fiber.Map{"status": http.StatusOK, "customer": customer})
}

func (h *Handler) AddUser(c *fiber.Ctx) {
	if h.db == nil {
		return
	}
	var customer models.Customer
	err := c.JSONP(&customer)
	if err != nil {
		c.JSON(fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
	}
	customer, err = h.db.AddUser(customer)
	if err != nil {
		c.JSON(fiber.Map{"status": http.StatusInternalServerError, "error": err.Error()})
		return
	}
	c.JSON(fiber.Map{"status": http.StatusOK, "customer": customer})
}

func (h *Handler) SignOut(c *fiber.Ctx) {
	if h.db == nil {
		return
	}
	p := c.Params("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		return
	}
	err = h.db.SignOut(id)
	if err != nil {
		c.JSON(fiber.Map{"status": http.StatusInternalServerError, "error": err.Error()})
		return
	}
}

func (h *Handler) Orders(c *fiber.Ctx) {
	if h.db == nil {
		return
	}
	p := c.Params("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		return
	}
	orders, err := h.db.CustomerOrdersById(id)
	if err != nil {
		c.JSON(fiber.Map{"status": http.StatusInternalServerError, "error": err.Error()})
		return
	}
	c.JSON(fiber.Map{"status": http.StatusOK, "orders": orders})
}

func (h *Handler) Charge(c *fiber.Ctx) {
	if h.db == nil {
		return
	}
}
