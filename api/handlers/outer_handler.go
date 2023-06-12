package handlers

import (
	"api-public-platform/pkg/service"

	"github.com/gin-gonic/gin"
)

type OuterHandler struct {
	outerService service.OuterService
}

func NewOuterHandler() *OuterHandler {
	return &OuterHandler{
		outerService: service.NewOuterHandler(),
	}
}

func (oh *OuterHandler) Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": oh.outerService.Hello(),
	})
}
