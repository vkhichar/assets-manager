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
	router.Handle("/assets", GetAllAssets(deps.assetService)).Methods("GET")
	router.Handle("/assets/{id:[0-9]+}", UpdateAssets(deps.assetService)).Methods("PUT")
	router.Handle("/assets/{id:[0-9]+}", DeleteAssets(deps.assetService)).Methods("DELETE")

	router.Handle("/assets", CreateAssetHandler(deps.assetService)).Methods("POST")
	router.Handle("/assets/{asset_id:[0-9]+}", FindAssetHandler(deps.assetService)).Methods("GET")

	//asset allocation routes
	router.Handle("/users/{user_id:[0-9]+}/allocate-asset", AllocateAssetHandler(deps.assetAllocService)).Methods("POST")
	router.Handle("/users/{user_id:[0-9]+}/assets/{asset_id:[0-9]+}", DeallocateAssetHandler(deps.assetAllocService)).Methods("DELETE")

	return router

}
