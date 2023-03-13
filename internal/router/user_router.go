package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/poeticalcode/gim/internal/service"
)

// 用户相关路由组
type userRouterGroup struct{}

// List 用户列表路由
func (u userRouterGroup) List(ctx *gin.Context) {
	list, err := service.UserService.FetchUserList()
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    -1, // 0 表示成功， -1 表示失败
			"message": "获取用户列表失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, list)
}
