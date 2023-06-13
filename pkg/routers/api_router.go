package routers

import (
	"api-public-platform/api/handlers"
	"api-public-platform/api/middlewares"
	"api-public-platform/api/request"

	"github.com/gin-gonic/gin"
)

type APIRouter struct {
	apiHandler handlers.APIHandler
}

func NewAPIRouter() *APIRouter {
	return &APIRouter{
		apiHandler: handlers.NewAPIHandler(),
	}
}

func (ar *APIRouter) loadAPIRouter(e *gin.RouterGroup) {
	api := e.Group("/apis")
	{
		api.POST("",
			middlewares.ValidateRequest(&request.APICreateRequest{}),
			ar.apiHandler.CreateAPI)
		api.PUT("",
			middlewares.ValidateRequest(&request.APIUpdateRequest{}),
			ar.apiHandler.UpdateAPI)
		api.DELETE("/:id",
			middlewares.Authorization("admin|superadmin"),
			ar.apiHandler.DeleteAPI)
		api.GET("/:id", ar.apiHandler.GetAPIByID)
		api.GET("/",
			middlewares.Authorization("admin|superadmin"),
			middlewares.ValidatePageQueryRequest(),
			ar.apiHandler.GetAllAPIs)
		api.POST("/enable/:id",
			middlewares.Authorization("admin|superadmin"),
			ar.apiHandler.EnableAPI)
		api.POST("/disable/:id",
			middlewares.Authorization("admin|superadmin"),
			ar.apiHandler.DisableAPI)
	}
}
