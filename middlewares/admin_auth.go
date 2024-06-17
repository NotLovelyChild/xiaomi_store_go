package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"
	"xiaomi_store/models"
	"xiaomi_store/mysql/xiaomi"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var filtrationPath = []string{
	"/admin/login",
	"/admin/doLogin",
	"/admin/captcha",
	"/admin/logout",
}

func AdminAuth(c *gin.Context) {
	// 判断用户是否登录
	sessions := sessions.Default(c)
	user := sessions.Get(models.ManagerSessionName)
	userString, ok := user.(string)
	if !ok {
		fmt.Println("session 转换失败")
		gotoLogin(c)
		return
	}
	var userInfo xiaomi.Manager
	err := json.Unmarshal([]byte(userString), &userInfo)
	if err != nil {
		fmt.Println("sessionJson 转换失败", err, userString)
		gotoLogin(c)
		return
	}
	c.Set("user", userInfo)
}

func gotoLogin(c *gin.Context) {
	path := c.Request.URL.Path
	fmt.Println("-----------", path)
	for _, v := range filtrationPath {
		if path == v {
			return
		}
	}
	c.HTML(http.StatusOK, "admin/public/fail.html", gin.H{
		"message":     "请先登录",
		"redirectURL": "/admin/login",
	})
	c.Abort()
}
