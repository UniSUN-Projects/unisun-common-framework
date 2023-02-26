package webservices

import (
	"fmt"
	"time"

	"github.com/UniSUN-Projects/unisun-common-framework/environment"
	"github.com/UniSUN-Projects/unisun-common-framework/web-services/rest-api/interfaces"
	"github.com/UniSUN-Projects/unisun-common-framework/web-services/rest-api/internals"
	"github.com/UniSUN-Projects/unisun-common-framework/web-services/rest-api/models"
	"github.com/gin-gonic/gin"
)

func Rest() interfaces.RestHandlerPort {
	environment.Default().Load(&models.Config)
	internals.Router = gin.Default()
	internals.Router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
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
	internals.Router.Use(gin.Recovery())
	return internals.NewSetHandle()
}
