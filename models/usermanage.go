package models

import "time"

// 权限修改参数模型
type UserPermissionParam struct {
	Id         uint64 `json:"id"`
	Permission int64  `json:"permission"`
}

// 用户信息查询参数模型
type GetUserInfoParam struct {
	Id       uint64 `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

// 用户列表查询参数模型
type UserQueryParam struct {
	Page       Page
	Permission int64 `form:"permission" binding:"omitempty"`
}

// 用户列表传输模型
type UserList struct {
	Id         uint64    `json:"id"`
	Email      string    `json:"email"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Permission int64     `json:"permission"`
	Completed  uint64    `json:"completed"` // 已完成的预约
	Failed     uint64    `json:"failed"`    //
	Status     time.Time `json:"status"`    // 恢复时间，上次违约时间+1月
}
