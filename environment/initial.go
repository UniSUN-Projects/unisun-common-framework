package environment

import "github.com/UniSUN-Projects/unisun-common-framework/environment/interfaces"

type environment struct {
	Factory interfaces.FactoryInterface
}

func Init(option interfaces.FactoryInterface) *environment {
	return &environment{
		Factory: option,
	}
}
