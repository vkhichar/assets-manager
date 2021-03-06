package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vkhichar/assets-manager/contract"
	"github.com/vkhichar/assets-manager/custom_errors"
	"github.com/vkhichar/assets-manager/domain"
	"github.com/vkhichar/assets-manager/service"
)

func FindAssetHandler(service service.AssetService) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		id := mux.Vars(r)["id"]
		rw.Header().Set("content-type", "application/json")
		//decode request
		fmt.Println("id is ", id)
		Id, err := strconv.Atoi(id)

		if err != nil {

			fmt.Println("error while decoding request for Find asset ", err.Error())
			//bad request from client
			responseBody, _ := json.Marshal(err.Error())
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(responseBody)
			return
		}
		//validate
		if Id < 0 {
			fmt.Println("error while validating request find asset", err.Error())
			//bad request from client
			responseBody, _ := json.Marshal(err.Error())
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(responseBody)
			return
		}

		//call service layer method
		asset, err := service.FindAsset(r.Context(), Id)
		if err != nil {
			fmt.Println("error while processing request for find asset", err.Error())
			//internal server error
			responseBytes, _ := json.Marshal(err.Error())
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write(responseBytes)
			return
		}
		//success
		responseBytes, err := json.Marshal(contract.FindAssetResponse{Id: asset.Id, Name: asset.Name, Category: asset.Category, Specification: asset.Specification, InitCost: asset.InitCost, Status: asset.Status})
		rw.WriteHeader(http.StatusOK)
		rw.Write(responseBytes)
	}
}

func GetAllAssets(service service.AssetService) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		rw.Header().Set("content-type", "application/json")
		//call service layer method
		assets, err := service.GetAssets()

		if err != nil {
			fmt.Print("error while processing request for get all assets", err.Error())

			//internal server error
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: err.Error()})
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write(responseBytes)
			return
		}
		//success
		responseBytes, err := json.Marshal(contract.ListAssetsResponse{Assets: assets})
		rw.WriteHeader(http.StatusOK)
		rw.Write(responseBytes)

	}
}
func UpdateAssets(service service.AssetService) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		//check request
		rw.Header().Set("content-type", "application/json")

		var req contract.UpadateAssetRequest
		err := json.NewDecoder(r.Body).Decode(&req)

		if err != nil {

			fmt.Println("error while decoding request for update asset ", err.Error())
			//bad request from client
			responseBody, _ := json.Marshal(contract.ErrorResponse{Error: err.Error()})
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(responseBody)
			return

		}

		//validate request
		err = req.Validate()
		if err != nil {

			fmt.Println("error while validating request for update asset ", err.Error())
			//bad request from client
			responseBody, _ := json.Marshal(contract.ErrorResponse{Error: err.Error()})
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(responseBody)
			return

		}
		//call service layer method
		assets, err := service.UpdateAsset(r.Context(), &req)
		if err == custom_errors.InvalidIdError {

			fmt.Println("error Invalid Id")
			//Bad Request
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: err.Error()})
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(responseBytes)
			return
		}
		if err != nil {

			fmt.Println("error while processing request for Update asset", err.Error())
			//internal server error
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: err.Error()})
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write(responseBytes)
			return

		}
		//success
		responseBytes, _ := json.Marshal(assets)
		rw.WriteHeader(http.StatusOK)
		rw.Write(responseBytes)
	}
}

func DeleteAssets(service service.AssetService) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		//validate request
		query := r.URL.Query()
		data := query.Get("id")

		id, err := strconv.Atoi(data)

		if err != nil || id < 0 {
			//Bad Request
			fmt.Println("invaid id")
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: custom_errors.InvalidIdError.Error()})
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(responseBytes)
			return

		}

		//call service layer method
		asset, err := service.DeleteAsset(r.Context(), id)
		if err == custom_errors.InvalidIdError {

			fmt.Println("error Invalid Id")
			//Bad Request
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: err.Error()})
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(responseBytes)
			return
		}
		if err != nil {

			fmt.Println("error occoured while processing Delete asset request")
			//internal server error
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: err.Error()})
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write(responseBytes)
			return
		}
		//success
		responseBytes, _ := json.Marshal(asset)
		rw.WriteHeader(http.StatusOK)
		rw.Write(responseBytes)

	}
}
func CreateAssetHandler(assetService service.AssetService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Set Content-Type for response
		w.Header().Set("Content-Type", "application/json")

		var req contract.CreateAssetRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			fmt.Printf("handler: error while decoding request for create asset: %s", err.Error())

			w.WriteHeader(http.StatusBadRequest)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "invalid request"})
			w.Write(responseBytes)
			return
		}

		err = req.Validate()
		if err != nil {
			fmt.Printf("handler: invalid request for createAsset, some fields are blank")

			w.WriteHeader(http.StatusBadRequest)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: err.Error()})
			w.Write(responseBytes)
			return
		}

		asset := &domain.Asset{
			Name:          req.Name,
			Category:      req.Category,
			Specification: req.Specification,
			InitCost:      req.InitCost,
		}

		asset, err = assetService.CreateAsset(r.Context(), asset)

		if err != nil {
			fmt.Printf("handler: error while creating asset, error: %s", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "something went wrong"})
			w.Write(responseBytes)
			return
		}

		w.WriteHeader(http.StatusOK)
		responseBytes, _ := json.Marshal(contract.CreateAssetResponse{
			ID:            asset.Id,
			Name:          asset.Name,
			Category:      asset.Category,
			Specification: asset.Specification,
			InitCost:      asset.InitCost,
			Status:        asset.Status,
		})
		w.Write(responseBytes)
	}
}
