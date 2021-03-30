package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/vkhichar/assets-manager/custom_errors"
	"github.com/vkhichar/assets-manager/repository"
)

type AssetAllocationService interface {
	AllocateAsset(context.Context, int, int, string) error
	DeallocateAsset(context.Context, int, int, string) error
}

type assetAllocationService struct {
	asset_alloc_repo repository.AssetsAllocationRepository
}

func NewAssetAllocationService(asset_repo repository.AssetsAllocationRepository) AssetAllocationService {
	return &assetAllocationService{asset_alloc_repo: asset_repo}
}

func (asset_alloc_service *assetAllocationService) AllocateAsset(context context.Context, user_id int, asset_id int, date string) error {

	err := asset_alloc_service.asset_alloc_repo.AllocateAsset(context, user_id, asset_id, date)

	if err == sql.ErrNoRows {
		return custom_errors.InvalidAssetOrUserIdError
	}
	if err == custom_errors.InvalidAssetStatusError {

		return custom_errors.InvalidAssetStatusError

	}
	if err != nil {
		return errors.New("something went wrong")
	}
	return nil
}

func (asset_alloc_service *assetAllocationService) DeallocateAsset(context context.Context, user_id int, asset_id int, date_dealloc string) error {

	err := asset_alloc_service.asset_alloc_repo.DeallocateAsset(context, user_id, asset_id, date_dealloc)

	if err == sql.ErrNoRows {
		return custom_errors.InvalidAllocationError
	}

	if err != nil {
		return errors.New("something went wrong")
	}

	return nil
}
