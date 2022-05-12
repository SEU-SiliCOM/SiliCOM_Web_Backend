package api

import (
	"SilicomAPPv0.3/models"
	"SilicomAPPv0.3/response"
	"SilicomAPPv0.3/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

var appointment service.AppointmentService

// 后台管理前端，预约创建
func CreateAppointment(c *gin.Context) {
	username := fmt.Sprint(c.Keys["username"])
	permission := u.GetUserPermission(username)
	if permission < 0 {
		response.Failed("权限不足", c)
		return
	}
	if u.GetUserFreeze(username).After(time.Now()) {
		response.Failed("账户冻结中", c)
		return
	}
	var param models.AppointmentCreateParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if activity.GetActivityStatus(param.Aid) != 0 {
		response.Failed("预约已结束", c)
		return
	}

	if permission < 1 || param.Uid == 0 {
		param.Uid = u.GetUserUid(username) // 只有管理员可以设置uid，普通用户从token中获取
	}
	if count := appointment.Create(param); count > 0 {
		response.Success("创建成功", count, c)
		return
	}
	response.Failed("创建失败", c)
}

// 后台管理前端，预约更新
func UpdateAppointment(c *gin.Context) {
	username := fmt.Sprint(c.Keys["username"])
	var param models.AppointmentUpdateParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	param.Uid = u.GetUserUid(username)
	if count := appointment.Update(param); count > 0 {
		response.Success("更新成功", count, c)
		return
	}
	response.Failed("更新失败", c)
}

// 后台管理前端，预约删除
func DeleteAppointment(c *gin.Context) {
	var param models.AppointmentDeleteParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := appointment.Delete(param); count > 0 {
		response.Success("删除成功", count, c)
		return
	}
	response.Failed("删除失败", c)
}

// 后台管理前端，预约查询
func GetAppointmentList(c *gin.Context) {
	username := fmt.Sprint(c.Keys["username"])
	if u.GetUserPermission(username) < 1 {
		response.Failed("权限不足", c)
		return
	}
	var param models.AppointmentQueryParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	appList, rows := appointment.GetList(param)
	response.SuccessPage("查询成功", appList, rows, c)
}

// 后台管理前端，签到
func SignAppointment(c *gin.Context) {
	username := fmt.Sprint(c.Keys["username"])
	if u.GetUserPermission(username) < 1 {
		response.Failed("权限不足", c)
		return
	}
	var param models.AppointmentSignParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := appointment.Sign(param); count > 0 {
		response.Success("签到成功", count, c)
		return
	}
	response.Failed("签到失败", c)
}
