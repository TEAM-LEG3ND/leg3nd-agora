package repository

import (
	"context"
	"fmt"
	"leg3nd-agora/ent"
	"leg3nd-agora/ent/account"
	"leg3nd-agora/internal/core/domain"
	"leg3nd-agora/internal/core/ports"
)

type AccountRepository struct {
	client *ent.Client
}

func ProvideAccountRepository(client *ent.Client) *AccountRepository {
	return &AccountRepository{client: client}
}

var _ ports.AccountRepository = (*AccountRepository)(nil)

func (r *AccountRepository) Save(ctx context.Context, ac *domain.Account) (*domain.Account, error) {
	savedAccount, err := r.client.Account.
		Create().
		SetEmail(ac.Email).
		SetNickname(ac.Nickname).
		SetFullName(ac.FullName).
		SetOauthProvider(account.OauthProvider(ac.OAuthProvider.String())).
		Save(ctx)
	domainAccount := domain.Account{
		Id:            savedAccount.ID,
		Email:         savedAccount.Email,
		Nickname:      savedAccount.Nickname,
		FullName:      savedAccount.FullName,
		OAuthProvider: domain.OAuthProvider(savedAccount.OauthProvider),
	}

	return &domainAccount, err
}

func (r *AccountRepository) FindById(ctx context.Context, accountId int64) (*domain.Account, error) {
	ac, err := r.client.Account.Query().Where(account.ID(accountId)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find account by id : %w", err)
	}

	domainAccount := domain.Account{
		Id:            ac.ID,
		Email:         ac.Email,
		Nickname:      ac.Nickname,
		FullName:      ac.FullName,
		OAuthProvider: domain.OAuthProvider(ac.OauthProvider),
	}

	return &domainAccount, nil
}

func (r *AccountRepository) FindByEmail(ctx context.Context, email string) (*domain.Account, error) {
	ac, err := r.client.Account.Query().Where(account.Email(email)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find account by email : %w", err)
	}

	domainAccount := domain.Account{
		Id:            ac.ID,
		Email:         ac.Email,
		Nickname:      ac.Nickname,
		FullName:      ac.FullName,
		OAuthProvider: domain.OAuthProvider(ac.OauthProvider),
	}

	return &domainAccount, nil
}
