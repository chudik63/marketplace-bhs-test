package service

import (
	"context"
	"marketplace-bhs-test/internal/entity"
	"marketplace-bhs-test/internal/repository"
)

type AssetService interface {
	CreateAsset(ctx context.Context, asset *entity.Asset) error
	DeleteAsset(ctx context.Context, id uint64) error
	BuyAsset(ctx context.Context, id, userID uint64) error
}

type assetService struct {
	repo repository.AssetRepository
}

func NewAssetService(repository repository.AssetRepository) AssetService {
	return &assetService{repository}
}

func (s *assetService) CreateAsset(ctx context.Context, asset *entity.Asset) error {
	return s.repo.Create(ctx, asset)
}

func (s *assetService) DeleteAsset(ctx context.Context, id uint64) error {
	return s.repo.Delete(ctx, id)
}

func (s *assetService) BuyAsset(ctx context.Context, id uint64, userID uint64) error {
	return s.repo.BuyAsset(ctx, userID, id)
}
