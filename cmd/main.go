package main

import (
	_ "github.com/go-sql-driver/mysql"
	"leg3nd-agora/internal/server"
	"log"
)

func main() {
	app, _, _ := server.New()
	port := ":8080"
	log.Fatal(app.Listen(port))
}
