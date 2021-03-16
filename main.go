package main

import (
	"fmt"
	"net/http"

	"github.com/vkhichar/assets-manager/config"
	"github.com/vkhichar/assets-manager/handler"
	"github.com/vkhichar/assets-manager/repository"
)

func main() {
	err := config.Init()
	if err != nil {
		fmt.Printf("main: error while initialising config: %s", err.Error())
		return
	}

	// initialise db connection
	repository.InitDB()

	handler.InitDependencies()
	handler.Routes()

	err = http.ListenAndServe(":"+config.GetAppPort(), handler.Routes())
	if err != nil {
		fmt.Printf("main: error while starting server: %s", err.Error())
		return
	}
}
