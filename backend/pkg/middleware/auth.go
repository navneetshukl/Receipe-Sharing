package middleware

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

func GenerateJWT(userID string) (string, error) {

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("error in loading the .env ", err)
		return "", err
	}
	secret := os.Getenv("JWT_SECRET_KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		log.Println("Error in signing the token ", err)
		return "", err
	}

	//? Save this JWT token to Cookie

	// c.SetSameSite(http.SameSiteLaxMode)
	// c.SetCookie("Authorization", tokenString, int(time.Hour*24*30), "/", "", false, true)

	return tokenString, nil
}

func AuthenticateJWT(c *gin.Context) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
		return
	}
	tokenString, err := c.Cookie("Authorization")
	secret := os.Getenv("SECRET")

	if err != nil {
		log.Println("Error in Getting the Tokenstring from cookie ", err)
		return

	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return

		}
		userID := claims["sub"].(string)
		if userID == "" {
			return

		}
		c.Set("user", userID)
		c.Next()

	} else {
		return

	}

}
