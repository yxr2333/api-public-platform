package routers

import (
	"api-public-platform/api/handlers"
	"api-public-platform/api/middlewares"
	"api-public-platform/api/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func loadUserRouter(e *gin.RouterGroup) {
	user := e.Group("/user")
	{
		user.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello World!",
			})
		})
		user.POST("/register",
			middlewares.ValidateRequest(&request.UserRegisterRequest{}),
			handlers.UserRegisterHandler)
		user.POST("/login",
			middlewares.ValidateRequest(&request.UserLoginRequest{}),
			handlers.UserLoginHandler)
		user.GET("/admin",
			middlewares.Authorization("admin|superadmin"),
		)

	}
}
