package repository

import (
	"context"
	"marketplace-bhs-test/internal/entity"

	"gorm.io/gorm"
)

type AssetRepository interface {
	Create(ctx context.Context, asset *entity.Asset) error
	Delete(ctx context.Context, id uint) error
}

type assetRepository struct {
	db *gorm.DB
}

func NewAssetRepository(db *gorm.DB) AssetRepository {
	return &assetRepository{db}
}

func (r *assetRepository) Create(ctx context.Context, asset *entity.Asset) error {
	return r.db.WithContext(ctx).Create(asset).Error
}

func (r *assetRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&entity.Asset{}, id).Error
}
