package user

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/navneetshukl/receipe-sharing/internal/core/user"
)

func handleError(c *fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, user.ErrAddingUser):
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "error in inserting the user",
		})

	case errors.Is(err, user.ErrHashingPassword):
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "error in hashing password",
		})

	case errors.Is(err, user.ErrInvalidPassword):
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "invalid password",
		})

	case errors.Is(err, user.ErrInvalidPhoneNumber):
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "phone number is invalid",
		})

	case errors.Is(err, user.ErrMissingField):
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "some of the fields are missing",
		})

	case errors.Is(err, user.ErrSomethingWentWrong):
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "something went wrong",
		})

	case errors.Is(err, user.ErrUserAlreadyExist):
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "user already exists for given detail",
		})

	case errors.Is(err, user.ErrUserNotFound):
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "user not found",
		})

	default:
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "something went wrong",
		})
	}
}
