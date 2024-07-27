package main

import (
	"github.com/navneetshukl/receipe-sharing/internal/adapter/persistence/db"
)

func main() {
	appDB := db.Connect()
	db.NewReceipeDatabase(appDB)

}
