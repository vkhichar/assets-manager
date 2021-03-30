package repository_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vkhichar/assets-manager/config"
	"github.com/vkhichar/assets-manager/domain"
	"github.com/vkhichar/assets-manager/repository"
)

func TestDBConnection(t *testing.T) {

	os.Setenv("APP_PORT", "9000")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USERNAME", "postgres")
	os.Setenv("DB_PASSWORD", "1234")
	os.Setenv("DB_NAME", "Asset_Manager")

	err := config.Init()
	repository.InitDB()

	assert.NoError(t, err)
}

func TestGetUser_ReturnsError(t *testing.T) {

	userRepo := repository.NewUserRepository()

	ctx := context.Background()
	id := 3
	dbuser, err := userRepo.GetUser(ctx, id)

	assert.Error(t, err)
	assert.Nil(t, dbuser)
}

func TestGetUser_ReturnsSuccess(t *testing.T) {
	userRepo := repository.NewUserRepository()

	ctx := context.Background()
	user := domain.User{
		ID:       1,
		Name:     "Raj",
		Email:    "Raj@gmail.com",
		Password: "raj",
		IsAdmin:  false,
	}
	dbuser, err := userRepo.GetUser(ctx, 1)

	assert.Nil(t, err)
	assert.Equal(t, dbuser, &user)
}

func TestInsertUser_ReturnsSuccess(t *testing.T) {
	//configEnvVars()
	config.Init()
	repository.InitDB()
	db := repository.GetDB()

	tx := db.MustBegin()
	tx.MustExec("DELETE FROM users;")
	tx.Commit()

	userRepo := repository.NewUserRepository()

	ctx := context.Background()

	user := domain.User{
		ID:       1,
		Name:     "sham",
		Email:    "sham123@gmail.com",
		Password: "sham",
		IsAdmin:  false,
	}
	dbuser, err := userRepo.InsertUser(ctx, "sham", "sham123@gmail.com", "sham", false)
	user.ID = dbuser.ID
	assert.Nil(t, err)
	assert.Equal(t, dbuser, &user)
}
func TestInsertUser_ReturnsError(t *testing.T) {

	userRepo := repository.NewUserRepository()

	ctx := context.Background()
	dbuser, err := userRepo.InsertUser(ctx, "sham", "sham12@gmail.com", "sham", false)

	assert.Error(t, err)
	assert.Nil(t, dbuser)
}
