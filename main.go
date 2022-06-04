package main

import (
	"github.com/soap51/boss51.tech/internal/api"
	"github.com/soap51/boss51.tech/internal/repositories"
	"github.com/soap51/boss51.tech/internal/services"
)

func main() {
	userRepository := repositories.NewMongoUserRepository()
	userService := services.NewUserService(userRepository)
	api.NewUserApi(userService)

}
