package api

import (
	"SilicomAPPv0.3/models"
	"SilicomAPPv0.3/response"
	"SilicomAPPv0.3/service"
	"github.com/gin-gonic/gin"
)

var cart service.CartService

// AppAddCart 微信小程序，添加购物车
func AppAddCart(c *gin.Context) {
	var param models.AppCartAddParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if cart.Add(param) {
		response.Success("添加成功", 1, c)
		return
	}
	response.Failed("已添加", c)
}

// AppDeleteCart 微信小程序，删除购物车中的商品
func AppDeleteCart(c *gin.Context) {
	var param models.AppCartDeleteParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := cart.Delete(param); count > 0 {
		response.Success("删除成功", count, c)
		return
	}
	response.Failed("删除失败", c)
}

// AppDeleteCart 微信小程序，清除购物车中的商品
func AppClearCart(c *gin.Context) {
	var param models.AppCartClearParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := cart.Clear(param); count > 0 {
		response.Success("清除成功", count, c)
		return
	}
	response.Failed("清除失败", c)
}

// AppGetCartInfo 微信小程序，获取购物车中信息
func AppGetCartInfo(c *gin.Context) {
	var param models.AppCartQueryParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	info := cart.GetInfo(param)
	if len(info.CartItem) == 0 {
		response.Success("购物车竟然是空的", info, c)
	}
	response.Success("查询成功", info, c)
}
