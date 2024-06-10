package routers

import (
	controller_admin "xiaomi_store/controllers/admin"
	"xiaomi_store/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminRouters(c *gin.Engine) {
	admin := c.Group("/admin", middlewares.AdminAuth)

	admin.GET("/", controller_admin.AdminController{}.Index)
	admin.GET("/welcome", controller_admin.AdminController{}.Welcome)

	admin.GET("/login", controller_admin.LoginController{}.Login)
	admin.GET("/captcha", controller_admin.LoginController{}.Captcha)
	admin.POST("/doLogin", controller_admin.LoginController{}.LoginPost)
	admin.GET("/logout", controller_admin.LoginController{}.Logout)

	admin.GET("/manager", controller_admin.ManagerController{}.Index)
	admin.GET("/manager/add", controller_admin.ManagerController{}.Add)

	admin.GET("/goods", controller_admin.GoodsController{}.Index)
	admin.GET("/goods/add", controller_admin.GoodsController{}.Add)

	admin.GET("/goods_type", controller_admin.GoodsTypeController{}.Index)
	admin.GET("/goods_type/add", controller_admin.GoodsTypeController{}.Add)

	admin.GET("/focus", controller_admin.FocusController{}.Index)
	admin.GET("/focus/add", controller_admin.FocusController{}.Add)
}
