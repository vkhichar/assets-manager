package service

import (
	"context"
	"errors"

	"github.com/vkhichar/assets-manager/domain"
	"github.com/vkhichar/assets-manager/repository"
)

var ErrInvalidEmailPassword = errors.New("invalid email or password")
var ErrDuplicateEmail = errors.New("this email is already registered")
var ErrNoSqlRow = errors.New("no value for this id")

type UserService interface {
	Login(ctx context.Context, email, password string) (user *domain.User, token string, err error)
	Register(ctx context.Context, name, email, password string, isAdmin bool) (user *domain.User, err error)
	GetUser(ctx context.Context, id int) (*domain.User, error)
}

type userService struct {
	userRepo repository.UserRepository
	tokenSvc TokenService
}

func NewUserService(repo repository.UserRepository, ts TokenService) UserService {
	return &userService{
		userRepo: repo,
		tokenSvc: ts,
	}
}

func (service *userService) Login(ctx context.Context, email, password string) (*domain.User, string, error) {
	user, err := service.userRepo.FindUser(ctx, email)
	if err != nil {
		return nil, "", err
	}

	if user == nil {
		return nil, "", ErrInvalidEmailPassword
	}

	if user.Password != password {
		return nil, "", ErrInvalidEmailPassword
	}

	claims := &Claims{UserID: user.ID, IsAdmin: user.IsAdmin}
	token, err := service.tokenSvc.GenerateToken(claims)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (service *userService) Register(ctx context.Context, name, email, password string, isAdmin bool) (*domain.User, error) {
	user, err := service.userRepo.InsertUser(ctx, name, email, password, isAdmin)

	if err != nil && err.Error() == ErrDuplicateEmail.Error() {
		return nil, ErrDuplicateEmail
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (service *userService) GetUser(ctx context.Context, id int) (*domain.User, error) {
	user, err := service.userRepo.GetUser(ctx, id)

	if err != nil && err.Error() == ErrNoSqlRow.Error() {
		return nil, ErrNoSqlRow
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}
