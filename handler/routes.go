package handler

import "net/http"

func Routes() {
	http.Handle("/ping", PingHandler())
	http.Handle("/login", LoginHandler(deps.userService))
	http.Handle("/user", CreateUserHandler(deps.userService))
	http.Handle("/getUser", GetUser(deps.userService))
}
