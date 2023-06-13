package handlers

import (
	"api-public-platform/pkg/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

type OuterHandler struct {
	outerService service.OuterService
}

func NewOuterHandler() OuterHandler {
	return OuterHandler{
		outerService: service.NewOuterHandler(),
	}
}

func (oh *OuterHandler) Hello(c *gin.Context) {
	fmt.Println("Hello函数被调用")
	c.JSON(200, gin.H{
		"msg": oh.outerService.Hello(),
	})
}
