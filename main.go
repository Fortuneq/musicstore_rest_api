package main

import (
	"github.com/alonelegion/musicstore_rest_api/controllers"
	"github.com/alonelegion/musicstore_rest_api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	// Подключение к базе данных
	models.ConnectDB()

	// Маршруты
	route.GET("/tracks", controllers.GetAllTracks)
	route.POST("/tracks", controllers.CreateTrack)
	route.GET("/tracks/:id", controllers.GetTrack)
	route.PATCH("/tracks/:id", controllers.UpdateTrack)
	route.DELETE("/tracks/:id", controllers.DeleteTrack)

	// Запуск сервера
	route.Run()
}
