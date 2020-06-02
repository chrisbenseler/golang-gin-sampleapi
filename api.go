package main

import (
	"errors"
	"fmt"
	"golang-gin-sampleapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

var books = []models.Book{
	models.Book{Title: "The Black Swan", Year: 2010, ID: "1"},
	models.Book{Title: "Skin in the Game", Year: 2012, ID: "2"},
}

func listBooksEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func createBookEndpoint(c *gin.Context) {
	var newBook models.Book

	if c.ShouldBind(&newBook) == nil {
		uID := uuid.Must(uuid.NewV4())
		newBook.ID = uID.String()
		books = append(books, newBook)
		c.JSON(http.StatusCreated, newBook)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{})
	}

}

func FindById(a []models.Book, id string) int {
	fmt.Println("id to search", id)
	for i, n := range a {
		if id == n.ID {
			return i
		}
	}
	return -1
}

func FindBookById(a []models.Book, id string) (models.Book, error) {
	fmt.Println("id to search", id)
	for _, n := range a {
		if id == n.ID {
			return n, nil
		}
	}
	return models.Book{}, errors.New("No book found")
	//	return nil
}

func updateBookEndpoint(c *gin.Context) {
	id := c.Param("id")
	bookIndex := FindById(books, id)

	if bookIndex == -1 {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	var newBook models.Book
	c.ShouldBind(&newBook)
	newBook.ID = books[bookIndex].ID

	books[bookIndex] = newBook

	c.JSON(http.StatusOK, newBook)
}

func RemoveBookByIndex(s []models.Book, index int) []models.Book {
	return append(s[:index], s[index+1:]...)
}

func deleteBookEndpoint(c *gin.Context) {
	id := c.Param("id")
	bookIndex := FindById(books, id)

	if bookIndex == -1 {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	books = RemoveBookByIndex(books, bookIndex)

	c.JSON(http.StatusNoContent, gin.H{})
}

func getBookEndpoint(c *gin.Context) {
	id := c.Param("id")
	book, err := FindBookById(books, id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}
	c.JSON(http.StatusOK, book)
}

func main() {

	router := gin.Default()

	booksRoutes := router.Group("/books")
	{
		booksRoutes.GET("/", listBooksEndpoint)
		booksRoutes.POST("/", createBookEndpoint)
		booksRoutes.PUT("/:id", updateBookEndpoint)
		booksRoutes.GET("/:id", getBookEndpoint)
		booksRoutes.DELETE("/:id", deleteBookEndpoint)
	}

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
