package environment

import _interface "unisun.com/backend/unisun-common-framework/environment/interface"

type environment struct {
	Factory _interface.FactoryInterface
}

func Init(option _interface.FactoryInterface) *environment {
	return &environment{
		Factory: option,
	}
}
