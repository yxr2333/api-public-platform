package middlewares

import (
	"api-public-platform/config"
	"api-public-platform/internal/db"
	"api-public-platform/pkg/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckApiOpenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		endpoint := strings.TrimPrefix(c.Request.URL.Path, config.ServerCfg.API.Outer.Prefix)
		method := c.Request.Method
		api := model.API{}
		if err := db.MySQLDB.Where("api_endpoint = ? AND request_method = ?", endpoint, method).First(&api).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "API not found",
			})
			c.Abort()
			return
		}
		if !api.IsOpen {
			c.JSON(200, gin.H{
				"msg":  "API not open",
				"code": 500,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
