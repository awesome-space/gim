package main

import (
	"github.com/gin-gonic/gin"
	"github.com/poeticalcode/gim/internal/initialize"
	"github.com/poeticalcode/gim/internal/router"
)

func main() {
	// 初始化日志
	initialize.InitLogger()
	// 初始化配置
	initialize.InitConfig()
	// 初始化数据库
	initialize.InitDB()

	server := gin.Default()
	router.RouterRegister(server)
	server.Run(":9090")
}
