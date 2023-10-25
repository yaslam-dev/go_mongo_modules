package user

import "github.com/Yasir900Aslam/go_mongo_modules/core"

type UserModule = core.Module

func New() *UserModule {
	return core.NewModule().
		SetName("User").
		SetDescription("User Module for authentication and authorization").
		AddController(NewAuthController()).
		AddEntity(User{})
}
