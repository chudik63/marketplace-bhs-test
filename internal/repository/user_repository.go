package repository

import (
	"context"
	"marketplace-bhs-test/internal/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	GetByName(ctx context.Context, name string) (*entity.User, error)
	GetByID(ctx context.Context, userID uint) (*entity.User, error)
	UpdateBalance(ctx context.Context, userID uint, newBalance float64) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepository) GetByName(ctx context.Context, name string) (*entity.User, error) {
	var user entity.User
	if err := r.db.WithContext(ctx).Where("username = ?", name).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByID(ctx context.Context, userID uint) (*entity.User, error) {
	var user entity.User

	if err := r.db.WithContext(ctx).First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateBalance(ctx context.Context, userID uint, newBalance float64) error {
	if err := r.db.WithContext(ctx).Model(&entity.User{}).Where("id = ?", userID).Update("balance", newBalance).Error; err != nil {
		return err
	}
	return nil
}
