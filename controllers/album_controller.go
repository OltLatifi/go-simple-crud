package controllers

import (
	"github.com/gin-gonic/gin"
	"example/web-service-gin/models" 
	"example/web-service-gin/config" 
)

func GetAlbums (c *gin.Context) {
	var albums []models.Album
	config.DB.Find(&albums)

	c.JSON(200, albums)
}

func MakeAlbum(c *gin.Context) {
	var body struct {
		Title string `json:"title" binding:"required"`
		Songs []int `json:"songs"`
	}

	var songs []models.Song

	err := c.Bind(&body)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	config.DB.Where("id IN (?)", body.Songs).Find(&songs)
	album := models.Album{Title: body.Title, Songs: songs}

	result := config.DB.Create(&album)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(201, album)
}

func FindAlbum(c *gin.Context){
	var album models.Album

	id := c.Param("id")
	result := config.DB.First(&album, &id)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.JSON(200, album)
}


func UpdateAlbum(c *gin.Context){
	var album models.Album
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

	id := c.Param("id")
	result := config.DB.First(&album, &id)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	album.Title = body.Title
	config.DB.Save(&album)
	
	c.JSON(200, album)
}