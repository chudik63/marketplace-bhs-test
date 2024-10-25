package service

import (
	"marketplace-bhs-test/internal/infrastructure/hash"
	"marketplace-bhs-test/internal/repository"
)

type SignInInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserService interface {
	SignUp(input SignInInput) error
}

type userService struct {
	hasher hash.PasswordHasher
	repo   repository.UserRepository
}

func NewUserService(hasher *hash.SHA1Hasher) *userService {
	return &userService{
		hasher: hasher,
	}
}

func (s *userService) SignUp(input SignInInput) error {
	passwordHash, err := s.hasher.Hash(input.Password)
	if err != nil {
		return err
	}

	return nil
}