package helper

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("error in hashing password ", err)
		return "", nil
	}

	return string(hashPassword), nil
}

func ComaprePassword(password, hashPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		log.Println("error in comparing password ", err)
		return err
	}
	return nil
}
