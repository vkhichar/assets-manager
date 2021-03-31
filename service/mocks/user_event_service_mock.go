package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/vkhichar/assets-manager/domain"
)

type MockUserEventService struct {
	mock.Mock
}

func (m *MockUserEventService) CreateUserEvent(ctx context.Context, user *domain.User) (string, error) {

	args := m.Called(ctx, user)

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

func (m *MockUserEventService) UpdateUserEvent(ctx context.Context, user *domain.User) (string, error) {
	args := m.Called(ctx, user)

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
