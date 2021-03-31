package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/vkhichar/assets-manager/contract"
	"github.com/vkhichar/assets-manager/domain"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Login(ctx context.Context, email, password string) (*domain.User, string, error) {
	return nil, "", nil
}

func (m *MockUserService) Register(ctx context.Context, name, email, password string, isAdmin bool) (*domain.User, error) {
	args := m.Called(ctx, name, email, password, isAdmin)

	var user *domain.User
	if args[0] != nil {
		user = args[0].(*domain.User)
	}

	var err error
	if args[1] != nil {
		err = args[1].(error)
	}

	return user, err
}

func (m *MockUserService) GetUser(ctx context.Context, id int) (*domain.User, error) {
	args := m.Called(ctx, id)

	var user *domain.User
	if args[0] != nil {
		user = args[0].(*domain.User)
	}

	var err error
	if args[1] != nil {
		err = args[1].(error)
	}

	return user, err
}

func (m *MockUserService) Update(ctx context.Context, id int, val contract.UpdateUserRequest) (*domain.User, error) {
	return nil, nil
}
