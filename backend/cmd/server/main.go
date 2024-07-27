package main

import (
	"log"

	"github.com/navneetshukl/receipe-sharing/internal/adapter/persistence/db"
	routes "github.com/navneetshukl/receipe-sharing/internal/interface"
	"github.com/navneetshukl/receipe-sharing/internal/interface/handler"
	receipe "github.com/navneetshukl/receipe-sharing/internal/usecase"
)

func main() {
	appDB := db.Connect()
	receipeRepo := db.NewReceipeDatabase(appDB)

	receipeUsecase := receipe.NewReceipeUseCase(receipeRepo)
	receipeHandler := handler.NewHandler(receipeUsecase)

	router := routes.SetUpRoutes(*receipeHandler)

	err:=router.Run(":8080")
	if err!=nil{
		log.Println("error in starting the server ",err)
		return
	}
	log.Println("server started succesfully at port 8080")

}
