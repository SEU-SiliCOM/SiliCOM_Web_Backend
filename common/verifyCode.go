package common

import (
	"SilicomAPPv0.3/global"
	"SilicomAPPv0.3/models"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func SendCode(toAddress string, flag bool) error {
	code := ""
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 6; i++ {
		code += strconv.Itoa(rd.Intn(10))
	}
	if !flag {
		if err := SendMail(toAddress, "欢迎注册，请查看您的验证码", code+"，有效期五分钟", "未注册的用户"); err != nil {
			return err
		}
	} else {
		if err := SendMail(toAddress, "密码找回，请查看您的验证码", code+"，有效期五分钟", "忘记密码的用户"); err != nil {
			return err
		}
	}

	// 验证码写入数据库
	verifyCode := models.Verifycode{
		Email:     toAddress,
		Code:      code,
		InvalidAt: time.Now().Add(5 * time.Minute),
	}
	global.Db.Create(&verifyCode)
	fmt.Println(code, " to ", toAddress)
	return nil
}

func VerifyCode(code string, email string) bool {
	var records []models.Verifycode
	global.Db.Table("verifycode").Select("`invalid_at`").Where("code = ? and email = ?", code, email).Scan(&records)
	for _, record := range records {
		if time.Now().Sub(record.InvalidAt) < 0 {
			return true
		}
	}
	return false
}
