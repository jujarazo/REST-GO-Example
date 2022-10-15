package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "4", Title: "Enema of the state", Artist: "Blink 182", Price: 49.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbum)

	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumById(c *gin.Context) {
	// Extract the album id from the gin context
	id := c.Param("id")

	// Loop through the albums slice to find the album with the same id
	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	// Format album not found message and send it with HTTP Status 404
	notFoundMessage := fmt.Sprintf("The album with the id: %v does not exist", id)

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": notFoundMessage})
}

// postAlbum adds an album to the album slice and returns it as JSON
func postAlbum(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to newAlbum
	// Check for error on binding
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the album to the slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusOK, newAlbum)
}
