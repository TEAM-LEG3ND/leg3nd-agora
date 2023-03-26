//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

var ServerSet = wire.NewSet(AppSet)

func New() (*fiber.App, func(), error) {
	wire.Build(ServerSet)
	return nil, nil, nil
}
