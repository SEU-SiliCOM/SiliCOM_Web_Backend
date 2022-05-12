package service

import (
	"SilicomAPPv0.3/common"
	"SilicomAPPv0.3/global"
	"SilicomAPPv0.3/models"
	"SilicomAPPv0.3/response"
	"github.com/gin-gonic/gin"
	"time"
)

type WebUserService struct {
}

// Login 用户登录信息验证
func (u *WebUserService) EmailVerify(email string, c *gin.Context) bool {
	var resNum = 0
	global.Db.Where("`email` = ?", email).Find(&models.UserInfo{}).Count(&resNum)
	if resNum > 0 {
		response.Failed("该邮箱已注册", c)
		return false
	}
	return true
}

// Login 用户登录信息验证
func (u *WebUserService) Login(param models.UserLoginParam) uint64 {
	var user models.UserInfo
	global.Db.Where("(username = ? or email = ? )and password = ?", param.Username, param.Email, param.Password).First(&user)
	return (user.Id)
}

// 获取用户权限
func (u *WebUserService) GetUserPermission(username string) int {
	var user models.UserInfo
	global.Db.Where("username = ?", username).First(&user)
	return int(user.Permission)
}

// 获取用户冻结时间
func (u *WebUserService) GetUserFreeze(username string) time.Time {
	var user models.UserInfo
	global.Db.Where("username = ?", username).First(&user)
	return user.Status
}

// 获取用户的uid
func (u *WebUserService) GetUserUid(username string) uint64 {
	var user models.UserInfo
	global.Db.Where("username = ?", username).First(&user)
	return user.Id
}

// Login 用户登录信息验证
func (u *WebUserService) PermissionUpdate(param models.UserPermissionParam) int64 {
	rows := global.Db.Model(&models.UserInfo{}).Where("id = ?", param.Id).Update("permission", param.Permission).RowsAffected
	return rows
}

// 获取用户信息
func (u *WebUserService) GetInfo(param models.GetUserInfoParam) (models.UserInfo, int64) {
	userInfo := models.UserInfo{
		Id:       param.Id,
		Email:    param.Email,
		Username: param.Username,
	}
	rows := global.Db.Model(&models.UserInfo{}).Where(userInfo).First(&userInfo).RowsAffected
	userInfo.Password = "******"
	return userInfo, rows
}

// 获取用户列表
func (u *WebUserService) GetList(param models.UserQueryParam) ([]models.UserList, int64) {
	userList := make([]models.UserList, 0)
	query := models.UserInfo{
		Permission: param.Permission,
	}
	rows := common.RestPage(param.Page, "user_info", query, &userList, &[]models.UserInfo{})
	for _, user := range userList {
		user.Password = "******"
	}
	return userList, rows
}
