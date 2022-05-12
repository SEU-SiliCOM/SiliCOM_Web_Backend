package models

// 数据库，商品数据映射模型
type Product struct {
	Id                uint64  `gorm:"primaryKey"         json:"id"`
	CategoryId        uint64  `gorm:"category_id"        json:"categoryId"`
	Title             string  `gorm:"title"              json:"title"`
	Description       string  `gorm:"description"        json:"description"`
	Price             float64 `gorm:"price"              json:"price"`
	Amount            uint64  `gorm:"amount"             json:"amount"`
	Sales             uint64  `gorm:"sales"              json:"sales"`
	MainImage         string  `gorm:"main_image"         json:"mainImage"`
	Delivery          string  `gorm:"delivery"           json:"delivery"`
	Assurance         string  `gorm:"assurance"          json:"assurance"`
	Name              string  `gorm:"name"               json:"name"`
	Weight            float32 `gorm:"weight"             json:"weight"`
	Brand             string  `gorm:"brand"              json:"brand"`
	Origin            string  `gorm:"origin"             json:"origin"`
	ShelfLife         uint64  `gorm:"shelf_life"         json:"shelfLife"`
	NetWeight         float32 `gorm:"net_weight"         json:"netWeight"`
	UseWay            string  `gorm:"use_way"            json:"useWay"`
	PackingWay        string  `gorm:"packing_way"        json:"packingWay"`
	StorageConditions string  `gorm:"storage_conditions" json:"storageConditions"`
	DetailImage       string  `gorm:"detail_image"       json:"detailImage"`
	Status            uint64  `gorm:"status"             json:"status"`
	Created           string  `gorm:"created"            json:"created"`
	Updated           string  `gorm:"updated"            json:"updated"`
}

// 后台管理前端，商品创建参数模型
type WebProductCreateParam struct {
	CategoryId        uint64  `json:"categoryId"         binding:"required,gt=0"`
	Title             string  `json:"title"              binding:"required"`
	Description       string  `json:"description"`
	Price             float64 `json:"price"              binding:"required,gt=0"`
	Amount            uint64  `json:"amount"             binding:"required,gt=0"`
	MainImage         string  `json:"mainImage"          binding:"required"`
	Delivery          string  `json:"delivery"           binding:"required"`
	Assurance         string  `json:"assurance"          binding:"required"`
	Name              string  `json:"name"               binding:"required"`
	Weight            float32 `json:"weight"             binding:"required"`
	Brand             string  `json:"brand"`
	Origin            string  `json:"origin"`
	ShelfLife         uint64  `json:"shelfLife"          binding:"required,gt=0"`
	NetWeight         float32 `json:"netWeight"          binding:"required,gt=0"`
	UseWay            string  `json:"useWay"`
	PackingWay        string  `json:"packingWay"`
	StorageConditions string  `json:"storageConditions"`
	DetailImage       string  `json:"detailImage"        binding:"required"`
	Status            uint64  `json:"status"             binding:"required,oneof=1 2"`
}

// 后台管理前端，商品删除参数模型
type WebProductDeleteParam struct {
	Id uint64 `form:"id" binding:"required,gt=0"`
}

// 后台管理前端，商品更新参数模型
type WebProductUpdateParam struct {
	Id                uint64  `json:"id"                 binding:"required,gt=0"`
	CategoryId        uint64  `json:"categoryId"         binding:"required,gt=0"`
	Title             string  `json:"title"              binding:"required"`
	Description       string  `json:"description"`
	Price             float64 `json:"price"              binding:"required,gt=0"`
	Amount            uint64  `json:"amount"             binding:"required,gt=0"`
	MainImage         string  `json:"mainImage"          binding:"required"`
	Delivery          string  `json:"delivery"           binding:"required"`
	Assurance         string  `json:"assurance"          binding:"required"`
	Name              string  `json:"name"               binding:"required"`
	Weight            float32 `json:"weight"             binding:"required"`
	Brand             string  `json:"brand"`
	Origin            string  `json:"origin"`
	ShelfLife         uint64  `json:"shelfLife"          binding:"required,gt=0"`
	NetWeight         float32 `json:"netWeight"          binding:"required,gt=0"`
	UseWay            string  `json:"useWay"`
	PackingWay        string  `json:"packingWay"`
	StorageConditions string  `json:"storageConditions"`
	DetailImage       string  `json:"detailImage"        binding:"required"`
	Status            uint64  `json:"status"             binding:"omitempty,oneof=1 2"`
}

// 后台管理前端，商品状态更新参数模型
type WebProductStatusUpdateParam struct {
	Id     uint64 `json:"id"     binding:"required,gt=0"`
	Status uint64 `json:"status" binding:"required,oneof=1 2"`
}

// 后台管理前端，商品信息查询参数模型
type WebProductInfoParam struct {
	Id uint64 `form:"id" binding:"required,gt=0"`
}

// 后台管理前端，商品列表查询参数模型
type WebProductListParam struct {
	Page       Page
	Id         uint64 `form:"id"         binding:"omitempty,gt=0"`
	CategoryId uint64 `form:"categoryId" binding:"omitempty,gt=0"`
	Title      string `form:"title"      binding:"omitempty"`
	Status     uint64 `form:"status"     binding:"omitempty,oneof=1 2"`
}

// 后台管理前端，商品信息传输模型
type WebProductInfo struct {
	Id                uint64  `json:"id"`
	CategoryId        uint64  `json:"categoryId"`
	Title             string  `json:"title"`
	Description       string  `json:"description"`
	Price             float64 `json:"price"`
	Amount            uint64  `json:"amount"`
	MainImage         string  `json:"mainImage"`
	Delivery          string  `json:"delivery"`
	Assurance         string  `json:"assurance"`
	Name              string  `json:"name"`
	Weight            float32 `json:"weight"`
	Brand             string  `json:"brand"`
	Origin            string  `json:"origin"`
	ShelfLife         uint64  `json:"shelfLife"`
	NetWeight         float32 `json:"netWeight"`
	UseWay            string  `json:"useWay"`
	PackingWay        string  `json:"packingWay"`
	StorageConditions string  `json:"storageConditions"`
	DetailImage       string  `json:"detailImage"`
}

// 后台管理前端，商品项传输模型
type WebProductItem struct {
	Id        uint64  `json:"id"`
	Title     string  `json:"title"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	MainImage string  `json:"mainImage"`
}

// 后台管理前端，商品列表传输模型
type WebProductList struct {
	Id        uint64  `json:"id"`
	Title     string  `json:"title"`
	Price     float64 `json:"price"`
	Amount    int     `json:"amount"`
	MainImage string  `json:"mainImage"`
	Sales     int     `json:"sales"`
	Status    int     `json:"status"`
	Created   string  `json:"created"`
}

// 微信小程序，商品列表参数模型
type AppProductListParam struct {
	CategoryId uint64 `form:"categoryId"`
}

// 微信小程序，商品搜索参数模型
type AppProductSearchParam struct {
	KeyWord string `form:"keyWord"`
}

// 微信小程序，商品详情参数模型
type AppProductDetailParam struct {
	ProductId uint64 `form:"productId"`
}

// 微信小程序，商品列表传输模型
type AppProductList struct {
	Id          uint64  `json:"id"`
	MainImage   string  `json:"mainImage"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Sales       int     `json:"sales"`
}

// 微信小程序，商品项传输模型
type AppProductItem struct {
	Id        uint64  `json:"id"`
	Title     string  `json:"title"`
	Price     float64 `json:"price"`
	MainImage string  `json:"mainImage"`
}

// 微信小程序，商品信息传输模型
type AppProductDetail struct {
	Id                uint64  `json:"id"`
	Title             string  `json:"title"`
	Price             float64 `json:"price"`
	Amount            uint64  `json:"amount"`
	Sales             uint64  `json:"sales"`
	MainImage         string  `json:"mainImage"`
	Delivery          string  `json:"delivery"`
	Assurance         string  `json:"assurance"`
	Name              string  `json:"name"`
	Brand             string  `json:"brand"`
	Origin            string  `json:"origin"`
	ShelfLife         uint64  `json:"shelfLife"`
	NetWeight         float32 `json:"netWeight"`
	UseWay            string  `json:"useWay"`
	PackingWay        string  `json:"packingWay"`
	StorageConditions string  `json:"storageConditions"`
	DetailImage       string  `json:"detailImage"`
}
