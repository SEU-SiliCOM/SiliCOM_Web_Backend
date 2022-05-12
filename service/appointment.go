package service

import (
	"SilicomAPPv0.3/common"
	"SilicomAPPv0.3/global"
	"SilicomAPPv0.3/models"
	"github.com/jinzhu/gorm"
	"time"
)

type AppointmentService struct {
}

// 创建预约
func (a *AppointmentService) Create(param models.AppointmentCreateParam) uint64 {
	var appointment models.Appointment
	result := global.Db.Where("uid = ? and aid = ?", param.Uid, param.Aid).First(&appointment)
	if result.RowsAffected > 0 {
		return appointment.Id
	}
	appointment = models.Appointment{
		Uid:     param.Uid,
		Aid:     param.Aid,
		Status:  0,
		Created: time.Now(),
	}
	global.Db.Create(&appointment)
	global.Db.Model(&models.Activity{}).Where("id = ?", param.Aid).UpdateColumn("total_appointment", gorm.Expr("total_appointment + ?", 1))
	return appointment.Id
}

// 更新预约
func (a *AppointmentService) Update(param models.AppointmentUpdateParam) int64 {
	var appointment models.Appointment
	appointment = models.Appointment{
		Uid:     param.Uid,
		Aid:     param.Aid,
		Status:  0,
		Updated: time.Now(),
	}
	return global.Db.Model(&appointment).Updates(appointment).RowsAffected
}

// 删除预约
func (a *AppointmentService) Delete(param models.AppointmentDeleteParam) int64 {
	return global.Db.Delete(&models.Appointment{}, param.Id).RowsAffected
}

// 查询预约
func (a *AppointmentService) GetList(param models.AppointmentQueryParam) ([]models.AppointmentList, int64) {
	appointmentList := make([]models.AppointmentList, 0)
	query := models.Appointment{
		Id:     param.Id,
		Uid:    param.Uid,
		Aid:    param.Aid,
		Status: param.Status,
	}
	rows := common.RestPage(param.Page, "appointment", query, &appointmentList, &[]models.Appointment{})
	return appointmentList, rows
}

// 签到
func (a *AppointmentService) Sign(param models.AppointmentSignParam) int64 {
	var appointment models.Appointment
	global.Db.Where("id = ?", param.Id).First(&appointment)
	global.Db.Model(&models.Activity{}).Where("id = ?", appointment.Aid).UpdateColumn("total_complete", gorm.Expr("total_complete + ?", 1))
	return global.Db.Model(&models.Appointment{}).Where("id = ?", param.Id).Update("status", 1).RowsAffected
}

//
