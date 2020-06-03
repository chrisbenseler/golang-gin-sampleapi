package main

import (
	"fmt"
	"golang-gin-sampleapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
	//"go.mongodb.org/mongo-driver/bson"
	"github.com/globalsign/mgo/bson"
)

var books = []models.Book{
	//models.Book{Title: "The Black Swan", Year: 2010, ID: "1"},
	//models.Book{Title: "Skin in the Game", Year: 2012, ID: "2"},
}

var connection = models.Db()

func listBooksEndpoint(c *gin.Context) {
	book := &models.Book{}
	var books = []models.Book{}
	find := connection.Collection("books").Find(bson.M{})

	for find.Next(book) {
		books = append(books, *book)
	}
	c.JSON(http.StatusOK, books)
}

func createBookEndpoint(c *gin.Context) {
	var newBook models.Book

	if c.ShouldBind(&newBook) == nil {
		connection.Collection("books").Save(&newBook)
		c.JSON(http.StatusCreated, newBook)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{})
	}

}

func FindById(a []models.Book, id string) int {
	fmt.Println("id to search", id)
	for i, _ := range a {
		//if id == n.ID {
		if id == "aaa" {
			return i
		}
	}
	return -1
}

func FindBookByID(id string) (models.Book, error) {
	book := &models.Book{}
	err := connection.Collection("books").FindById(bson.ObjectIdHex(id), book)
	return *book, err
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

	books[bookIndex] = newBook

	c.JSON(http.StatusOK, newBook)
}

func RemoveBookByID(id string) error {
	err := connection.Collection("books").DeleteOne(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

func deleteBookEndpoint(c *gin.Context) {
	id := c.Param("id")

	err := RemoveBookByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

func getBookEndpoint(c *gin.Context) {
	id := c.Param("id")
	book, err := FindBookByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
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
