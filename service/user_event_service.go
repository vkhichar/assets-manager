package service

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/vkhichar/assets-manager/domain"
)

type EventService interface {
	CreateUserEvent(context.Context, *domain.User) (string, error)
	UpdateUserEvent(context.Context, *domain.User) (string, error)
}

type eventService struct {
}

func NewEventService() EventService {
	return &eventService{}
}

func (es *eventService) CreateUserEvent(ctx context.Context, user *domain.User) (string, error) {
	url := "http://localhost:9035/events"
	requestByte, _ := json.Marshal(user)

	requestReader := bytes.NewReader(requestByte)

	req, err := http.NewRequest("POST", url, requestReader)

	if err != nil {
		return "", err
	}
	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (es *eventService) UpdateUserEvent(ctx context.Context, user *domain.User) (string, error) {
	url := "http://localhost:9035/events/"
	url = url + strconv.Itoa(user.ID)

	requestByte, _ := json.Marshal(user)

	requestReader := bytes.NewReader(requestByte)

	req, err := http.NewRequest("POST", url, requestReader)

	if err != nil {
		return "", err
	}
	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
