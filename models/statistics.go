package models

// 后台管理前端，数据总览统计传输模型
type WebDataOverviewInfo struct {
	GoodsCount int     `json:"goodsCount"`
	OrderCount int     `json:"orderCount"`
	Amount     float64 `json:"amount"`
}

// 后台管理前端，今日订单数据统计传输模型
type WebTodayOrderInfo struct {
	Data [5]int `json:"data"`
}

// 后台管理前端，本周数据总览统计传输模型
type WebWeekOrderInfo struct {
	Orders [7]int     `json:"orders"`
	Amount [7]float64 `json:"amount"`
}
