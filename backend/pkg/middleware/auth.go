package middleware

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

var secret = []byte("U23#$65@*&53@dert$&3!@#$31dD")

func CreateJwtCookie(userID string, c *fiber.Ctx) bool {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": userID,
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		log.Println("failed to create token ", err)
		return false
	} else {
		log.Println("Token created")
		// c.SetCookie("Authorise", tokenString, 3600, "/", "", false, true)
		// c.SetSameSite(http.SameSiteNoneMode)

		c.Cookie(&fiber.Cookie{
			Name:     "authorise",
			Value:    tokenString,
			MaxAge:   3600,
			Domain:   "",
			Path:     "/",
			Secure:   false,
			HTTPOnly: true,
			SameSite: "None",
		})

		return true
	}
}

func RetrieveJwtToken(c *fiber.Ctx) (string, error) {
	cookie := c.Cookies("authorise")

	headers := c.GetReqHeaders()

	log.Println("headers in retrieve jwt is ", headers)

	if len(cookie) == 0 {
		return "", errors.New("cookie not found")
	} else {
		token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return secret, nil
		})

		if err != nil {
			return "", err
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := claims["user"].(string)
			return userID, nil
		} else {
			return "", fmt.Errorf("invalid token")
		}
	}
}

func ValidateCookie(c *fiber.Ctx) bool {
	cookie := c.Cookies("authorise")

	headers := c.GetReqHeaders()

	log.Println("headers in validate cookie is  ", headers)

	if len(cookie) == 0 || cookie == "" {
		fmt.Println("cookie not found in validate cookie")
		return false
	} else {
		return true
	}
}

func ValidateJwt(c *fiber.Ctx) error {
	valid := ValidateCookie(c)
	if !valid {
		return c.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"error": "user not login",
		})
	}

	userID, err := RetrieveJwtToken(c)
	if err != nil {
		log.Println("error in retrieving token", err)
		return c.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"error": "cookie retrieving failed",
		})
	}

	log.Println("UserID is ", userID)

	c.Locals("userID", userID)
	return c.Next()
}
