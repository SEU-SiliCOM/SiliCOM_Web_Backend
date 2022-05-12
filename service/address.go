package service

import (
	"SilicomAPPv0.3/common"
	"SilicomAPPv0.3/global"
	"SilicomAPPv0.3/models"
	"fmt"
)

type AddressService struct {
}

// Add 微信小程序，添加地址
func (a *AddressService) Add(param models.AppAddressAddParam) int64 {
	address := models.Address{
		UserId:          param.UserId,
		Name:            param.Name,
		Mobile:          param.Mobile,
		PostalCode:      param.PostalCode,
		Province:        param.Province,
		City:            param.City,
		District:        param.District,
		DetailedAddress: param.DetailedAddress,
		IsDefault:       param.IsDefault,
		Created:         common.NowTime(),
	}
	if address.IsDefault == 1 {
		var addressId uint64
		row := global.Db.Table("address").Select("id").
			Where("is_default = ? and user_id = ?", 1, address.UserId).Take(&addressId).RowsAffected
		if row > 0 {
			global.Db.Table("address").Where("id = ?", addressId).
				Update("is_default", 2)
			return global.Db.Create(&address).RowsAffected
		}
		return global.Db.Create(&address).RowsAffected
	}
	return global.Db.Create(&address).RowsAffected
}

// Delete 微信小程序，删除地址
func (a *AddressService) Delete(id uint64) int64 {
	return global.Db.Delete(&models.Address{}, id).RowsAffected
}

// Update 微信小程序，更新地址
func (a *AddressService) Update(param models.AppAddressUpdateParam) int64 {
	address := models.Address{
		Id:              param.Id,
		UserId:          param.UserId,
		Name:            param.Name,
		Mobile:          param.Mobile,
		PostalCode:      param.PostalCode,
		Province:        param.Province,
		City:            param.City,
		District:        param.District,
		DetailedAddress: param.DetailedAddress,
		IsDefault:       param.IsDefault,
		Updated:         common.NowTime(),
	}
	if address.IsDefault == 1 {
		var addressId uint64
		row := global.Db.Table("address").Select("id").
			Where("is_default = ? and user_id = ?", 1, address.UserId).Take(&addressId).RowsAffected
		fmt.Println(addressId)
		if row > 0 {
			global.Db.Table("address").Where("id = ?", addressId).
				Update("is_default", 2)
			return global.Db.Updates(&address).RowsAffected
		}
		return global.Db.Updates(&address).RowsAffected
	}
	return global.Db.Updates(&address).RowsAffected
}

// GetId 微信小程序，获取地址Id
func (a *AddressService) GetId(uid string) uint64 {
	var id uint64
	global.Db.Table("address").Select("id").
		Where("is_default = ? and user_id = ?", 1, uid).Take(&id)
	return id
}

// GetInfo 微信小程序，获取更新信息
func (a *AddressService) GetInfo(id uint64) models.AppAddressUpdateInfo {
	var updateInfo models.AppAddressUpdateInfo
	global.Db.Table("address").First(&updateInfo, id)
	return updateInfo
}

// GetList 微信小程序，获取地址列表
func (a *AddressService) GetList(uid string) []models.AppAddressList {
	aList := make([]models.AppAddressList, 0)
	global.Db.Table("address").Where("user_id = ?", uid).Order("is_default").Find(&aList)
	return aList
}
