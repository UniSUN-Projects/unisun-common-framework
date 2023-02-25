package restapi

import "github.com/UniSUN-Projects/unisun-common-framework/rest-api/interfaces"

type restApi struct {
	Request interfaces.RequestInterface
}

func Init(option interfaces.RequestInterface) *restApi {
	return &restApi{
		Request: option,
	}
}
