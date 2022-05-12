package models

// 数据库，订单数据映射模型
type Order struct {
	Id          uint64  `gorm:"primaryKey"`
	ProductItem string  `gorm:"product_item"`
	TotalPrice  float64 `gorm:"total_price"`
	Status      string  `gorm:"status"`
	AddressId   uint64  `gorm:"address_id"`
	UserId      string  `gorm:"user_id"`
	NickName    string  `gorm:"nick_name"`
	Created     string  `gorm:"created"`
	Updated     string  `gorm:"updated"`
}

// 后台管理前端，订单删除参数模型
type WebOrderDeleteParam struct {
	Id uint64 `form:"id"`
}

// 后台管理前端，订单更新参数模型
type WebOrderUpdateParam struct {
	Id     uint64 `json:"id"`
	Status string `json:"status"`
}

// 后台管理前端，订单列表查询参数模型
type WebOrderListParam struct {
	Page   Page
	Id     uint64 `form:"id"`
	Status string `form:"status"`
}

// 后台管理前端，订单列表查询参数模型
type WebOrderDetailParam struct {
	Id uint64 `form:"id"`
}

// 后台管理前端，订单列表传输模型
type WebOrderList struct {
	Id         uint64  `json:"id"`
	NickName   string  `json:"nickName"`
	Status     string  `json:"status"`
	TotalPrice float64 `json:"totalPrice"`
	Created    string  `json:"created"`
}

// 后台管理前端，订单详情传输模型
type WebOrderDetail struct {

	// 订单信息
	Id         uint64  `json:"id"`
	Created    string  `json:"created"`
	NickName   string  `json:"nickName"`
	Status     string  `json:"status"`
	TotalPrice float64 `json:"totalPrice"`

	// 地址信息
	Name            string `json:"name"`
	Mobile          string `json:"mobile"`
	PostalCode      int    `json:"postalCode"`
	Province        string `json:"province"`
	City            string `json:"city"`
	District        string `json:"district"`
	DetailedAddress string `json:"detailedAddress"`

	// 商品列表信息
	ProductItem []WebProductItem `json:"productItem"`
}

// 微信小程序，订单提交参数模型
type AppOrderSubmitParam struct {
	UserId   string `form:"userId"    json:"userId"`
	NickName string `form:"nickName"  json:"nickName"`
	Status   string `form:"status"    json:"status"`
}

// 微信小程序，订单查询参数模型
type AppOrderQueryParam struct {
	UserId string `form:"userId"    json:"userId"`
	Status string `form:"status"    json:"status"`
}

// 微信小程序，订单列表传输模型
type AppOrderListInfo struct {
	Id          uint64           `json:"id"`
	Status      string           `json:"status"`
	TotalPrice  float64          `json:"totalPrice"`
	ProductItem []AppProductItem `json:"productItem"`
}
