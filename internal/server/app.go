package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"leg3nd-agora/internal/api"
	"leg3nd-agora/internal/core/service"
	"leg3nd-agora/internal/repository"
	"log"
)

var AppSet = wire.NewSet(ProvideApp, api.Set, service.Set, repository.Set)

func ProvideApp(accountHandler *api.AccountHandlers) *fiber.App {
	app := fiber.New()

	v1 := app.Group("/v1")

	v1.Post("", accountHandler.CreateAccount)
	v1.Get("/:id", accountHandler.FindAccountById)

	log.Println("Application is starting...")
	return app
}