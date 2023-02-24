package interfaces

import "github.com/gin-gonic/gin"

type HandlerPort interface {
	New(_method string, _path string, _handlers gin.HandlerFunc)
}
