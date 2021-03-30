package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockAssetAllocationService struct {
	mock.Mock
}

func (mock *MockAssetAllocationService) AllocateAsset(context context.Context, user_id int, asset_id int, allocationDate string) error {

	args := mock.Called(context, user_id, asset_id, allocationDate)

	var err error
	if args[0] != nil {
		err = args[0].(error)
	}
	return err
}
func (mock *MockAssetAllocationService) DeallocateAsset(context context.Context, user_id int, asset_id int, deallocationDate string) error {
	args := mock.Called(context, user_id, asset_id, deallocationDate)

	var err error
	if args[0] != nil {
		err = args[0].(error)
	}
	return err
}
