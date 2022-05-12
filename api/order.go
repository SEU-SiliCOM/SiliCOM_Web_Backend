package api

import (
	"SilicomAPPv0.3/models"
	"SilicomAPPv0.3/response"
	"SilicomAPPv0.3/service"
	"github.com/gin-gonic/gin"
)

var webOrder service.WebOrderService
var appOrder service.AppOrderService

// WebDeleteOrder 后台管理前端，删除订单
func WebDeleteOrder(c *gin.Context) {
	var param models.WebOrderDeleteParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := webOrder.Delete(param); count > 0 {
		response.Success("删除成功", count, c)
		return
	}
	response.Failed("删除失败", c)
}

// WebUpdateOrder 后台管理前端，更新订单
func WebUpdateOrder(c *gin.Context) {
	var param models.WebOrderUpdateParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := webOrder.Update(param); count > 0 {
		response.Success("更新成功", count, c)
		return
	}
	response.Failed("更新失败", c)
}

// WebGetOrderList 后台管理前端，获取订单列表
func WebGetOrderList(c *gin.Context) {
	var param models.WebOrderListParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	productList, rows := webOrder.GetList(param)
	response.SuccessPage("查询成功", productList, rows, c)
}

// WebGetOrderDetail 后台管理前端，获取订单详情
func WebGetOrderDetail(c *gin.Context) {
	var param models.WebOrderDetailParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	productDetail := webOrder.GetDetail(param)
	response.Success("查询成功", productDetail, c)
}

// AppCreateOrder 微信小程序，提交订单
func AppCreateOrder(c *gin.Context) {
	var param models.AppOrderSubmitParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := appOrder.Submit(param); count > 0 {
		response.Success("提交成功", count, c)
		return
	}
	response.Failed("提交失败", c)
}

// AppGetOrderList 微信小程序，获取订单列表信息
func AppGetOrderList(c *gin.Context) {
	var param models.AppOrderQueryParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	orderList := appOrder.GetList(param)
	response.Success("查询成功", orderList, c)
}
