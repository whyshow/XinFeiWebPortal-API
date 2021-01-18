package controllers

import (
	"XinFeiWebPortal-API/models"
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/astaxie/beego"
)

type MailController struct {
	beego.Controller
}

// 发送一封邮件
func (c *MailController) SendEmaiOne() {
	sendemail := models.Sendemail{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &sendemail)
	if err := models.SendGoMail([]string{sendemail.Addressee}, sendemail.Subject, sendemail.Mailtext); err == nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "发送成功"}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": err.Error()}
	}
	c.ServeJSON()
}

type Sendemails struct {
	Subject  string
	Mailtext string
}
type Emails struct {
	student_name string
	student_mail string
}

// 群发邮件
func (c *MailController) SendEmais() {
	var nums = 0
	sendemails := Sendemails{}
	// 获取解析后的群发名单
	//读取Excel
	xlsx, err := excelize.OpenFile("excel/model/群发模板.xlsx")
	//是否打开时出现错误 如果错误则退出
	if err != nil {
		fmt.Println(err)
		c.Data["json"] = map[string]interface{}{"code": 0, "message": err.Error()}
	} else {
		//打开Excel 中的 Sheet1表
		rows, err := xlsx.GetRows("Sheet1")
		if err != nil {
			c.Data["json"] = map[string]interface{}{"code": 0, "message": err.Error()}
			c.ServeJSON()
		}
		emails := []Emails{}
		email := Emails{}
		//遍历这张表格有多少行数据，并一行行循环取出
		for index, row := range rows {
			// 将对应的数据放进切片中
			if index != 0 {
				email.student_name = row[0] //这一行数据中的第一列数据
				email.student_mail = row[1] //这一行数据中的第二列数据
				emails = append(emails, email)
			}
		}
		json.Unmarshal(c.Ctx.Input.RequestBody, &sendemails)

		for _, rows := range emails {

			err := models.SendGoMail([]string{rows.student_mail}, rows.student_name+"同学:"+sendemails.Subject, sendemails.Mailtext)
			if err == nil {
				nums = nums + 1
			}
		}

		c.Data["json"] = map[string]interface{}{"code": 1, "message": "发送完成", "nums": nums}
	}
	c.ServeJSON()
}

// 定时群发邮件
func (c *MailController) TimerSendEmals() {

}
