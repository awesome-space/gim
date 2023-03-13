package service

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/poeticalcode/gim/internal/global"
	"github.com/poeticalcode/gim/internal/model"
	"github.com/poeticalcode/gim/internal/tool"
)

type userService struct {
}

// FindUserList 获取用户列表
func (u userService) FindUserList() ([]*model.User, error) {
	var list []*model.User
	if tx := global.DB.Find(&list); tx.RowsAffected == 0 {
		return nil, errors.New(fmt.Sprintf("获取用户列表失败 %v", tx.Error))
	}
	return list, nil
}

// FindUserByUserNameAndPwd 通过用户名和密码查找用户
func (u userService) FindUserByUserNameAndPwd(userName string, password string) (*model.User, error) {
	user := model.User{}
	if tx := global.DB.Where("user_name = ? and pass_word=?", userName, password).First(&user); tx.RowsAffected == 0 {
		return nil, errors.New(fmt.Sprintf("未查询到记录 %v", tx.Error))
	}
	// token 加密
	t := strconv.Itoa(int(time.Now().Unix()))
	temp := tool.Md5encoder(t)
	if tx := global.DB.Model(&user).Where("id = ?", user.ID).Update("identity", temp); tx.RowsAffected == 0 {
		return nil, errors.New(fmt.Sprintf("写入identity失败 %v", tx.Error))
	}
	return &user, nil
}

// CheckUserExistByUserName 根据用户名判断用户是否存在
func (u userService) CheckUserExistByUserName(userName string) bool {
	if tx := global.DB.Select("id").Where("user_name = ?", userName); tx.RowsAffected == 0 {
		return false
	}
	return true
}

// createUser 创建用户
func (u userService) createUser(user *model.User) (*model.User, error) {
	if tx := global.DB.Save(user); tx.RowsAffected == 0 {
		return nil, tx.Error
	}
	return user, nil
}
