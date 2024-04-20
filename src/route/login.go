package route

import (
	"go-init/src/controller"

	"github.com/gin-gonic/gin"
)

// {功能} 生成登录路由
// {参数} 路由实例
// {返回} 无
func generateLoginRouter(router *gin.Engine) {
	userG := router.Group("user")

	userG.GET("login", controller.GetLogin)
}
