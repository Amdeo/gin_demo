package core

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	g "server/global"
)

func initServer(router *gin.Engine) {
	//// 初始化Router
	if err := router.Run(":" + g.Config.System.Port); err != nil {
		panic(fmt.Sprintf("startup service failed, err:%v\n\n", err))
	}
}

func RunServer() {
	// 初始化Redis
	Redis()

	// 初始化DB
	g.DB = InitDB()
	if g.DB != nil {
		MigrateTables(g.DB)
		db, _ := g.DB.DB()
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				panic("close db failed")
			}
		}(db)
	}

	//// 初始化Router
	router := SetupRouter()
	initServer(router)
}
