package middlewares

import (
	"api-public-platform/config"
	"api-public-platform/internal/db"
	"api-public-platform/pkg/model"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func validateAPIToken(apiToken string, apiEndpoint string, requestMethod string) (bool, error) {
	var user model.User
	if err := db.MySQLDB.Preload("Permissions").Where("api_token = ?", apiToken).First(&user).Error; err != nil {
		return false, fmt.Errorf("invalid api token")
	}
	// 假设每个API都有一个唯一的endpoint和请求方法method组合
	var api model.API
	if err := db.MySQLDB.Where("api_endpoint = ? AND request_method = ?", apiEndpoint, requestMethod).First(&api).Error; err != nil {
		return false, fmt.Errorf("invalid api endpoint")
	}

	// 首先检查用户是否有这个API的权限
	for _, permission := range user.Permissions {
		if permission.APIID == api.ID {
			return true, nil
		}
	}

	// 如果用户没有直接的权限，那就需要检查用户的角色是否有权限访问这个API
	var role model.Role
	if err := db.MySQLDB.Where("id = ?", user.RoleID).First(&role).Error; err != nil {
		// 如果查询出错或者没有找到匹配的角色，那么就说明他没有权限访问这个API
		return false, fmt.Errorf("invalid role")
	}

	// 检查用户的角色是否有这个API的权限
	for _, permission := range role.Permissions {
		if permission.APIID == api.ID {
			return true, nil
		}
	}

	return false, fmt.Errorf("you don't have permission to access this resource")

}

func ValidateAPITokenAndPermissions() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("t")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "API Token required",
			})
			c.Abort()
			return
		}
		urlPath := c.Request.URL.Path
		endpoint := strings.TrimPrefix(urlPath, config.ServerCfg.API.Outer.Prefix)
		// 检查APIToken是否有效
		isValid, err := validateAPIToken(token, endpoint, c.Request.Method)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}
		if !isValid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid api token",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
