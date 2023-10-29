// Package modules bootstrap all of the modules in your application
package modules

import (
	"github.com/Yasir900Aslam/go_mongo_modules/core"
	user "github.com/Yasir900Aslam/go_mongo_modules/modules/user"
)

type AppModule = core.Module

func NewAppModule() *AppModule {
	return core.NewModule().
		SetName("App").
		AddImport(user.New())
}
