package environment

import (
	"log"
	"strings"

	_interface "github.com/UniSUN-Projects/unisun-common-framework/environment/interface"
	_model "github.com/UniSUN-Projects/unisun-common-framework/environment/model"

	"github.com/spf13/viper"
)

type ConfigEnvironment struct {
	Type   string
	Name   string
	Path   string
	Option _model.Option
}

func Init(_name string, _type string, _path string, _option _model.Option) *ConfigEnvironment {
	return &ConfigEnvironment{
		Type:   _type,
		Name:   _name,
		Path:   _path,
		Option: _option,
	}
}

type LoadFunc struct {
	Load _interface.LoadInterface
}

func New(load _interface.LoadInterface) *LoadFunc {
	return &LoadFunc{
		Load: load,
	}
}

func (c *ConfigEnvironment) Load(model *any) {
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
