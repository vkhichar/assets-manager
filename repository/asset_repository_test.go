package repository_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vkhichar/assets-manager/config"
	"github.com/vkhichar/assets-manager/contract"
	"github.com/vkhichar/assets-manager/custom_errors"
	"github.com/vkhichar/assets-manager/domain"
	"github.com/vkhichar/assets-manager/repository"
)

func TestDbConnection(t *testing.T) {

	os.Setenv("APP_PORT", "9000")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USERNAME", "postgres")
	os.Setenv("DB_PASSWORD", "12345")
	os.Setenv("DB_NAME", "asset_management")

	err := config.Init()
	repository.InitDB()

	assert.NoError(t, err)
}

func TestFindAsset_When_ReturnsError(t *testing.T) {

	ctx := context.Background()

	assetRepo := repository.NewAssetRepository()

	asset, err := assetRepo.FindAsset(ctx, -1)

	assert.Error(t, err)
	assert.Nil(t, asset)
	assert.Equal(t, custom_errors.InvalidIdError.Error(), err.Error())

}

func TestFindAsset_When_ReturnsAsset(t *testing.T) {

	ctx := context.Background()

	assetRepo := repository.NewAssetRepository()
	asset, err := assetRepo.FindAsset(ctx, 4)

	assert.NoError(t, err)
	assert.NotNil(t, asset)
}

func TestGetAllAssets_When_ReturnsListOfAssets(t *testing.T) {

	assetRepo := repository.NewAssetRepository()
	var expected_list []domain.Asset

	returned_list, err := assetRepo.GetAllAssets()

	assert.NoError(t, err)
	assert.IsType(t, expected_list, returned_list)
}

func TestUpdateAssets_When_ReturnsError(t *testing.T) {

	ctx := context.Background()
	assetRepo := repository.NewAssetRepository()

	asset_request := &contract.UpadateAssetRequest{
		Id:       -1,
		Name:     "hp",
		Category: "laptop",
		InitCost: 1000,
		Status:   0,
	}
	asset, err := assetRepo.UpdateAsset(ctx, asset_request)

	assert.Error(t, err)
	assert.Nil(t, asset)

}

func TestUpdateAssets_When_ReturnsAsset(t *testing.T) {

	ctx := context.Background()

	asset_request := &contract.UpadateAssetRequest{
		Id:       4,
		Name:     "hp",
		Category: "laptop",
		InitCost: 1000,
		Status:   0,
	}

	assetRepo := repository.NewAssetRepository()
	asset_response, err := assetRepo.UpdateAsset(ctx, asset_request)

	assert.NoError(t, err)
	assert.NotNil(t, asset_response)

}

func TestDeleteAssets_When_ReturnsError(t *testing.T) {

	ctx := context.Background()
	assetRepo := repository.NewAssetRepository()

	asset, err := assetRepo.DeleteAsset(ctx, -1)

	assert.Error(t, err)
	assert.Nil(t, asset)
	assert.Equal(t, custom_errors.InvalidIdError.Error(), err.Error())

}

func TestDeleteAssets_When_ReturnsAsset(t *testing.T) {

	ctx := context.Background()
	assetRepo := repository.NewAssetRepository()

	asset, err := assetRepo.DeleteAsset(ctx, 4)

	assert.NoError(t, err)
	assert.NotNil(t, asset)

}
