package models

// 微信小程序，商品收藏添加参数模型
type AppCollectionAddParam struct {
	ProductId uint64 `form:"productId" binding:"required"`
	UserId    string `form:"userId" binding:"required"`
}

// 微信小程序，商品收藏删除参数模型
type AppCollectionDeleteParam struct {
	UserId string `form:"userId" binding:"required"`
}

// 微信小程序，商品收藏查询参数模型
type AppCollectionQueryParam struct {
	UserId string `form:"userId" binding:"required"`
}

// 微信小程序，商品收藏列表传输模型
type AppCollectionList struct {
	Id        uint64  `json:"id"`
	MainImage string  `json:"mainImage"`
	Title     string  `json:"title"`
	Price     float64 `json:"price"`
}
