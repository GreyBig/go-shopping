package initialize

import (
	"fmt"

	"shopping/common/global"
	"shopping/common/middleware"
	"shopping/controller"

	"github.com/gin-gonic/gin"
)

func Router() {

	engine := gin.Default()

	// 开启跨域
	engine.Use(middleware.Cors())

	// 静态资源请求映射
	engine.Static("/image", global.Config.Upload.SavePath)

	// 后台管理员前端接口
	web := engine.Group("/web")

	{
		// 用户登录API
		web.GET("/captcha", controller.WebGetCaptcha)
		web.POST("/login", controller.WebUserLogin)

		// 开启JWT认证,以下接口需要认证成功才能访问
		web.Use(middleware.JwtAuth())

	}

	// 启动、监听端口
	post := fmt.Sprintf(":%s", global.Config.Server.Post)
	if err := engine.Run(post); err != nil {
		fmt.Printf("server start error: %s", err)
	}
}
