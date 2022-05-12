package models

// 微信小程序，购物车添加参数模型
type AppCartAddParam struct {
	ProductId    uint64 `form:"productId"`
	ProductCount uint64 `form:"productCount"`
	UserId       string `form:"userId"`
}

// 微信小程序，购物车删除参数模型
type AppCartDeleteParam struct {
	ProductId string `form:"productId"`
	UserId    string `form:"userId"`
}

// 微信小程序，购物车清除参数模型
type AppCartClearParam struct {
	UserId string `form:"userId"`
}

// 微信小程序，购物车信息查询参数模型
type AppCartQueryParam struct {
	ProductId uint64 `form:"productId"`
	UserId    string `form:"userId"`
}

// 微信小程序，购物车商品项传输模型
type AppCartItem struct {
	Id        uint64  `gorm:"primaryKey" json:"id"`
	MainImage string  `gorm:"image_url"  json:"mainImage"`
	Title     string  `gorm:"title"      json:"title"`
	Price     float64 `gorm:"price"      json:"price"`
	Count     int     `gorm:"count"      json:"count"`
}

// 购物车信息传输模型
type AppCartInfo struct {
	CartItem   []AppCartItem `json:"cartItem"`
	TotalPrice float64       `json:"totalPrice"`
}
