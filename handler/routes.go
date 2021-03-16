package handler

import (
	"net/http"
	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	router := mux.NewRouter()
	router.Handle("/ping", PingHandler())
	router.Handle("/login", LoginHandler(deps.userService))
	router.Handle("/users", CreateUserHandler(deps.userService).Method("POST"))
	router.Handle("/users", GetUser(deps.userService).Method("GET"))
}
