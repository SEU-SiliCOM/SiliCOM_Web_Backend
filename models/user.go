package models

import (
	"time"
)

// 数据库中用户信息参数
type UserInfo struct {
	Id         uint64 `gorm:"primaryKey"`
	Email      string `gorm:"unique"`
	Username   string `gorm:"unique;index"`
	Password   string
	Permission int64
	Completed  uint64    // 已完成的预约
	Failed     uint64    //
	Status     time.Time // 恢复时间，上次违约时间+1月
}

// 注册申请邮箱验证码参数
type EmailParam struct {
	Email string `json:"email"`
	Type  bool   `json:"type"`
}

// 数据库中邮箱验证码参数
type Verifycode struct {
	Id        uint64 `gorm:"primaryKey"`
	Email     string
	Code      string
	InvalidAt time.Time
}

// 注册参数
type UserRegisterParam struct {
	Email      string `json:"email"`
	Verifycode string `json:"code"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

// 找回密码参数
type UserFindParam struct {
	Email      string `json:"email"`
	Verifycode string `json:"code"`
	Password   string `json:"password"`
}

// 登录参数
type UserLoginParam struct {
	Email        string `json:"email"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	CaptchaId    string `json:"captchaId"`
	CaptchaValue string `json:"captchaValue"`
}

// 后台管理前端，用户信息传输模型
type WebUserInfo struct {
	Uid   uint64 `json:"uid"`
	Token string `json:"token"`
}

// 用户名查询模型
type FindUsernameParam struct {
	Email string `json:"email"`
}

// 微信小程序，用户登录凭证校验模型
type AppCode2Session struct {
	Code      string
	AppId     string
	AppSecret string
}

// 微信小程序，凭证校验后返回的JSON数据包模型
type AppCode2SessionJson struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    uint64 `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}
