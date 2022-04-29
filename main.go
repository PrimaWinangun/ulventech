package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"ulventech/controller"
	"ulventech/repository"
	"ulventech/usecase"
)

func init() {
	if _, err := os.Stat("upload"); os.IsNotExist(err) {
		err := os.Mkdir("upload", 0777)
		if err != nil {
			log.Fatal("failed to create directory")
		}
	}
}

func main() {
	// initiate repository layer
	wcRepo := repository.NewWcRepository()
	vtRepo := repository.NewVtRepository()

	// initiate use case layer
	wcUseCase := usecase.NewWcUsecase(wcRepo)
	vtUseCase := usecase.NewVTUseCase(vtRepo)

	//initiate delivery layer
	gin.SetMode("debug")
	rg := gin.Default()
	v1 := rg.Group("api/v1/")
	controller.NewWordCount(v1, wcUseCase)
	controller.NewValidTime(v1, vtUseCase)
	err := rg.Run(":9000")
	if err != nil {
		log.Fatal("Failed to run server")
	}
}
