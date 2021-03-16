package handler

import (
	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	router := mux.NewRouter()
	router.Handle("/ping", PingHandler())
	router.Handle("/login", LoginHandler(deps.userService))
	router.Handle("/users", CreateUserHandler(deps.userService)).Methods("POST")
	router.Handle("/users", GetUser(deps.userService)).Methods("GET")
	router.Handle("/assets", CreateAssetHandler(deps.assetService)).Methods("POST")
	router.Handle("/assets/{id}", FindAssetHandler(deps.assetService)).Methods("GET")

	return router

}
