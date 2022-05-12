package service

import (
	"SilicomAPPv0.3/common"
	"SilicomAPPv0.3/global"
	"SilicomAPPv0.3/models"
	"fmt"
	"github.com/olivere/elastic/v7"
	"reflect"
	"strconv"
)

type WebProductService struct {
}

type AppProductService struct {
}

// Create 后台管理前端，创建商品
func (p *WebProductService) Create(param models.WebProductCreateParam) int64 {
	product := models.Product{
		CategoryId:        param.CategoryId,
		Title:             param.Title,
		Description:       param.Description,
		Price:             param.Price,
		Amount:            param.Amount,
		MainImage:         param.MainImage,
		Delivery:          param.Delivery,
		Assurance:         param.Assurance,
		Name:              param.Name,
		Weight:            param.Weight,
		Brand:             param.Brand,
		Origin:            param.Origin,
		ShelfLife:         param.ShelfLife,
		NetWeight:         param.NetWeight,
		UseWay:            param.UseWay,
		PackingWay:        param.PackingWay,
		StorageConditions: param.StorageConditions,
		DetailImage:       param.DetailImage,
		Status:            param.Status,
		Created:           common.NowTime(),
	}
	rows := global.Db.Create(&product).RowsAffected
	records := global.Db.First(&product, product.Id).RowsAffected
	if records > 0 {
		id := strconv.FormatUint((product.Id), 10) // Id的十进制的字符串表示
		result, err := global.Es.Index().Index("product").Id(id).BodyJson(product).Do(ctx)
		if err != nil {
			fmt.Println(err)
		}
		return result.PrimaryTerm
	}
	return rows
}

// Delete 后台管理前端，删除商品
func (p *WebProductService) Delete(param models.WebProductDeleteParam) int64 {
	rows := global.Db.Delete(&models.Product{}, param.Id).RowsAffected
	if rows > 0 {
		id := strconv.FormatUint(param.Id, 10)
		_, err := global.Es.Delete().Index("product").Id(id).Do(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	return rows
}

// Update 后台管理前端，更新商品
func (p *WebProductService) Update(param models.WebProductUpdateParam) int64 {
	product := models.Product{
		Id:                param.Id,
		CategoryId:        param.CategoryId,
		Title:             param.Title,
		Description:       param.Description,
		Price:             param.Price,
		Amount:            param.Amount,
		MainImage:         param.MainImage,
		Delivery:          param.Delivery,
		Assurance:         param.Assurance,
		Name:              param.Name,
		Weight:            param.Weight,
		Brand:             param.Brand,
		Origin:            param.Origin,
		ShelfLife:         param.ShelfLife,
		NetWeight:         param.NetWeight,
		UseWay:            param.UseWay,
		PackingWay:        param.PackingWay,
		StorageConditions: param.StorageConditions,
		DetailImage:       param.DetailImage,
		Status:            param.Status,
		Updated:           common.NowTime(),
	}
	rows := global.Db.Model(&product).Updates(product).RowsAffected
	records := global.Db.First(&product, product.Id).RowsAffected
	if records > 0 {
		id := strconv.FormatUint(param.Id, 10)
		_, err := global.Es.Update().Index("product").Id(id).Doc(product).Do(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	return rows
}

// Update 后台管理前端，更新商品状态
func (p *WebProductService) UpdateStatus(param models.WebProductStatusUpdateParam) int64 {
	product := models.Product{
		Id:     param.Id,
		Status: param.Status,
	}
	rows := global.Db.Model(&product).Update("status", product.Status).RowsAffected
	records := global.Db.First(&product, product.Id).RowsAffected
	if records > 0 {
		id := strconv.FormatUint(param.Id, 10)
		_, err := global.Es.Update().Index("product").Id(id).Doc(product).Do(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	return rows
}

// GetInfo 后台管理前端，获取商品信息
func (p *WebProductService) GetInfo(param models.WebProductInfoParam) models.WebProductInfo {
	var product models.WebProductInfo
	global.Db.Table("product").First(&product, param.Id)
	return product
}

// GetList 后台管理前端，获取商品列表
func (p *WebProductService) GetList(param models.WebProductListParam) ([]models.WebProductList, int64) {
	query := &models.Product{
		Id:         param.Id,
		CategoryId: param.CategoryId,
		Title:      param.Title,
		Status:     param.Status,
	}
	productList := make([]models.WebProductList, 0)
	rows := common.RestPage(param.Page, "product", query, &productList, &[]models.Product{})
	return productList, rows
}

// GetList 微信小程序，获取商品列表
func (p *AppProductService) GetList(param models.AppProductListParam) []models.AppProductList {
	productList := make([]models.AppProductList, 0)
	if param.CategoryId != 0 {
		category := models.Category{
			ParentId: param.CategoryId,
		}
		var categoryIds []uint64
		global.Db.Table("category").Select("id").Where(category).Find(&categoryIds)
		if len(categoryIds) > 0 {
			global.Db.Table("product").Where("status = 1 and category_id in ?", categoryIds).Find(&productList)
		}
		return productList
	}
	global.Db.Table("product").Where("status = 1").Find(&productList)
	return productList
}

// GetSearchList 微信小程序，搜索商品
func (p *AppProductService) GetSearchList(param models.AppProductSearchParam) []models.AppProductList {
	productList := make([]models.AppProductList, 0)
	phraseQuery := elastic.NewMatchPhraseQuery("title", param.KeyWord)
	searchResult, err := global.Es.Search().Index("product").Query(phraseQuery).Do(ctx)
	if err != nil {
		fmt.Println(err)
	}
	for _, item := range searchResult.Each(reflect.TypeOf(models.AppProductList{})) {
		t := item.(models.AppProductList)
		productList = append(productList, t)
	}
	return productList
}

// GetDetail 微信小程序，获取商品详情
func (p *AppProductService) GetDetail(param models.AppProductDetailParam) models.AppProductDetail {
	var productDetail models.AppProductDetail
	global.Db.Table("product").First(&productDetail, param.ProductId)
	return productDetail
}
