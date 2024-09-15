package api

import (
	"github.com/fitness-tracker/internal/app"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service app.ServiceInterface
}

func NewHandler(service *app.Service) *Handler {
    return &Handler{service: service}
}

func (h *Handler) HelloWorld(c *fiber.Ctx) error {
    greeting := h.service.GetGreeting()
    return c.SendString(greeting)
}