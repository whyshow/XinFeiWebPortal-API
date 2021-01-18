package controllers

import (
	"XinFeiWebPortal-API/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

type HonorController struct {
	beego.Controller
}

// 增加荣誉
func (c *HonorController) InsertHonor() {
	honor := models.Xinfei_honor{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &honor); err == nil {
		if b, err := models.InsertHonor(honor); b {
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "添加成功"}
		} else {
			c.Data["json"] = map[string]interface{}{"code": -1, "message": err.Error()}
		}
	} else {
		c.Data["json"] = map[string]interface{}{"code": -1, "message": "参数解析错误!"}
	}
	c.ServeJSON()
}

// 删除荣誉
func (c *HonorController) DeleteHonor() {
	honor := models.Xinfei_honor{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &honor); err == nil {
		if msg, err := models.DeleteHonor(honor.Honor_id); err == nil {
			c.Data["json"] = map[string]interface{}{"code": 1, "message": msg}
		} else {
			c.Data["json"] = map[string]interface{}{"code": 0, "message": msg}
		}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "参数解析错误!"}
	}
	c.ServeJSON()
}

// 根据分类获取荣誉列表
func (c *HonorController) GetHonorList() {
	list_num := 7
	//
	if n, err := c.GetInt("n"); err == nil {
		list_num = n
	}
	page := 1
	p, err := c.GetInt("p")
	if err == nil {
		page = p
	}
	cl := c.GetString("c")
	if cl == "" {
		c.Data["json"] = map[string]interface{}{"code": -1, "message": "参数解析错误!"}
	} else {
		honorlist, num, err := models.SeleteHonorList(cl, list_num, page)
		if num > 0 {
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "查询完成", "total": num, "result": honorlist}
		} else {
			if err == nil {
				c.Data["json"] = map[string]interface{}{"code": -1, "message": "检索为空", "result": err}
			} else {
				c.Data["json"] = map[string]interface{}{"code": -1, "message": "检索错误", "result": err.Error()}
			}
		}
	}
	c.ServeJSON()
}

// 获取荣所有誉列表 不分类
func (c *HonorController) GetHonorAllList() {
	list_num := 6
	if n, err := c.GetInt("n"); err == nil {
		list_num = n
	}
	p, err := c.GetInt("p")
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": -1, "message": err.Error()}
	} else {
		if result, nums, err := models.SeleteAllHonorList(list_num, p); err == nil {
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "查询完成", "total": nums, "result": result}
		} else {
			c.Data["json"] = map[string]interface{}{"code": -1, "message": "缺少参数"}
		}
	}
	c.ServeJSON()
}

// 根据id获取荣誉列表
func (c *HonorController) GetHonor() {
	honor := models.Xinfei_honor{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &honor); err == nil {
		h, msg, err := models.SeleteHonor(honor.Honor_id)
		if err == nil {
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "查询完成", "result": h}
		} else {
			c.Data["json"] = map[string]interface{}{"code": -1, "message": msg}
		}
	} else {
		c.Data["json"] = map[string]interface{}{"code": -1, "message": "参数解析错误!"}
	}
	c.ServeJSON()
}
