package db

import (
	"context"
	"log"

	"github.com/navneetshukl/receipe-sharing/internal/core/receipe"
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

func (rd *ReceipeDatabase) InsertReceipe(data *receipe.Receipe) error {

	_, err := rd.db.Collection("receipe").InsertOne(context.Background(), data)
	if err != nil {
		log.Println("error in inserting to mongodb ", err)
		return err
	}
	log.Println("Inserted successfully")
	return nil

}
