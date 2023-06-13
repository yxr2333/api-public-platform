package routers

import (
	"api-public-platform/api/handlers"
	"api-public-platform/api/middlewares"
	"api-public-platform/api/request"
	"api-public-platform/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	userHandler handlers.UserHandler
}

func NewUserRouter() *UserRouter {
	return &UserRouter{
		userHandler: handlers.NewUserHandler(),
	}
}

func (ur *UserRouter) loadUserRouter(e *gin.RouterGroup) {
	user := e.Group("/user")
	{
		user.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello World!",
			})
		})
		user.POST("/register",
			middlewares.ValidateRequest(&request.UserRegisterRequest{}),
			ur.userHandler.UserRegisterHandler)
		user.POST("/login",
			middlewares.ValidateRequest(&request.UserLoginRequest{}),
			ur.userHandler.UserLoginHandler)
		user.GET("/admin",
			middlewares.Authorization("admin|superadmin"),
		)
		// 添加创建用户的路由
		user.POST("/",
			middlewares.Authorization("admin|superadmin"),
			middlewares.ValidateRequest(&request.UserCreateRequest{}),
			ur.userHandler.CreateUserHandler)
		// 添加更新用户的路由
		user.PUT("/:id",
			middlewares.Authorization("any"),
			middlewares.ValidateRequest(&model.User{}),
			ur.userHandler.UpdateUserHandler)
		// 添加删除用户的路由
		user.DELETE("/:id",
			middlewares.Authorization("admin|superadmin"),
			ur.userHandler.DeleteUserHandler)
		// 添加生成API令牌的路由
		user.POST("/:id/token/generate",
			middlewares.Authorization("any"),
			ur.userHandler.GenerateAPITokenHandler)
		// 添加更新API令牌的路由
		user.PUT("/:id/token/update",
			middlewares.Authorization("any"),
			ur.userHandler.UpdateAPITokenHandler)

	}
}
