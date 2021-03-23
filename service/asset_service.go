package service

import (
	"context"
	"errors"

	"github.com/vkhichar/assets-manager/domain"
	"github.com/vkhichar/assets-manager/repository"
)

var ErrInvalidId = errors.New("Invalid ID")

type AssetService interface {
	CreateAsset(ctx context.Context, asset *domain.Asset) (*domain.Asset, error)
	FindAsset(ctx context.Context, id int) (*domain.Asset, error)
}

type assetService struct {
	assetRepo repository.AssetRepository
}

func NewAssetService(repo repository.AssetRepository) AssetService {
	return &assetService{
		assetRepo: repo,
	}
}

func (service *assetService) CreateAsset(ctx context.Context, asset *domain.Asset) (*domain.Asset, error) {
	asset, err := service.assetRepo.CreateAsset(ctx, asset)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

func (service *assetService) FindAsset(ctx context.Context, id int) (*domain.Asset, error) {
	asset, err := service.assetRepo.FindAsset(ctx, id)

	if err != nil {
		return nil, err
	}
	if asset == nil {
		return nil, ErrInvalidId
	}
	return asset, nil
}
