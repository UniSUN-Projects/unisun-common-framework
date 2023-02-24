package webservices

import (
	"fmt"
	"time"

	"github.com/UniSUN-Projects/unisun-common-framework/web-services/constant"
	"github.com/UniSUN-Projects/unisun-common-framework/web-services/external"
	"github.com/UniSUN-Projects/unisun-common-framework/web-services/interfaces"
	"github.com/gin-gonic/gin"
)

func NewPortHandle(port interfaces.HandlerPort) *interfaces.HandlerPort {
	return &port
}

func Init() *interfaces.HandlerPort {
	constant.Router = gin.Default()
	constant.Router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	constant.Router.Use(gin.Recovery())
	SetHandleStruct := external.NewSetHandle()
	return NewPortHandle(SetHandleStruct)
}

func Start() {
	constant.Router.Run()
}

func StartWithPort(_port string) {
	constant.Router.Run(":" + _port)
}
