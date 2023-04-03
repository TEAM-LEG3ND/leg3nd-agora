package main

import (
	"github.com/joho/godotenv"
	"leg3nd-agora/internal/server"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app, _, _ := server.New()
	port := ":8080"
	log.Fatal(app.Listen(port))
}
