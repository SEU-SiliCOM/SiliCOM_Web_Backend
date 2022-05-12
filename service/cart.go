package service

import (
	"SilicomAPPv0.3/global"
	"SilicomAPPv0.3/models"
	"strconv"
	"strings"
)

type CartService struct {
}

// Add 微信小程序，添加购物车
func (c *CartService) Add(param models.AppCartAddParam) bool {
	key := strings.Join([]string{"user", param.UserId, "cart"}, ":")
	pid := strconv.Itoa(int(param.ProductId))
	return global.Rdb.HSetNX(ctx, key, pid, param.ProductCount).Val()
}

// Delete 微信小程序，删除购物车中的某项商品
func (c *CartService) Delete(param models.AppCartDeleteParam) int64 {
	key := strings.Join([]string{"user", param.UserId, "cart"}, ":")
	return global.Rdb.HDel(ctx, key, param.ProductId).Val()
}

// Clear 微信小程序，清楚购物车中的商品
func (c *CartService) Clear(param models.AppCartClearParam) int64 {
	key := strings.Join([]string{"user", param.UserId, "cart"}, ":")
	pidsAndCounts := global.Rdb.HGetAll(ctx, key).Val()
	var rows int64
	for id, _ := range pidsAndCounts {
		rows += global.Rdb.HDel(ctx, key, id).Val()
	}
	return rows
}

// GetInfo 微信小程序，获取购物车信息
func (c *CartService) GetInfo(param models.AppCartQueryParam) models.AppCartInfo {
	var cartInfo models.AppCartInfo
	key := strings.Join([]string{"user", param.UserId, "cart"}, ":")
	productIdsAndCounts := global.Rdb.HGetAll(ctx, key).Val()
	productIds := make([]uint64, 0)
	idsAndCounts := make(map[uint64]int, 0)
	for pid, pcount := range productIdsAndCounts {
		id, _ := strconv.Atoi(pid)
		count, _ := strconv.Atoi(pcount)
		productIds = append(productIds, uint64(id))
		idsAndCounts[uint64(id)] = count
	}
	if len(productIds) > 0 {
		global.Db.Table("product").Find(&cartInfo.CartItem, productIds)
		for index, item := range cartInfo.CartItem {
			cartInfo.CartItem[index].Count = idsAndCounts[item.Id]
			cartInfo.TotalPrice = cartInfo.TotalPrice + item.Price*float64(idsAndCounts[item.Id])
		}
		return cartInfo
	}
	return cartInfo
}
