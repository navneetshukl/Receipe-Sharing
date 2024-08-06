package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	Email     string             `json:"email" bson:"email"`
	Mobile    string             `json:"mobile" bson:"mobile"`
	Password  string             `json:"password" bson:"password"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

type UserUseCaseImpl interface {
	AddUser(user *User) error
	LoginUser(loginData *LoginUser) (string, error)
}
