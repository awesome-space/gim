package config

import "fmt"

// MySQLConfig mysql信息配置 MySQL
type MySQLConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"name" json:"Name"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}

// Dsn 获取 MySQL 的连接信息
func (s MySQLConfig) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", s.User,
		s.Password, s.Host, s.Port, s.Name)
}
