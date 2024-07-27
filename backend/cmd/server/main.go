package main

import (
	"github.com/navneetshukl/receipe-sharing/internal/adapter/persistence/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	appDB := db.Connect()
	db.NewReceipeDatabase(appDB)
	
}
