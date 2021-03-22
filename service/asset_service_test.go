package service_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vkhichar/assets-manager/contract"
	"github.com/vkhichar/assets-manager/domain"
	mocks "github.com/vkhichar/assets-manager/repository/mock"
	"github.com/vkhichar/assets-manager/service"
)

func TestFindAsset_When_ReturnsError(t *testing.T) {

	ctx := context.Background()
	id := 100
	mockAssetRepo := &mocks.MockRepo{}

	mockAssetRepo.On("FindAsset", ctx, id).Return(nil, errors.New("something went wrong"))

	assetService := service.NewAssetService(mockAssetRepo)
	asset, err := assetService.FindAsset(ctx, id)

	assert.Error(t, err)
	assert.Equal(t, "something went wrong", err.Error())
	assert.Nil(t, asset)

}
func TestFindAsset_When_ReturnsAsset(t *testing.T) {

	ctx := context.Background()
	id := 1
	asset := &domain.Asset{
		Id:       1,
		Name:     "test asset",
		Category: "testing",
		InitCost: 100,
		Status:   0,
	}
	mockAssetRepo := &mocks.MockRepo{}

	mockAssetRepo.On("FindAsset", ctx, id).Return(asset, nil)

	assetService := service.NewAssetService(mockAssetRepo)
	asset_ret, err := assetService.FindAsset(ctx, id)

	assert.Equal(t, asset, asset_ret)
	assert.Nil(t, err)

}

func TestGetAssets__When_ReturnsError(t *testing.T) {

	mockAssetRepo := &mocks.MockRepo{}

	mockAssetRepo.On("GetAllAssets").Return(nil, errors.New("something went wrong"))
	assetService := service.NewAssetService(mockAssetRepo)
	assets, err := assetService.GetAssets()

	assert.Nil(t, assets)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "something went wrong")

}

func TestGetAssets_When_ReturnsAssets(t *testing.T) {

	assets := make([]domain.Asset, 3)
	for i := 0; i < 3; i++ {
		assets[0] = domain.Asset{Id: i, Name: fmt.Sprintf("test_user%d", i), Category: "testing", InitCost: 0, Status: 0}
	}

	mockAssetRepo := &mocks.MockRepo{}
	mockAssetRepo.On("GetAllAssets").Return(assets, nil)
	assetService := service.NewAssetService(mockAssetRepo)
	assets_ret, err := assetService.GetAssets()

	assert.Nil(t, err)
	assert.NoError(t, err)
	assert.Equal(t, assets, assets_ret)
}

func TestUpdateAsset_When_ReturnsError(t *testing.T) {

	ctx := context.Background()

	mockAssetRepo := &mocks.MockRepo{}
	var asset_req *contract.UpadateAssetRequest

	mockAssetRepo.On("UpdateAsset", ctx, asset_req).Return(nil, errors.New("something went wrong"))
	assetService := service.NewAssetService(mockAssetRepo)
	asset, err := assetService.UpdateAsset(ctx, nil)

	assert.Nil(t, asset)
	assert.Error(t, err)
	assert.Equal(t, "something went wrong", err.Error())
}

func TestUpdateAsset_When_ReturnsAsset(t *testing.T) {

	ctx := context.Background()

	mockAssetRepo := &mocks.MockRepo{}
	asset := &domain.Asset{Id: 1, Name: "test_user", Category: "testing", InitCost: 0, Status: 0}
	var asset_req *contract.UpadateAssetRequest
	mockAssetRepo.On("UpdateAsset", ctx, asset_req).Return(asset, nil)

	assetService := service.NewAssetService(mockAssetRepo)
	asset_res, err := assetService.UpdateAsset(ctx, nil)

	assert.Nil(t, err)
	assert.NoError(t, err)
	assert.Equal(t, asset, asset_res)

}

func TestDeleteAsset_When_ReturnsError(t *testing.T) {

	ctx := context.Background()
	id := 100
	mockAssetRepo := &mocks.MockRepo{}

	mockAssetRepo.On("DeleteAsset", ctx, id).Return(nil, errors.New("something went wrong"))

	assetService := service.NewAssetService(mockAssetRepo)
	asset, err := assetService.DeleteAsset(ctx, id)

	assert.Error(t, err)
	assert.Equal(t, "something went wrong", err.Error())
	assert.Nil(t, asset)

}
func TestDeleteAsset_When_ReturnsAsset(t *testing.T) {

	ctx := context.Background()
	id := 1
	specification := json.RawMessage("{test:test}")
	asset := &domain.Asset{
		Id:            1,
		Name:          "test asset",
		Category:      "testing",
		Specification: &specification,
		InitCost:      100,
		Status:        0,
	}
	mockAssetRepo := &mocks.MockRepo{}

	mockAssetRepo.On("DeleteAsset", ctx, id).Return(asset, nil)

	assetService := service.NewAssetService(mockAssetRepo)
	asset_ret, err := assetService.DeleteAsset(ctx, id)

	assert.Equal(t, asset, asset_ret)
	assert.Nil(t, err)

}
