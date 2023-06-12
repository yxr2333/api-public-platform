package db

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MySQLDB *gorm.DB
var err error
var DSN = "root:123456@tcp(127.0.0.1:3306)/mydb3?charset=utf8mb4&parseTime=True&loc=Local"

func ConnectDatabase() (*gorm.DB, error) {
	MySQLDB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               DSN,
		DefaultStringSize: 255,
	}))
	if err != nil {
		return nil, err
	}
	sqlDB, err := MySQLDB.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetConnMaxIdleTime(time.Minute * 5)
	sqlDB.SetMaxIdleConns(10)

	return MySQLDB, nil
}
