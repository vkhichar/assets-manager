package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/vkhichar/assets-manager/contract"
	"github.com/vkhichar/assets-manager/domain"
)

type MockAssetService struct {
	mock.Mock
}

func (mock *MockAssetService) CreateAsset(ctx context.Context, asset *domain.Asset) (*domain.Asset, error) {
	args := mock.Called(ctx, asset)

	var newAsset *domain.Asset
	if args[0] != nil {
		newAsset = args[0].(*domain.Asset)
	}
	var err error
	if args[1] != nil {
		err = args[1].(error)
	}

	return newAsset, err
}

func (mock *MockAssetService) FindAsset(context context.Context, id int) (*domain.Asset, error) {

	args := mock.Called(context, id)
	var asset *domain.Asset
	if args[0] != nil {
		asset = args[0].(*domain.Asset)
	}
	var err error
	if args[1] != nil {
		err = args[1].(error)
	}
	return asset, err

}
func (mock *MockAssetService) GetAssets() ([]domain.Asset, error) {

	args := mock.Called()
	var asset []domain.Asset
	if args[0] != nil {
		asset = args[0].([]domain.Asset)
	}
	var err error
	if args[1] != nil {
		err = args[1].(error)
	}
	return asset, err

}
func (mock *MockAssetService) UpdateAsset(context context.Context, asset *contract.UpadateAssetRequest) (*domain.Asset, error) {

	args := mock.Called(context, asset)
	var ret_asset *domain.Asset
	if args[0] != nil {
		ret_asset = args[0].(*domain.Asset)
	}
	var err error
	if args[1] != nil {
		err = args[1].(error)
	}
	return ret_asset, err

}
func (mock *MockAssetService) DeleteAsset(context context.Context, id int) (*domain.Asset, error) {

	args := mock.Called(context, id)
	var asset *domain.Asset
	if args[0] != nil {

		asset = args[0].(*domain.Asset)

	}
	var err error
	if args[1] != nil {
		err = args[1].(error)
	}
	return asset, err

}
