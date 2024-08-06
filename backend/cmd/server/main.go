package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/navneetshukl/receipe-sharing/internal/adapter/persistence/db"
	routes "github.com/navneetshukl/receipe-sharing/internal/interface"
	receipeHand "github.com/navneetshukl/receipe-sharing/internal/interface/handler/receipe"
	userHand "github.com/navneetshukl/receipe-sharing/internal/interface/handler/user"

	"github.com/navneetshukl/receipe-sharing/internal/usecase/receipe"
	"github.com/navneetshukl/receipe-sharing/internal/usecase/user"
	"github.com/navneetshukl/receipe-sharing/pkg/logging"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}
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

	err := router.Listen(":8080")
	if err != nil {
		log.Println("error in starting the server ", err)
		return
	}
	log.Println("server started succesfully at port 8080")

}
