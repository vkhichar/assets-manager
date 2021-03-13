package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/vkhichar/assets-manager/domain"
)

var ErrDuplicateEmail = errors.New("this email is already registered")

const (
	getUserByEmailQuery = "SELECT id, name, email, password, is_admin FROM users WHERE email= $1"
	registerUser        = "INSERT INTO users (name,email,password, is_admin) values ($1, $2, $3, $4)"
	getUserByID         = "SELECT id, name, email, password, is_admin from users where id = $1"
)

type UserRepository interface {
	FindUser(ctx context.Context, email string) (*domain.User, error)
	InsertUser(ctx context.Context, name, email, password string, isAdmin bool) (*domain.User, error)
	GetUser(ctx context.Context, id int) (*domain.User, error)
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepository() UserRepository {
	return &userRepo{
		db: GetDB(),
	}
}

func (repo *userRepo) FindUser(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	err := repo.db.Get(&user, getUserByEmailQuery, email)
	if err == sql.ErrNoRows {
		fmt.Printf("repository: couldn't find user for email: %s", email)

		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *userRepo) InsertUser(ctx context.Context, name, email, password string, isAdmin bool) (*domain.User, error) {
	var user domain.User

	err := repo.db.Get(&user, getUserByEmailQuery, email)

	if err == nil {
		return nil, ErrDuplicateEmail
	}
	_, err = repo.db.Exec(registerUser, name, email, password, isAdmin)

	if err != nil {
		return nil, err
	}

	repo.db.Get(&user, getUserByEmailQuery, email)

	return &user, nil
}

func (repo *userRepo) GetUser(ctx context.Context, id int) (*domain.User, error) {
	var user domain.User

	err := repo.db.Get(&user, getUserByID, id)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
