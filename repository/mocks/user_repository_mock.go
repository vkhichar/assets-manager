package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/vkhichar/assets-manager/contract"
	"github.com/vkhichar/assets-manager/domain"
)

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) FindUser(ctx context.Context, email string) (*domain.User, error) {
	return nil, nil
}

func (m *MockUserRepo) InsertUser(ctx context.Context, name, email, password string, isAdmin bool) (*domain.User, error) {
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

func (m *MockUserRepo) GetUser(ctx context.Context, id int) (*domain.User, error) {
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

func (m *MockUserRepo) UpdateUser(ctx context.Context, id int, val contract.UpdateUserRequest) (*domain.User, error) {
	return nil, nil
}
