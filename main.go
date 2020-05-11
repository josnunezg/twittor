package main

import (
	"log"

	"github.com/josnunezg/twittor/bd"
	"github.com/josnunezg/twittor/handlers"
)

func main() {
	if bd.CheckingConnection() == 0 {
		log.Fatal("Without connection to the BD")
		return
	}

	handlers.Handlers()
}
