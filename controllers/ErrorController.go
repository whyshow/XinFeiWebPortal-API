package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

/**
  该控制器处理页面错误请求
*/
type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error401() {
	c.Data["content"] = map[string]interface{}{"status": 401, "message": "需要进行身份验证"}
	c.ServeJSON()
}
func (c *ErrorController) Error403() {
	c.Data["content"] = map[string]interface{}{"status": 403, "message": "服务器拒绝请求"}
	c.ServeJSON()
}
func (c *ErrorController) Error404() {
	fmt.Println("123")
	c.Data["content"] = map[string]interface{}{"status": 404, "message": "找不到请求对象"}
	c.ServeJSON()
}

func (c *ErrorController) Error500() {
	c.Data["content"] = map[string]interface{}{"status": 500, "message": "服务器错误"}
	c.ServeJSON()
}
func (c *ErrorController) Error503() {
	c.Data["content"] = map[string]interface{}{"status": 503, "message": "服务器目前无法使用(由于超载或停机维护)"}
	c.ServeJSON()
}
