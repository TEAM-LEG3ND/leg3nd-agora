package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"leg3nd-agora/internal/core/domain"
	"leg3nd-agora/internal/core/ports"
	"log"
	"math/rand"
)

type AccountService struct {
	repository ports.AccountRepository
}

// Check implementation of port interface
var _ ports.AccountService = (*AccountService)(nil)

func ProvideAccountService(repository ports.AccountRepository) *AccountService {
	return &AccountService{repository: repository}
}

func (s *AccountService) CreateAccount(ctx context.Context, email string, nickname string, fullName string, oAuthProvider domain.OAuthProvider) (*domain.Account, error) {
	if nickname == "" {
		nickname = getRandomNickname()
	}
	log.Println(nickname)
	newAccount := domain.NewAccount(email, nickname, fullName, oAuthProvider)

	savedAccount, err := s.repository.Save(ctx, newAccount)
	if err != nil {
		return nil, fmt.Errorf("error creating new account: %w", err)
	}

	return savedAccount, nil
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

func getRandomNickname() string {
	colors := []string{"빨강", "주황", "노랑", "초록", "파랑", "보라", "누런", "갈색", "남색", "하늘", "하양", "검정", "까망", "누런"}
	foods := []string{"마라샹궈", "알리오올리오", "김피탕", "짜장면", "짬뽕", "탕수육", "떡볶이", "양꼬치"}
	uuidStr := uuid.New().String()

	nickname := fmt.Sprintf("%s%s-%s", colors[rand.Intn(len(colors))], foods[rand.Intn(len(foods))], uuidStr)

	return nickname
}
