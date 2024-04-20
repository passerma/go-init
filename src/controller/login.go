package controller

import (
	"go-init/src/util"

	"github.com/gin-gonic/gin"
)

// {功能} 登录
// {参数} 路由
// {返回} 无
func GetLogin(ctx *gin.Context) {
	var res map[string]string
	ctx.JSON(200, util.SendSusModel(res))
}
