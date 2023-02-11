package environment

import _interface "com/unisun/core-framework/environment/interface"

type environment struct {
	Factory _interface.FactoryInterface
}

func Init(option _interface.FactoryInterface) *environment {
	return &environment{
		Factory: option,
	}
}
