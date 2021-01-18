package middleware

import (
	"github.com/astaxie/beego"
	"gopkg.in/gomail.v2"
	"strconv"
)

type Mail struct {
	// 邮件服务器地址
	MAIL_HOST string
	// 端口
	MAIL_PORT int
	// 发送邮件用户账号
	MAIL_USER string
	// 授权密码
	MAIL_PWD string
}

var _Mail Mail

func init() {
	_Mail.MAIL_HOST = beego.AppConfig.String("mail.host")
	_Mail.MAIL_PORT, _ = strconv.Atoi(beego.AppConfig.String("mail.port"))
	_Mail.MAIL_USER = beego.AppConfig.String("mail.user")
	_Mail.MAIL_PWD = beego.AppConfig.String("mail.password")

}
func SendEMail(mailAddress []string, subject string, body string) error {

	m := gomail.NewMessage()
	nickname := beego.AppConfig.String("mail.name")
	m.SetHeader("From", m.FormatAddress("message@ccit.club", nickname))
	// 发送给多个用户
	m.SetHeader("To", mailAddress...)
	// 设置邮件主题
	m.SetHeader("Subject", subject)
	// 设置邮件正文
	m.SetBody("text/html", body)
	d := gomail.NewDialer(_Mail.MAIL_HOST, _Mail.MAIL_PORT, _Mail.MAIL_USER, _Mail.MAIL_PWD)
	// 发送邮件
	err := d.DialAndSend(m)
	return err
}
