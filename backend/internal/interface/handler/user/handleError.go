package user

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/navneetshukl/receipe-sharing/internal/core/user"
)

func handleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, user.ErrAddingUser):
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error in inserting the user",
		})

	case errors.Is(err, user.ErrHashingPassword):
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error in hashing password",
		})

	case errors.Is(err, user.ErrInvalidPassword):
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "invalid password",
		})

	case errors.Is(err, user.ErrInvalidPhoneNumber):
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "phone number is invalid",
		})

	case errors.Is(err, user.ErrMissingField):
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "some of the fields are missing",
		})

	case errors.Is(err, user.ErrSomethingWentWrong):
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})

	case errors.Is(err, user.ErrUserAlreadyExist):
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "user already exists for given detail",
		})

	case errors.Is(err, user.ErrUserNotFound):
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "user not found",
		})

	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})

	}
}
