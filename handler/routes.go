package handler

import "net/http"

//Routes handles the routing
func Routes() {
	http.Handle("/ping", PingHandler())
	http.Handle("/login", LoginHandler(deps.userService))
	http.Handle("/user", CreateUserHandler(deps.userService))
}
