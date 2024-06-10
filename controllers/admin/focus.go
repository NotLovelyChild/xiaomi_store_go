package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FocusController struct {
	BaseController
}

func (f FocusController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/focus/index.html", gin.H{})
}

func (f FocusController) Add(c *gin.Context) {
    c.HTML(http.StatusOK, "admin/focus/add.html", gin.H{})
}
