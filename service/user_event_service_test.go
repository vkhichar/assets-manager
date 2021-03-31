package service_test

import (
	"context"
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vkhichar/assets-manager/domain"
	"github.com/vkhichar/assets-manager/service"
	"gopkg.in/h2non/gock.v1"
)

func TestCreateUserEvent_ReturnSuccess(t *testing.T) {
	defer gock.Off()
	eventService := service.NewEventService()
	ctx := context.Background()
	user := domain.User{
		ID:       1,
		Name:     "jay",
		Email:    "jay",
		Password: "jay",
		IsAdmin:  false,
	}

	gock.New("http://localhost:9035").
		Post("/events").
		Reply(200).
		JSON(user).
		BodyString(strconv.Itoa(user.ID))

	id, err := eventService.CreateUserEvent(ctx, &user)
	assert.Nil(t, err)
	assert.Equal(t, "1", id)
}

func TestCreateUserEvent_ReturnError(t *testing.T) {
	defer gock.Off()
	eventService := service.NewEventService()
	ctx := context.Background()
	user := domain.User{
		ID:       1,
		Name:     "jay",
		Email:    "jay",
		Password: "jay",
		IsAdmin:  false,
	}

	gock.New("http://localhost:9035").
		Post("/events").
		ReplyError(errors.New("user not created"))

	_, err := eventService.CreateUserEvent(ctx, &user)
	assert.NotNil(t, err)
	//assert.JSONEq(t, expErr, id)
}

func TestUpdateUserEvent_ReturnSuccess(t *testing.T) {
	defer gock.Off()
	eventService := service.NewEventService()
	ctx := context.Background()
	user := domain.User{
		ID:       1,
		Name:     "jay",
		Email:    "jay",
		Password: "jay",
		IsAdmin:  false,
	}

	gock.New("http://localhost:9035").
		Post("/events").
		Reply(200).
		JSON(user).
		BodyString(strconv.Itoa(user.ID))

	id, err := eventService.CreateUserEvent(ctx, &user)
	assert.Nil(t, err)
	assert.Equal(t, "1", id)
}

func TestUpdateUserEvent_ReturnError(t *testing.T) {
	defer gock.Off()
	eventService := service.NewEventService()
	ctx := context.Background()
	user := domain.User{
		ID:       1,
		Name:     "jay",
		Email:    "jay",
		Password: "jay",
		IsAdmin:  false,
	}

	gock.New("http://localhost:9035").
		Post("/events").
		ReplyError(errors.New("user not updated"))

	_, err := eventService.CreateUserEvent(ctx, &user)
	assert.NotNil(t, err)
}
