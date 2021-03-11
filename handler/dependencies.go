package handler

import (
	"github.com/vkhichar/assets-manager/repository"
	"github.com/vkhichar/assets-manager/service"
)

type dependencies struct {
	userService service.UserService
}

var deps dependencies

func InitDependencies() {
	userRepo := repository.NewUserRepository()
	plainTokenService := service.NewPlainTokenService()

	userService := service.NewUserService(userRepo, plainTokenService)
	deps.userService = userService
}
