package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func (b *BaseController) Prepare(c *gin.Context, redirectURL string) {
	
}

func (b *BaseController) Success(c *gin.Context, message, redirectURL string) {
	c.HTML(http.StatusOK, "admin/public/success.html", gin.H{
		"message": message,
		"redirectURL": redirectURL,
	})
}

func (b *BaseController) Fail(c *gin.Context, message, redirectURL string) {
	c.HTML(http.StatusOK, "admin/public/fail.html", gin.H{
		"message": message,
		"redirectURL": redirectURL,
	})
}
