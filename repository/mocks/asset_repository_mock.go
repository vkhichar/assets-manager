package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
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

func (m *MockAssetRepo) FindAsset(ctx context.Context, id int) (*domain.Asset, error) {
	args := m.Called(ctx, id)

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
