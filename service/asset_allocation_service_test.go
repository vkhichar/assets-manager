package service_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vkhichar/assets-manager/custom_errors"
	"github.com/vkhichar/assets-manager/repository/mocks"
	"github.com/vkhichar/assets-manager/service"
)

func TestAllocateAsset_When_InvalidUserOrAssetId(t *testing.T) {

	user_id, asset_id, allocationDate := -10, -20, "2000-12-12"

	mockRepo := &mocks.MockAssetAllocationRepo{}
	mockRepo.On("AllocateAsset", context.Background(), user_id, asset_id, allocationDate).Return(sql.ErrNoRows)

	service := service.NewAssetAllocationService(mockRepo)

	err := service.AllocateAsset(context.Background(), user_id, asset_id, allocationDate)

	assert.Error(t, err)
	assert.Equal(t, custom_errors.InvalidAssetOrUserIdError, err)
}

func TestAllocateAsset_When_InvalidStatusError(t *testing.T) {

	user_id, asset_id, allocationDate := 10, 20, "2000-12-12"

	mockRepo := &mocks.MockAssetAllocationRepo{}
	mockRepo.On("AllocateAsset", context.Background(), user_id, asset_id, allocationDate).Return(custom_errors.InvalidAssetStatusError)

	service := service.NewAssetAllocationService(mockRepo)

	err := service.AllocateAsset(context.Background(), user_id, asset_id, allocationDate)

	assert.Error(t, err)
	assert.Equal(t, custom_errors.InvalidAssetStatusError, err)

}

func TestAllocateAsset_When_Success(t *testing.T) {

	user_id, asset_id, allocationDate := 10, 20, "2000-12-12"

	mockRepo := &mocks.MockAssetAllocationRepo{}
	mockRepo.On("AllocateAsset", context.Background(), user_id, asset_id, allocationDate).Return(nil)

	service := service.NewAssetAllocationService(mockRepo)

	err := service.AllocateAsset(context.Background(), user_id, asset_id, allocationDate)

	assert.NoError(t, err)
	assert.Nil(t, err)

}

func TestDeallocateAsset_When_InvaidAllocationError(t *testing.T) {

	user_id, asset_id, deallocationDate := 10, 20, "2000-12-12"

	mockRepo := &mocks.MockAssetAllocationRepo{}
	mockRepo.On("DeallocateAsset", context.Background(), user_id, asset_id, deallocationDate).Return(sql.ErrNoRows)

	service := service.NewAssetAllocationService(mockRepo)

	err := service.DeallocateAsset(context.Background(), user_id, asset_id, deallocationDate)

	assert.Error(t, err)
	assert.Equal(t, custom_errors.InvalidAllocationError, err)

}

func TestDeallocateAsset_When_Success(t *testing.T) {
	user_id, asset_id, deallocationDate := 10, 20, "2000-12-12"

	mockRepo := &mocks.MockAssetAllocationRepo{}
	mockRepo.On("DeallocateAsset", context.Background(), user_id, asset_id, deallocationDate).Return(nil)

	service := service.NewAssetAllocationService(mockRepo)

	err := service.DeallocateAsset(context.Background(), user_id, asset_id, deallocationDate)

	assert.NoError(t, err)
	assert.Nil(t, err)
}
