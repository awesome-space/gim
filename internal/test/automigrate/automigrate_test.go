package automigrate

import (
	"testing"

	"github.com/poeticalcode/gim/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	dsn := "root:he..123456@tcp(150.158.95.91:3306)/gim?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	return db
}

// 测试用户表的自动迁移
func TestUserAutoMiGrate(t *testing.T) {
	db := initDB()
	if db == nil {
		t.Fatal("db == nil")
	}
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.User{})
}
