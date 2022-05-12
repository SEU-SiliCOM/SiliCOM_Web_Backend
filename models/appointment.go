package models

import "time"

// 数据库，预约数据映射模型
type Appointment struct {
	Id      uint64 `gorm:"primaryKey"`
	Uid     uint64
	Aid     uint64
	Status  uint64 // 0未完成 1已完成
	Created time.Time
	Updated time.Time
}

// 后台管理前端，预约创建参数模型
type AppointmentCreateParam struct {
	Uid uint64 `json:"uid"`
	Aid uint64 `json:"aid" binding:"required"`
}

// 后台管理前端，预约更新参数模型
type AppointmentUpdateParam struct {
	Id  uint64 `json:"id" binding:"required"`
	Uid uint64 `json:"uid" binding:"-"`
	Aid uint64 `json:"aid" binding:"omitempty"`
}

// 后台管理前端，预约删除参数模型
type AppointmentDeleteParam struct {
	Id uint64 `json:"id" binding:"required"`
}

// 活动信息查询参数模型
type AppointmentQueryParam struct {
	Page   Page
	Id     uint64 `form:"id"       binding:"omitempty"`
	Uid    uint64 `form:"uid"   binding:"omitempty"`
	Aid    uint64 `form:"aid"     binding:"omitempty"`
	Status uint64 `form:"status" binding:"omitempty"`
}

// 活动列表传输模型
type AppointmentList struct {
	Id     uint64 `json:"id"`
	Aid    uint64 `json:"aid"`
	Uid    uint64 `json:"uid"`
	Status uint64 `json:"status"`
}

// 签到信息传输模型
type AppointmentSignParam struct {
	Id uint64 `json:"id" binding:"required"`
}
