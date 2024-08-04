package user

import (
	"log"
	"net/http"
	"time"

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
			handleError(c, err)
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
			handleError(c, err)
			return
		}

		log.Println("login user is ", loginUser)

		jwtToken, userID, err := uh.userUsecaseImpl.LoginUser(&loginUser)
		if err != nil {
			handleError(c, err)
			return
		}

		//save jwt token to cookie
		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("Authorization", jwtToken, int(time.Hour*24*30), "/", "", false, true)
		value, err := c.Cookie("Authorization")
		if err != nil {
			log.Println("Error in getting cookie value in login ", err)
		} else {
			log.Println("Cookie value in login is ", value)
		}
		c.JSON(200, gin.H{
			"token":  jwtToken,
			"userId": userID,
		})
	}
}

func (uh *UserHandler) AuthHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		value, _ := c.Get("user")
		log.Println("Value is ", value)
		c.JSON(http.StatusOK, gin.H{
			"message": "user authenticated",
		})
	}
}
