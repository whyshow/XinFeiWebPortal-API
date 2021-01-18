package controllers

import (
	"XinFeiWebPortal-API/middleware"
	"XinFeiWebPortal-API/models"
	"XinFeiWebPortal-API/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

type ArticleController struct {
	beego.Controller
}

var num = 10 //每页文章显示10行
/**
 * 文章列表首页
 */
func (c *ArticleController) ArticleHome() {
	t := time.Now()
	list_num := 6
	if numlist, err := c.GetInt("num"); err == nil {
		list_num = numlist
	}
	if p, err := c.GetInt("p"); err == nil {
		if article, nums, err := models.ArticleSeleteAll(p, list_num, "DESC", false, false); err == nil {
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "请求成功", "time": utils.Millisecond(time.Since(t)), "result": article, "page": utils.Paginator(p, list_num, nums), "total": nums}
		} else {
			c.Data["json"] = map[string]interface{}{"code": 0, "message": err, "time": utils.Millisecond(time.Since(t))}
		}
	} else {
		if article, nums, err := models.ArticleSeleteAll(1, num, "DESC", false, false); err == nil {
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "请求成功", "time": utils.Millisecond(time.Since(t)), "result": article, "page": utils.Paginator(p, list_num, nums), "total": nums}
		} else {
			c.Data["json"] = map[string]interface{}{"code": 0, "message": err, "time": utils.Millisecond(time.Since(t))}
		}
	}
	c.ServeJSON()

}

// 添加一篇文章
func (c *ArticleController) ArticleAddOne() {
	t := time.Now()
	article := models.Xinfei_article{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &article); err == nil {
		fmt.Println(article)
		nameRune := []rune(article.Article_text)
		// 文章文本取前100字
		if len(nameRune) > 150 {
			article.Article_text = string(nameRune[0:150])
		} else {
			article.Article_text = string(nameRune[0:len(nameRune)])
		}
		if err := models.ArticleInsertOne(article); err == nil {
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "添加成功", "time": utils.Millisecond(time.Since(t)), "result": err}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"code": 0, "message": "错误", "time": utils.Millisecond(time.Since(t)), "result": err}
			c.ServeJSON()
		}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "参数错误", "time": utils.Millisecond(time.Since(t)), "result": err}
		c.ServeJSON()
	}
}

//修改文章
func (c *ArticleController) ArticleAlter() {
	t := time.Now()
	article := models.Xinfei_article{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &article); err == nil {
		nameRune := []rune(article.Article_text)
		// 文章文本取前100字
		if len(nameRune) > 150 {
			article.Article_text = string(nameRune[0:150])
		} else {
			article.Article_text = string(nameRune[0:len(nameRune)])
		}
		if _, err := models.ArticleUpdateOne(article); err == nil {
			atlJson, _ := json.Marshal(article)
			middleware.RedisExec("Set", article.Article_id, atlJson)
			middleware.RedisExpire(article.Article_id, 1)
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "修改成功", "time": utils.Millisecond(time.Since(t)), "result": err}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"code": 0, "message": "错误", "time": utils.Millisecond(time.Since(t)), "result": err.Error()}
			c.ServeJSON()
		}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "参数错误", "time": utils.Millisecond(time.Since(t)), "result": err}
		c.ServeJSON()
	}
}

//根据文章id删除一篇文章
func (c *ArticleController) ArticleDeleteOne() {
	t := time.Now()
	article := models.Xinfei_article{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &article)
	if _, err := models.ArticleDeleteOne(article.Article_id); err == nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "删除成功", "time": utils.Millisecond(time.Since(t)), "result": err}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "删除失败", "time": utils.Millisecond(time.Since(t)), "result": err}
		c.ServeJSON()
	}
}

//根据文章id修改一篇文章
func (c *ArticleController) ArticleAlterOne() {

}

// 查询文章详情内容
func (c *ArticleController) ArticleQueryOne() {
	t := time.Now()
	if article, err := models.ArticleSeleteOne(c.Ctx.Input.Param(":id")); err == nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "查询成功", "time": utils.Millisecond(time.Since(t)), "result": article}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "查询失败", "time": utils.Millisecond(time.Since(t)), "result": err}
		c.ServeJSON()
	}
}

// 查询所有文章(按日期查询)
func (c *ArticleController) ArticleQueryAllDate() {
	t := time.Now()
	if p, err := c.GetInt("p"); err == nil {
		if article, nums, err := models.ArticleSeleteAll(p, num, "DESC", false, false); err == nil {
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "查询成功", "time": utils.Millisecond(time.Since(t)), "result": article, "page": utils.Paginator(p, num, nums)}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"code": 0, "message": "查询失败", "time": utils.Millisecond(time.Since(t)), "result": err}
			c.ServeJSON()
		}
	} else {
		if article, nums, err := models.ArticleSeleteAll(1, num, "DESC", false, false); err == nil {
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "查询成功", "time": utils.Millisecond(time.Since(t)), "result": article, "page": utils.Paginator(p, num, nums)}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"code": 0, "message": "查询失败", "time": utils.Millisecond(time.Since(t)), "result": err}
			c.ServeJSON()
		}
	}
}

// 查询所有文章(按热度查询)
func (c *ArticleController) ArticleQueryAllHot() {
	t := time.Now()
	if p, err := c.GetInt("p"); err == nil {
		if article, nums, err := models.ArticleSeleteAll(p, num, "DESC", true, false); err == nil {
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "查询成功", "time": utils.Millisecond(time.Since(t)), "result": article, "page": utils.Paginator(p, num, nums)}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"code": 0, "message": "查询失败", "time": utils.Millisecond(time.Since(t)), "result": err}
			c.ServeJSON()
		}
	} else {
		if article, nums, err := models.ArticleSeleteAll(1, num, "DESC", true, false); err == nil {
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "查询成功", "time": utils.Millisecond(time.Since(t)), "result": article, "page": utils.Paginator(p, num, nums)}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"code": 0, "message": "查询失败", "time": utils.Millisecond(time.Since(t)), "result": err}
			c.ServeJSON()
		}
	}
}

// 上下架一篇文章
//func (c * ArticleController) ArticleActivityOne() {
//	t := time.Now()
//
//	if mun,err:= models.ArticleActivityOne(c.Ctx.Input.Param(":id"));err == nil{
//		c.Data["json"] = map[string]interface{}{"code": mun, "message": "上架成功", "time": utils.Millisecond(time.Since(t)), "result": nil}
//		c.ServeJSON()
//	}else {
//		c.Data["json"] = map[string]interface{}{"code": mun, "message": "下架成功", "time": utils.Millisecond(time.Since(t)), "result": nil}
//		c.ServeJSON()
//	}
//
//}
