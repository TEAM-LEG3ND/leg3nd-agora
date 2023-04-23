package service

import (
	"context"
	"fmt"
	"leg3nd-agora/internal/core/domain"
	"leg3nd-agora/internal/core/ports"
	"log"
)

type AccountService struct {
	repository ports.AccountRepository
}

// Check implementation of port interface
var _ ports.AccountService = (*AccountService)(nil)

func ProvideAccountService(repository ports.AccountRepository) *AccountService {
	return &AccountService{repository: repository}
}

func (s *AccountService) CreateAccount(ctx context.Context, email string, fullName string, oAuthProvider domain.OAuthProvider) (*domain.Account, error) {
	// TODO: add proper validation
	newAccount := domain.NewAccount(email, fullName, oAuthProvider)

	savedAccount, err := s.repository.Create(ctx, newAccount)
	if err != nil {
		return nil, fmt.Errorf("error creating new account: %w", err)
	}

	return savedAccount, nil
}

func (s *AccountService) UpdateAccount(
	ctx context.Context,
	id int64,
	email *string,
	fullName *string,
	nickname *string,
	oAuthProvider *domain.OAuthProvider,
	status *domain.Status,
) (*domain.Account, error) {
	account, err := s.repository.FindById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error occurred when find user by id: %d", id)
	}

	if email != nil {
		account.Email = *email
	}
	if fullName != nil {
		account.FullName = *fullName
	}
	if nickname != nil {
		account.Nickname = nickname
	}
	if oAuthProvider != nil {
		account.OAuthProvider = *oAuthProvider
	}
	if status != nil {
		account.Status = *status
	}

	updatedAccount, err := s.repository.Update(ctx, account)
	if err != nil {
		return nil, fmt.Errorf("error while updating new account: %w", err)
	}

	return updatedAccount, nil
}

func (s *AccountService) FindAccountById(ctx context.Context, id int64) (*domain.Account, error) {
	if acc, err := s.repository.FindById(ctx, id); err != nil {
		return nil, fmt.Errorf("error finding account by id: %w", err)
	} else if acc == nil {
		return nil, fmt.Errorf("entity not found")
	} else {
		return acc, nil
	}
}

func (s *AccountService) FindAccountByEmail(ctx context.Context, email string) (*domain.Account, error) {
	if acc, err := s.repository.FindByEmail(ctx, email); err != nil {
		return nil, fmt.Errorf("error finding account by email: %v", err)
	} else if acc == nil {
		return nil, fmt.Errorf("entity not found by email: %s", email)
	} else {
		return acc, nil
	}
}

func (s *AccountService) checkNicknameExists(ctx context.Context, nickname string) (bool, error) {
	exist, err := s.repository.ExistByNickname(ctx, nickname)
	if err != nil {
		log.Println(err)
		return true, err
	}
	return exist, nil
}
