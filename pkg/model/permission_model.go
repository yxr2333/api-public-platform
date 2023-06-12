package model

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	APIID                 uint   `gorm:"not null;unique_index"`
	API                   API    `gorm:"foreignkey:APIID"`
	PermissionDescription string `gorm:"type:varchar(255)"`
}
