package external

import (
	"github.com/UniSUN-Projects/unisun-common-framework/environment/constant"
	"github.com/UniSUN-Projects/unisun-common-framework/environment/internal"
	"github.com/UniSUN-Projects/unisun-common-framework/environment/models"
)

func (_option *OptionConfig) Default(model interface{}) {
	var option models.Option
	init := internal.Init(constant.DEFAULT_NAME, constant.DEFAULT_TYPE, constant.DEFAULT_PATH, option)
	internal.New(init).Load.Load(&model)
}

func (_option *OptionConfig) Custom(model interface{}) {
	init := internal.Init(_option.Name, _option.Type, _option.Path, _option.Option)
	internal.New(init).Load.Load(&model)
}
