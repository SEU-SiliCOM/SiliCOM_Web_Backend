package service

import (
	"SilicomAPPv0.3/common"
	"SilicomAPPv0.3/global"
	"SilicomAPPv0.3/models"
	"strconv"
	"strings"
)

var cart CartService
var address AddressService

type WebOrderService struct {
}

type AppOrderService struct {
}

// Delete 后台管理前端，删除订单
func (o *WebOrderService) Delete(param models.WebOrderDeleteParam) int64 {
	return global.Db.Delete(&models.Order{}, param.Id).RowsAffected
}

// Update 后台管理前端，更新订单
func (o *WebOrderService) Update(param models.WebOrderUpdateParam) int64 {
	order := models.Order{
		Id:      param.Id,
		Status:  param.Status,
		Updated: common.NowTime(),
	}
	return global.Db.Model(&order).Updates(order).RowsAffected
}

// GetList 后台管理前端，获取订单列表
func (o *WebOrderService) GetList(param models.WebOrderListParam) ([]models.WebOrderList, int64) {
	orderList := make([]models.WebOrderList, 0)
	query := &models.Order{
		Id:     param.Id,
		Status: param.Status,
	}
	rows := common.RestPage(param.Page, "order", query, &orderList, &[]models.Order{})
	return orderList, rows
}

// GetDetail 后台管理前端，获取订单详情
func (o *WebOrderService) GetDetail(param models.WebOrderDetailParam) (od models.WebOrderDetail) {
	var order models.Order
	var address models.Address
	var productItem []models.WebProductItem

	// 查询订单信息与地址信息
	global.Db.First(&order, param.Id)
	global.Db.First(&address, order.AddressId)

	// 查询订单中包含的商品信息
	idList := strings.Split(order.ProductItem, ",")
	productIdList := make([]uint64, 0)
	for _, id := range idList {
		pid, _ := strconv.Atoi(id)
		if pid != 0 {
			productIdList = append(productIdList, uint64(pid))
		}
	}
	global.Db.Table("product").Find(&productItem, productIdList)
	orderDetail := models.WebOrderDetail{
		Id:              order.Id,
		Created:         order.Created,
		NickName:        order.NickName,
		Status:          order.Status,
		TotalPrice:      order.TotalPrice,
		Name:            address.Name,
		Mobile:          address.Mobile,
		PostalCode:      address.PostalCode,
		Province:        address.Province,
		City:            address.City,
		District:        address.District,
		DetailedAddress: address.DetailedAddress,
		ProductItem:     productItem,
	}
	return orderDetail
}

// Submit 微信小程序，提交订单
func (o *AppOrderService) Submit(param models.AppOrderSubmitParam) int64 {
	info := cart.GetInfo(models.AppCartQueryParam{UserId: param.UserId})
	pids := make([]string, 0)
	for _, item := range info.CartItem {
		pids = append(pids, strconv.Itoa(int(item.Id)))
	}
	pidsItem := strings.Join(pids, ",")
	order := models.Order{
		ProductItem: pidsItem,
		TotalPrice:  info.TotalPrice,
		Status:      param.Status,
		AddressId:   address.GetId(param.UserId),
		UserId:      param.UserId,
		NickName:    param.NickName,
		Created:     common.NowTime(),
	}
	return global.Db.Create(&order).RowsAffected
}

// GetList 微信小程序，获取订单列表
func (o *AppOrderService) GetList(param models.AppOrderQueryParam) []models.AppOrderListInfo {

	// 查询满足特定条件的订单
	orderList := make([]models.Order, 0)
	query := &models.Order{
		UserId: param.UserId,
		Status: param.Status,
	}
	global.Db.Table("order").Where(query).Find(&orderList)
	orderInfoList := make([]models.AppOrderListInfo, 0)
	for _, order := range orderList {
		var orderInfo models.AppOrderListInfo
		productItem := make([]models.AppProductItem, 0)
		orderInfo.Id = order.Id
		orderInfo.Status = order.Status
		orderInfo.TotalPrice = order.TotalPrice

		// 查询商品信息
		sList := strings.Split(order.ProductItem, ",")
		idList := make([]uint64, 0)
		for _, pid := range sList {
			id, _ := strconv.Atoi(pid)
			if id != 0 {
				i := uint64(id)
				idList = append(idList, i)
			}
		}
		global.Db.Table("product").Find(&productItem, idList)
		orderInfo.ProductItem = productItem
		orderInfoList = append(orderInfoList, orderInfo)
	}
	return orderInfoList
}
