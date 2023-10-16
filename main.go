package main

import (
    "github.com/gin-gonic/gin"

	"example/web-service-gin/controllers"
	"example/web-service-gin/config"
	"example/web-service-gin/models"
)

func main() {
    //Database
	db := config.ConnectDB()
	db.Table("albums").AutoMigrate(&models.Album{})
	db.Table("songs").AutoMigrate(&models.Song{})

    router := gin.Default()

	albums := router.Group("/albums")
	{
		albums.GET("/", controllers.GetAlbums)
		albums.GET("/:id", controllers.FindAlbum)
		albums.PUT("/:id", controllers.UpdateAlbum)
		albums.POST("/", controllers.MakeAlbum)
	}

	songs := router.Group("/songs")
	{
		// songs.GET("/", controllers.Getsongs)
		// songs.PUT("/:id", controllers.UpdateAlbum)
		songs.GET("/:id", controllers.FindSong)
		songs.POST("/", controllers.MakeSong)
	}

    router.Run("localhost:8000")
}