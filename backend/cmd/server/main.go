package main

import (
	"log"

	"github.com/navneetshukl/receipe-sharing/internal/adapter/persistence/db"
	routes "github.com/navneetshukl/receipe-sharing/internal/interface"
	"github.com/navneetshukl/receipe-sharing/internal/interface/handler"
	"github.com/navneetshukl/receipe-sharing/internal/usecase/receipe"
	"github.com/navneetshukl/receipe-sharing/internal/usecase/user"
)

func main() {
	appDB := db.Connect()
	receipeRepo := db.NewReceipeDatabase(appDB)

	receipeUsecase := receipe.NewReceipeUseCase(receipeRepo)
	receipeHandler := handler.NewReceipeHandler(receipeUsecase)

	userRepo := db.NewUserDatabase(appDB)

	userUseCase := user.NewUserUseCase(userRepo)
	userHandler := handler.NewUserHandler(userUseCase)

	router := routes.SetUpRoutes(*receipeHandler, *userHandler)

	err := router.Run(":8080")
	if err != nil {
		log.Println("error in starting the server ", err)
		return
	}
	log.Println("server started succesfully at port 8080")

}
