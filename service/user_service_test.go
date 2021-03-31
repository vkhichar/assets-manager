package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vkhichar/assets-manager/contract"
	"github.com/vkhichar/assets-manager/domain"
	mock "github.com/vkhichar/assets-manager/repository/mocks"
	"github.com/vkhichar/assets-manager/service"
	serviceMock "github.com/vkhichar/assets-manager/service/mocks"
)

func TestGetUser_ReturnErrNoSqlRow(t *testing.T) {
	ctx := context.Background()
	id := 1

	mockUserRepo := &mock.MockUserRepo{}

	mockUserRepo.On("GetUser", ctx, id).Return(nil, errors.New("no value for this id"))

	userService := service.NewUserService(mockUserRepo, nil, nil)

	user, err := userService.GetUser(ctx, id)

	assert.Error(t, err)
	assert.Equal(t, "no value for this id", err.Error())
	assert.Nil(t, user)
}

func TestGetUser_ReturnError(t *testing.T) {
	ctx := context.Background()
	id := 1

	mockUserRepo := &mock.MockUserRepo{}

	mockUserRepo.On("GetUser", ctx, id).Return(nil, errors.New("some db error"))

	userService := service.NewUserService(mockUserRepo, nil, nil)

	user, err := userService.GetUser(ctx, id)

	assert.Error(t, err)
	assert.Equal(t, "some db error", err.Error())
	assert.Nil(t, user)
}

func TestGetUser_ReturnSuccess(t *testing.T) {
	ctx := context.Background()
	id := 1
	user := domain.User{
		ID:       1,
		Name:     "Roy",
		Email:    "roy@gmail.com",
		Password: "roy",
		IsAdmin:  true,
	}

	mockUserRepo := &mock.MockUserRepo{}
	mockUserEventService := &serviceMock.MockUserEventService{}

	mockUserRepo.On("GetUser", ctx, id).Return(&user, nil)

	userService := service.NewUserService(mockUserRepo, nil, mockUserEventService)

	dbuser, err := userService.GetUser(ctx, id)

	assert.Nil(t, err)
	assert.Equal(t, &user, dbuser)
}

func TestInsertUser_ReturnErrDuplicateEmail(t *testing.T) {
	ctx := context.Background()
	name := "shiva"
	email := "shiva@gmail.com"
	password := "shiva"
	isAdmin := false

	mockUserRepo := &mock.MockUserRepo{}

	mockUserRepo.On("InsertUser", ctx, name, email, password, isAdmin).Return(nil, errors.New("this email is already registered"))
	userService := service.NewUserService(mockUserRepo, nil, nil)

	dbuser, err := userService.Register(ctx, name, email, password, isAdmin)

	assert.Error(t, err)
	assert.Equal(t, "this email is already registered", err.Error())
	assert.Nil(t, dbuser)
}

func TestInsertUser_ReturnError(t *testing.T) {
	ctx := context.Background()
	name := "shiva"
	email := "shiva@gmail.com"
	password := "shiva"
	isAdmin := false

	mockUserRepo := &mock.MockUserRepo{}

	mockUserRepo.On("InsertUser", ctx, name, email, password, isAdmin).Return(nil, errors.New("some db error"))
	userService := service.NewUserService(mockUserRepo, nil, nil)

	dbuser, err := userService.Register(ctx, name, email, password, isAdmin)

	assert.Error(t, err)
	assert.Equal(t, "some db error", err.Error())
	assert.Nil(t, dbuser)
}

func TestInsertUser_ReturnSuccess(t *testing.T) {
	ctx := context.Background()
	name := "shiva"
	email := "shiva@gmail.com"
	password := "shiva"
	isAdmin := false

	user := domain.User{
		ID:       2,
		Name:     "shiva",
		Email:    "shiva@gmail.com",
		Password: "shiva",
		IsAdmin:  false,
	}

	mockUserRepo := &mock.MockUserRepo{}
	mockUserEventService := &serviceMock.MockUserEventService{}

	mockUserRepo.On("InsertUser", ctx, name, email, password, isAdmin).Return(&user, nil)
	mockUserEventService.On("CreateUserEvent", ctx, &user).Return("", nil)
	userService := service.NewUserService(mockUserRepo, nil, mockUserEventService)

	dbuser, err := userService.Register(ctx, name, email, password, isAdmin)

	assert.Nil(t, err)
	assert.Equal(t, &user, dbuser)
}

func TestUpdateUser_ReturnError(t *testing.T) {
	ctx := context.Background()
	updateUser := contract.UpdateUserRequest{
		Name:     "shiva",
		Email:    "shiva@gmail.com",
		Password: "shiva",
		IsAdmin:  false,
	}
	user := domain.User{
		ID:       2,
		Name:     "shiva",
		Email:    "shiva@gmail.com",
		Password: "shiva",
		IsAdmin:  false,
	}

	mockUserRepo := &mock.MockUserRepo{}
	mockUserEventService := &serviceMock.MockUserEventService{}

	mockUserEventService.On("UpdateUserEvent", ctx, &user).Return("", nil)
	mockUserRepo.On("UpdateUser", ctx, user.ID, updateUser).Return(nil, errors.New("some db error"))
	userService := service.NewUserService(mockUserRepo, nil, mockUserEventService)

	dbuser, err := userService.Update(ctx, user.ID, updateUser)

	assert.Error(t, err)
	assert.Equal(t, "some db error", err.Error())
	assert.Nil(t, dbuser)
}
func TestUpdateUser_ReturnSuccess(t *testing.T) {
	ctx := context.Background()

	user := domain.User{
		ID:       2,
		Name:     "shiva",
		Email:    "shiva@gmail.com",
		Password: "shiva",
		IsAdmin:  false,
	}

	updateUser := contract.UpdateUserRequest{
		Name:     "shiva",
		Email:    "shiva@gmail.com",
		Password: "shiva",
		IsAdmin:  false,
	}

	mockUserRepo := &mock.MockUserRepo{}
	mockUserEventService := &serviceMock.MockUserEventService{}

	mockUserRepo.On("UpdateUser", ctx, user.ID, updateUser).Return(&user, nil)
	mockUserEventService.On("UpdateUserEvent", ctx, &user).Return("", nil)
	userService := service.NewUserService(mockUserRepo, nil, mockUserEventService)

	dbuser, err := userService.Update(ctx, user.ID, updateUser)

	assert.Nil(t, err)
	assert.Equal(t, &user, dbuser)
}
