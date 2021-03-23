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
	"github.com/vkhichar/assets-manager/contract"
	"github.com/vkhichar/assets-manager/domain"
	"github.com/vkhichar/assets-manager/handler"
	mockService "github.com/vkhichar/assets-manager/service/mocks"
)

func TestCreateAssetHandler_When_BadRequest(t *testing.T) {
	ctx := context.Background()

	obj := domain.Asset{
		InitCost: 13000,
		Category: "Mobile",
	}

	requestByte, _ := json.Marshal(obj)

	requestReader := bytes.NewReader(requestByte)

	req, err := http.NewRequest("POST", "/assets", requestReader)
	if err != nil {
		t.Fatal(err)
	}

	expectedErr := string(`{"error":"name is required"}`)

	mockAssetService := &mockService.MockAssetService{}
	mockAssetService.On("CreateAsset", ctx, &obj).Return(nil, nil)

	resp := httptest.NewRecorder()
	handler := handler.CreateAssetHandler(mockAssetService)

	handler.ServeHTTP(resp, req)

	assert.JSONEq(t, expectedErr, resp.Body.String())

}

func TestCreateAssetHandler_When_InvalidRequest(t *testing.T) {
	ctx := context.Background()

	obj := domain.Asset{
		Name:     "Mi A1",
		InitCost: 13000,
		Category: "Mobile",
	}

	requestByte, _ := json.Marshal(obj)

	requestReader := bytes.NewReader(requestByte)

	req, err := http.NewRequest("POST", "/assets", requestReader)
	if err != nil {
		t.Fatal(err)
	}

	expectedErr := string(`{"error":"something went wrong"}`)

	mockAssetService := &mockService.MockAssetService{}
	mockAssetService.On("CreateAsset", ctx, &obj).Return(nil, errors.New("something went wrong"))

	resp := httptest.NewRecorder()
	handler := handler.CreateAssetHandler(mockAssetService)

	handler.ServeHTTP(resp, req)

	assert.JSONEq(t, expectedErr, resp.Body.String())

}

func TestCreateAssetHandler_When_Success(t *testing.T) {
	ctx := context.Background()

	obj := domain.Asset{
		Name:     "Mi A1",
		InitCost: 13000,
		Category: "Mobile",
	}

	requestByte, _ := json.Marshal(obj)

	requestReader := bytes.NewReader(requestByte)

	req, err := http.NewRequest("POST", "/assets", requestReader)
	if err != nil {
		t.Fatal(err)
	}

	mockAssetService := &mockService.MockAssetService{}
	mockAssetService.On("CreateAsset", ctx, &obj).Return(&obj, nil)

	resp := httptest.NewRecorder()
	handler := handler.CreateAssetHandler(mockAssetService)

	handler.ServeHTTP(resp, req)

	expectedJson := string(`{"id":0,"name":"Mi A1","category":"Mobile","specification":null,"initCost":13000,"status":0}`)
	assert.JSONEq(t, resp.Body.String(), expectedJson)

}

func TestFindAssetHandler_When_InvalidId(t *testing.T) {
	ctx := context.Background()

	id := 0

	req, err := http.NewRequest("GET", "/assets/0", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp := httptest.NewRecorder()

	expectedErr := string(`{"error":"invalid id"}`)

	mockAssetService := mockService.MockAssetService{}
	mockAssetService.On("FindAsset", ctx, id).Return(nil, errors.New("invalid id"))
	handler := handler.FindAssetHandler(&mockAssetService)

	handler.ServeHTTP(resp, req)

	assert.JSONEq(t, expectedErr, resp.Body.String())

}

func TestFindAssetHandler_When_Success(t *testing.T) {
	ctx := context.Background()

	id := 1

	req, err := http.NewRequest("GET", "/assets/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp := httptest.NewRecorder()

	obj := contract.CreateAssetResponse{
		ID:            0,
		Name:          "Mi A1",
		InitCost:      13000,
		Category:      "Mobile",
		Status:        0,
		Specification: nil,
	}

	expectedAsset, _ := json.Marshal(contract.CreateAssetResponse{
		ID:            0,
		Name:          "Mi A1",
		InitCost:      13000,
		Category:      "Mobile",
		Status:        0,
		Specification: nil,
	})

	mockAssetService := mockService.MockAssetService{}
	mockAssetService.On("FindAsset", ctx, id).Return(&obj, nil)
	handler := handler.FindAssetHandler(&mockAssetService)

	handler.ServeHTTP(resp, req)

	assert.JSONEq(t, string(expectedAsset), resp.Body.String())

}
