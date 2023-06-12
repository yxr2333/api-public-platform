package middlewares

import (
	"api-public-platform/internal/db"
	"api-public-platform/pkg/model"
	"fmt"
)

func ValidateAPIToken(apiToken string, apiEndpoint string, requestMethod string) (bool, error) {
	var user model.User
	if err := db.MySQLDB.Where("api_token = ?", apiToken).First(&user).Error; err != nil {
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
