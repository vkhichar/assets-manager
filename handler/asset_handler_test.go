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

func TestUpdateAsset_When_BadRequest(t *testing.T) {

	context := context.Background()

	asset := &domain.Asset{
		Id: -100,
	}
	requestBytes, _ := json.Marshal(asset)
	req, err := http.NewRequest("PUT", "assets/update", bytes.NewReader(requestBytes))

	if err != nil {
		t.Fatal(err)
	}

	mockAssetService := &mockService.MockAssetService{}
	mockAssetService.On("UpdateAsset", context, asset).Return(nil, errors.New("invalid id"))
	expected_error := string(`{"error":"invalid id"}`)

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.UpdateAssets(mockAssetService))
	handler.ServeHTTP(resp, req)

	assert.Equal(t, expected_error, resp.Body.String())
}

func TestUpdateAsset_When_InternalServerError(t *testing.T) {

	context := context.Background()

	asset := &contract.UpadateAssetRequest{
		Id:       1,
		Name:     "hp",
		Category: "laptop",
		InitCost: 42134,
		Status:   0,
	}

	requestBytes, _ := json.Marshal(asset)
	req, _ := http.NewRequest("PUT", "assets/update", bytes.NewReader(requestBytes))

	mockAssetService := &mockService.MockAssetService{}
	mockAssetService.On("UpdateAsset", context, asset).Return(nil, errors.New("something went wrong"))
	expected_error := string(`{"error":"something went wrong"}`)

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.UpdateAssets(mockAssetService))
	handler.ServeHTTP(resp, req)

	assert.Equal(t, expected_error, resp.Body.String())

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

func TestUpdateAsset_When_Success(t *testing.T) {

	context := context.Background()

	asset_request := &contract.UpadateAssetRequest{
		Id:       1,
		Name:     "hp",
		Category: "laptop",
		InitCost: 1000,
		Status:   0,
	}
	asset_expected := domain.Asset{
		Id:       1,
		Name:     "hp",
		Category: "laptop",
		InitCost: 1000,
		Status:   0,
	}
	requestBytes, _ := json.Marshal(asset_request)
	req, err := http.NewRequest("PUT", "assets/update", bytes.NewReader(requestBytes))

	if err != nil {
		t.Fatal(err)
	}

	mockAssetService := &mockService.MockAssetService{}
	mockAssetService.On("UpdateAsset", context, asset_request).Return(&asset_expected, nil)
	expectedResp, _ := json.Marshal(asset_expected)

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.UpdateAssets(mockAssetService))
	handler.ServeHTTP(resp, req)

	assert.Equal(t, bytes.NewBuffer(expectedResp).String(), resp.Body.String())
}

func TestDeleteAssets_When_BadRequest(t *testing.T) {

	context := context.Background()

	req, _ := http.NewRequest("DELETE", "assets/delete?id=asdf", nil)

	mockAssetService := &mockService.MockAssetService{}
	mockAssetService.On("DeleteAsset", context, 1).Return(nil, errors.New("invalid id"))
	expected_error := string(`{"error":"invalid id"}`)

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.DeleteAssets(mockAssetService))
	handler.ServeHTTP(resp, req)

	assert.Equal(t, expected_error, resp.Body.String())

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
func TestDeleteAssets_When_InternalServerError(t *testing.T) {

	context := context.Background()

	req, err := http.NewRequest("DELETE", "assets/delete?id=1", nil)

	if err != nil {
		t.Fatal(err)
	}

	mockAssetService := &mockService.MockAssetService{}
	mockAssetService.On("DeleteAsset", context, 1).Return(nil, errors.New("something went wrong"))
	expected_error := string(`{"error":"something went wrong"}`)

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.DeleteAssets(mockAssetService))
	handler.ServeHTTP(resp, req)

	assert.Equal(t, expected_error, resp.Body.String())
}

func TestDeleteAssets_When_Success(t *testing.T) {

	context := context.Background()

	asset_expected := &domain.Asset{
		Id:       1,
		Name:     "hp",
		Category: "laptop",
		InitCost: 1000,
		Status:   0,
	}
	req, err := http.NewRequest("DELETE", "asset/delete?id=1", nil)

	if err != nil {
		t.Fatal(err)
	}

	mockAssetService := &mockService.MockAssetService{}
	mockAssetService.On("DeleteAsset", context, 1).Return(asset_expected, nil)
	expectedResp, _ := json.Marshal(asset_expected)

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.DeleteAssets(mockAssetService))
	handler.ServeHTTP(resp, req)

	assert.Equal(t, bytes.NewBuffer(expectedResp).String(), resp.Body.String())
}

func TestGetAllAssets_When_InternalServerError(t *testing.T) {

	req, err := http.NewRequest("GET", "assets/all", nil)

	if err != nil {
		t.Fatal()
	}

	mockAssetService := &mockService.MockAssetService{}
	mockAssetService.On("GetAssets").Return(nil, errors.New("something went wrong"))

	expected_error := string(`{"error":"something went wrong"}`)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.GetAllAssets(mockAssetService))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, expected_error, rr.Body.String())
}

// func TestGetAllAssets_When_Success(t *testing.T) {

// 	list_assets := make([]domain.Asset, 3)

// 	for i := 0; i < 3; i++ {
// 		list_assets[i] = domain.Asset{Id: i, Name: fmt.Sprintf("test_user%d", i), Category: "testing", InitCost: 0, Status: 0}
// 	}

// 	req, err := http.NewRequest("GET", "assets/all", nil)

// 	if err != nil {
// 		t.Fatal()
// 	}

// 	mockAssetService := &mockService.MockAssetService{}
// 	mockAssetService.On("GetAssets").Return(list_assets, nil)
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(handler.GetAllAssets(mockAssetService))
// 	handler.ServeHTTP(rr, req)
// 	m := make(map[string]interface{})
// 	m["Assets"] = list_assets

// 	expected_list, _ := json.Marshal(m)
// 	assert.JSONEq(t, string(expected_list), rr.Body.String())

// }

// func TestFindAssetHandler_When_InvalidId(t *testing.T) {
// 	ctx := context.Background()

// 	id := 0

// 	req, err := http.NewRequest("GET", "/assets/0", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	resp := httptest.NewRecorder()

// 	expectedErr := string(`{"error":"invalid id"}`)

// 	mockAssetService := mockService.MockAssetService{}
// 	mockAssetService.On("FindAsset", ctx, id).Return(nil, errors.New("invalid id"))
// 	handler := handler.FindAssetHandler(&mockAssetService)

// 	handler.ServeHTTP(resp, req)

// 	assert.Equal(t, expectedErr, resp.Body.String())

// }

// func TestFindAssetHandler_When_Success(t *testing.T) {
// 	ctx := context.Background()

// 	id := 1

// 	req, err := http.NewRequest("GET", "/assets/{id}", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	resp := httptest.NewRecorder()

// 	obj := contract.CreateAssetResponse{
// 		ID:            0,
// 		Name:          "Mi A1",
// 		InitCost:      13000,
// 		Category:      "Mobile",
// 		Status:        0,
// 		Specification: nil,
// 	}

// 	expectedAsset, _ := json.Marshal(contract.CreateAssetResponse{
// 		ID:            0,
// 		Name:          "Mi A1",
// 		InitCost:      13000,
// 		Category:      "Mobile",
// 		Status:        0,
// 		Specification: nil,
// 	})

// 	mockAssetService := mockService.MockAssetService{}
// 	mockAssetService.On("FindAsset", ctx, id).Return(&obj, nil)
// 	handler := handler.FindAssetHandler(&mockAssetService)

// 	handler.ServeHTTP(resp, req)

// 	assert.JSONEq(t, string(expectedAsset), resp.Body.String())

// }
