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
	Router.Handle(_method, _path, _handlers)
}
