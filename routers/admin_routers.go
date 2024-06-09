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
	admin.POST("/doLogin", controller_admin.LoginController{}.LoginPost)

	admin.GET("/manager", controller_admin.ManagerController{}.Index)
	admin.GET("/manager/add", controller_admin.ManagerController{}.Add)
}
