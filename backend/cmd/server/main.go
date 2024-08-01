package main

import (
	"log"

	"github.com/navneetshukl/receipe-sharing/internal/adapter/persistence/db"
	routes "github.com/navneetshukl/receipe-sharing/internal/interface"
	receipeHand "github.com/navneetshukl/receipe-sharing/internal/interface/handler/receipe"
	userHand "github.com/navneetshukl/receipe-sharing/internal/interface/handler/user"

	"github.com/navneetshukl/receipe-sharing/internal/usecase/receipe"
	"github.com/navneetshukl/receipe-sharing/internal/usecase/user"
	"github.com/navneetshukl/receipe-sharing/pkg/logging"
)

func main() {
	appDB := db.Connect()
	logs := logging.NewLogging()
	receipeRepo := db.NewReceipeDatabase(appDB)

	receipeUseCase := receipe.NewReceipeUseCase(receipeRepo, logs)
	receipeHandler := receipeHand.NewReceipeHandler(receipeUseCase)

	userRepo := db.NewUserDatabase(appDB)
	userUseCase := user.NewUserUseCase(userRepo, logs)
	userHandler := userHand.NewUserHandler(userUseCase)

	router := routes.SetUpRoutes(*receipeHandler, *userHandler)

	err := router.Run(":8080")
	if err != nil {
		log.Println("error in starting the server ", err)
		return
	}
	log.Println("server started succesfully at port 8080")

}
