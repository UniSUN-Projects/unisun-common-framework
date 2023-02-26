package internals

import (
	"log"
	"strings"

	"github.com/UniSUN-Projects/unisun-common-framework/environment/models"
	"github.com/spf13/viper"
)

type ConfigEnvironment struct {
	Type   string
	Name   string
	Path   string
	Option models.Option
}

func NewLoad(_name string, _type string, _path string, _loadEnv bool) *ConfigEnvironment {
	return &ConfigEnvironment{
		Type: _type,
		Name: _name,
		Path: _path,
		Option: models.Option{
			LoadENV: _loadEnv,
		},
	}
}

func (c *ConfigEnvironment) Load(model any) {
	v := viper.New()
	v.SetConfigName(c.Name)
	v.SetConfigType(c.Type)
	v.AddConfigPath(c.Path)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if c.Option.LoadENV {
		v.AutomaticEnv()
	}
	if err := v.ReadInConfig(); err != nil {
		log.Panic(err)
	}
	if err := v.Unmarshal(model); err != nil {
		log.Panic(err)
	}
}
