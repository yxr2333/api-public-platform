package routers

import (
	"api-public-platform/api/middlewares"

	"github.com/gin-gonic/gin"
)

type Router struct {
	userRouter  *UserRouter
	apiRouter   *APIRouter
	outerRouter *OuterRouter
}

func NewRouter() *Router {
	return &Router{
		userRouter:  NewUserRouter(),
		apiRouter:   NewAPIRouter(),
		outerRouter: NewOuterRouter(),
	}
}

func (router *Router) SetUpRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api/internal/v1")
	{
		router.userRouter.loadUserRouter(api)
		router.apiRouter.loadAPIRouter(api)
	}
	public := r.Group("/api/public/v1")
	{
		public.Use(middlewares.Authorization("any"))
		public.Use(middlewares.APICallHistoryMiddleware())
		router.outerRouter.loadOuterRouter(public)
	}
	return r
}
