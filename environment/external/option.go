package external

import (
	"github.com/UniSUN-Projects/unisun-common-framework/environment/constant"
	"github.com/UniSUN-Projects/unisun-common-framework/environment/internals"
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

func (*OptionHandle) Option() *internals.ConfigEnvironment {
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
		return internals.NewLoad(stringsUtils(constant.NAME, op.Name).(string), stringsUtils(constant.TYPE, op.Type).(string), stringsUtils(constant.PATH, op.Path).(string), stringsUtils(constant.LOAD_ENV, op.Option.LoadENV).(bool))
	} else {
		return internals.NewLoad(constant.DEFAULT_NAME, constant.DEFAULT_TYPE, constant.DEFAULT_PATH, constant.DEFAULT_LOAD_ENV)
	}
}

func stringsUtils(key string, value interface{}) interface{} {
	if value != "" || value == true {
		return value
	} else {
		var _value interface{}
		switch key {
		case constant.NAME:
			_value = constant.DEFAULT_NAME
		case constant.TYPE:
			_value = constant.DEFAULT_TYPE
		case constant.PATH:
			_value = constant.DEFAULT_PATH
		case constant.LOAD_ENV:
			_value = constant.DEFAULT_LOAD_ENV
		}
		return _value
	}
}
