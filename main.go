package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John", Price: 36.99},
	{ID: "2", Title: "Green Car", Artist: "Doe", Price: 26.22},
	{ID: "3", Title: "Red Rabbit", Artist: "Alice", Price: 33.11},
	{ID: "4", Title: "Brown guitar", Artist: "Gore", Price: 44.77},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbum)

	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context) {
	var newAlbum album

	err := c.BindJSON(&newAlbum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "All fields are required",
		})
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusOK, newAlbum)
}

func getAlbumById(c *gin.Context) {
	albumId := c.Param("id")
	for _, album := range albums {
		if album.ID == albumId {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Not Found item"})
}
