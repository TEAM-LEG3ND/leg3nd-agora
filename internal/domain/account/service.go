package account

import "fmt"

type Service struct {
	repository Repository
}

func ProvideService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) CreateAccount(email string, nickname string, fullName string, oAuthProvider OAuthProvider) (*Account, error) {
	newAccount := NewAccount(email, nickname, fullName, oAuthProvider)

	if err := s.repository.Save(newAccount); err != nil {
		return nil, fmt.Errorf("error creating new account: %w", err)
	}

	return newAccount, nil
}

func (s *Service) FindAccountById(id int64) (*Account, error) {
	if acc, err := s.repository.FindById(id); err != nil {
		return nil, fmt.Errorf("error finding account by id: %w", err)
	} else if acc == nil {
		return nil, fmt.Errorf("entity not found")
	} else {
		return acc, nil
	}
}
