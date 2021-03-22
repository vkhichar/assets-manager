package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/vkhichar/assets-manager/contract"
	"github.com/vkhichar/assets-manager/domain"
)

type MockRepo struct {
	mock.Mock
}

func (mock *MockRepo) FindAsset(context context.Context, id int) (*domain.Asset, error) {

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
func (mock *MockRepo) GetAllAssets() ([]domain.Asset, error) {

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

func (mock *MockRepo) UpdateAsset(ctx context.Context, asset *contract.UpadateAssetRequest) (*domain.Asset, error) {

	args := mock.Called(ctx, asset)

	var asset_ret *domain.Asset
	if args[0] != nil {
		asset_ret = args[0].(*domain.Asset)
	}
	var err error
	if args[1] != nil {
		err = args[1].(error)
	}

	return asset_ret, err
}
func (mock *MockRepo) DeleteAsset(ctx context.Context, id int) (*domain.Asset, error) {

	args := mock.Called(ctx, id)
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
