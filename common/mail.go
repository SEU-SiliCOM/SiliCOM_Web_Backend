package common

import (
	"SilicomAPPv0.3/global"
	"github.com/go-gomail/gomail"
)

func SendMail(toAddress, title, content, toName string) error {
	if len(toName) == 0 {
		toName = "未注册的用户"
	}
	m := gomail.NewMessage()
	e := global.Config.Email
	m.SetAddressHeader("From", e.Address, e.Name)                // 发件人
	m.SetHeader("To", m.FormatAddress(toAddress, "未注册的用户"))      // 收件人
	m.SetHeader("Subject", title)                                // 主题
	m.SetBody("text/html", content)                              // 正文
	d := gomail.NewDialer(e.Host, e.Port, e.Address, e.Password) // 发送邮件服务器、端口、发件人账号、发件人密码
	return d.DialAndSend(m)
}
