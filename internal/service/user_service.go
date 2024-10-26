package service

import (
	"context"
	"marketplace-bhs-test/internal/auth"
	"marketplace-bhs-test/internal/entity"
	"marketplace-bhs-test/internal/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	SignUp(ctx context.Context, input *SignUpInput) error
	SignIn(ctx context.Context, input *SignInInput) (Tokens, error)
	UpdateBalance(ctx context.Context, userId uint64, count float64) error
}

type userService struct {
	repo         repository.UserRepository
	tokenManager auth.TokenManager

	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewUserService(repository repository.UserRepository, tokenManager auth.TokenManager, accessTokenTTL, refreshTokenTTL time.Duration) *userService {
	return &userService{
		repo:            repository,
		tokenManager:    tokenManager,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
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

func (s *userService) SignIn(ctx context.Context, input *SignInInput) (Tokens, error) {
	user, err := s.repo.GetByName(ctx, input.Name)
	if err != nil {
		return Tokens{}, err
	}
	passwordOK := bcrypt.CompareHashAndPassword([]byte(user.Password_hash), []byte(input.Password))
	if passwordOK != nil {
		return Tokens{}, passwordOK
	}

	var tokens Tokens

	tokens.AccessToken, err = s.tokenManager.NewJWT(user.ID, s.accessTokenTTL)
	if err != nil {
		return tokens, err
	}
	tokens.AccessTokenTTL = s.accessTokenTTL

	tokens.RefreshToken, err = s.tokenManager.NewRefreshToken()
	if err != nil {
		return tokens, err
	}
	tokens.RefreshTokenTTL = s.refreshTokenTTL

	return tokens, nil
}

func (s *userService) UpdateBalance(ctx context.Context, userID uint64, count float64) error {
	user, err := s.repo.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	return s.repo.UpdateBalance(ctx, userID, user.Balance+count)
}
