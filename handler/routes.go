package handler

import (
	"net/http"
	"github.com/gorilla/mux"
)

<<<<<<< HEAD
func Routes() {
	http.Handle("/ping", PingHandler())
	http.Handle("/login", LoginHandler(deps.userService))
	http.Handle("/assets", CreateAssetHandler(deps.assetService))
	http.Handle("/asset", FindAssetHandler(deps.assetService))
=======
func Routes() *mux.Router {
	router := mux.NewRouter()
	router.Handle("/ping", PingHandler())
	router.Handle("/login", LoginHandler(deps.userService))
	router.Handle("/users", CreateUserHandler(deps.userService).Method("POST"))
	router.Handle("/users", GetUser(deps.userService).Method("GET"))
>>>>>>> ee27ae2d82de8635a945000574666bac17d6d7f0
}
