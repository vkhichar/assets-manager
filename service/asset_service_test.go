package service_test

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vkhichar/assets-manager/domain"
	mockRepo "github.com/vkhichar/assets-manager/repository/mocks"
	"github.com/vkhichar/assets-manager/service"
)

func TestAssetService_CreateAsset_When_CreateAssetReturnsError(t *testing.T) {
	ctx := context.Background()

	obj := domain.Asset{
		ID:            655,
		Status:        0,
		Category:      "Laptops",
		Specification: &json.RawMessage{},
	}

	mockAssetRepo := &mockRepo.MockAssetRepo{}

	mockAssetRepo.On("CreateAsset", ctx, &obj).Return(nil, errors.New("some db error"))

	assetService := service.NewAssetService(mockAssetRepo)
	asset, err := assetService.CreateAsset(ctx, &obj)

	assert.Error(t, err)
	assert.Nil(t, asset)
}

func TestAssetService_CreateAsset_Success(t *testing.T) {
	ctx := context.Background()

	obj := domain.Asset{
		ID:            1,
		Name:          "Ideapad",
		Category:      "Laptop",
		Status:        0,
		InitCost:      50000,
		Specification: &json.RawMessage{},
	}

	mockAssetRepo := &mockRepo.MockAssetRepo{}

	mockAssetRepo.On("CreateAsset", ctx, &obj).Return(&obj, nil)

	assetService := service.NewAssetService(mockAssetRepo)
	dbAsset, err := assetService.CreateAsset(ctx, &obj)

	assert.NoError(t, err)
	assert.Equal(t, &obj, dbAsset)
}

func TestAssetService_FindAsset_Returns_error(t *testing.T) {
	ctx := context.Background()
	id := 1
	mockAssetRepo := &mockRepo.MockAssetRepo{}
	mockAssetRepo.On("FindAsset", ctx, id).Return(nil, errors.New("invalid id"))
	assetService := service.NewAssetService(mockAssetRepo)

	asset, err := assetService.FindAsset(ctx, id)

	assert.Error(t, err)
	assert.Equal(t, "invalid id", err.Error())
	assert.Nil(t, asset)

}

func TestAssetService_FindAsset_Returns_Success(t *testing.T) {
	ctx := context.Background()
	obj := domain.Asset{
		ID:            1,
		Name:          "Ideapad",
		Category:      "Laptop",
		Status:        0,
		InitCost:      50000,
		Specification: &json.RawMessage{},
	}
	id := 1
	mockAssetRepo := &mockRepo.MockAssetRepo{}
	mockAssetRepo.On("FindAsset", ctx, id).Return(&obj, nil)
	assetService := service.NewAssetService(mockAssetRepo)

	asset, err := assetService.FindAsset(ctx, id)

	assert.Nil(t, err)
	assert.NotNil(t, asset)

}
