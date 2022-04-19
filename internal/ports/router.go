package ports

import "github.com/gin-gonic/gin"

// Router defines the interface for managing router library
type Router interface {
	GET(uri string, fn func(ctx *gin.Context))
	POST(uri string, fn func(ctx *gin.Context))
	PUT(uri string, fn func(ctx *gin.Context))
	DELETE(uri string, fn func(ctx *gin.Context))
	SERVE() error
	GROUP(path string) *gin.RouterGroup
}
