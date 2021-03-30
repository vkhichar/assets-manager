package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockAssetAllocationRepo struct {
	mock.Mock
}

func (mockRepo *MockAssetAllocationRepo) AllocateAsset(context context.Context, user_id int, asset_id int, allocationDate string) error {

	vars := mockRepo.Called(context, user_id, asset_id, allocationDate)

	var err error
	if vars[0] != nil {
		err = vars[0].(error)
	}
	return err
}

func (mockRepo *MockAssetAllocationRepo) DeallocateAsset(context context.Context, user_id int, asset_id int, deallocationDate string) error {
	vars := mockRepo.Called(context, user_id, asset_id, deallocationDate)

	var err error
	if vars[0] != nil {
		err = vars[0].(error)
	}
	return err
}
