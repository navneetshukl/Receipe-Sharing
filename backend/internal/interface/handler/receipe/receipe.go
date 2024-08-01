package receipe

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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

func (h *ReceipeHandler) CreateReceipeHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		var resp receipe.Receipe

		err := c.ShouldBindJSON(&resp)
		if err != nil {
			log.Println("error in reading the body ", err)
			c.JSON(http.StatusInternalServerError, gin.H{
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
