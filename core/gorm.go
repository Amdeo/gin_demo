package core

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"server/global"
)

func InitDB() (db *gorm.DB) {
	// 连接数据库
	cfg := global.GConfig.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Dbname)
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}
	return DB
}

func Tables(db *gorm.DB) {
	err := db.AutoMigrate(
	// model
	)
	if err != nil {
		zap.L().Error("register table failed", zap.Any("err", err))
		os.Exit(0)
	}
	zap.L().Info("register table success")
}
