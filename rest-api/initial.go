package restapi

import _interface "com/unisun/core-framework/rest-api/interface"

type restApi struct {
	Request _interface.RequestInterface
}

func Init(option _interface.RequestInterface) *restApi {
	return &restApi{
		Request: option,
	}
}
