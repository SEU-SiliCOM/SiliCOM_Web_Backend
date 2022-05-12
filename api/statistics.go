package api

import (
	"SilicomAPPv0.3/response"
	"SilicomAPPv0.3/service"
	"github.com/gin-gonic/gin"
)

var webStatistics service.WebStatisticsService

// WebGetDataOverviewInfo 后台管理前端，获取数据总览统计信息
func WebGetDataOverviewInfo(c *gin.Context) {
	overviewInfo := webStatistics.GetDataOverviewInfo()
	response.Success("查询成功", overviewInfo, c)
}

// WebGetTodayOrderDataInfo 后台管理前端，获取今日订单数据统计信息
func WebGetTodayOrderDataInfo(c *gin.Context) {
	todayInfo := webStatistics.GetTodayDataInfo()
	response.Success("查询成功", todayInfo, c)
}

// WebGetWeekDataInfo 后台管理前端，获取本周数据统计信息
func WebGetWeekDataInfo(c *gin.Context) {
	weekInfo := webStatistics.GetWeekDataInfo()
	response.Success("查询成功", weekInfo, c)
}
