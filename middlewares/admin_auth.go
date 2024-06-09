package middlewares

import "github.com/gin-gonic/gin"

func AdminAuth(c *gin.Context) {
	c.Next()
}
