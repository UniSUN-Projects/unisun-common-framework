package externalservices

import (
	"github.com/UniSUN-Projects/unisun-common-framework/environment"
	"github.com/UniSUN-Projects/unisun-common-framework/external-services/rest-api/interfaces"
	"github.com/UniSUN-Projects/unisun-common-framework/external-services/rest-api/internals"
	"github.com/UniSUN-Projects/unisun-common-framework/external-services/rest-api/models"
)

func Rest(_method string, _url string, _body []byte) interfaces.RequestInterface {
	environment.Default().Load(&models.Config)
	return internals.NewOptionsHandle(_method, _url, _body)
}
