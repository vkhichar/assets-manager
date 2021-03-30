package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vkhichar/assets-manager/contract"
	"github.com/vkhichar/assets-manager/custom_errors"
	"github.com/vkhichar/assets-manager/service"
)

func AllocateAssetHandler(alloc_service service.AssetAllocationService) http.HandlerFunc {

	return func(rw http.ResponseWriter, r *http.Request) {

		rw.Header().Set("content-type", "application/json")

		user_id, err := strconv.Atoi(mux.Vars(r)["user_id"])

		//validate user_id
		if err != nil {
			//bad request
			responseBody, _ := json.Marshal(contract.ErrorResponse{Error: custom_errors.InvalidAssetOrUserIdError.Error()})
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(responseBody)
			return
		}

		var req contract.AllocateAssetRequest

		err = json.NewDecoder(r.Body).Decode(&req)
		//decode request
		if err != nil {
			//bad request
			responseBody, _ := json.Marshal(contract.ErrorResponse{Error: err.Error()})
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(responseBody)
			return
		}
		//validate request

		err = req.Validate()
		if err != nil {
			//invalid request
			responseBody, _ := json.Marshal(contract.ErrorResponse{Error: err.Error()})
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(responseBody)
			return
		}

		//call service layer method
		err = alloc_service.AllocateAsset(r.Context(), user_id, req.Asset_id, req.Date)

		if err == custom_errors.InvalidAssetOrUserIdError || err == custom_errors.InvalidAssetStatusError {

			responseBody, _ := json.Marshal(contract.ErrorResponse{Error: err.Error()})
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(responseBody)
			return

		}

		if err != nil {
			//internal server error
			responseBody, _ := json.Marshal(contract.ErrorResponse{Error: err.Error()})
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write(responseBody)
			return
		}

		responseBody, _ := json.Marshal("{success : asset allocated successfully}")
		rw.WriteHeader(http.StatusOK)
		rw.Write(responseBody)
	}
}

func DeallocateAssetHandler(service service.AssetAllocationService) http.HandlerFunc {

	return func(rw http.ResponseWriter, r *http.Request) {

		user_id, err := strconv.Atoi(mux.Vars(r)["user_id"])
		asset_id, err2 := strconv.Atoi(mux.Vars(r)["asset_id"])

		rw.Header().Set("content-type", "application/json")
		if err != nil || err2 != nil || asset_id < 0 || user_id < 0 {
			//bad request
			fmt.Println("error occoured while validating request for Deallocate Asset ", err)

			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: custom_errors.InvalidAssetOrUserIdError.Error()})
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(responseBytes)
			return
		}

		var req contract.DeallocateAssetRequest
		//decode request
		err = json.NewDecoder(r.Body).Decode(&req)

		if err != nil {
			//bad request
			fmt.Println("error occoured while decoding request for Deallocate Asset ", err.Error())

			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: err.Error()})
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(responseBytes)
			return
		}

		//validate request
		err = req.Validate()
		if err != nil {
			//bad request
			fmt.Println("error occoured while decoding request for Deallocate Asset ", err.Error())

			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: err.Error()})
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(responseBytes)
			return
		}

		//call service layer method
		err = service.DeallocateAsset(r.Context(), user_id, asset_id, req.Date)

		if err == custom_errors.InvalidAllocationError {
			//bad request
			fmt.Println("error occoured while Processing request for Deallocate Asset ", err.Error())

			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: err.Error()})
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(responseBytes)
			return
		}

		if err != nil {
			//Internal server error
			fmt.Println("error occoured while Processing request for Deallocate Asset ", err.Error())

			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: err.Error()})
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write(responseBytes)
			return
		}

		//success
		responseBytes, _ := json.Marshal("{success : asset deallocated successfully}")
		rw.WriteHeader(http.StatusOK)
		rw.Write(responseBytes)
		return
	}
}
