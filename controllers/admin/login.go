package admin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"xiaomi_store/models"
	"xiaomi_store/mysql/xiaomi"
	"xiaomi_store/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	BaseController
}

func (l LoginController) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login/login.html", gin.H{})
}

func (l LoginController) Captcha(c *gin.Context) {
	// 生成验证码
	id, b64s, err := models.MakeCaptcha()
	if err != nil {
		c.String(http.StatusBadRequest, "验证码生成失败")
		return
	}
	// 返回验证码图片给前端
	c.JSON(http.StatusOK, gin.H{
		"code":           200,
		"captcha_id":     id,
		"captcha_base64": b64s,
	})
}

func (l LoginController) LoginPost(c *gin.Context) {
	captchaId := c.PostForm("captchaId")
	verifyValue := c.PostForm("verifyValue")
	manager := &xiaomi.Manager{
		Username: c.PostForm("username"),
		Password: utils.MD5(c.PostForm("password")),
	}
	if models.VerifyCaptcha(captchaId, verifyValue) {
		// 验证码正确
		err := manager.FindWithUserNameAndPassWord()
		if err != nil {
			fmt.Println("mysql error:", err)
			l.Fail(c, "用户名或密码错误", "/admin/login")
			return
		}
		if manager.ID == 0 {
			fmt.Println("mysql not find", manager)
			l.Fail(c, "用户名或密码错误", "/admin/login")
			return
		}
		fmt.Println("login success:", manager)
		jsonData, _ := json.Marshal(manager)
		sessions := sessions.Default(c)
		sessions.Set(models.ManagerSessionName, string(jsonData))
		sessions.Save()
		l.Success(c, "验证通过", "/admin")
	} else {
		// 验证码错误
		l.Fail(c, "验证码错误", "/admin/login")
	}
}

func (l LoginController) Logout(c *gin.Context) {
	sessions := sessions.Default(c)
	// user := sessions.Get(models.ManagerSessionName)
	sessions.Delete(models.ManagerSessionName)
	sessions.Save()
	l.Success(c, "退出登录", "/admin/login")
}
