package user

import (
	"errors"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/navneetshukl/receipe-sharing/internal/core/user"
	"github.com/navneetshukl/receipe-sharing/pkg/middleware"
)

type UserHandler struct {
	userUsecaseImpl user.UserUseCaseImpl
}

func NewUserHandler(uc user.UserUseCaseImpl) *UserHandler {
	return &UserHandler{
		userUsecaseImpl: uc,
	}
}

func (uh *UserHandler) CreateUserHandler(c *fiber.Ctx) error {

	var userDet user.User
	err := c.BodyParser(&userDet)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "Invalid request body",
		})

	}
	err = uh.userUsecaseImpl.AddUser(&userDet)
	if err != nil {
		return handleError(c, err)
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "user registered successfully",
	})

}

func (uh *UserHandler) LoginUserHandler(c *fiber.Ctx) error {

	var loginUser user.LoginUser
	err := c.BodyParser(&loginUser)
	if err != nil {
		return handleError(c, err)

	}

	log.Println("login user is ", loginUser)

	userID, err := uh.userUsecaseImpl.LoginUser(&loginUser)
	if err != nil {
		return handleError(c, err)

	}

	//save jwt token to cookie

	// c.Cookie(&fiber.Cookie{
	// 	Name:     "auth",
	// 	Value:    jwtToken,
	// 	MaxAge:   3600,
	// 	Domain:   "",
	// 	Path:     "/",
	// 	Secure:   false,
	// 	HTTPOnly: true,
	// 	SameSite: "None",
	// })

	isToken := middleware.CreateJwtCookie(userID, c)
	if !isToken {
		return handleError(c, errors.New("token not created"))
	}

	log.Println("token created is ", isToken)

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "user login successfully",
		"userID":  userID,
	})

}

func (uh *UserHandler) AuthHandler(c *fiber.Ctx) error {

	value := c.Locals("userID")
	log.Println("Value is ", value)
	headers := c.GetReqHeaders()
	log.Println("Headers is ", headers)

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "user authenticated",
		"email":   value,
	})
}
