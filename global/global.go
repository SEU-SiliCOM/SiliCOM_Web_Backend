package global

import (
	"SilicomAPPv0.3/config"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"github.com/olivere/elastic/v7"
)

var (
	Config config.Config
	Db     *gorm.DB
	Es     *elastic.Client
	Rdb    *redis.Client
)
