package restapi

import _interface "unisun.com/backend/unisun-common-framework/rest-api/interface"

type restApi struct {
	Request _interface.RequestInterface
}

func Init(option _interface.RequestInterface) *restApi {
	return &restApi{
		Request: option,
	}
}
