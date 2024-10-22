package main

import (
	"net/http"

	"github.com/gin-gonic/gin" // добавить в зависимости go get .
)

// структура альбомов
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// срез с перечнем альбомов
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// postAlbums добавляет альбом из JSON полученный в теле запроса.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON привязывает полученный JSON в newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Добавляет новый альбом в срез
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbums отправляет список альбомов как JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums) // getAlbums()- это () передача результата функции
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

// Write a handler to return a specific item.....
