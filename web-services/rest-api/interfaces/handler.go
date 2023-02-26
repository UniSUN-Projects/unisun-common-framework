package interfaces

import "github.com/gin-gonic/gin"

type RestHandlerPort interface {
	Start()
	NewServices(_method string, _path string, _handlers gin.HandlerFunc)
}
