package main

import (
	"log"
)

func main() {
	storage, err := NewPostgresStorage()

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("%+v\n", storage)

	server := NewAPIServer(":3000", storage)
	server.Run()
}