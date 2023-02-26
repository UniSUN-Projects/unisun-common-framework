package internals

import (
	"fmt"

	"github.com/UniSUN-Projects/unisun-common-framework/web-services/rest-api/models"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

type RestHandleStruct struct{}

func NewSetHandle() *RestHandleStruct {
	return &RestHandleStruct{}
}

func (*RestHandleStruct) Start() {
	Router.Run(fmt.Sprintf(":%s", models.Config.AppConfig.Port))
}

func (*RestHandleStruct) NewServices(_method string, _path string, _handlers gin.HandlerFunc) {
	var path string
	if models.Config.AppConfig.Context != "" {
		path = fmt.Sprintf("%s%s", models.Config.AppConfig.Context, _path)
	}
	Router.Handle(_method, path, _handlers)
}
