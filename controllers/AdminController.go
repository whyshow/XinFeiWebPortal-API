package controllers

import (
	"XinFeiWebPortal-API/middleware"
	"XinFeiWebPortal-API/models"
	"XinFeiWebPortal-API/utils"
	"encoding/json"
	"github.com/astaxie/beego"
	"time"
)

type AdminController struct {
	beego.Controller
}

// 验证登录token是否正确
func (c *AdminController) Token() {
	// 如果能请求到这个方法说明是已经登录过的
	c.Data["json"] = map[string]interface{}{"status": 200, "message": "登录正常!"}
	c.ServeJSON()
}
func (c *AdminController) RedirectHome() {
	c.Redirect("/admin", 302)

}

// 基础信息统计显示页面
func (c *AdminController) Welcome() {
	t := time.Now()
	sy := models.GetSystemStatus()
	c.Data["UserCount"] = models.GetUserCount()
	c.Data["ArticleCount"] = models.GetArticleCount()
	c.Data["System"] = sy
	c.Data["time"] = utils.Millisecond(time.Since(t))
	c.ServeJSON()
}

// 获取系统状态信息
func (c *AdminController) GetSystemStatus() {
	t := time.Now()
	sy := models.GetSystemStatus()
	c.Data["json"] = map[string]interface{}{"code": 1, "message": "请求成功", "time": utils.Millisecond(time.Since(t)), "result": sy}
	c.ServeJSON()
}

// 管理员登录操作
func (c *AdminController) AdminLogin() {
	t := time.Now()
	var admin models.Xinfei_Admin
	json.Unmarshal(c.Ctx.Input.RequestBody, &admin)
	if bl, msg, admin := models.AdminLogin(admin); bl {
		// 登录成功
		// 生成token
		token := utils.MD5(utils.Random(5))
		// 将登录成功生成的token写进redis中
		middleware.RedisExec("Set", admin.Admin_name, token)
		// 设置过期时间
		middleware.RedisExpire(admin.Admin_name, 30)
		message := map[string]interface{}{"name": admin.Admin_name, "token": token}
		//加入session
		//c.SetSession("admin_id", admin.Admin_id)
		c.Data["json"] = map[string]interface{}{"code": 1, "message": message, "result": "/admin/", "time": utils.Millisecond(time.Since(t))}
		c.ServeJSON()
	} else {
		// 登录失败
		c.Data["json"] = map[string]interface{}{"code": 0, "message": msg, "result": nil, "time": utils.Millisecond(time.Since(t))}
		c.ServeJSON()
	}
}

// 退出登录
func (c *AdminController) ExitLogin() {
	c.DelSession("admin_id")
	c.Data["json"] = map[string]interface{}{"code": 1, "message": "退出成功", "result": "/admin/login"}
	c.ServeJSON()

}

// 重新启动系统
// 测试中
func (c *AdminController) RebootSystem() {

}
