package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/navneetshukl/receipe-sharing/internal/interface/handler"
)

func SetUpRoutes(receipeHandler handler.Handler) *gin.Engine {
	router := gin.Default()

	// CORS configuration
	corsConfig := cors.Config{
		AllowOrigins: []string{"http://localhost:5173"}, // Replace with your frontend's URL
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
	}
	router.Use(cors.New(corsConfig))

	router.POST("/api/receipe/add", receipeHandler.CreateReceipe())
	return router

}
