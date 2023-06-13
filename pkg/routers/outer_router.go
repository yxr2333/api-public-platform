package routers

import (
	"api-public-platform/api/handlers"
	"api-public-platform/api/middlewares"

	"github.com/gin-gonic/gin"
)

type OuterRouter struct {
	outerHandler handlers.OuterHandler
}

func NewOuterRouter() *OuterRouter {
	return &OuterRouter{
		outerHandler: handlers.NewOuterHandler(),
	}
}

func (or *OuterRouter) loadOuterRouter(e *gin.RouterGroup) {
	e.GET("/hello",
		middlewares.ValidateAPITokenAndPermissions(),
		or.outerHandler.Hello)
}
