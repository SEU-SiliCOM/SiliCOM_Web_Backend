package models

import (
	"time"
)

// 数据库，活动信息数据映射模型，活动添加模型
type Activity struct {
	Id               uint64 `gorm:"primaryKey"`
	Type             uint64
	Name             string
	Status           uint64 // 0可预约 1已结束
	TotalAppointment uint64
	TotalComplete    uint64
	StartTime        time.Time
	Created          time.Time
	Updated          time.Time
}

type ActivityAddParam struct {
	Type      uint64 `json:"type" binding:"required,oneof=1 2"`
	Name      string `json:"name" binding:"required"`
	StartTime string `json:"startTime"`
}

// 活动信息删除模型
type ActivityDeleteParam struct {
	Id uint64 `json:"actId" binding:"required"`
}

// 活动信息更改模型
type ActivityUpdateParam struct {
	Id        uint64 `json:"Id" binding:"required"`
	Type      uint64 `json:"type" binding:"omitempty,oneof=1 2"`
	Name      string `json:"name"`
	StartTime string `json:"startTime"`
}

// 活动信息查询参数模型
type ActivityQueryParam struct {
	Page      Page
	Id        uint64 `form:"id"       binding:"omitempty,gt=0"`
	Type      uint64 `form:"type"   binding:"omitempty,oneof=1 2"`
	Name      string `form:"name"     binding:"omitempty"`
	Status    uint64 `form:"status" binding:"omitempty"`
	StartTime string `form:"startTime" binding:"omitempty"`
}

// 活动列表传输模型
type ActivityList struct {
	Id               uint64 `json:"id"`
	Type             uint64 `json:"type"`
	Name             string `json:"name"`
	Status           uint64 `json:"status"`
	TotalAppointment uint64 `json:"totalAppointment"`
	TotalComplete    uint64 `json:"totalComplete"`
	StartTime        string `json:"startTime"`
}

// 活动信息完成模型
type ActivityCompleteParam struct {
	Id uint64 `json:"id" binding:"required"`
}
