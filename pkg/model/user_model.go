package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName    string `gorm:"type:varchar(100);unique_index"`
	Password    string `gorm:"size:255"`
	Avatar      string `gorm:"type:varchar(255)"`
	Gender      string `gorm:"type:varchar(10)"`
	Email       string `gorm:"type:varchar(100);unique_index"`
	APIToken    string `gorm:"type:varchar(255);unique_index"`
	RoleID      uint
	Role        Role         `gorm:"foreignkey:RoleID"`
	Permissions []Permission `gorm:"many2many:user_permissions;"` // 增加了一个用户和权限的关联，方便直接控制用户的权限
}

type UserAPI struct {
	gorm.Model
	UserID   uint `gorm:"not null;foreignKey"`
	APIID    uint `gorm:"not null;foreignKey"`
	IsBanned bool `gorm:"not null"`
	User     User `gorm:"foreignKey:UserID;references:ID"`
	API      API  `gorm:"foreignKey:APIID;references:ID"`
}
