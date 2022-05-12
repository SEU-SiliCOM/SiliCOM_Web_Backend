package service

import (
	"SilicomAPPv0.3/global"
	"SilicomAPPv0.3/models"
	"strconv"
	"strings"
)

type CollectionService struct {
}

// Add 微信小程序，添加收藏的商品
func (c *CollectionService) Add(param models.AppCollectionAddParam) int64 {
	key := strings.Join([]string{"user", param.UserId, "collect"}, ":")
	return global.Rdb.SAdd(ctx, key, param.ProductId).Val()
}

// Delete 微信小程序，删除收藏的商品
func (c *CollectionService) Delete(param models.AppCollectionDeleteParam) int64 {
	key := strings.Join([]string{"user", param.UserId, "collect"}, ":")
	pids := global.Rdb.SMembers(ctx, key).Val()
	return global.Rdb.SRem(ctx, key, pids).Val()
}

// GetList 微信小程序，获取收藏的商品列表
func (c *CollectionService) GetList(param models.AppCollectionQueryParam) []models.AppCollectionList {
	var items []models.AppCollectionList
	key := strings.Join([]string{"user", param.UserId, "collect"}, ":")
	pids := global.Rdb.SMembers(ctx, key).Val()
	productIds := make([]uint64, 0)
	for _, pid := range pids {
		id, _ := strconv.Atoi(pid)
		productIds = append(productIds, uint64(id))
	}
	if len(productIds) > 0 {
		global.Db.Table("product").Find(&items, productIds)
		return items
	}
	return items
}
