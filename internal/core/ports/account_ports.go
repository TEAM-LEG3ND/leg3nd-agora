package ports

import (
	"context"
	"leg3nd-agora/internal/core/domain"
)

type AccountService interface {
	CreateAccount(ctx context.Context, email string, nickname string, fullName string, oAuthProvider domain.OAuthProvider) (*domain.Account, error)
	FindAccountById(ctx context.Context, id int64) (*domain.Account, error)
	FindAccountByEmail(ctx context.Context, email string) (*domain.Account, error)
}

type AccountRepository interface {
	Save(ctx context.Context, ac *domain.Account) (*domain.Account, error)
	FindById(ctx context.Context, accountId int64) (*domain.Account, error)
	FindByEmail(ctx context.Context, email string) (*domain.Account, error)
}
