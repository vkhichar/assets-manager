package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/vkhichar/assets-manager/contract"
	"github.com/vkhichar/assets-manager/domain"
)

type MockAssetRepo struct {
	mock.Mock
}

func (m *MockAssetRepo) CreateAsset(ctx context.Context, asset *domain.Asset) (*domain.Asset, error) {
	args := m.Called(ctx, asset)

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

func (mock *MockAssetRepo) FindAsset(context context.Context, id int) (*domain.Asset, error) {

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
func (mock *MockAssetRepo) GetAllAssets() ([]domain.Asset, error) {

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

func (mock *MockAssetRepo) UpdateAsset(ctx context.Context, asset *contract.UpadateAssetRequest) (*domain.Asset, error) {

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
func (mock *MockAssetRepo) DeleteAsset(ctx context.Context, id int) (*domain.Asset, error) {

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
