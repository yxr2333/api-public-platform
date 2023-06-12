package routers

import "github.com/gin-gonic/gin"

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api/internal/v1")
	{
		loadUserRouter(api)
		loadAPIRouter(api)
	}
	public := r.Group("/api/public/v1")
	{
		loadOuterRouter(public)
	}
	return r
}
