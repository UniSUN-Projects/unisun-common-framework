package environment

import _model "github.com/UniSUN-Projects/unisun-common-framework/environment/model"

type OptionConfig struct {
	Name   string
	Type   string
	Path   string
	Option _model.Option
}

func Option() *OptionConfig {
	return &OptionConfig{
		Name: "",
		Type: "",
		Path: "",
		Option: _model.Option{
			LoadENV: false,
		},
	}
}
