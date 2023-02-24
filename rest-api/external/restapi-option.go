package restapi

import (
	model "github.com/UniSUN-Projects/unisun-common-framework/rest-api/model"
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
