package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Year  int    `json:"year"`
}

func listBooksEndpoint(c *gin.Context) {

	var books = []book{
		book{Title: "The Black Swan", Year: 2010, ID: 1},
		book{Title: "Skin in the Game", Year: 2012, ID: 2},
	}

	fmt.Println("First Length:", len(books))

	c.JSON(http.StatusOK, books)
}

func main() {
	router := gin.Default()

	booksRoutes := router.Group("/books")
	{
		booksRoutes.GET("/", listBooksEndpoint)
	}

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
