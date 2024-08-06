package receipe

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/navneetshukl/receipe-sharing/internal/core/receipe"
)

func handlerError(c *fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, receipe.ErrAddingReceipe):
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error adding recipe",
		})

	default:
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Something went wrong",
		})
	}
}
