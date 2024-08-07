package receipe

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/navneetshukl/receipe-sharing/internal/core/receipe"
)

type ReceipeHandler struct {
	receipeUsecaseImpl receipe.ReceipeUseCaseImpl
}

func NewReceipeHandler(ru receipe.ReceipeUseCaseImpl) *ReceipeHandler {
	return &ReceipeHandler{
		receipeUsecaseImpl: ru,
	}
}

func (h *ReceipeHandler) CreateReceipeHandler(c *fiber.Ctx) error {

	var resp receipe.Receipe

	userID := c.Locals("userID")

	err := c.BodyParser(&resp)
	if err != nil {
		return handlerError(c, err)
	}

	err = h.receipeUsecaseImpl.AddReceipe(userID.(string), &resp)
	if err != nil {
		return handlerError(c, err)
	}

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "data inserted successfully",
		"data":    []string{},
	})

}
