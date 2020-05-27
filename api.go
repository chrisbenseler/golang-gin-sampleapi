package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type Book struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Year  int    `json:"year"`
}

var books = []Book{
	Book{Title: "The Black Swan", Year: 2010, ID: "1"},
	Book{Title: "Skin in the Game", Year: 2012, ID: "2"},
}

func listBooksEndpoint(c *gin.Context) {

	c.JSON(http.StatusOK, books)
}

func createBookEndpoint(c *gin.Context) {
	var newBook Book

	if c.ShouldBind(&newBook) == nil {
		uID := uuid.Must(uuid.NewV4())
		newBook.ID = uID.String()
		books = append(books, newBook)
		c.JSON(http.StatusCreated, newBook)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{})
	}

}

func main() {

	router := gin.Default()

	booksRoutes := router.Group("/books")
	{
		booksRoutes.GET("/", listBooksEndpoint)
		booksRoutes.POST("/", createBookEndpoint)
	}

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
