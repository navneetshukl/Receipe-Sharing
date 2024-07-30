package user

import (
	"log"

	"github.com/navneetshukl/receipe-sharing/internal/adapter/persistence/db"
	"github.com/navneetshukl/receipe-sharing/internal/adapter/persistence/ports"
	"github.com/navneetshukl/receipe-sharing/internal/core/user"
	"github.com/navneetshukl/receipe-sharing/pkg/helper"
	"github.com/navneetshukl/receipe-sharing/pkg/middleware"
)

type UserUseCase struct {
	User ports.UserRepo
}

func NewUserUseCase(user ports.UserRepo) *UserUseCase {
	return &UserUseCase{
		User: user,
	}
}

// AddUser function will add the user
func (uc *UserUseCase) AddUser(data *user.User) error {
	if len(data.Name) == 0 || len(data.Email) == 0 || len(data.Password) == 0 || len(data.Phone) == 0 {
		log.Println("some fields are missing")
		return user.ErrMissingField
	}

	if len(data.Phone) != 10 {
		log.Println("phone number is not valid")
		return user.ErrInvalidPhoneNumber
	}

	// check if user with particular email or phone is present

	userDet, err := uc.User.FindUserByEmailOrPhone(data.Phone, true)
	if err != nil {
		if err != db.ErrDocumentNotFound {
			log.Println("error in finding user by phone", err)
			return user.ErrSomethingWentWrong
		}

	}

	if userDet != nil {
		log.Println("user with the same  phone already exists")
		return user.ErrUserAlreadyExist
	}

	userDet, err = uc.User.FindUserByEmailOrPhone(data.Email, false)
	if err != nil {
		if err != db.ErrDocumentNotFound {
			log.Println("error in finding user by email", err)
			return user.ErrSomethingWentWrong
		}

	}

	if userDet != nil {
		log.Println("user with the same email  already exists")
		return user.ErrUserAlreadyExist
	}

	hashPassword, err := helper.HashPassword(data.Password)
	if err != nil {
		log.Println("error in hashing the password")
		return user.ErrHashingPassword
	}

	data.Password = hashPassword
	err = uc.User.InsertUser(data)
	if err != nil {
		log.Println("error in adding the user ", err)
		return user.ErrAddingUser
	}

	return nil

}

func (uc *UserUseCase) LoginUser(loginData *user.LoginUser) (string, string, error) {
	if loginData.Email == "" || len(loginData.Email) == 0 {
		log.Println("email is missing")
		return "", "", user.ErrMissingField
	}

	loginUser, err := uc.User.FindUserByEmailOrPhone(loginData.Email, false)
	if err != nil {
		if err == db.ErrDocumentNotFound {
			return "", "", user.ErrUserNotFound
		}
		return "", "", user.ErrSomethingWentWrong
	}

	// Check if loginUser is nil
	if loginUser == nil {
		log.Println("loginUser is nil")
		return "", "", user.ErrUserNotFound
	}

	// Check if loginUser.Password is empty
	if loginUser.Password == "" {
		log.Println("password is missing for user")
		return "", "", user.ErrUserNotFound
	}
	err = helper.ComaprePassword(loginData.Password, loginUser.Password)
	if err != nil {
		log.Println("password doesnot match ", err)
		return "", "", user.ErrInvalidPassword
	}

	token, err := middleware.GenerateJWT(loginUser.ID.Hex())
	if err != nil {
		log.Println("error in generating jwt ", err)
		return "", "", user.ErrSomethingWentWrong
	}
	return token, loginUser.ID.Hex(), nil

}
