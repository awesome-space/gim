package initialize

import (
	"github.com/poeticalcode/gim/internal/global"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func InitConfig() {
	//实例化对象
	v := viper.New()
	configFile := "./config/config.yaml"
	//读取配置文件
	v.SetConfigFile(configFile)
	//读入文件
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	// 反序列化
	if err := v.Unmarshal(&global.ServerConfig); err != nil {
		panic(err)
	}
	zap.S().Info("配置信息", global.ServerConfig)
}
