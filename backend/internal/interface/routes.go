package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/navneetshukl/receipe-sharing/internal/interface/handler/receipe"
	"github.com/navneetshukl/receipe-sharing/internal/interface/handler/user"
	"github.com/navneetshukl/receipe-sharing/pkg/middleware"
)

func SetUpRoutes(receipeHandler receipe.ReceipeHandler, userHandler user.UserHandler) *gin.Engine {
	router := gin.Default()

	// CORS configuration
	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Replace with your frontend's URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}
	router.Use(cors.New(corsConfig))

	router.POST("/api/receipe/add", receipeHandler.CreateReceipeHandler())
	router.POST("/api/user/register", userHandler.CreateUserHandler())
	router.POST("/api/user/login", userHandler.LoginUserHandler())

	// authenticated route

	router.GET("/api/auth", middleware.AuthenticateJWT, userHandler.AuthHandler())
	return router

}
