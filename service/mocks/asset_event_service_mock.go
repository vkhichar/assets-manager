package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/vkhichar/assets-manager/domain"
)

type MockAssetEventService struct {
	mock.Mock
}

func (m *MockAssetEventService) PostCreateAssetEvent(ctx context.Context, asset *domain.Asset) (string, error) {

	args := m.Called(ctx, asset)

	var newUser string
	if args[0] != nil {
		newUser = args[0].(string)
	}
	var err error
	if args[1] != nil {
		err = args[1].(error)
	}
	return newUser, err
}

func (m *MockAssetEventService) PostUpdateAssetEvent(ctx context.Context, asset *domain.Asset) (string, error) {
	args := m.Called(ctx, asset)

	var newUser string
	if args[0] != nil {
		newUser = args[0].(string)
	}
	var err error
	if args[1] != nil {
		err = args[1].(error)
	}
	return newUser, err
}
