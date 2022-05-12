package api

import (
	"SilicomAPPv0.3/models"
	"SilicomAPPv0.3/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 权限管理
func ManageUser(c *gin.Context) {
	username := fmt.Sprint(c.Keys["username"])
	if u.GetUserPermission(username) < 2 {
		response.Failed("权限不足", c)
		return
	}
	var param models.UserPermissionParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := u.PermissionUpdate(param); count > 0 {
		response.Success("修改成功", count, c)
		return
	}
	response.Failed("修改失败", c)
}

// 拉黑用户
func BlackUser(c *gin.Context) {
	username := fmt.Sprint(c.Keys["username"])
	if u.GetUserPermission(username) < 1 {
		response.Failed("权限不足", c)
		return
	}
	var param models.UserPermissionParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	param.Permission = -1
	if count := u.PermissionUpdate(param); count > 0 {
		response.Success("拉黑成功", count, c)
		return
	}
	response.Failed("拉黑失败", c)
}

// 用户查询
func GetUserInfo(c *gin.Context) {
	username := fmt.Sprint(c.Keys["username"])
	if u.GetUserPermission(username) < 1 {
		response.Failed("权限不足", c)
		return
	}
	var param models.GetUserInfoParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if userInfo, rows := u.GetInfo(param); rows > 0 {
		response.Success("查询成功", userInfo, c)
		return
	}
	response.Failed("查询失败", c)
}

// 用户列表
func GetUserList(c *gin.Context) {
	username := fmt.Sprint(c.Keys["username"])
	if u.GetUserPermission(username) < 1 {
		response.Failed("权限不足", c)
		return
	}
	var param models.UserQueryParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	userList, rows := u.GetList(param)
	response.SuccessPage("查询成功", userList, rows, c)

}
