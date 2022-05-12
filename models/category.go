package models

// 数据库,类目数据映射模型
type Category struct {
	Id       uint64 `gorm:"primaryKey"`
	Name     string `gorm:"name"`
	ParentId uint64 `gorm:"parent_id"`
	Level    uint64 `gorm:"level"`
	Sort     uint64 `gorm:"sort"`
	Created  string `gorm:"created"`
	Updated  string `gorm:"updated"`
}

// 后台管理前端，类目创建参数模型
type WebCategoryCreateParam struct {
	Name     string `json:"name"     binding:"required"`
	ParentId uint64 `json:"parentId" binding:"required,gt=0"` //gt表示大于
	Level    uint64 `json:"level"    binding:"required,oneof=1 2 3"`
	Sort     uint64 `json:"sort"     binding:"required,gt=0"`
}

// 后台管理前端，类目删除参数模型
type WebCategoryDeleteParam struct {
	Id uint64 `form:"id" binding:"required,gt=0"`
}

// 后台管理前端，类目更新参数模型
type WebCategoryUpdateParam struct {
	Id   uint64 `json:"id"       binding:"required,gt=0"`
	Name string `json:"name"     binding:"required"`
	Sort uint64 `json:"sort"     binding:"required,gt=0"`
}

// 后台管理前端，类目查询参数模型
type WebCategoryQueryParam struct {
	Page     Page
	Id       uint64 `form:"id"       binding:"omitempty,gt=0"`
	Name     string `form:"name"     binding:"omitempty"`
	ParentId uint64 `form:"parentId" binding:"omitempty,gt=0"`
	Level    uint64 `form:"level"    binding:"omitempty,oneof=1 2 3"`
}

// 后台管理前端，类目列表传输模型
type WebCategoryList struct {
	Id       uint64 `json:"id"`
	Name     string `json:"name"`
	ParentId uint64 `json:"parentId"`
	Level    uint64 `json:"level"`
	Sort     uint64 `json:"sort"`
}

// 后台管理前端，类目选项传输模型
type WebCategoryOption struct {
	Value    uint64              `json:"value"`
	Label    string              `json:"label"`
	Children []WebCategoryOption `json:"children"`
}

// 微信小程序，类目选项传输模型
type AppCategoryOption struct {
	Id   uint64 `json:"id"`
	Text string `json:"text"`
}
