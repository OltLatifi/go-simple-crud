package controllers

import (
	"github.com/gin-gonic/gin"
	"example/web-service-gin/models" 
	"example/web-service-gin/config" 
)

// func GetSongsFromAlbum (c *gin.Context) {
// 	var song []models.Album
// 	config.DB.Find(&albums)

// 	c.JSON(200, albums)
// }

func MakeSong(c *gin.Context) {
	var body struct {
		Title string `json:"title" binding:"required"`
	}

	err := c.Bind(&body)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	song := models.Song{Title: body.Title}

	result := config.DB.Create(&song)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(201, song)
}

func FindSong(c *gin.Context){
	var song models.Song

	id := c.Param("id")
	result := config.DB.First(&song, &id)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.JSON(200, song)
}


// func UpdateAlbum(c *gin.Context){
// 	var album models.Album
// 	var body struct {
// 		Title string `json:"title" binding:"required"`
// 	}

// 	err := c.Bind(&body)
// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	id := c.Param("id")
// 	result := config.DB.First(&album, &id)

// 	if result.Error != nil {
// 		c.JSON(500, gin.H{
// 			"error": result.Error.Error(),
// 		})
// 		return
// 	}

// 	album.Title = body.Title
// 	config.DB.Save(&album)
	
// 	c.JSON(200, album)
// }