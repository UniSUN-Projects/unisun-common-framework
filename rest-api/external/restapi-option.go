package restapi

import (
	model "unisun.com/backend/unisun-common-framework/rest-api/model"
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
