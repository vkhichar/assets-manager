package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/vkhichar/assets-manager/contract"
	"github.com/vkhichar/assets-manager/custom_errors"
	"github.com/vkhichar/assets-manager/handler"
	mockService "github.com/vkhichar/assets-manager/service/mocks"
)

func TestAllocateAsset_When_BadRequest(t *testing.T) {

	reqBody, _ := json.Marshal(contract.AllocateAssetRequest{
		Asset_id: -1,
		Date:     "2020-12-12",
	})

	req, err := http.NewRequest("POST", "users/{user_id}/allocate-asset", bytes.NewReader(reqBody))

	if err != nil {

		t.Fatal()
	}
	vars := map[string]string{
		"user_id": "11",
	}
	req = mux.SetURLVars(req, vars)

	mockService := &mockService.MockAssetAllocationService{}
	mockService.On("AllocateAsset", context.Background(), -1, -1, "2020-12-12").Return(nil)

	respRec := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.AllocateAssetHandler(mockService))
	handler.ServeHTTP(respRec, req)

	expectedErr, _ := json.Marshal(contract.ErrorResponse{Error: custom_errors.InvalidAssetOrUserIdError.Error()})

	assert.Equal(t, http.StatusBadRequest, respRec.Code)
	assert.Equal(t, string(expectedErr), respRec.Body.String())

}

func TestAllocateAsset_When_BadRequest_InvalidDateFormat(t *testing.T) {

	reqBody, _ := json.Marshal(contract.AllocateAssetRequest{
		Asset_id: 1,
		Date:     "2020/12/12",
	})

	req, err := http.NewRequest("POST", "users/{user_id}/allocate-asset", bytes.NewReader(reqBody))

	if err != nil {

		t.Fatal()
	}
	vars := map[string]string{
		"user_id": "1",
	}
	req = mux.SetURLVars(req, vars)
	mockService := &mockService.MockAssetAllocationService{}
	mockService.On("AllocateAsset", context.Background(), 1, 1, "2020/12/12").Return(nil)

	respRec := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.AllocateAssetHandler(mockService))
	handler.ServeHTTP(respRec, req)

	expectedErr, _ := json.Marshal(contract.ErrorResponse{Error: custom_errors.InvalidDateFormatError.Error()})

	assert.Equal(t, http.StatusBadRequest, respRec.Code)
	assert.Equal(t, string(expectedErr), respRec.Body.String())

}

func TestAllocateAsset_When_InternalServerError(t *testing.T) {

	reqBody, _ := json.Marshal(contract.AllocateAssetRequest{
		Asset_id: 1,
		Date:     "2020-12-12",
	})

	req, err := http.NewRequest("POST", "users/{user_id}/allocate-asset", bytes.NewReader(reqBody))

	if err != nil {

		t.Fatal()
	}
	vars := map[string]string{
		"user_id": "1",
	}
	req = mux.SetURLVars(req, vars)
	mockService := &mockService.MockAssetAllocationService{}
	mockService.On("AllocateAsset", req.Context(), 1, 1, "2020-12-12").Return(errors.New("something went wrong"))

	respRec := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.AllocateAssetHandler(mockService))
	handler.ServeHTTP(respRec, req)

	expectedErr, _ := json.Marshal(contract.ErrorResponse{Error: "something went wrong"})

	assert.Equal(t, http.StatusInternalServerError, respRec.Code)
	assert.Equal(t, string(expectedErr), respRec.Body.String())

}

func TestAllocateAsset_When_Success(t *testing.T) {

	reqBody, _ := json.Marshal(contract.AllocateAssetRequest{
		Asset_id: 1,
		Date:     "2020-12-12",
	})

	req, err := http.NewRequest("POST", "users/{user_id}/allocate-asset", bytes.NewReader(reqBody))

	if err != nil {

		t.Fatal()
	}
	vars := map[string]string{
		"user_id": "1",
	}
	req = mux.SetURLVars(req, vars)
	mockService := &mockService.MockAssetAllocationService{}
	mockService.On("AllocateAsset", req.Context(), 1, 1, "2020-12-12").Return(nil)

	respRec := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.AllocateAssetHandler(mockService))
	handler.ServeHTTP(respRec, req)

	expectedResp, _ := json.Marshal("{success : asset allocated successfully}")

	assert.Equal(t, http.StatusOK, respRec.Code)
	assert.Equal(t, string(expectedResp), respRec.Body.String())

}
func TestDeallocateAsset_When_BadRequest(t *testing.T) {

	reqBody, _ := json.Marshal(contract.DeallocateAssetRequest{
		Date: "2020-12-12",
	})

	req, err := http.NewRequest("DELETE", "users/{user_id}/assets/{asset_id}", bytes.NewReader(reqBody))

	if err != nil {

		t.Fatal()
	}
	vars := map[string]string{
		"user_id":  "-1",
		"asset_id": "-1",
	}
	req = mux.SetURLVars(req, vars)

	mockService := &mockService.MockAssetAllocationService{}
	mockService.On("DeallocateAsset", req.Context(), -1, -1, "2020-12-12").Return(nil)

	respRec := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.DeallocateAssetHandler(mockService))
	handler.ServeHTTP(respRec, req)

	expectedErr, _ := json.Marshal(contract.ErrorResponse{Error: custom_errors.InvalidAssetOrUserIdError.Error()})

	assert.Equal(t, http.StatusBadRequest, respRec.Code)
	assert.Equal(t, string(expectedErr), respRec.Body.String())

}

func TestDeallocateAsset_When_BadRequest_InvalidDateFormat(t *testing.T) {

	reqBody, _ := json.Marshal(contract.DeallocateAssetRequest{
		Date: "2020/12/12",
	})

	req, err := http.NewRequest("DELETE", "users/{user_id}/assets/{asset_id}", bytes.NewReader(reqBody))

	if err != nil {

		t.Fatal()
	}
	vars := map[string]string{
		"user_id":  "1",
		"asset_id": "1",
	}
	req = mux.SetURLVars(req, vars)

	mockService := &mockService.MockAssetAllocationService{}
	mockService.On("DeallocateAsset", req.Context(), -1, -1, "2020/12/12").Return(nil)

	respRec := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.DeallocateAssetHandler(mockService))
	handler.ServeHTTP(respRec, req)

	expectedErr, _ := json.Marshal(contract.ErrorResponse{Error: custom_errors.InvalidDateFormatError.Error()})

	assert.Equal(t, http.StatusBadRequest, respRec.Code)
	assert.Equal(t, string(expectedErr), respRec.Body.String())

}

func TestDeallocateAsset_When_InternalServerError(t *testing.T) {

	reqBody, _ := json.Marshal(contract.DeallocateAssetRequest{
		Date: "2020-12-12",
	})

	req, err := http.NewRequest("DELETE", "users/{user_id}/assets/{asset_id}", bytes.NewReader(reqBody))

	if err != nil {

		t.Fatal()
	}
	vars := map[string]string{
		"user_id":  "1",
		"asset_id": "1",
	}
	req = mux.SetURLVars(req, vars)

	mockService := &mockService.MockAssetAllocationService{}
	mockService.On("DeallocateAsset", req.Context(), 1, 1, "2020-12-12").Return(errors.New("something went wrong"))

	respRec := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.DeallocateAssetHandler(mockService))
	handler.ServeHTTP(respRec, req)

	expectedErr, _ := json.Marshal(contract.ErrorResponse{Error: "something went wrong"})

	assert.Equal(t, http.StatusInternalServerError, respRec.Code)
	assert.Equal(t, string(expectedErr), respRec.Body.String())

}

func TestDeallocateAsset_When_Success(t *testing.T) {

	reqBody, _ := json.Marshal(contract.DeallocateAssetRequest{
		Date: "2020-12-12",
	})

	req, err := http.NewRequest("DELETE", "users/{user_id}/assets/{asset_id}", bytes.NewReader(reqBody))

	if err != nil {

		t.Fatal()
	}
	vars := map[string]string{
		"user_id":  "1",
		"asset_id": "1",
	}
	req = mux.SetURLVars(req, vars)

	mockService := &mockService.MockAssetAllocationService{}
	mockService.On("DeallocateAsset", req.Context(), 1, 1, "2020-12-12").Return(nil)

	respRec := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.DeallocateAssetHandler(mockService))
	handler.ServeHTTP(respRec, req)

	expectedErr, _ := json.Marshal("{success : asset deallocated successfully}")

	assert.Equal(t, http.StatusOK, respRec.Code)
	assert.Equal(t, string(expectedErr), respRec.Body.String())

}
