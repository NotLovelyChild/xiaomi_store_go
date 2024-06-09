package main

import (
	"xiaomi_store/routers"

	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
	// 加载模板
	r.LoadHTMLGlob("templates/**/**/*")
	//配置静态web目录   第一个参数表示路由, 第二个参数表示映射的目录
	r.Static("/static", "./static")
	// 中间件
	// cookie 引擎
	// 路由
	routers.AdminRouters(r)

	r.Run(":8082")
}
