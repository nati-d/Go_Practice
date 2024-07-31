package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Album struct {
	Id     int
	Title  string
	Artist string
	Price  string
}

var albums = []Album{
	{Id: 1, Title: "Blue Train", Artist: "John Coltrane", Price: "$20.00"},
	{Id: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: "$21.00"},
	{Id: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: "$22.00"},
}

func getAlbums(c *gin.Context) {
	c.JSON(200, albums)

}

func addAlbums(c *gin.Context) {
	var newAlbum Album

	err := c.BindJSON(&newAlbum)
	if err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.JSON(200, newAlbum)

}

func getSingleAlbum(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, album := range albums {

		if id == album.Id {
			c.JSON(200, album)
			return
		}

	}
	c.JSON(404, gin.H{"error": "Album not found"})
}

func updateAlbum(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var updatedAlbum Album
	err := c.BindJSON(&updatedAlbum)
	if err != nil {
		return
	}

	for i, album := range albums {
		if album.Id == id {
			if updatedAlbum.Title != "" {
				albums[i].Title = updatedAlbum.Title
			}

			if updatedAlbum.Artist != "" {
				albums[i].Artist = updatedAlbum.Artist
			}

			if updatedAlbum.Price != "" {
				albums[i].Price = updatedAlbum.Price

				
			}
			c.JSON(200, gin.H{"message": "Album updated"})
			return
		}
	}
	c.JSON(404, gin.H{"error": "Album not found"})

}


func deleteItem (c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))

	for i, album := range albums {
		if album.Id == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.JSON(200, gin.H{"message": "Album removed"})
			return
		}
	}
	c.JSON(404, gin.H{"error": "Album not found"})
}

func main() {

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", addAlbums)
	router.GET("/albums/:id", getSingleAlbum)
	router.PUT("/albums/:id", updateAlbum)
	router.DELETE("/albums/:id", deleteItem)

	defer router.Run("localhost:8080")

}
