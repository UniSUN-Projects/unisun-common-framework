package restapi

import (
	model "com/unisun/core-framework/rest-api/model"
)

type OptionConfig struct {
	Method string
	Url    string
	Body   []byte
	Option model.Option `mapstructure:"option"`
}

func New() *OptionConfig {
	return &OptionConfig{}
}
