package main

import (
	"fmt"
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

func FindById(a []Book, id string) int {
	fmt.Println("id to search", id)
	for i, n := range a {
		if id == n.ID {
			return i
		}
	}
	return -1
}

func updateBookEndpoint(c *gin.Context) {
	id := c.Param("id")
	bookIndex := FindById(books, id)

	if bookIndex == -1 {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}
	c.JSON(http.StatusOK, books[bookIndex])
}

func main() {

	router := gin.Default()

	booksRoutes := router.Group("/books")
	{
		booksRoutes.GET("/", listBooksEndpoint)
		booksRoutes.POST("/", createBookEndpoint)
		booksRoutes.PUT("/:id", updateBookEndpoint)

	}

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
