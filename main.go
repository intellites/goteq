package main

import (
	"log"

	"github.com/intellites/goteq/config"
)

func main() {
	// Server starting
	log.Println("[INIT] Server starting")

	// Load server config
	config.NewServer()

	// Server started
	log.Println("[INIT] Server started")
}
