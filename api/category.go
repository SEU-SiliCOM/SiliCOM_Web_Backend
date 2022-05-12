package api

import (
	"SilicomAPPv0.3/models"
	"SilicomAPPv0.3/response"
	"SilicomAPPv0.3/service"
	"github.com/gin-gonic/gin"
)

var webCategory service.WebCategoryService
var appCategory service.AppCategoryService

// 后台管理前端，创建类目
func WebCreateCategory(c *gin.Context) {
	var param models.WebCategoryCreateParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := webCategory.Create(param); count > 0 {
		response.Success("创建成功", count, c)
		return
	}
	response.Failed("创建失败", c)
}

// 后台管理前端，删除类目
func WebDeleteCategory(c *gin.Context) {
	var param models.WebCategoryDeleteParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := webCategory.Delete(param); count > 0 {
		response.Success("删除成功", count, c)
		return
	}
	response.Failed("删除失败", c)
}

// 后台管理前端，更新类目
func WebUpdateCategory(c *gin.Context) {
	var param models.WebCategoryUpdateParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := webCategory.Update(param); count > 0 {
		response.Success("更新成功", count, c)
		return
	}
	response.Failed("更新失败", c)
}

// WebGetCategoryList 后台管理前端，获取类目列表
func WebGetCategoryList(c *gin.Context) {
	var param models.WebCategoryQueryParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	cateList, rows := webCategory.GetList(param)
	response.SuccessPage("查询成功", cateList, rows, c)
}

// WebGetCategoryOption 后台管理前端，获取类目选项
func WebGetCategoryOption(c *gin.Context) {
	option := webCategory.GetOption()
	response.Success("查询成功", option, c)
}

// AppGetCategoryOption 微信小程序，获取类目选项
func AppGetCategoryOption(c *gin.Context) {
	option := appCategory.GetOption()
	response.Success("查询成功", option, c)
}
