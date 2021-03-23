package handler

import (
	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	router := mux.NewRouter()
	//user routes
	router.Handle("/ping", PingHandler())
	router.Handle("/login", LoginHandler(deps.userService))
	router.Handle("/users", CreateUserHandler(deps.userService)).Methods("POST")
	router.Handle("/users", GetUser(deps.userService)).Methods("GET")

	//asset routes
	//	router.Handle("/asset/find", FindAssetHandler(deps.assetService)).Methods("GET")
	router.Handle("/assets/all", GetAllAssets(deps.assetService)).Methods("GET")
	router.Handle("/assets/update", UpdateAssets(deps.assetService)).Methods("PUT")
	router.Handle("/assets/delete", DeleteAssets(deps.assetService)).Methods("DELETE")

	router.Handle("/assets", CreateAssetHandler(deps.assetService)).Methods("POST")
	router.Handle("/assets/{id}", FindAssetHandler(deps.assetService)).Methods("GET")

	return router

}
