package routers

import (
	"comment/api"
)

func (RG RouterGroup) UserRouter() {
	userApi := api.Groups.UserApi
	RG.Router.GET("/user_get_login", userApi.UserGetLoginView)
	RG.Router.POST("/user_login", userApi.UsernameLoginView)
	RG.Router.POST("/user_register", userApi.UserRegisterView)
}
