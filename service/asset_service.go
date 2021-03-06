package service

import (
	"context"
	"errors"

	"github.com/vkhichar/assets-manager/contract"
	"github.com/vkhichar/assets-manager/domain"
	"github.com/vkhichar/assets-manager/repository"
)

var ErrInvalidId = errors.New("Invalid ID")

type AssetService interface {
	CreateAsset(ctx context.Context, asset *domain.Asset) (*domain.Asset, error)
	FindAsset(context.Context, int) (*domain.Asset, error)
	GetAssets() ([]domain.Asset, error)
	UpdateAsset(ctx context.Context, asset *contract.UpadateAssetRequest) (*domain.Asset, error)
	DeleteAsset(ctx context.Context, id int) (*domain.Asset, error)
}

type assetService struct {
	assetRepo repository.AssetRepository
}

func NewAssetService(repo repository.AssetRepository) AssetService {

	return &assetService{
		assetRepo: repo,
	}

}

func (service *assetService) FindAsset(ctx context.Context, id int) (*domain.Asset, error) {

	asset, err := service.assetRepo.FindAsset(ctx, id)

	if err != nil {
		return nil, err
	}
	if asset == nil {
		return nil, errors.New("error while processing for asset")
	}

	return asset, nil
}

func (service *assetService) GetAssets() ([]domain.Asset, error) {
	assets, err := service.assetRepo.GetAllAssets()

	if err != nil {
		return nil, err
	}
	return assets, nil
}

func (service *assetService) UpdateAsset(ctx context.Context, asset *contract.UpadateAssetRequest) (*domain.Asset, error) {

	assets, err := service.assetRepo.UpdateAsset(ctx, asset)

	if err != nil {
		return nil, err
	}
	return assets, nil

}

func (service *assetService) CreateAsset(ctx context.Context, asset *domain.Asset) (*domain.Asset, error) {
	asset, err := service.assetRepo.CreateAsset(ctx, asset)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

func (service *assetService) DeleteAsset(ctx context.Context, id int) (*domain.Asset, error) {

	assets, err := service.assetRepo.DeleteAsset(ctx, id)

	if err != nil {
		return nil, err
	}

	return assets, nil

}
