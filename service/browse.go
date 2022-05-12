package service

import (
	"SilicomAPPv0.3/global"
	"SilicomAPPv0.3/models"
	"strconv"
	"strings"
	"time"
)

type BrowseService struct {
}

// Save 微信小程序，保存浏览记录
func (c *BrowseService) Save(param models.AppBrowseSaveParam) float64 {
	key := strings.Join([]string{"user", param.UserId, "browse"}, ":")
	now := float64(time.Now().Unix())
	pid := strconv.Itoa(int(param.ProductId))
	if count := global.Rdb.ZIncrBy(ctx, key, now, pid).Val(); count > 0 {
		return count
	}
	return 0
}

// Delete 微信小程序，删除浏览记录
func (c *BrowseService) Delete(param models.AppBrowseDeleteParam) int64 {
	key := strings.Join([]string{"user", param.UserId, "browse"}, ":")
	return global.Rdb.ZRem(ctx, key, param.ProductIds).Val()
}

// GetList 微信小程序，获取浏览记录列表
func (c *BrowseService) GetList(param models.AppBrowseQueryParam) []models.AppBrowseRecordList {
	key := strings.Join([]string{"user", param.UserId, "browse"}, ":")

	// 删除30天前的浏览记录
	before30Day := time.Now().AddDate(0, 0, -30).Unix()
	global.Rdb.ZRemRangeByScore(ctx, key, "0", strconv.FormatInt(before30Day, 10))

	// 获取近30天所有浏览记录的商品id和Scores, 查询浏览记录的商品列表的信息
	scoresAndMembers := global.Rdb.ZRevRangeWithScores(ctx, key, 0, -1).Val()
	var browseList []models.AppBrowseRecordList
	for _, v := range scoresAndMembers {
		var browseItem models.AppBrowseItem
		browseTime := time.Unix(int64(v.Score), 0).Format("01月02日")
		global.Db.Table("product").Find(&browseItem, v.Member)
		browseInfo := models.AppBrowseRecordList{BrowseTime: browseTime, BrowseItem: browseItem}
		browseList = append(browseList, browseInfo)
	}
	return browseList
}
