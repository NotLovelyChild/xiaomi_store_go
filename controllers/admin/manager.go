package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ManagerController struct {
	BaseController
}

func (m ManagerController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/manager/index.html", gin.H{})
}

func (m ManagerController) Add(c *gin.Context) {
    c.HTML(http.StatusOK, "admin/manager/add.html", gin.H{})
}
