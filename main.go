package main

import (
	"github.com/PrimaWinangun/ulventech/controller"
	_ "github.com/PrimaWinangun/ulventech/docs"
	"github.com/PrimaWinangun/ulventech/repository"
	"github.com/PrimaWinangun/ulventech/usecase"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os"
)

//go:generate swag init --parseDependency --parseInternal

func init() {
	if _, err := os.Stat("upload"); os.IsNotExist(err) {
		err := os.Mkdir("upload", 0777)
		if err != nil {
			log.Fatal("failed to create directory")
		}
	}
}

// @title Ulventech Technical Test
// @version 1.0
// @description This is a simple application for ulventech technical test.
// @termsOfService http://swagger.io/terms/

// @contact.name Putu Prima Winangun
// @contact.email primawinangun@gmail.com

// @license.name MIT

// @host localhost:9000
// @BasePath /
// @schemes http
func main() {
	// initiate repository layer
	wcRepo := repository.NewWcRepository()
	vtRepo := repository.NewVtRepository()

	// initiate use case layer
	wcUseCase := usecase.NewWcUseCase(wcRepo)
	vtUseCase := usecase.NewVTUseCase(vtRepo)

	//initiate gin
	gin.SetMode("debug")
	rg := gin.Default()
	v1 := rg.Group("api/v1/")

	// initiate delivery layer
	controller.NewWordCount(v1, wcUseCase)
	controller.NewValidTime(v1, vtUseCase)

	// swagger setting
	url := ginSwagger.URL("http://localhost:9000/swagger/doc.json")
	rg.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// start server
	err := rg.Run(":9000")
	if err != nil {
		log.Fatal("Failed to run server")
	}
}
