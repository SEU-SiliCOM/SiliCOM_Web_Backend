package service

import (
	"SilicomAPPv0.3/common"
	"SilicomAPPv0.3/global"
	"SilicomAPPv0.3/models"
	"github.com/jinzhu/gorm"
	"time"
)

type ActivityService struct {
}

// 创建活动
func (a *ActivityService) Create(param models.ActivityAddParam) uint64 {
	actTime, _ := time.ParseInLocation("2006-01-02 15:04:05", param.StartTime, time.Local)
	activity := models.Activity{
		Type:      param.Type,
		Name:      param.Name,
		StartTime: actTime,
		Created:   time.Now(),
	}
	global.Db.Create(&activity)
	return activity.Id
}

// 删除活动
func (a *ActivityService) Delete(param models.ActivityDeleteParam) int64 {
	rows := global.Db.Delete(&models.Activity{}, param.Id).RowsAffected
	return rows
}

// 更新活动
func (a *ActivityService) Update(param models.ActivityUpdateParam) int64 {
	startTime, _ := time.ParseInLocation("2006-01-02 15:04:05", param.StartTime, time.Local)
	activity := models.Activity{
		Id:        param.Id,
		Type:      param.Type,
		Name:      param.Name,
		StartTime: startTime,
		Updated:   time.Now(),
	}
	rows := global.Db.Model(&activity).Updates(activity).RowsAffected
	return rows
}

// 查询活动
func (a *ActivityService) GetList(param models.ActivityQueryParam) ([]models.ActivityList, int64) {
	activityList := make([]models.ActivityList, 0)
	startTime, _ := time.ParseInLocation("2006-01-02 15:04:05", param.StartTime, time.Local)
	query := &models.Activity{
		Id:        param.Id,
		Type:      param.Type,
		Name:      param.Name,
		StartTime: startTime,
	}
	rows := common.RestPage(param.Page, "activity", query, &activityList, &[]models.Activity{})
	return activityList, rows
}

// 完成活动，已完成
func (a *ActivityService) Complete(param models.ActivityCompleteParam) int64 {
	type user struct {
		Uid uint64
	}
	var users []user
	var userId []uint64
	rows := global.Db.Model(&models.Appointment{}).Select("uid").Where("status = ? and aid = ?", uint64(1), param.Id).Scan(&users).RowsAffected
	for _, user := range users {
		userId = append(userId, user.Uid)
	}
	global.Db.Debug().Model(&models.UserInfo{}).Where("id IN (?)", userId).UpdateColumn("completed", gorm.Expr("completed + ?", 1))

	var usersFailed []user
	var userFailedId []uint64
	rows += global.Db.Model(&models.Appointment{}).Select("uid").Where("status = ? and aid = ?", uint64(0), param.Id).Scan(&usersFailed).RowsAffected
	for _, user := range usersFailed {
		userFailedId = append(userFailedId, user.Uid)
	}
	global.Db.Debug().Model(&models.UserInfo{}).Where("id IN (?)", userFailedId).UpdateColumn("failed", gorm.Expr("failed + ?", 1))
	global.Db.Model(&models.UserInfo{}).Where("id IN (?)", userFailedId).UpdateColumn("status", time.Now().Add(7*24*time.Hour))
	global.Db.Model(&models.Activity{}).Where("id = ?", param.Id).UpdateColumn("status", uint64(1))
	return rows
}

// 获取活动状态
func (a *ActivityService) GetActivityStatus(aid uint64) uint64 {
	var activity models.Activity
	global.Db.Where("id = ?", aid).First(&activity)
	return activity.Status
}
