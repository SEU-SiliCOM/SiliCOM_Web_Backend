package initialize

import (
	"SilicomAPPv0.3/global"
	"SilicomAPPv0.3/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

func Mysql() {
	m := global.Config.Mysql
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.Username, m.Password, m.Ip, m.Port, m.Database)
	db, err := gorm.Open("mysql", dsn)
	global.Db = db
	if err != nil {
		fmt.Printf("mysql error: %s", err)
		return
	}
	sqlDb := db.DB()
	if err != nil {
		fmt.Printf("mysql error2: %s", err)
	}
	db.SingularTable(true)
	//if global.Config.Rm_rf {
	//	db.DropTableIfExists(&models.Verifycode{}, &models.Category{}, &models.Activity{}, &models.Address{})
	//}
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Verifycode{}, &models.UserInfo{}, &models.Activity{}, &models.Appointment{})
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour)
}
