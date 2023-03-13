package model

import (
	"time"
)

type User struct {
	BaseModel
	UserName      string     `gorm:"comment:用户名"`  // 用户名
	Password      string     `gorm:"comment:密码"`   // 密码
	Avatar        string     `gorm:"comment:头像"`   // 头像
	Gender        string     `gorm:"comment:性别"`   // 性别
	Phone         string     `gorm:"comment:手机号"`  // 手机号
	Email         string     `gorm:"comment:邮箱"`   // 邮箱
	OnlineState   bool       `gorm:"comment:在线状态"` // 在线状态
	DeviceInfo    string     `gorm:"comment:登录设备"` //	登录设备
	LoginTime     *time.Time `gorm:"comment:登录时间"` // 登录时间
	HeartBeatTime *time.Time `gorm:"comment:心跳时间"` // 心跳时间
	LoginOutTime  *time.Time `gorm:"comment:退出时间"` // 退出时间
}

// TableName 表名
func (u *User) TableName() string {
	return "im_user"
}
