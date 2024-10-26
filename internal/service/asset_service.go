package service

import "marketplace-bhs-test/internal/repository"

type AssetService interface {
}

type assetService struct {
}

func NewAssetService(repository repository.AssetRepository) AssetService {
	return &assetService{}
}
