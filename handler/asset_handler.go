package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vkhichar/assets-manager/contract"
	"github.com/vkhichar/assets-manager/domain"
	"github.com/vkhichar/assets-manager/service"
)

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
			ID:            asset.ID,
			Name:          asset.Name,
			Category:      asset.Category,
			Specification: asset.Specification,
			InitCost:      asset.InitCost,
			Status:        asset.Status,
		})
		w.Write(responseBytes)
	}
}

func FindAssetHandler(assetService service.AssetService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Set content type for response
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])

		if err != nil {
			fmt.Printf("handler: invalid request")
			w.WriteHeader(http.StatusBadRequest)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "Invalid Id"})
			w.Write(responseBytes)
			return
		}

		asset, err := assetService.FindAsset(r.Context(), id)

		if err == service.ErrInvalidId {
			fmt.Printf("handler: invalid id,  Id: %v", id)

			w.WriteHeader(http.StatusNotFound)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "no asset with such id found"})
			w.Write(responseBytes)
			return
		}

		if err != nil {
			fmt.Printf("handler: error while fetching asset: %s", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "something went wrong"})
			w.Write(responseBytes)
			return
		}

		w.WriteHeader(http.StatusOK)
		responseBytes, _ := json.Marshal(contract.CreateAssetResponse{
			ID:            asset.ID,
			Name:          asset.Name,
			Category:      asset.Category,
			Specification: asset.Specification,
			InitCost:      asset.InitCost,
			Status:        asset.Status,
		})
		w.Write(responseBytes)

	}
}
