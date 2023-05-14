package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"leg3nd-agora/internal/api"
	"leg3nd-agora/internal/core/service"
	"leg3nd-agora/internal/repository"
	"leg3nd-agora/internal/server/middleware/internalchecker"
	"log"
)

var AppSet = wire.NewSet(ProvideApp, api.Set, service.Set, repository.Set)

func ProvideApp(accountHandler *api.AccountHandlers) *fiber.App {
	app := fiber.New()
	app.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.SendString("pong")
	})

	v1 := app.Group("/v1")
	v1.Get("/test", accountHandler.TestGateway)

	internal := app.Group("/internal")
	internal.Use(internalchecker.New())
	internalV1 := internal.Group("/v1")
	internalV1Account := internalV1.Group("/account")

	internalV1Account.Post("", accountHandler.CreateAccount)
	internalV1Account.Get("/:id", accountHandler.FindAccountById)
	internalV1Account.Patch("/:id", accountHandler.UpdateAccount)
	internalV1Account.Get("/email/:email", accountHandler.FindAccountByEmail)

	log.Println("Application is starting...")
	return app
}
