package repository

import (
	"context"
	"fmt"
	"leg3nd-agora/ent"
	"leg3nd-agora/ent/account"
	"leg3nd-agora/internal/core/domain"
	"leg3nd-agora/internal/core/ports"
	"log"
)

type AccountRepository struct {
	client *ent.Client
}

func ProvideAccountRepository(client *ent.Client) *AccountRepository {
	return &AccountRepository{client: client}
}

var _ ports.AccountRepository = (*AccountRepository)(nil)

func (r *AccountRepository) Create(ctx context.Context, ac *domain.Account) (*domain.Account, error) {
	savedAccount, err := r.client.Account.
		Create().
		SetEmail(ac.Email).
		SetNillableNickname(ac.Nickname).
		SetFullName(ac.FullName).
		SetOauthProvider(account.OauthProvider(ac.OAuthProvider.String())).
		SetStatus("draft").
		Save(ctx)
	if err != nil {
		log.Println("err", err)
		return nil, err
	}
	domainAccount := domain.Account{
		Id:            savedAccount.ID,
		Email:         savedAccount.Email,
		Nickname:      savedAccount.Nickname,
		FullName:      savedAccount.FullName,
		OAuthProvider: domain.OAuthProvider(savedAccount.OauthProvider),
		Status:        domain.Status(savedAccount.Status),
	}

	return &domainAccount, nil
}

func (r *AccountRepository) Update(ctx context.Context, ac *domain.Account) (*domain.Account, error) {
	oAuthProvider, err := convertOAuthProvider(ac.OAuthProvider)
	if err != nil {
		return nil, err
	}
	status, err := convertStatus(ac.Status)
	if err != nil {
		return nil, err
	}

	updatedAccount, err := r.client.Account.
		UpdateOneID(ac.Id).
		SetEmail(ac.Email).
		SetFullName(ac.FullName).
		SetNillableNickname(ac.Nickname).
		SetOauthProvider(oAuthProvider).
		SetStatus(status).
		Save(ctx)

	if err != nil {
		log.Printf("error occurred while updating account, %v", err)
	}
	domainAccount := &domain.Account{
		Id:            updatedAccount.ID,
		Email:         updatedAccount.Email,
		Nickname:      updatedAccount.Nickname,
		FullName:      updatedAccount.FullName,
		OAuthProvider: domain.OAuthProvider(updatedAccount.OauthProvider),
		Status:        domain.Status(updatedAccount.Status),
	}

	return domainAccount, nil
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
		Status:        domain.Status(ac.Status),
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
		Status:        domain.Status(ac.Status),
	}

	return &domainAccount, nil
}

func (r *AccountRepository) ExistByNickname(ctx context.Context, nickname string) (bool, error) {
	exist, err := r.client.Account.Query().Where(account.Nickname(nickname)).Exist(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to check exist account by nickname : %w", err)
	}
	return exist, nil
}

func convertOAuthProvider(domainOAuthProvider domain.OAuthProvider) (account.OauthProvider, error) {
	switch domainOAuthProvider {
	case domain.Google:
		return account.OauthProviderGoogle, nil
	default:
		return "", fmt.Errorf("convert OAuthProvider error occurred, no such OAuthProvider")
	}
}

func convertStatus(domainStatus domain.Status) (account.Status, error) {
	switch domainStatus {
	case domain.Draft:
		return account.StatusDraft, nil
	case domain.Ok:
		return account.StatusOk, nil
	case domain.Suspended:
		return account.StatusSuspended, nil
	case domain.Withdraw:
		return account.StatusWithdraw, nil
	default:
		return "", fmt.Errorf("convert Status error occurred, no such Status")
	}
}
