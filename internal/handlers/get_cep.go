package handlers

import "github.com/gofiber/fiber/v2"

type Handlers struct{}

func (h *Handlers) GetCep(c *fiber.Ctx) error {
	return c.JSON("{message: Hello, World!}")
}
