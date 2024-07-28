package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/navneetshukl/receipe-sharing/internal/core/user"
)

type UserHandler struct {
	userUsecaseImpl user.UserUseCaseImpl
}

func NewUserHandler(uc user.UserUseCaseImpl) *UserHandler {
	return &UserHandler{
		userUsecaseImpl: uc,
	}
}

func (uh *UserHandler) CreateUserHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		var userDet user.User
		err := c.ShouldBindJSON(&userDet)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}
		err = uh.userUsecaseImpl.AddUser(&userDet)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error adding user"})
			return
		}
		c.JSON(201, userDet)

	}
}

func (uh *UserHandler) LoginUserHandler() func(c *gin.Context) {
	return func(c *gin.Context) {

		var loginUser user.LoginUser
		err := c.ShouldBindJSON(&loginUser)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		jwtToken, userID, err := uh.userUsecaseImpl.LoginUser(&loginUser)
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid username or password"})
			return
		}
		c.JSON(200, gin.H{
			"token":  jwtToken,
			"userId": userID,
		})
	}
}
