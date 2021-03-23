package handler

import (
	"github.com/vkhichar/assets-manager/repository"
	"github.com/vkhichar/assets-manager/service"
)

type dependencies struct {
	userService  service.UserService
	assetService service.AssetService
}

var deps dependencies

func InitDependencies() {

	//user dependencies
	userRepo := repository.NewUserRepository()
	plainTokenService := service.NewPlainTokenService()
	userService := service.NewUserService(userRepo, plainTokenService)

	//asset depedencies
	assetRepo := repository.NewAssetRepository()
	assetService := service.NewAssetService(assetRepo)

	deps.assetService = assetService
	deps.userService = userService

}
