package routers

import (
	"api-public-platform/api/handlers"
	"api-public-platform/api/middlewares"
	"api-public-platform/api/request"

	"github.com/gin-gonic/gin"
)

var apiHandler = handlers.NewAPIHandler()

func loadAPIRouter(e *gin.RouterGroup) {
	api := e.Group("/apis")
	{
		api.POST("",
			middlewares.ValidateRequest(&request.APICreateRequest{}),
			apiHandler.CreateAPI)
		api.PUT("",
			middlewares.ValidateRequest(&request.APIUpdateRequest{}),
			apiHandler.UpdateAPI)
		api.DELETE("/:id",
			middlewares.Authorization("admin|superadmin"),
			apiHandler.DeleteAPI)
		api.GET("/:id", apiHandler.GetAPIByID)
		api.GET("/",
			middlewares.Authorization("admin|superadmin"),
			middlewares.ValidatePageQueryRequest(),
			apiHandler.GetAllAPIs)
		api.POST("/enable/:id",
			middlewares.Authorization("admin|superadmin"),
			apiHandler.EnableAPI)
		api.POST("/disable/:id",
			middlewares.Authorization("admin|superadmin"),
			apiHandler.DisableAPI)
	}
}
