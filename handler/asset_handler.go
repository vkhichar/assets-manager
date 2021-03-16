package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/vkhichar/assets-manager/contract"
	"github.com/vkhichar/assets-manager/service"
)

func FindAssetHandler(service service.AssetService) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		rw.Header().Set("content-type", "application/json")
		//decode request
		var req contract.FindAssetRequest
		err := json.NewDecoder(r.Body).Decode(&req)

		if err != nil {

			fmt.Println("error while decoding request for Find asset ", err.Error())
			//bad request from client
			responseBody, _ := json.Marshal(err.Error())
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(responseBody)
			return
		}
		//validate
		if err = req.Validate(); err != nil {
			fmt.Println("error while validating request find asset", err.Error())
			//bad request from client
			responseBody, _ := json.Marshal(err.Error())
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(responseBody)
			return
		}

		//call service layer method
		asset, err := service.FindAsset(r.Context(), req.Id)
		if err != nil {
			fmt.Println("error while processing request for find asset", err.Error())
			//internal server error
			responseBytes, _ := json.Marshal(err.Error())
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write(responseBytes)
			return
		}

		responseBytes, err := json.Marshal(contract.FindAssetResponse{Id: asset.Id, Name: asset.Name, Category: asset.Category, Specification: *asset.Specification, InitCost: asset.InitCost, Status: asset.Status})
		rw.WriteHeader(http.StatusOK)
		rw.Write(responseBytes)
	}
}

func GetAllAssets(service service.AssetService) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		rw.Header().Set("content-type", "application/json")

		assets, err := service.GetAssets()

		if err != nil {
			fmt.Print("error while processing request for get all assets", err.Error())
			responseBytes, _ := json.Marshal(err.Error())
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write(responseBytes)
		}

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
			responseBody, _ := json.Marshal(err.Error())
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(responseBody)
			return

		}

		//validate request
		err = req.Validate()
		if err != nil {

			fmt.Println("error while validating request for update asset ", err.Error())
			//bad request from client
			responseBody, _ := json.Marshal(err.Error())
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(responseBody)
			return

		}
		//call service layer method
		asset, err := service.UpdateAsset(r.Context(), &req)

		if err != nil {

			fmt.Println("error while processing request for Update asset", err.Error())
			//internal server error
			responseBytes, _ := json.Marshal(err.Error())
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write(responseBytes)
			return

		}
		if asset == nil {
			fmt.Println("error while processing request for Update asset is nil", err.Error())
			//internal server error
			responseBytes, _ := json.Marshal(err.Error())
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write(responseBytes)
			return
		}
		responseBytes, _ := json.Marshal(asset)
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

			fmt.Println("invaid id ")
			rw.WriteHeader(http.StatusBadRequest)
			return

		}

		//call service layer method
		asset, err := service.DeleteService(r.Context(), id)
		if err != nil {

			fmt.Println("error occoured while processing Delete asset request")

			responseBytes, _ := json.Marshal(err.Error())
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write(responseBytes)
			return
		}

		responseBytes, _ := json.Marshal(asset)
		rw.WriteHeader(http.StatusOK)
		rw.Write(responseBytes)
	}
}
