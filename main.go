package main

import (
	"github.com/gin-gonic/gin"
	"github.com/poeticalcode/gim/internal/initialize"
)

func main() {
	// 初始化日志
	initialize.InitLogger()
	// 初始化数据库
	initialize.InitDB()

	server := gin.Default()
	server.Run(":9090")
}
