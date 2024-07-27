package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/navneetshukl/receipe-sharing/internal/core/receipe"
)

type Handler struct {
	receipeUsecaseImpl receipe.ReceipeUseCaseImpl
}

func NewHandler(ru receipe.ReceipeUseCaseImpl) *Handler {
	return &Handler{
		receipeUsecaseImpl: ru,
	}
}

func (h *Handler) CreateReceipe() func(c *gin.Context) {
	return func(c *gin.Context) {
		var resp receipe.Receipe

		err := c.ShouldBindJSON(&resp)
		if err != nil {
			log.Println("error in reading the body ", err)
			c.JSON(http.StatusInternalServerError,gin.H{
				"error": "invalid json",
			})
			return
		}

		err = h.receipeUsecaseImpl.AddReceipe(resp)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "data inserted successfully",
		})
	}
}
