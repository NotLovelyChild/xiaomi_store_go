package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GoodsTypeController struct {
    BaseController
}

func (g GoodsTypeController) Index (c *gin.Context){
	c.HTML(http.StatusOK, "admin/goods_type/index.html", gin.H{})
}

func (g GoodsTypeController) Add (c *gin.Context){
    c.HTML(http.StatusOK, "admin/goods_type/add.html", gin.H{})
}
