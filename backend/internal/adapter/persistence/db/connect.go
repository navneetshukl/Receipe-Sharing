package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectToDB() (*mongo.Database, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("error in loading the env ", err)
		return nil, err
	}

	mongo_uri := os.Getenv("CONNECTION_STRING")
	clientOptions := options.Client().ApplyURI(mongo_uri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	log.Println("connected to mongoDB")
	return client.Database("mydb"), nil
}

func Connect() *mongo.Database {
	db, err := connectToDB()

	if err != nil {
		log.Println("error in connecting to mongodb ", err)
		return nil
	}
	return db
}
