package db

import (
	"api-public-platform/config"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MySQLDB *gorm.DB
var err error

func ConnectDatabase(cfg *config.ServerConfig) (*gorm.DB, error) {
	DSN := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DataSource.Username,
		cfg.DataSource.Password,
		cfg.DataSource.Host,
		cfg.DataSource.Port,
		cfg.DataSource.Database)
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
