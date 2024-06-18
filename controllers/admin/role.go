package admin

import (
	"net/http"
	"strconv"
	"xiaomi_store/mysql/xiaomi"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	BaseController
}

func (r RoleController) Index(c *gin.Context) {
	// 获取所有角色
	roleList, err := xiaomi.FindALLRoles()
	if err != nil {
		r.Fail(c, "获取角色失败失败"+err.Error(), "/admin")
		return
	}
	c.HTML(http.StatusOK, "admin/role/index.html", gin.H{
		"roleList": roleList,
	})
}

func (r RoleController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/role/add.html", gin.H{})
}

func (r RoleController) DoAdd(c *gin.Context) {
	title := c.PostForm("title")
	desc := c.PostForm("description")
	if title == "" {
		r.Fail(c, "角色名不能为空", "/admin/role/add")
		return
	}
	role := &xiaomi.Role{
		Title:       title,
		Description: desc,
		Status:      true,
	}
	err := role.Create()
	if err != nil {
		r.Fail(c, "添加角色失败:"+err.Error(), "/admin/role/add")
		return
	}
	r.Success(c, "添加角色成功", "/admin/role")
}

func (r RoleController) Edit(c *gin.Context) {
	id := c.Query("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		r.Fail(c, "获取角色失败:"+err.Error(), "/admin/role")
		return
	}
	role := &xiaomi.Role{
		ID: idInt,
	}
	err = role.FindByID()
	if err != nil {
		r.Fail(c, "角色不存在:"+err.Error(), "/admin/role")
		return
	}
	c.HTML(http.StatusOK, "admin/role/edit.html", gin.H{
		"role": role,
	})
}

func (r RoleController) DoEdit(c *gin.Context) {
	id := c.Query("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		r.Fail(c, "获取角色失败:"+err.Error(), "/admin/role")
		return
	}
	title := c.PostForm("title")
	desc := c.PostForm("description")
	if title == "" {
		r.Fail(c, "角色名不能为空", "/admin/role/edit?id="+id)
		return
	}
	role := &xiaomi.Role{
		ID: idInt,
	}
	err = role.FindByID()
	if err != nil {
		r.Fail(c, "角色不存在:"+err.Error(), "/admin/role")
		return
	}
	role.Title = title
	role.Description = desc
	err = role.Update()
	if err != nil {
		r.Fail(c, "修改角色失败:"+err.Error(), "/admin/role/edit?id="+id)
		return
	}
	r.Success(c, "修改角色成功", "/admin/role")
}

func (r RoleController) Delete(c *gin.Context) {
	id := c.Query("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		r.Fail(c, "获取角色失败:"+err.Error(), "/admin/role")
		return
	}
	role := &xiaomi.Role{
		ID: idInt,
	}
	err = role.FindByID()
	if err != nil {
		r.Fail(c, "角色不存在:"+err.Error(), "/admin/role")
		return
	}
	err = role.Delete()
	if err != nil {
		r.Fail(c, "删除角色失败:"+err.Error(), "/admin/role")
		return
	}
	r.Success(c, "删除角色成功", "/admin/role")
}
