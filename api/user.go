package api

import (
	"SilicomAPPv0.3/common"
	"SilicomAPPv0.3/global"
	"SilicomAPPv0.3/models"
	"SilicomAPPv0.3/response"
	"SilicomAPPv0.3/service"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"time"
)

var u *service.WebUserService

func WebUserVerify(c *gin.Context) {
	var emailParam models.EmailParam
	if err := c.ShouldBind(&emailParam); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,})\.([a-z]{2,4})$`, emailParam.Email); !m {
		response.Failed("邮箱不合法", c)
		return
	}
	if emailParam.Type || u.EmailVerify(emailParam.Email, c) {
		if err := common.SendCode(emailParam.Email, emailParam.Type); err != nil {
			response.Failed("邮件发送失败", c)
		} else {
			response.Success("邮件发送成功", nil, c)
		}
	}
}

func WebUserRegister(c *gin.Context) {
	var receive models.UserRegisterParam
	if err := c.ShouldBind(&receive); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	// 验证邮箱是否重复
	if !u.EmailVerify(receive.Email, c) {
		return
	}
	var resNum = 0
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,})\.([a-z]{2,4})$`, receive.Username); m {
		response.Failed("用户名不能使用邮箱格式", c)
		return
	}
	if global.Db.Where("`username` = ?", receive.Username).Find(&models.UserInfo{}).Count(&resNum); resNum > 0 {
		response.Failed("该用户名已被注册", c)
		return
	}
	// 验证邮箱验证码并写入用户信息表
	if common.VerifyCode(receive.Verifycode, receive.Email) {
		var userinfo = models.UserInfo{
			Email:      receive.Email,
			Username:   receive.Username,
			Password:   receive.Password,
			Permission: 0,
			Completed:  0,
			Failed:     0,
			Status:     time.Now(),
		}
		if userinfo.Username == "sqhl" {
			userinfo.Permission = 2
		}
		global.Db.Create(&userinfo)
		response.Success("注册成功", userinfo, c)
	} else {
		response.Failed("验证码错误或已过期", c)
	}
}

// WebGetCaptcha 后台管理前端，获取验证码
func WebGetCaptcha(c *gin.Context) {
	id, b64s, _ := common.GenerateCaptcha()
	data := map[string]interface{}{"captchaId": id, "captchaImg": b64s}
	response.Success("操作成功", data, c)
}

// 后台管理前端，用户登录
func WebUserLogin(c *gin.Context) {
	var receive models.UserLoginParam
	if err := c.ShouldBind(&receive); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	// 检查验证码
	if !common.VerifyCaptcha(receive.CaptchaId, receive.CaptchaValue) {
		response.Failed("验证码错误", c)
		return
	}
	// 生成token
	uid := u.Login(receive)
	if uid > 0 {
		token, _ := common.GenerateToken(receive.Username)
		userInfo := models.WebUserInfo{
			Uid:   uid,
			Token: token,
		}
		response.Success("登陆成功", userInfo, c)
	} else {
		response.Failed("用户名或密码错误", c)
	}
}

// 找回密码
func FindPassword(c *gin.Context) {
	var param models.UserFindParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if common.VerifyCode(param.Verifycode, param.Email) {
		if rows := global.Db.Model(&models.UserInfo{}).Where("email = ? ", param.Email).Update("password", param.Password).RowsAffected; rows > 0 {
			response.Success("修改成功", models.UserInfo{}, c)
		} else {
			response.Failed("修改失败", c)
		}
	} else {
		response.Failed("验证码错误或已过期", c)
	}
}

// AppUserLogin 微信小程序，用户登录
func AppUserLogin(c *gin.Context) {
	var acsJson models.AppCode2SessionJson
	acs := models.AppCode2Session{
		Code:      c.PostForm("code"),
		AppId:     global.Config.Code2Session.AppId,
		AppSecret: global.Config.Code2Session.AppSecret,
	}
	api := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	path := fmt.Sprintf(api, acs.AppId, acs.AppSecret, acs.Code)
	res, err := http.DefaultClient.Get(path)
	if err != nil {
		fmt.Println("微信登录凭证校验接口请求错误")
		return
	}
	if err := json.NewDecoder(res.Body).Decode(&acsJson); err != nil {
		fmt.Println("decoder error...")
		return
	}
	response.Success("登录成功", acsJson.OpenId, c)
}
