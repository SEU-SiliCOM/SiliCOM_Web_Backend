package models

// 数据库，收货地址数据映射模型
type Address struct {
	Id              uint64 `gorm:"primaryKey"`
	UserId          string `gorm:"user_id"`
	Name            string `gorm:"name"`
	Mobile          string `gorm:"mobile"`
	PostalCode      int    `gorm:"postal_code"`
	Province        string `gorm:"province"`
	City            string `gorm:"city"`
	District        string `gorm:"district"`
	DetailedAddress string `gorm:"detailed_address"`
	IsDefault       int    `gorm:"is_default"`
	Created         string `gorm:"created"`
	Updated         string `gorm:"updated"`
}

// 微信小程序，收货地址添加参数模型
type AppAddressAddParam struct {
	UserId          string `form:"userId"`
	Name            string `form:"name"`
	Mobile          string `form:"mobile"`
	PostalCode      int    `form:"postalCode"`
	Province        string `form:"province"`
	City            string `form:"city"`
	District        string `form:"district"`
	DetailedAddress string `form:"detailedAddress"`
	IsDefault       int    `form:"isDefault"`
}

// 微信小程序，收货地址更新参数模型
type AppAddressUpdateParam struct {
	Id              uint64 `form:"id"`
	UserId          string `form:"userId"`
	Name            string `form:"name"`
	Mobile          string `form:"mobile"`
	PostalCode      int    `form:"postalCode"`
	Province        string `form:"province"`
	City            string `form:"city"`
	District        string `form:"district"`
	DetailedAddress string `form:"detailedAddress"`
	IsDefault       int    `form:"isDefault"`
}

// 微信小程序，收货地址删除参数模型
type AppAddressDeleteParam struct {
	AddressId uint64 `form:"addressId"`
}

// 微信小程序，收货地址信息参数模型
type AppAddressInfoParam struct {
	AddressId uint64 `form:"addressId"`
}

// 微信小程序，收货地址列表参数模型
type AppAddressListParam struct {
	UserId string `form:"userId" json:"userId"`
}

// 微信小程序，收货地址列表传输模型
type AppAddressList struct {
	Id              uint64 `json:"id"`
	Name            string `json:"name"`
	Mobile          string `json:"mobile"`
	PostalCode      int    `json:"postalCode"`
	Province        string `json:"province"`
	City            string `json:"city"`
	District        string `json:"district"`
	DetailedAddress string `json:"detailedAddress"`
	IsDefault       int    `json:"isDefault"`
}

// 微信小程序，收货地址更新信息传输模型
type AppAddressUpdateInfo struct {
	Id              uint64 `json:"id"`
	Name            string `json:"name"`
	Mobile          string `json:"mobile"`
	PostalCode      int    `json:"postalCode"`
	Province        string `json:"province"`
	City            string `json:"city"`
	District        string `json:"district"`
	DetailedAddress string `json:"detailedAddress"`
	IsDefault       int    `json:"isDefault"`
}
