package main

import (
    "github.com/gin-gonic/gin"

	"example/web-service-gin/controllers"
	"example/web-service-gin/config"
	"example/web-service-gin/models"
)



// routes
func main() {
    //Database
	db := config.ConnectDB()

	db.Table("albums").AutoMigrate(&models.Album{})

    router := gin.Default()
    router.GET("/albums", controllers.GetAlbums)
    router.GET("/albums/:id", controllers.FindAlbum)
	router.PUT("/albums/:id", controllers.UpdateAlbum)
	router.POST("/albums", controllers.MakeAlbum)

    router.Run("localhost:8000")
}