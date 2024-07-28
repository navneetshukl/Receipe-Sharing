package ports

import "github.com/navneetshukl/receipe-sharing/internal/core/user"

type UserRepo interface {
	InsertUser(user *user.User) error
	FindUserByEmailOrPhone(email string, flag bool) (*user.User, error) //flag->true phone, flag->false email
}
