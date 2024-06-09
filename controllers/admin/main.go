package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
	BaseController
}

func (m AdminController) Index(c *gin.Context) {
	// 权限判断
	c.HTML(http.StatusOK, "admin/main/index.html", gin.H{})
}

func (m AdminController) Welcome(c *gin.Context){
	c.HTML(http.StatusOK, "admin/main/welcome.html", gin.H{})
}
