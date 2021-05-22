package core

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
	"server/apps/Rc/model"
	g "server/global"
)

func InitDB() (db *gorm.DB) {
	// 连接数据库
	cfg := g.Config.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Dbname)
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user`
		},
	})
	if err != nil {
		return nil
	}
	return DB
}

func MigrateTables(db *gorm.DB) {
	err := db.AutoMigrate(
		// model
		model.User{},
	)
	if err != nil {
		zap.L().Error("register table failed", zap.Any("err", err))
		os.Exit(0)
	}
	zap.L().Info("register table success")
}
