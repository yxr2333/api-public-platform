package middlewares

import (
	"api-public-platform/pkg/security"
	"api-public-platform/pkg/service"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var roleService service.RoleService = service.NewRoleService()

func Authorization(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roles := strings.Split(requiredRole, "|")
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "Unauthorized",
			})
			c.Abort()
			return
		}

		jwtService := security.NewJWTService()
		tokenStr, err := jwtService.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "Invalid token",
			})
			c.Abort()
			return
		}
		claims, ok := tokenStr.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "Invalid token",
			})
			c.Abort()
			return
		}
		userId, ok := claims["userId"].(float64)
		if !ok {
			fmt.Printf("userId: %v\n", claims["userId"])
			c.JSON(http.StatusForbidden, gin.H{
				"msg": "You don't have permission to access this resource",
			})
			c.Abort()
			return
		}
		role, err := roleService.GetRoleByUserId(uint(userId))
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"msg": "You don't have permission to access this resource",
			})
			c.Abort()
			return
		}
		for _, r := range roles {
			if r == role.RoleName {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "You don't have permission to access this resource",
		})
		c.Abort()
	}
}
