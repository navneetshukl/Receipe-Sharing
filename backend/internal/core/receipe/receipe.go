package receipe

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Receipe struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	UserID      primitive.ObjectID `json:"user_id" bson:"user_id"`
	Name        string             `json:"name" bson:"name"`
	Ingredients []string           `json:"ingredients" bson:"ingredients"`
	Description string             `json:"description" bson:"description"`
	Created_At  time.Time          `json:"created_at" bson:"creadted_at"`
}

type ReceipeUseCaseImpl interface {
	AddReceipe(data Receipe) error
}
