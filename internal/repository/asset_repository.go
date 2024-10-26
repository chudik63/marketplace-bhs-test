package repository

import (
	"context"
	"errors"
	"marketplace-bhs-test/internal/entity"

	"gorm.io/gorm"
)

type AssetRepository interface {
	Create(ctx context.Context, asset *entity.Asset) error
	Delete(ctx context.Context, id uint64) error
	BuyAsset(ctx context.Context, userID, assetID uint64) error
}

type assetRepository struct {
	db       *gorm.DB
	userRepo UserRepository
}

func NewAssetRepository(db *gorm.DB, userRepo UserRepository) AssetRepository {
	return &assetRepository{
		db:       db,
		userRepo: userRepo,
	}
}

func (r *assetRepository) Create(ctx context.Context, asset *entity.Asset) error {
	return r.db.WithContext(ctx).Create(asset).Error
}

func (r *assetRepository) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Unscoped().Delete(&entity.Asset{}, id).Error
}

func (r *assetRepository) BuyAsset(ctx context.Context, userID, assetID uint64) error {
	return r.db.WithContext(ctx).Transaction(func(db *gorm.DB) error {
		var asset entity.Asset

		if err := db.First(&asset, assetID).Error; err != nil {
			return err
		}

		user, err := r.userRepo.GetByID(ctx, userID)
		if err != nil {
			return err
		}

		if user.Balance < asset.Price {
			return errors.New("low balance")
		}

		if err := r.userRepo.UpdateBalance(ctx, userID, user.Balance-asset.Price); err != nil {
			return err
		}

		if err := db.Model(&asset).Update("user_id", userID).Error; err != nil {
			return err
		}

		return nil
	})
}
