package models

import (
	"log"

	"github.com/go-bongo/bongo"
)

func Db() *bongo.Connection {
	config := &bongo.Config{
		ConnectionString: "localhost",
		Database:         "bongotest",
	}

	connection, err := bongo.Connect(config)

	if err != nil {
		log.Fatal(err)
	}

	return connection
}
