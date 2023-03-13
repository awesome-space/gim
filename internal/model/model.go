package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint           `gorm:"primaryKey;comment:主键"` // 主键
	CreatedAt time.Time      `gorm:"comment:创建时间"`          // 创建时间
	UpdatedAt time.Time      `gorm:"comment:更新时间"`          // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index;comment:删除时间"`    // 删除时间
}
