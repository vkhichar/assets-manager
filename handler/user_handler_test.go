package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vkhichar/assets-manager/domain"
	"github.com/vkhichar/assets-manager/handler"
	mock "github.com/vkhichar/assets-manager/service/mocks"
)

func TestGetUser_ReturnsError(t *testing.T) {
	ctx := context.Background()
	id := 1

	mockUserService := &mock.MockUserService{}
	expectedErr := string(`{"error":"something went wrong"}`)
	mockUserService.On("GetUser", ctx, id).Return(nil, errors.New("something went wrong"))

	req, err := http.NewRequest("GET", "/users?id=1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := handler.GetUser(mockUserService)

	handler.ServeHTTP(rr, req)

	//assert.Error(t, err)
	assert.Equal(t, expectedErr, rr.Body.String())
}

func TestGetUser_ReturnsNoRowErr(t *testing.T) {
	ctx := context.Background()
	id := 1

	mockUserService := &mock.MockUserService{}
	expectedErr := string(`{"error":"No data present for this id"}`)
	mockUserService.On("GetUser", ctx, id).Return(nil, errors.New("no value for this id"))

	req, err := http.NewRequest("GET", "/users?id=1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := handler.GetUser(mockUserService)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, expectedErr, rr.Body.String())
}

func TestGetUser_ReturnsNoId(t *testing.T) {
	ctx := context.Background()
	id := 1

	mockUserService := &mock.MockUserService{}
	expectedErr := string(`{"error":"No user id is provided"}`)
	mockUserService.On("GetUser", ctx, id).Return(nil, errors.New("something went wrong"))

	req, err := http.NewRequest("GET", "/users?id=", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := handler.GetUser(mockUserService)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, expectedErr, rr.Body.String())
}

func TestCreateUser_ReturnsError(t *testing.T) {
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

	requestByte, _ := json.Marshal(user)

	requestReader := bytes.NewReader(requestByte)

	mockUserService := &mock.MockUserService{}

	expectedErr := string(`{"error":"something went wrong"}`)

	mockUserService.On("Register", ctx, name, email, password, isAdmin).Return(nil, errors.New("something went wrong"))

	req, err := http.NewRequest("POST", "/users", requestReader)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := handler.CreateUserHandler(mockUserService)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, expectedErr, rr.Body.String())
}

func TestCreateUser_ReturnsDuplicateErr(t *testing.T) {
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

	requestByte, _ := json.Marshal(user)

	requestReader := bytes.NewReader(requestByte)

	mockUserService := &mock.MockUserService{}

	expectedErr := string(`{"error":"email already exist"}`)

	mockUserService.On("Register", ctx, name, email, password, isAdmin).Return(nil, errors.New("this email is already registered"))

	req, err := http.NewRequest("POST", "/users", requestReader)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := handler.CreateUserHandler(mockUserService)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, expectedErr, rr.Body.String())
}

func TestCreateUser_ReturnsSuccess(t *testing.T) {
	ctx := context.Background()
	name := "krish"
	email := "krish@gmail.com"
	password := "krish"
	isAdmin := false

	user := domain.User{
		ID:       2,
		Name:     "krish",
		Email:    "krish@gmail.com",
		Password: "krish",
		IsAdmin:  false,
	}

	requestByte, _ := json.Marshal(user)

	requestReader := bytes.NewReader(requestByte)

	mockUserService := &mock.MockUserService{}

	mockUserService.On("Register", ctx, name, email, password, isAdmin).Return(&user, nil)

	req, err := http.NewRequest("POST", "/users", requestReader)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := handler.CreateUserHandler(mockUserService)

	handler.ServeHTTP(rr, req)

	assert.Nil(t, err)
	assert.NotEmpty(t, rr.Body.String())
}
