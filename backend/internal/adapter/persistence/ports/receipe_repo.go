package ports

import "go.mongodb.org/mongo-driver/bson/primitive"

type ReceipeRepo interface {
	InsertReceipe(userID primitive.ObjectID,name,description string,ingredients []string)error
}