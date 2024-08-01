package receipe

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/navneetshukl/receipe-sharing/internal/core/receipe"
)

func handlerError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, receipe.ErrAddingReceipe):
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error adding receipe",
		})

	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong",
		})

	}
}
