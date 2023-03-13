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

// FetchUserList 获取用户列表
func (u userService) FetchUserList() ([]*model.User, error) {
	var list []*model.User
	if tx := global.DB.Find(&list); tx.RowsAffected == 0 {
		return nil, errors.New(fmt.Sprintf("获取用户列表失败 %v", tx.Error))
	}
	return list, nil
}

// 通过用户名和密码查找用户
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
