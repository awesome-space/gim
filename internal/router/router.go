package router

import (
	"github.com/gin-gonic/gin"
	"github.com/poeticalcode/gim/internal/middleware"
)

var (
	userRouter = userRouterGroup{}
)

// RouterRegister 路由注册
func RouterRegister(engine *gin.Engine) *gin.Engine {
	//v1版本
	v1 := engine.Group("v1")
	// 用户相关信息均需要鉴权
	user := v1.Group("user", middleware.JWTAuth())
	{
		user.GET("/list", userRouter.List)
	}
	return engine
}
