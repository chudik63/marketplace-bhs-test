package service

import (
	"context"
	"errors"
	"marketplace-bhs-test/internal/auth"
	"marketplace-bhs-test/internal/entity"
	"marketplace-bhs-test/internal/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	SignUp(ctx context.Context, input *SignUpInput) error
	SignIn(ctx context.Context, input *SignUpInput) (Tokens, error)
}

type userService struct {
	repo         repository.UserRepository
	tokenManager auth.TokenManager

	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{
		repo: repository,
	}
}

func (s *userService) SignUp(ctx context.Context, input *SignUpInput) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &entity.User{
		Username:      input.Name,
		Password_hash: string(hashedPassword),
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *userService) SignIn(ctx context.Context, input *SignUpInput) (Tokens, error) {
	user, err := s.repo.GetByCredentials(ctx, input.Name)
	passwordOK := bcrypt.CompareHashAndPassword([]byte(user.Password_hash), []byte(input.Password))
	if err != nil || passwordOK != nil {
		return Tokens{}, errors.New("wrong name or password")
	}

	return Tokens{}, nil
}
