package datasource

import (
	"github.com/UniSUN-Projects/unisun-common-framework/datasource/interfaces"
	"github.com/UniSUN-Projects/unisun-common-framework/datasource/internals"
	"github.com/UniSUN-Projects/unisun-common-framework/datasource/models"
	"github.com/UniSUN-Projects/unisun-common-framework/environment"
)

func Start() interfaces.DbConnectInterface {
	environment.Default().Load(&models.Config)
	if models.Config.AppConfig.Datasource.Host == "" || models.Config.AppConfig.Datasource.Username == "" || models.Config.AppConfig.Datasource.Password == "" {
		panic("Check your property datasource!")
	}
	return internals.NewOptionHandle()
}
