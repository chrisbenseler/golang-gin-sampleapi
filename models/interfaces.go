package models

import "github.com/go-bongo/bongo"

type Book struct {
	bongo.DocumentBase `bson:",inline"`
	//ID                 string `json:"id"`
	Title string `json:"title"`
	Year  int    `json:"year"`
}

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
