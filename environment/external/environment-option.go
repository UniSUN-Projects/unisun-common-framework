package environment

import _model "github.com/UniSUN-Projects/unisun-common-framework/environment/model"

type OptionConfig struct {
	Name   string
	Type   string
	Path   string
	Option _model.Option
}

var OptionSet map[string]interface{}

func SetOption(key string, value string) {
	OptionSet[key] = value
}

func DelOption(key string) {
	delete(OptionSet, key)
}

func Option() *OptionConfig {
	index := len(OptionSet)
	if index > 0 {
		op := OptionConfig{}
		for k, v := range OptionSet {
			switch k {
			case "name":
				op.Name = v.(string)
			case "type":
				op.Type = v.(string)
			case "path":
				op.Path = v.(string)
			case "loadENV":
				op.Option.LoadENV = v.(bool)
			}
		}
		return &op
	} else {
		return &OptionConfig{
			Name: "",
			Type: "",
			Path: "",
			Option: _model.Option{
				LoadENV: false,
			},
		}
	}
}
