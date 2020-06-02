package models

import "github.com/go-bongo/bongo"

type Book struct {
	bongo.DocumentBase `bson:",inline"`
	//ID                 string `json:"id"`
	Title string `json:"title"`
	Year  int    `json:"year"`
}
