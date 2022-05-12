package api

import (
	"SilicomAPPv0.3/models"
	"SilicomAPPv0.3/response"
	"SilicomAPPv0.3/service"
	"github.com/gin-gonic/gin"
)

var browse service.BrowseService

// AppSaveBrowseRecord 微信小程序，保存浏览记录
func AppSaveBrowseRecord(c *gin.Context) {
	var param models.AppBrowseSaveParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := browse.Save(param); count > 0 {
		response.Success("保存成功", count, c)
		return
	}
	response.Failed("保存失败", c)
}

// AppDeleteBrowseRecord 微信小程序，删除浏览记录
func AppDeleteBrowseRecord(c *gin.Context) {
	var param models.AppBrowseDeleteParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := browse.Delete(param); count > 0 {
		response.Success("删除成功", count, c)
		return
	}
	response.Failed("删除失败", c)
}

// AppGetBrowseRecordList 微信小程序，获取浏览记录列表
func AppGetBrowseRecordList(c *gin.Context) {
	var param models.AppBrowseQueryParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	browseRecordList := browse.GetList(param)
	response.Success("查询成功", browseRecordList, c)
}
