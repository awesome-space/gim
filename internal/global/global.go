package global

import (
	"github.com/poeticalcode/gim/internal/config"
	"gorm.io/gorm"
)

var (
	// 数据库连接
	DB *gorm.DB
	// 服务配置
	ServerConfig config.ServerConfig
)
