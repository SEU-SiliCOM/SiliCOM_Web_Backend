package service

import (
	"SilicomAPPv0.3/common"
	"SilicomAPPv0.3/global"
	"SilicomAPPv0.3/models"
	"time"
)

type WebCategoryService struct {
}
type AppCategoryService struct {
}

// 创建类目
func (c *WebCategoryService) Create(param models.WebCategoryCreateParam) uint64 {
	var category models.Category
	result := global.Db.Where("name = ?", param.Name).First(&category)
	if result.RowsAffected > 0 {
		return category.Id
	}
	category = models.Category{
		Name:     param.Name,
		ParentId: param.ParentId,
		Level:    param.Level,
		Sort:     param.Sort,
		Created:  common.NowTime(),
	}
	global.Db.Create(&category)
	return category.Id
}

// 删除类目
func (c *WebCategoryService) Delete(param models.WebCategoryDeleteParam) int64 {
	var count int64
	var rec func(id uint64) // 递归删除子类
	rec = func(id uint64) {
		var pid2 []models.Category
		count += global.Db.Delete(&models.Category{}, id).RowsAffected
		global.Db.Where("parent_id = ?", id).Find(&pid2) //删除子类
		if len(pid2) == 0 {
			return
		} else {
			for _, kid := range pid2 {
				rec(kid.Id)
			}
		}
	}
	rec(param.Id)
	return count
}

// 更新类目
func (c *WebCategoryService) Update(param models.WebCategoryUpdateParam) int64 {
	category := models.Category{
		Id:      param.Id,
		Name:    param.Name,
		Sort:    param.Sort,
		Updated: common.NowTime(),
	}
	return global.Db.Model(&category).Updates(category).RowsAffected
}

// GetList 后台管理前端，获取类目列表
func (c *WebCategoryService) GetList(param models.WebCategoryQueryParam) ([]models.WebCategoryList, int64) {
	categoryList := make([]models.WebCategoryList, 0)
	query := &models.Category{
		Id:       param.Id,
		Name:     param.Name,
		Level:    param.Level,
		ParentId: param.ParentId,
	}
	rows := common.RestPage(param.Page, "category", query, &categoryList, &[]models.Category{})
	return categoryList, rows
}

// GetOption 后台管理前端，获取类目选项
func (c *WebCategoryService) GetOption() (option []models.WebCategoryOption) {
	selectList := make([]models.WebCategoryList, 0)
	global.Db.Table("category").Find(&selectList)
	return getTreeOptions(1, selectList)
}

// GetOption 微信小程序，获取类目选项
func (c *AppCategoryService) GetOption() []models.AppCategoryOption {
	var category []models.Category
	optionList := make([]models.AppCategoryOption, 0)
	global.Db.Table("category").Where("parent_id = ?", 1).Find(&category)
	for _, c := range category {
		optionList = append(optionList, models.AppCategoryOption{Id: c.Id, Text: c.Name})
	}
	return optionList
}

// 获取树形结构的选项
func getTreeOptions(id uint64, cateList []models.WebCategoryList) (option []models.WebCategoryOption) {
	optionList := make([]models.WebCategoryOption, 0)
	for _, opt := range cateList {
		if opt.ParentId == id && (opt.Level == 1 || opt.Level == 2 || opt.Level == 3) {
			if opt.Level == 2 {
				time.Sleep(1 * time.Second)
			}
			option := models.WebCategoryOption{
				Value:    opt.Id,
				Label:    opt.Name,
				Children: getTreeOptions(opt.Id, cateList),
			}
			if opt.Level == 3 {
				option.Children = nil
			}
			optionList = append(optionList, option)
		}
	}
	return optionList
}
