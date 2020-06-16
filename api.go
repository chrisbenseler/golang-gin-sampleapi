package main

import (
	"errors"
	"net/http"
	"strings"

	"./models"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
)

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

func FindBookByID(id string) (models.Book, error) {
	book := &models.Book{}
	err := connection.Collection("books").FindById(bson.ObjectIdHex(id), book)
	return *book, err
}

func updateBookEndpoint(c *gin.Context) {
	id := c.Param("id")
	var newBook models.Book

	if c.ShouldBind(&newBook) == nil {
		newBook.SetId(bson.ObjectIdHex(id))
		connection.Collection("books").Save(&newBook)
		c.JSON(http.StatusCreated, newBook)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{})
	}
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

var token = "111"

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.Request.Header.Get("Authorization")

		if authorizationHeader == "" {
			c.JSON(http.StatusUnauthorized, errors.New("Not authenticated"))
			c.Abort()
			return
		}

		if strings.Split(authorizationHeader, "Bearer ")[1] == token {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, errors.New("Not authorized"))
			c.Abort()
			return
		}

	}
}

func signinEndpoint(c *gin.Context) {
	users := []models.User{
		{
			Name:     "teste",
			Password: "teste",
		},
	}

	var signinUser models.User

	if err := c.ShouldBindJSON(&signinUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	println(signinUser.Name, signinUser.Password)

	isAuth := false
	for _, n := range users {
		println(n.Name)
		if signinUser.Name == n.Name && signinUser.Password == n.Password {
			isAuth = true
		}
	}

	if isAuth == true {
		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{})
	}

}

func main() {

	router := gin.Default()

	router.POST("/signin", signinEndpoint)

	booksRoutes := router.Group("/books")
	{
		booksRoutes.GET("/", TokenAuthMiddleware(), listBooksEndpoint)
		booksRoutes.POST("/", createBookEndpoint)
		booksRoutes.PUT("/:id", updateBookEndpoint)
		booksRoutes.GET("/:id", getBookEndpoint)
		booksRoutes.DELETE("/:id", deleteBookEndpoint)
	}

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
