package environment

import (
	"github.com/UniSUN-Projects/unisun-common-framework/environment/constant"
	"github.com/UniSUN-Projects/unisun-common-framework/environment/external"
	"github.com/UniSUN-Projects/unisun-common-framework/environment/interfaces"
	"github.com/UniSUN-Projects/unisun-common-framework/environment/internals"
)

var LoadInterface interfaces.LoadInterface
var OptionInterface interfaces.OptionInterface

func Default() interfaces.LoadInterface {
	external.Initial()
	return internals.NewLoad(constant.DEFAULT_NAME, constant.DEFAULT_TYPE, constant.DEFAULT_PATH, constant.DEFAULT_LOAD_ENV)
}

func New() *external.OptionHandle {
	external.Initial()
	return external.NewOptionHandle()
}
