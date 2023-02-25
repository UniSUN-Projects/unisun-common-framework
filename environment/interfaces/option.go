package interfaces

import "github.com/UniSUN-Projects/unisun-common-framework/environment/internal"

type OptionInterface interface {
	SetOption(key string, value string)
	DelOption(key string)
	Option() *internal.ConfigEnvironment
}
