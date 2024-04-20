package route

import (
	"go-init/src/conf"
	"go-init/src/log"
	"go-init/src/middleware"
	"go-init/src/util"

	"github.com/gin-gonic/gin"
)

// {功能} 设置中间件
// {参数} 路由实例
// {返回} 无
func setMiddleware(r *gin.Engine) {
	if util.IsDEV() {
		r.Use(middleware.UseCors)
	}
	r.Use(middleware.UseLog)
}

// {功能} 生成路由
// {参数} 无
// {返回} 无
func Init() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.ForwardedByClientIP = true

	setMiddleware(router)

	generateLoginRouter(router)

	port := conf.GetConf("port", "4006")

	log.ComLoggerFmt("server start at port ", port)

	if err := router.Run(":" + port); err != nil {
		panic(err)
	}
}
