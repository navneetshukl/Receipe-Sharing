package db

import (
	"context"
	"log"
	"time"

	"github.com/navneetshukl/receipe-sharing/internal/core/receipe"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ReceipeDatabase struct {
	db *mongo.Database
}

func NewReceipeDatabase(db *mongo.Database) *ReceipeDatabase {
	return &ReceipeDatabase{
		db: db,
	}
}

func (rd *ReceipeDatabase) InsertReceipe(userID primitive.ObjectID, name, description string, ingredients []string) error {
	var results receipe.Receipe

	results.UserID = userID
	results.Description = description
	results.Ingredients = ingredients
	results.Name = name
	results.Created_At = time.Now()

	_, err := rd.db.Collection("receipe").InsertOne(context.Background(), results)
	if err != nil {
		log.Println("error in inserting to mongodb ",err)
		return err
	}
	log.Println("Inserted successfully")
	return nil

}
