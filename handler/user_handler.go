package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vkhichar/assets-manager/contract"
	"github.com/vkhichar/assets-manager/service"
)

func LoginHandler(userService service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Set Content-Type for response
		w.Header().Set("Content-Type", "application/json")

		var req contract.LoginRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			fmt.Printf("handler: error while decoding request for login: %s", err.Error())

			w.WriteHeader(http.StatusBadRequest)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "invalid request"})
			w.Write(responseBytes)
			return
		}

		err = req.Validate()
		if err != nil {
			fmt.Printf("handler: invalid request for email: %s", req.Email)

			w.WriteHeader(http.StatusBadRequest)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: err.Error()})
			w.Write(responseBytes)
			return
		}

		user, token, err := userService.Login(r.Context(), req.Email, req.Password)
		if err == service.ErrInvalidEmailPassword {
			fmt.Printf("handler: invalid email or password for email: %s", req.Email)

			w.WriteHeader(http.StatusUnauthorized)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "invalid email or password"})
			w.Write(responseBytes)
			return
		}

		if err != nil {
			fmt.Printf("handler: error while logging in for email: %s, error: %s", req.Email, err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "something went wrong"})
			w.Write(responseBytes)
			return
		}

		w.WriteHeader(http.StatusOK)
		responseBytes, _ := json.Marshal(contract.LoginResponse{IsAdmin: user.IsAdmin, Token: token})
		w.Write(responseBytes)
	}
}

//RegisterHandler will handle registration
func CreateUserHandler(userService service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// Set Content-Type for response
		w.Header().Set("Content-Type", "application/json")

		var req contract.CreateUserRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			fmt.Printf("handler: error while decoding request for register: %s", err.Error())

			w.WriteHeader(http.StatusBadRequest)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "invalid request"})
			w.Write(responseBytes)
			return
		}

		err = req.Validate()
		if err != nil {
			fmt.Printf("handler: empty filed provided")

			w.WriteHeader(http.StatusBadRequest)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "Empty field provided"})
			w.Write(responseBytes)
			return
		}

		user, err := userService.Register(r.Context(), req.Name, req.Email, req.Password, req.IsAdmin)
		if err != nil && err.Error() == service.ErrDuplicateEmail.Error() {
			fmt.Printf("handler: email already exist: %s", req.Email)

			w.WriteHeader(http.StatusConflict)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "email already exist"})
			w.Write(responseBytes)
			return
		}

		if err != nil {
			fmt.Printf("handler: error while registering for email: %s, error: %s", req.Email, err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "something went wrong"})
			w.Write(responseBytes)
			return
		}

		w.WriteHeader(http.StatusOK)
		responseBytes, _ := json.Marshal(contract.CreateUserResponse{ID: user.ID, Name: user.Name, Email: user.Email, IsAdmin: user.IsAdmin})
		w.Write(responseBytes)
	}
}

func GetUser(userService service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		idStrr := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idStrr)
		if err != nil {
			fmt.Printf("handler: no id provided")
			w.WriteHeader(http.StatusBadRequest)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "No user id is provided"})
			w.Write(responseBytes)
			return
		}

		user, err := userService.GetUser(r.Context(), id)

		if err != nil && err.Error() == service.ErrNoSqlRow.Error() {
			fmt.Printf("handler: No value for id : %d, error: %s", id, err.Error())
			w.WriteHeader(http.StatusOK)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "No data present for this id"})
			w.Write(responseBytes)
			return
		}

		if err != nil {
			fmt.Printf("handler:error while fetching data")
			w.WriteHeader(http.StatusInternalServerError)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "something went wrong"})
			w.Write(responseBytes)
			return
		}

		w.WriteHeader(http.StatusOK)
		responseBytes, _ := json.Marshal(contract.GetUserResponse{Name: user.Name, Email: user.Email, IsAdmin: user.IsAdmin})
		w.Write(responseBytes)
	}
}

func UpdateUserHandler(userServce service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var req contract.CreateUserRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			fmt.Printf("handler: error while decoding request for update: %s", err.Error())

			w.WriteHeader(http.StatusBadRequest)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "invalid request"})
			w.Write(responseBytes)
			return
		}

		err = req.ValidateUpdate()
		if err != nil {
			fmt.Printf("handler: empty filed provided")

			w.WriteHeader(http.StatusBadRequest)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "Empty field provided"})
			w.Write(responseBytes)
			return
		}
		idStrr := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idStrr)
		if err != nil {
			fmt.Printf("handler: no id provided %s", idStrr)
			w.WriteHeader(http.StatusBadRequest)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "No user id is provided"})
			w.Write(responseBytes)
			return
		}
		val := contract.UpdateUserRequest{
			Name:     req.Name,
			Email:    req.Email,
			Password: req.Password,
			IsAdmin:  req.IsAdmin,
		}
		user, err := userServce.Update(r.Context(), id, val)
		if err != nil {
			fmt.Printf("handler: error while registering for email: %s, error: %s", req.Email, err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "something went wrong"})
			w.Write(responseBytes)
			return
		}

		w.WriteHeader(http.StatusOK)
		responseBytes, _ := json.Marshal(contract.CreateUserResponse{ID: user.ID, Name: user.Name, Email: user.Email, IsAdmin: user.IsAdmin})
		w.Write(responseBytes)
	}
}
