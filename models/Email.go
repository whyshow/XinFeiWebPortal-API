package models

import (
	"github.com/astaxie/beego"
	"gopkg.in/gomail.v2"
	"strconv"
)

type Email struct {
	// 邮件服务器地址
	MAIL_HOST string
	// 端口
	MAIL_PORT int
	// 发送邮件用户账号
	MAIL_USER string
	// 授权密码
	MAIL_PWD string
}
type Sendemail struct {
	Addressee string
	Subject   string
	Mailtext  string
}

var emai Email

func init() {
	emai.MAIL_HOST = beego.AppConfig.String("mail.host")
	emai.MAIL_PORT, _ = strconv.Atoi(beego.AppConfig.String("mail.port"))
	emai.MAIL_USER = beego.AppConfig.String("mail.user")
	emai.MAIL_PWD = beego.AppConfig.String("mail.password")
}

/*
title 使用gomail发送邮件
@param []string mailAddress 收件人邮箱
@param string subject 邮件主题
@param string body 邮件内容
@return error
*/
func SendGoMail(mailAddress []string, subject string, body string) error {
	m := gomail.NewMessage()
	nickname := beego.AppConfig.String("mail.name")
	m.SetHeader("From", m.FormatAddress("message@ccit.club", nickname))
	// 发送给多个用户
	m.SetHeader("To", mailAddress...)
	// 设置邮件主题
	m.SetHeader("Subject", subject)
	// 设置邮件正文
	m.SetBody("text/html", body)
	d := gomail.NewDialer(emai.MAIL_HOST, emai.MAIL_PORT, emai.MAIL_USER, emai.MAIL_PWD)
	// 发送邮件
	err := d.DialAndSend(m)
	return err
}
