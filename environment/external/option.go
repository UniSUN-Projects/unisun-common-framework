package external

import (
	"github.com/UniSUN-Projects/unisun-common-framework/environment/constant"
	"github.com/UniSUN-Projects/unisun-common-framework/environment/internal"
	"github.com/UniSUN-Projects/unisun-common-framework/environment/models"
)

var OptionSet map[string]interface{}

func Initial() {
	OptionSet = make(map[string]interface{})
}

type OptionHandle struct{}

func NewOptionHandle() *OptionHandle {
	return &OptionHandle{}
}

func (*OptionHandle) SetOption(key string, value string) {
	OptionSet[key] = value
}

func (*OptionHandle) DelOption(key string) {
	delete(OptionSet, key)
}

type OptionConfig struct {
	Name   string
	Type   string
	Path   string
	Option models.Option
}

func (*OptionHandle) Option() *internal.ConfigEnvironment {
	index := len(OptionSet)
	if index > 0 {
		op := OptionConfig{}
		for k, v := range OptionSet {
			switch k {
			case constant.NAME:
				op.Name = v.(string)
			case constant.TYPE:
				op.Type = v.(string)
			case constant.PATH:
				op.Path = v.(string)
			case constant.LOAD_ENV:
				op.Option.LoadENV = v.(bool)
			}
		}
		return internal.NewLoad(op.Name, op.Type, op.Path, op.Option.LoadENV)
	} else {
		return internal.NewLoad(constant.DEFAULT_NAME, constant.DEFAULT_TYPE, constant.DEFAULT_PATH, constant.DEFAULT_LOAD_ENV)
	}
}
