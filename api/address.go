package api

import (
	"SilicomAPPv0.3/models"
	"SilicomAPPv0.3/response"
	"SilicomAPPv0.3/service"
	"github.com/gin-gonic/gin"
)

var address service.AddressService

// AppAddAddress 微信小程序，添加地址
func AppAddAddress(c *gin.Context) {
	var param models.AppAddressAddParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := address.Add(param); count > 0 {
		response.Success("添加成功", count, c)
		return
	}
	response.Failed("添加失败", c)
}

// AppDeleteAddress 微信小程序，删除地址
func AppDeleteAddress(c *gin.Context) {
	var key models.AppAddressDeleteParam
	if err := c.ShouldBind(&key); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := address.Delete(key.AddressId); count > 0 {
		response.Success("删除成功", count, c)
		return
	}
	response.Failed("删除失败", c)
}

// AppUpdateAddress 微信小程序，更新地址
func AppUpdateAddress(c *gin.Context) {
	var param models.AppAddressUpdateParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	count := address.Update(param)
	if count > 0 {
		response.Success("更新成功", count, c)
		return
	}
	response.Failed("更新失败", c)
}

// AppGetAddressUpdateInfo 微信小程序，获取地址更新信息
func AppGetAddressUpdateInfo(c *gin.Context) {
	var key models.AppAddressInfoParam
	if err := c.ShouldBind(&key); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	info := address.GetInfo(key.AddressId)
	response.Success("查询成功", info, c)
}

// AppGetAddressList 微信小程序，获取地址列表
func AppGetAddressList(c *gin.Context) {
	var param models.AppAddressListParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	addressList := address.GetList(param.UserId)
	response.Success("查询成功", addressList, c)
}
