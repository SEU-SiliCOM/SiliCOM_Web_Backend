package models

// Page 页面请求参数
type Page struct {
	PageNum  int `form:"pageNum"  json:"pageNum"  binding:"required,gt=0"`
	PageSize int `form:"pageSize" json:"pageSize" binding:"required,gt=0"`
}
