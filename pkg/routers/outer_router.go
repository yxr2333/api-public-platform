package routers

import (
	"api-public-platform/api/handlers"

	"github.com/gin-gonic/gin"
)

var outerHandler = handlers.NewOuterHandler()

func loadOuterRouter(e *gin.RouterGroup) {
	e.GET("/hello", outerHandler.Hello)
}
