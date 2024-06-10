package admin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"xiaomi_store/models"
	"xiaomi_store/mysql/xiaomi"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AdminController struct {
	BaseController
}

func (m AdminController) Index(c *gin.Context) {
	// 权限判断
	// 判断用户是否登录
	sessions := sessions.Default(c)
	user := sessions.Get(models.ManagerSessionName)
	userString, ok := user.(string)
	if !ok {
		fmt.Println("session 转换失败")
		m.Fail(c, "请先登录", "/admin/login")
		return
	}
	var userInfo xiaomi.Manager
	err := json.Unmarshal([]byte(userString), &userInfo)
	if err != nil {
		fmt.Println("sessionJson 转换失败", err, userString)
		m.Fail(c, "请先登录", "/admin/login")
		return
	} 
	c.HTML(http.StatusOK, "admin/main/index.html", gin.H{
		"userName": userInfo.Username,
	})
}

func (m AdminController) Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/main/welcome.html", gin.H{})
}
