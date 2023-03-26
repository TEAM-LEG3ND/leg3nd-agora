package service

import (
	"github.com/google/wire"
	"leg3nd-agora/internal/core/ports"
)

var Set = wire.NewSet(ProvideAccountService, wire.Bind(new(ports.AccountService), new(*AccountService)))
