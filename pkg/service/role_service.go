package service

import (
	"api-public-platform/internal/db"
	"api-public-platform/pkg/model"
)

type RoleService interface {
	GetRoleById(roleId uint) (model.Role, error)
	GetRoleByUserId(userId uint) (model.Role, error)
}

func NewRoleService() RoleService {
	return &RoleServiceImpl{}
}

type RoleServiceImpl struct{}

func (rs *RoleServiceImpl) GetRoleById(roleId uint) (model.Role, error) {
	var role model.Role
	err := db.MySQLDB.First(&role, roleId).Error
	if err != nil {
		return role, err
	}
	return role, nil
}

func (rs *RoleServiceImpl) GetRoleByUserId(userId uint) (model.Role, error) {
	var role model.Role
	err := db.MySQLDB.Table("users").Select("roles.*").Joins("left join roles on users.role_id = roles.id").Where("users.id = ?", userId).Scan(&role).Error
	if err != nil {
		return role, err
	}
	return role, nil
}
