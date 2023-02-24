package external

import (
	"github.com/UniSUN-Projects/unisun-common-framework/web-services/constant"
	"github.com/gin-gonic/gin"
)

type SetHandleStruct struct{}

func NewSetHandle() *SetHandleStruct {
	return &SetHandleStruct{}
}

func (*SetHandleStruct) New(_method string, _path string, _handlers gin.HandlerFunc) {
	constant.Router.Handle(_method, _path, _handlers)
}
