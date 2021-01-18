package controllers

import (
	"XinFeiWebPortal-API/middleware"
	"XinFeiWebPortal-API/models"
	"XinFeiWebPortal-API/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"os"
	"time"
)

type ApiController struct {
	beego.Controller
}

var jsons = map[string]interface{}{}

func MakeJson() {

}

//查询文章详情
func (c *ApiController) QueryOneArticle() {
	// 接收id
	id := c.Ctx.Input.Param(":id")
	result, err := middleware.RedisExec("Get", id)
	//判断是否下架
	if err == nil {
		// 反序列化
		article := models.Xinfei_article{}
		json.Unmarshal([]byte(result), &article)
		if article.Article_display == "1" {
			fmt.Println("redis")
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "查询成功", "result": article}
		} else {
			c.Data["json"] = map[string]interface{}{"code": -1, "message": "动态已下架", "result": nil}
		}
	} else {
		if article, err := models.ArticleSeleteOne(id); err == nil {
			if article.Article_display == "0" {
				c.Data["json"] = map[string]interface{}{"code": -1, "message": "动态已下架", "result": nil}
			} else {
				fmt.Println("mysql")
				atlJson, _ := json.Marshal(article)
				middleware.RedisExec("Set", id, atlJson)
				middleware.RedisExpire(id, 1)
				c.Data["json"] = map[string]interface{}{"code": 1, "message": "查询成功", "result": article}
			}
		} else {
			c.Data["json"] = map[string]interface{}{"code": 0, "message": "查询失败", "result": err}
		}
	}
	c.ServeJSON()
}

//查询文章列表（按日期）
func (c *ApiController) QueryArticleList() {
	t := time.Now()
	if p, err := c.GetInt("p"); err == nil {
		if article, nums, err := models.ArticleSeleteAll(p, 6, "DESC", false, true); err == nil {
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "查询成功", "time": utils.Millisecond(time.Since(t)), "result": article, "page": utils.Paginator(p, 6, nums)}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"code": 0, "message": "查询失败", "time": utils.Millisecond(time.Since(t)), "result": err}
			c.ServeJSON()
		}
	} else {
		if article, nums, err := models.ArticleSeleteAll(1, 6, "DESC", false, true); err == nil {
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "查询成功", "time": utils.Millisecond(time.Since(t)), "result": article, "page": utils.Paginator(p, 6, nums)}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"code": 0, "message": "查询失败", "time": utils.Millisecond(time.Since(t)), "result": err}
			c.ServeJSON()
		}
	}
}

//查询文章列表（按热度值）
func (c *ApiController) QueryArticleListHot() {
	t := time.Now()
	if p, err := c.GetInt("p"); err == nil {
		if article, nums, err := models.ArticleSeleteAll(p, 6, "DESC", true, true); err == nil {
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "查询成功", "time": utils.Millisecond(time.Since(t)), "result": article, "page": utils.Paginator(p, 6, nums)}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"code": 0, "message": "查询失败", "time": utils.Millisecond(time.Since(t)), "result": err}
			c.ServeJSON()
		}
	} else {
		if article, nums, err := models.ArticleSeleteAll(1, 6, "DESC", true, true); err == nil {
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "查询成功", "time": utils.Millisecond(time.Since(t)), "result": article, "page": utils.Paginator(p, 6, nums)}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"code": 0, "message": "查询失败", "time": utils.Millisecond(time.Since(t)), "result": err}
			c.ServeJSON()
		}
	}
}

var ab = models.AppBar{}

// 获取导航栏信息
func (c *ApiController) GetAppBar() {
	t := time.Now()
	if ab.Appbar == nil {
		var appb = models.AppBar{}
		func() {}()
		{
			f, err := os.Open("./conf/appbar.json")
			defer f.Close()
			if err != nil {
				return
			}
			bytes, _ := ioutil.ReadAll(f)
			err = json.Unmarshal(bytes, &appb)
			if err != nil {
				fmt.Println("ERROR: ", err.Error())
				return
			}
			ab = appb
		}
	}
	c.Data["json"] = map[string]interface{}{"code": 1, "message": "查询成功", "data": ab, "time": utils.Millisecond(time.Since(t))}
	c.ServeJSON()
}

// 获取荣誉分类
func (c *ApiController) GetHonorClassify() {
	if hc, err := models.HonorClassify(); err == nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "查询完成", "result": hc}
	} else {
		c.Data["json"] = map[string]interface{}{"code": -1, "message": err.Error()}
	}
	c.ServeJSON()
}

// 获取成员列表
func (c *ApiController) MemberAll() {
	grade := c.GetString("grade")
	if grade == "" {
		user, _, err := models.UserGetAllInfo("", 0, 0, "")
		if err == nil {
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "查询成功", "result": user}
		} else {
			c.Data["json"] = map[string]interface{}{"code": 0, "message": "查询失败", "result": err}
		}
	} else {
		if user, _, err := models.UserGetWhereAllInfo(grade); err == nil {
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "查询成功", "result": user}
		} else {
			c.Data["json"] = map[string]interface{}{"code": 0, "message": "查询失败", "result": err}
		}
	}
	c.ServeJSON()
}

// 查询成员入学年级
func (c *ApiController) MemberGrade() {
	grade, max, min, err := models.UserGradeList()
	if err == nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "查询成功", "result": grade, "max": max, "min": min}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "查询失败", "result": err, "max": nil, "min": nil}
	}
	c.ServeJSON()
}
