package main

import (
	"xiaomi_store/config"
	"xiaomi_store/routers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
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
	redis_addr := config.Config.Section("redis").Key("host").String() + ":" + config.Config.Section("redis").Key("port").String()
	store, _ := redis.NewStore(10, "tcp", redis_addr, config.Config.Section("redis").Key("password").String(), []byte("secret"))
	sessionsExpires, _ := config.Config.Section("session").Key("expires").Int()
	store.Options(sessions.Options{
		MaxAge:   sessionsExpires,
		HttpOnly: true,
	})
	r.Use(sessions.Sessions("mysession", store))

	// 路由
	routers.AdminRouters(r)

	r.Run(":8082")
}
