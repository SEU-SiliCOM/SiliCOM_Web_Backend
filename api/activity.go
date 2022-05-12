package api

import (
	"SilicomAPPv0.3/models"
	"SilicomAPPv0.3/response"
	"SilicomAPPv0.3/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

var activity service.ActivityService

// 后台管理前端，创建活动
func CreateActivity(c *gin.Context) {

	username := fmt.Sprint(c.Keys["username"])
	if u.GetUserPermission(username) < 1 {
		response.Failed("权限不足", c)
		return
	}
	var param models.ActivityAddParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}

	if count := activity.Create(param); count > 0 {
		response.Success("创建成功", count, c)
		return
	}
	response.Failed("创建失败", c)
}

// 后台管理前端，删除活动
func DeleteActivity(c *gin.Context) {
	username := fmt.Sprint(c.Keys["username"])
	if u.GetUserPermission(username) < 1 {
		response.Failed("权限不足", c)
		return
	}
	var param models.ActivityDeleteParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := activity.Delete(param); count > 0 {
		response.Success("删除成功", count, c)
		return
	}
	response.Failed("删除失败", c)
}

// 后台管理前端，更新活动
func UpdateActivity(c *gin.Context) {
	username := fmt.Sprint(c.Keys["username"])
	if u.GetUserPermission(username) < 1 {
		response.Failed("权限不足", c)
		return
	}
	var param models.ActivityUpdateParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := activity.Update(param); count > 0 {
		response.Success("更新成功", count, c)
		return
	}
	response.Failed("更新失败", c)
}

// 后台管理前端，获取活动列表
func GetActivityList(c *gin.Context) {
	var param models.ActivityQueryParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	actList, rows := activity.GetList(param)
	response.SuccessPage("查询成功", actList, rows, c)
}

// 后台管理前端，处理签到结果，更新用户信息
func CompleteActivity(c *gin.Context) {
	username := fmt.Sprint(c.Keys["username"])
	if u.GetUserPermission(username) < 1 {
		response.Failed("权限不足", c)
		return
	}

	var param models.ActivityCompleteParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if activity.GetActivityStatus(param.Id) == 1 {
		response.Failed("活动已完成", c)
		return
	}
	if count := activity.Complete(param); count > 0 {
		response.Success("完成成功", count, c)
		return
	}
	response.Failed("完成失败", c)
}
