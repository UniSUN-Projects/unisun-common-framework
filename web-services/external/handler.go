package external

import (
	"github.com/gin-gonic/gin"
	"unisun.com/backend/unisun-common-framework/web-services/constant"
)

type SetHandleStruct struct{}

func NewSetHandle() *SetHandleStruct {
	return &SetHandleStruct{}
}

func (*SetHandleStruct) New(_method string, _path string, _handlers gin.HandlerFunc) {
	constant.Router.Handle(_method, _path, _handlers)
}
