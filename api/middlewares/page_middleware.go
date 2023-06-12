package middlewares

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ValidatePageQueryRequest() func(c *gin.Context) {
	return func(c *gin.Context) {
		page := c.DefaultQuery("page", "1")
		size := c.DefaultQuery("siz", "10")
		page_int, err := strconv.Atoi(page)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "page type error",
			})
			c.Abort()
			return
		}
		size_int, err := strconv.Atoi(size)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "size type error",
			})
			c.Abort()
			return
		}
		if page_int < 1 {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "page must be greater than 0",
			})
			c.Abort()
			return
		}
		if size_int < 1 {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "size must be greater than 0",
			})
			c.Abort()
			return
		}
		c.Set("page", page_int)
		c.Set("size", size_int)
		c.Next()
	}
}
