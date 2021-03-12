package handler

import "net/http"

func Routes() {
	http.Handle("/ping", PingHandler())
	http.Handle("/login", LoginHandler(deps.userService))
	http.Handle("/assets", CreateAssetHandler(deps.assetService))
	http.Handle("/asset", FindAssetHandler(deps.assetService))
}
