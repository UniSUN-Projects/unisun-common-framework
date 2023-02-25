package internal

import (
	"log"
	"strings"

	"github.com/UniSUN-Projects/unisun-common-framework/environment/interfaces"
	"github.com/UniSUN-Projects/unisun-common-framework/environment/models"
	"github.com/spf13/viper"
)

type ConfigEnvironment struct {
	Type   string
	Name   string
	Path   string
	Option models.Option
}

func Init(_name string, _type string, _path string, _option models.Option) *ConfigEnvironment {
	return &ConfigEnvironment{
		Type:   _type,
		Name:   _name,
		Path:   _path,
		Option: _option,
	}
}

type LoadFunc struct {
	Load interfaces.LoadInterface
}

func New(load interfaces.LoadInterface) *LoadFunc {
	return &LoadFunc{
		Load: load,
	}
}

func (c *ConfigEnvironment) Load(model *interface{}) {
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
