package external

import "github.com/UniSUN-Projects/unisun-common-framework/rest-api/models"

type OptionConfig struct {
	Method string
	Url    string
	Body   []byte
	Option models.Option `mapstructure:"option"`
}

func New() *OptionConfig {
	return &OptionConfig{}
}
