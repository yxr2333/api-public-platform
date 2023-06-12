package middlewares

import (
	"api-public-platform/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidateRequest(data interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindJSON(data); err != nil {
			errs, ok := err.(validator.ValidationErrors)
			if !ok {
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": err.Error(),
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": errs.Translate(utils.Trans),
			})
			c.Abort()
			return
		}

		c.Set("reqBody", data)
	}
}
