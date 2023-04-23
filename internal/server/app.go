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
	app.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.SendString("pong")
	})
	v1 := app.Group("/v1")

	v1.Post("", accountHandler.CreateAccount)
	v1.Get("/:id", accountHandler.FindAccountById)
	v1.Patch("/:id", accountHandler.UpdateAccount)
	v1.Get("/email/:email", accountHandler.FindAccountByEmail)

	log.Println("Application is starting...")
	return app
}
