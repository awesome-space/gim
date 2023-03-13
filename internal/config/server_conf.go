package config

// ServerConfig gim 服务相关配置
type ServerConfig struct {
	Port        int `mapstructure:"port" json:"port"`
	MySQLConfig `mapstructure:"mysql" json:"mysql"`
	RedisConfig `mapstructure:"redis" json:"redis"`
}
