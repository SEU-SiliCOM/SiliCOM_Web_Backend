package models

// 微信小程序，浏览记录保存参数模型
type AppBrowseSaveParam struct {
	ProductId uint64 `form:"productId" binding:"required"`
	UserId    string `form:"userId" binding:"required"`
}

// 微信小程序，浏览记录删除参数模型
type AppBrowseDeleteParam struct {
	ProductIds []string `form:"productIds" binding:"required"`
	UserId     string   `form:"userId" binding:"required"`
}

// 微信小程序，浏览记录查询参数模型
type AppBrowseQueryParam struct {
	UserId string `form:"userId" binding:"required"`
}

// 微信小程序，浏览记录列表商品项传输模型
type AppBrowseItem struct {
	Id        uint64  `json:"id"`
	MainImage string  `json:"mainImage"`
	Title     string  `json:"title"`
	Price     float64 `json:"price"`
}

// 微信小程序，浏览记录列表传输模型
type AppBrowseRecordList struct {
	BrowseTime string        `json:"browseTime"`
	BrowseItem AppBrowseItem `json:"browseItem"`
}
