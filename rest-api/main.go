package main

import (
	"github.com/affrianr/auto-reload-waterwind/rest-api/config"
	"github.com/affrianr/auto-reload-waterwind/rest-api/handler"
	"github.com/affrianr/auto-reload-waterwind/rest-api/repository"
	"github.com/affrianr/auto-reload-waterwind/rest-api/routes"
	"github.com/affrianr/auto-reload-waterwind/rest-api/usecase"
)

func main() {
	config.ConnectDB()

	waterWindRepo := repository.NewWaterWindRepository(config.DB)
	waterWindUsecase := usecase.NewWaterWindUsecase(waterWindRepo)
	waterWindHandler := handler.NewWaterWindHandler(waterWindUsecase)

	router := routes.SetupRouter(waterWindHandler)

	router.Run(":8080")

}
