package admin

import (
	"fmt"
	"net/http"
	"strconv"
	"xiaomi_store/mysql/xiaomi"
	"xiaomi_store/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ManagerController struct {
	BaseController
}

func (m ManagerController) Index(c *gin.Context) {
	// 获取所有管理员
	managers, err := xiaomi.FindAllManager()
	if err != nil {
		m.Fail(c, "获取管理员列表失败", "/admin/index")
		return
	}
	fmt.Println(managers)
	c.HTML(http.StatusOK, "admin/manager/index.html", gin.H{
		"managers": managers,
	})
}

func (m ManagerController) Add(c *gin.Context) {
	roles, err := xiaomi.FindALLRoles()
	if err != nil {
		m.Fail(c, "获取角色列表失败", "/admin/manager/index")
		return
	}
	c.HTML(http.StatusOK, "admin/manager/add.html", gin.H{
		"roleList": roles,
	})
}

func (m ManagerController) DoAdd(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")
	mobile := c.PostForm("mobile")
	roleId := c.PostForm("role_id")

	manager := &xiaomi.Manager{
		Username: username,
		Email:    email,
		Mobile:   mobile,
	}
	// 判断用户名密码长度
	if len(username) < 2 || len(password) < 6 {
		m.Fail(c, "用户名密码长度不符合要求", "/admin/manager/add")
		return
	}
	// 查重
	err := manager.FindWithUserName()
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			m.Fail(c, "查询失败"+err.Error(), "/admin/manager/add")
			return
		}
	}
	if manager.ID > 0 {
		m.Fail(c, "该用户名已存在", "/admin/manager/add")
		return
	}
	// 判断手机号是否符合要求
	if len(mobile) != 11 {
		m.Fail(c, "手机号不符合要求", "/admin/manager/add")
		return
	}
	// RoleID 判断
	roleIdInt, err := strconv.Atoi(roleId)
	if err != nil {
		m.Fail(c, "角色ID不符合要求", "/admin/manager/add")
		return
	}
	role := xiaomi.Role{
		ID: int64(roleIdInt),
	}
	err = role.FindByID()
	if err != nil {
		m.Fail(c, "角色ID不存在", "/admin/manager/add")
		return
	}
	// 添加管理员
	manager.RoleID = role.ID
	manager.Password = utils.MD5(password)
	err = manager.Create()
	if err != nil {
		m.Fail(c, "添加管理员失败"+err.Error(), "/admin/manager/add")
		return
	}
	m.Success(c, "添加管理员成功", "/admin/manager")
}

func (m ManagerController) Edit(c *gin.Context) {
	managerID := c.Query("id")
	managerIDInt, err := strconv.Atoi(managerID)
	if err != nil {
		m.Fail(c, "管理员ID不符合要求"+err.Error(), "/admin/manager")
		return
	}
	manager := &xiaomi.Manager{
		ID: int64(managerIDInt),
	}
	err = manager.FindWithID()
	if err != nil {
		m.Fail(c, "管理员ID不存在"+err.Error(), "/admin/manager")
		return
	}
	roles, err := xiaomi.FindALLRoles()
	if err != nil {
		m.Fail(c, "获取角色列表失败", "/admin/manager")
		return
	}
	c.HTML(http.StatusOK, "admin/manager/edit.html", gin.H{
		"manager":  manager,
		"roleList": roles,
	})
}

func (m ManagerController) DoEdit(c *gin.Context) {
	managerID := c.Query("id")
	managerIDInt, err := strconv.Atoi(managerID)
	if err != nil {
		fmt.Println(err, managerID)
		m.Fail(c, "管理员ID不符合要求"+err.Error(), "/admin/manager")
		return
	}
	manager := &xiaomi.Manager{
		ID: int64(managerIDInt),
	}
	err = manager.FindWithID()
	if err != nil {
		m.Fail(c, "管理员ID不存在"+err.Error(), "/admin/manager")
		return
	}
	password := c.PostForm("password")
	if len(password) > 0 {
		manager.Password = utils.MD5(password)
	}
	email := c.PostForm("email")
	mobile := c.PostForm("mobile")
	roleId := c.PostForm("role_id")

	roleIdInt, err := strconv.Atoi(roleId)
	if err != nil {
		m.Fail(c, "角色ID不符合要求", "/admin/manager")
		return
	}
	role := &xiaomi.Role{
		ID: int64(roleIdInt),
	}
	err = role.FindByID()
	if err != nil {
		m.Fail(c, "角色ID不存在", "/admin/manager")
		return
	}
	manager.Email = email
	manager.Mobile = mobile
	manager.RoleID = role.ID
	err = manager.Update()
	if err != nil {
		m.Fail(c, "更新管理员失败"+err.Error(), "/admin/manager")
		return
	}
	m.Success(c, "更新管理员成功", "/admin/manager")
}

func (m ManagerController) Delete(c *gin.Context) {
	managerID := c.Query("id")
	managerIDInt, err := strconv.Atoi(managerID)
	if err != nil {
		fmt.Println(err, managerID)
		m.Fail(c, "管理员ID不符合要求"+err.Error(), "/admin/manager")
		return
	}
	manager := &xiaomi.Manager{
		ID: int64(managerIDInt),
	}
	err = manager.FindWithID()
	if err != nil {
		m.Fail(c, "管理员ID不存在"+err.Error(), "/admin/manager")
		return
	}

	err = manager.Delete()
	if err != nil {
		m.Fail(c, "删除管理员失败"+err.Error(), "/admin/manager")
		return
	}
	m.Success(c, "删除管理员成功", "/admin/manager")
}
