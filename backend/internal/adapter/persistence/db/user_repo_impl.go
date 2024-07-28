package db

import (
	"context"

	"github.com/navneetshukl/receipe-sharing/internal/core/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserDatabase struct {
	db *mongo.Database
}

func NewUserDatabase(db *mongo.Database) *UserDatabase {
	return &UserDatabase{
		db: db,
	}
}

func (ud *UserDatabase) InsertUser(user *user.User) error {

	user.ID = primitive.NewObjectID()
	_, err := ud.db.Collection("user").InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return nil

}

// flag->true phone, flag->false email
func (ud *UserDatabase) FindUserByEmailOrPhone(data string, flag bool) (*user.User, error) {
	var filter primitive.M
	if flag {
		filter = bson.M{
			"phone": data,
		}

	} else {
		filter = bson.M{
			"email": data,
		}
	}
	var user user.User
	err := ud.db.Collection("user").FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil,err
	}
	return &user, nil
}
