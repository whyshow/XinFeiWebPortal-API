package main

import (
	"XinFeiWebPortal-API/controllers"
	"XinFeiWebPortal-API/middleware"
	"XinFeiWebPortal-API/models"
	_ "XinFeiWebPortal-API/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

func init() {
	url := beego.AppConfig.String("jdbc.url")
	port := beego.AppConfig.String("jdbc.port")
	databasename := beego.AppConfig.String("jdbc.databasename")
	username := beego.AppConfig.String("jdbc.username")
	password := beego.AppConfig.String("jdbc.password")
	orm.RegisterDataBase("default", "mysql", username+":"+password+"@tcp("+url+":"+port+")/"+databasename+"?charset=utf8")
	orm.RegisterModel(new(models.Xinfei_Admin))
	orm.RegisterModel(new(models.Xinfei_Day))
	orm.RegisterModel(new(models.Xinfei_user))
	orm.RegisterModel(new(models.Xinfei_article))
	orm.RegisterModel(new(models.Xinfei_honor))
}

func main() {
	//Config()
	//注册新的错误设置 404/401/502
	beego.ErrorController(&controllers.ErrorController{})
	//InsertFilter是提供一个过滤函数
	// 跨域
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "usertoken", "username", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "usertoken", "username", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	//注册过滤器  判断在用户访问admin后台时是否处于登录状态
	beego.InsertFilter("/admin/*", beego.BeforeRouter, FilterUser)
	beego.Run()
}

//管理后台的登录过滤中间件
var FilterUser = func(ctx *context.Context) {
	//判断是否是移动端
	keywords := []string{"Android", "iPhone", "iPod", "Mobile", "Windows Phone", "MQQBrowser"}
	for i := 0; i < len(keywords); i++ {
		if strings.Contains(ctx.Request.UserAgent(), keywords[i]) {
			fmt.Println("你现在使用的是" + keywords[i] + "端在访问，请使用PC端访问!")
			ctx.RenderMethodResult(map[string]interface{}{"status": 401, "message": "你现在使用的是" + keywords[i] + "端在访问，请使用PC端访问!"})
			return
		}
	}
	// 获取head头数据
	token := ctx.Request.Header["Usertoken"]
	name := ctx.Request.Header["Username"]
	//fmt.Println("数据类型为"+reflect.TypeOf(name[0]).String())
	// 判断redis中是否存在 key/value
	if token[0] == "null" || name[0] == "null" {
		ctx.RenderMethodResult(map[string]interface{}{"status": 401, "message": "请先登录"})
	} else if tokens, err := middleware.RedisExec("Get", name[0]); err != nil {
		ctx.RenderMethodResult(map[string]interface{}{"status": 401, "message": "请先登录"})
	} else {
		if tokens != token[0] {
			ctx.RenderMethodResult(map[string]interface{}{"status": 401, "message": "登录已过期,请重新登录!"})
		}
	}
}
