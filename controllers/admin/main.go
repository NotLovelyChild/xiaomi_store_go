package admin

import (
	"net/http"
	"xiaomi_store/mysql/xiaomi"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
	BaseController
}

func (m AdminController) Index(c *gin.Context) {
	// 权限判断
	user, ok := c.Get("user")
	if !ok {
		m.Fail(c, "用户不存在", "/admin/login")
		return
	}
	userInfo, ok := user.(xiaomi.Manager)
	if !ok {
		m.Fail(c, "用户未找到", "/admin/login")
		return
	}
	c.HTML(http.StatusOK, "admin/main/index.html", gin.H{
		"userName": userInfo.Username,
	})
}

func (m AdminController) Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/main/welcome.html", gin.H{})
}
