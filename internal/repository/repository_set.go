package repository

import (
	"github.com/google/wire"
	"leg3nd-agora/internal/core/ports"
)

var Set = wire.NewSet(ProvideAccountRepository, ProvideClient, wire.Bind(new(ports.AccountRepository), new(*AccountRepository)))
