package model

import (
	"time"

	"gorm.io/gorm"
)

type API struct {
	gorm.Model
	APIDescription string     `gorm:"type:varchar(255)"`
	APIEndpoint    string     `gorm:"type:varchar(255);unique_index"`
	RequestMethod  string     `gorm:"type:varchar(50)"`
	IsOpen         bool       `gorm:"not null"` // 增加了一个字段，表示API是否开放
	CallCount      uint       `gorm:"not null"` // 增加了一个字段，表示API被调用的次数
	LastCalled     *time.Time // 增加了一个字段，表示API最后被调用的时间
}

type APICallHistory struct {
	gorm.Model
	APIID        uint      `gorm:"not null"`
	API          API       `gorm:"foreignkey:APIID"`
	CalledBy     uint      `gorm:"not null"`
	CalledByUser User      `gorm:"ForeignKey:CalledBy;AssociationForeignKey:ID"`
	CalledAt     time.Time `gorm:"not null"`
	CallStatus   string    `gorm:"type:varchar(50)"`
	CallResponse string    `gorm:"type:varchar(255)"`
}
