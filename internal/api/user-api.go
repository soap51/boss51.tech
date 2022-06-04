package api

import (
	"github.com/gin-gonic/gin"
	"github.com/soap51/boss51.tech/internal/port/driving"
)

type userApi struct {
	userService driving.IUserService
}

func NewUserApi(userService driving.IUserService) *userApi {
	return &userApi{
		userService: userService,
	}
}

func (api *userApi) Register(g gin.Context) {

}

func (api *userApi) Login(g gin.Context) {

}

func (api *userApi) ForgotPassword(g gin.Context) {

}

func (api *userApi) VerifyForgotPassword(g gin.Context) {

}

func (api *userApi) ChangePassword(g gin.Context) {

}
