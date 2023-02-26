package interfaces

import "github.com/UniSUN-Projects/unisun-common-framework/environment/internals"

type OptionInterface interface {
	SetOption(key string, value string)
	DelOption(key string)
	Option() *internals.ConfigEnvironment
}
