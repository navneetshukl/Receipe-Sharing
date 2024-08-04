package user

import (
	"errors"
	"time"

	"github.com/navneetshukl/receipe-sharing/internal/adapter/persistence/db"
	"github.com/navneetshukl/receipe-sharing/internal/adapter/persistence/ports"
	"github.com/navneetshukl/receipe-sharing/internal/core/user"
	"github.com/navneetshukl/receipe-sharing/pkg/helper"
	"github.com/navneetshukl/receipe-sharing/pkg/logging"
	"github.com/navneetshukl/receipe-sharing/pkg/middleware"
)

type UserUseCase struct {
	User ports.UserRepo
	Logs logging.LogService
}

func NewUserUseCase(user ports.UserRepo, logs logging.LogService) *UserUseCase {
	return &UserUseCase{
		User: user,
		Logs: logs,
	}
}

// AddUser function will add the user
func (uc *UserUseCase) AddUser(data *user.User) error {
	if len(data.Name) == 0 || len(data.Email) == 0 || len(data.Password) == 0 || len(data.Mobile) == 0 {
		uc.Logs.ErrorLog("Some fields are missing ", nil)
		return user.ErrMissingField
	}

	// uncomment this code in production

	// if len(data.Mobile) != 10 {
	// 	uc.Logs.ErrorLog("phone number is not valid", nil)
	// 	return user.ErrInvalidPhoneNumber
	// }

	// check if user with particular email or phone is present

	userDet, err := uc.User.FindUserByEmailOrPhone(data.Mobile, true)
	if err != nil {
		if err != db.ErrDocumentNotFound {
			uc.Logs.ErrorLog("FindUserByEmailOrPhone ", err)
			return user.ErrSomethingWentWrong
		}

	}

	if userDet != nil {
		uc.Logs.ErrorLog("user with the same  phone already exists", nil)
		return user.ErrUserAlreadyExist
	}

	userDet, err = uc.User.FindUserByEmailOrPhone(data.Email, false)
	if err != nil {
		uc.Logs.ErrorLog("FindUserByEmailOrPhone ", err)

		if err != db.ErrDocumentNotFound {
			return user.ErrSomethingWentWrong
		}

	}

	if userDet != nil {
		uc.Logs.ErrorLog("user with the same email  already exists ", errors.New("user already exist"))
		return user.ErrUserAlreadyExist
	}

	hashPassword, err := helper.HashPassword(data.Password)
	if err != nil {
		uc.Logs.ErrorLog("HashPassword ", nil)
		return user.ErrHashingPassword
	}

	data.Password = hashPassword
	data.CreatedAt = time.Now()
	err = uc.User.InsertUser(data)
	if err != nil {
		uc.Logs.ErrorLog("InsertUser ", err)
		return user.ErrAddingUser
	}

	return nil

}

func (uc *UserUseCase) LoginUser(loginData *user.LoginUser) (string, string, error) {
	if loginData.Email == "" || len(loginData.Email) == 0 {
		uc.Logs.ErrorLog("email is missing", nil)
		return "", "", user.ErrMissingField
	}

	loginUser, err := uc.User.FindUserByEmailOrPhone(loginData.Email, false)
	if err != nil {
		uc.Logs.ErrorLog("FindUserByEmailOrPhone ", err)
		if err == db.ErrDocumentNotFound {
			return "", "", user.ErrUserNotFound
		}
		return "", "", user.ErrSomethingWentWrong
	}

	// Check if loginUser is nil
	if loginUser == nil {
		uc.Logs.ErrorLog("loginUser is nil", errors.New("user is not present"))
		return "", "", user.ErrUserNotFound
	}

	// Check if loginUser.Password is empty
	if loginUser.Password == "" {
		uc.Logs.ErrorLog("password is missing for user ", errors.New("password is missing"))
		return "", "", user.ErrUserNotFound
	}
	err = helper.ComaprePassword(loginData.Password, loginUser.Password)
	if err != nil {
		uc.Logs.ErrorLog("ComaprePassword ", err)
		return "", "", user.ErrInvalidPassword
	}

	token, err := middleware.GenerateJWT(loginUser.ID.Hex())
	if err != nil {
		uc.Logs.ErrorLog("GenerateJWT ", err)
		return "", "", user.ErrSomethingWentWrong
	}
	return token, loginUser.ID.Hex(), nil

}
