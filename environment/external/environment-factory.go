package environment

import (
	constant "github.com/UniSUN-Projects/unisun-common-framework/environment/constant"
	initial "github.com/UniSUN-Projects/unisun-common-framework/environment/internal"
	_model "github.com/UniSUN-Projects/unisun-common-framework/environment/model"
)

func (_option *OptionConfig) Default(model *any) {
	var option _model.Option
	init := initial.Init(constant.DEFAULT_NAME, constant.DEFAULT_TYPE, constant.DEFAULT_PATH, option)
	initial.New(init).Load.Load(model)
}

func (_option *OptionConfig) Custom(model *any) {
	init := initial.Init(_option.Name, _option.Type, _option.Path, _option.Option)
	initial.New(init).Load.Load(model)
}
