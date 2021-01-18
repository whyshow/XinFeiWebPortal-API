package controllers

import (
	"XinFeiWebPortal-API/models"
	"XinFeiWebPortal-API/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

type UserController struct {
	beego.Controller
}

// 使用GET的形式进行请求 返回前端的数据包括json格式的用户列表，分页信息。
func (c *UserController) UserHome() {
	t := time.Now()
	list_num := 6
	if numlist, err := c.GetInt("num"); err == nil {
		list_num = numlist
	}
	if p, err := c.GetInt("p"); err == nil {
		//查询用户列表打包成json数据
		if user, num, err := models.UserGetAllInfo("", p, list_num, "DESC"); err == nil {
			//制作分页数据
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "请求成功", "time": utils.Millisecond(time.Since(t)), "result": user, "page": utils.Paginator(p, 6, num), "total": num}
		} else {
			c.Data["json"] = map[string]interface{}{"code": 0, "message": err, "time": utils.Millisecond(time.Since(t))}
		}
	} else {
		//查询用户列表打包成json数据
		if user, num, err := models.UserGetAllInfo("", 1, list_num, "DESC"); err == nil {
			//制作分页数据
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "请求成功", "time": utils.Millisecond(time.Since(t)), "result": user, "page": utils.Paginator(1, 6, num), "total": num}
			//显示前端页面
		} else {
			c.Data["json"] = map[string]interface{}{"code": 0, "message": err, "time": utils.Millisecond(time.Since(t))}
		}
	}
	c.ServeJSON()
}

// 内容：添加用户
// 接收正确的json数据后会自动解析数据并添加到数据库中 返回json格式的数据
func (c *UserController) UserAdd() {
	t := time.Now()
	user := models.Xinfei_user{}
	//json数据封装到user对象中
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &user); err == nil {
		//校验数据
		if user.User_account == "" || user.User_name == "" || user.User_motto == "" {
			c.Data["json"] = map[string]interface{}{"code": -1, "message": "增加失败", "time": utils.Millisecond(time.Since(t)), "result": "信息不完整"}
			c.ServeJSON()
			return
		}
		if err := models.UserInsertOne(user); err == nil {
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "增加成功", "time": utils.Millisecond(time.Since(t)), "result": "/admin/member"}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"code": -1, "message": "增加失败", "time": utils.Millisecond(time.Since(t)), "result": err}
			c.ServeJSON()
		}
	} else {
		c.Data["json"] = map[string]interface{}{"code": -1, "message": "增加失败", "time": utils.Millisecond(time.Since(t)), "result": err}
		c.ServeJSON()
	}
}

// Excel 批量导入成员信息
func (c *UserController) ExcelAddUser() {
	//接收
	f, _, err := c.GetFile("file")
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 0, "msg": "未接收到文件!", "data": nil}
		c.ServeJSON()
		return
	}
	defer f.Close()
	//判断文件夹是否存在，如果不存在就创建
	utils.IsDir("excel/model/")
	err = c.SaveToFile("file", "excel/model/成员信息.xlsx") // 保存位置在 static/upload, 没有文件夹要先创建
	if err == nil {
		// 开始处理excel信息
		//xlsx, err := excelize.OpenFile("excel/model/群发模板.xlsx")
		if err != nil {
			c.Data["json"] = map[string]interface{}{"code": 0, "msg": "表格信息错误", "data": nil}
			c.ServeJSON()
			return
		}
		//打开Excel 中的 Sheet1表
		//	rows,err :=xlsx.GetRows("Sheet1")
		//定义个结构体模型
		//	user := models.Xinfei_user{}

		c.Data["json"] = map[string]interface{}{"code": 1, "msg": "上传成功", "data": nil}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "msg": "上传失败", "data": nil}
		c.ServeJSON()
	}

}

// 内容：根据学号删除指定的人的信息
func (c *UserController) UserDelete() {
	t := time.Now()
	user := models.Xinfei_user{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	if num, err := models.UserDeleteOne(user.User_account); num == 1 {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "删除成功", "time": utils.Millisecond(time.Since(t)), "result": err}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "删除失败", "time": utils.Millisecond(time.Since(t)), "result": err}
		c.ServeJSON()
	}
}

// 修改用户密码
func (c *UserController) UserAlterPassword() {
	t := time.Now()
	if num, err := models.UserUpdatePassword(c.GetString("account"), c.GetString("password")); err == nil {
		c.Data["json"] = map[string]interface{}{"code": num, "message": "修改成功", "time": utils.Millisecond(time.Since(t)), "result": err}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "修改失败", "time": utils.Millisecond(time.Since(t)), "result": err}
		c.ServeJSON()
	}
}

// 用户的账号激活或者禁用
func (c *UserController) UserActivate() {
	t := time.Now()
	activate, _ := c.GetBool("activate", false)
	account := c.GetString("account")
	if code, err := models.UserActivate(account, activate); code == 1 {
		c.Data["json"] = map[string]interface{}{"code": code, "message": "处理成功", "time": utils.Millisecond(time.Since(t)), "result": err}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"code": code, "message": "处理失败", "time": utils.Millisecond(time.Since(t)), "result": err}
		c.ServeJSON()
	}
}

// 内容：修改用户信息
func (c *UserController) UserAlterInfo() {
	t := time.Now()
	user := models.Xinfei_user{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &user); err == nil {
		fmt.Println(user)

		//校验数据
		if user.User_account == "" || user.User_name == "" || user.User_motto == "" {
			c.Data["json"] = map[string]interface{}{"code": -1, "message": "修改失败", "time": utils.Millisecond(time.Since(t)), "result": "信息不完整"}
			c.ServeJSON()
			return
		}
		if _, err := models.UserUpdateOne(user); err == nil {
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "修改成功", "time": utils.Millisecond(time.Since(t)), "result": err}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"code": -1, "message": "修改失败", "time": utils.Millisecond(time.Since(t)), "result": err}
			c.ServeJSON()
		}
	} else {
		c.Data["json"] = map[string]interface{}{"code": -1, "message": "修改失败", "time": utils.Millisecond(time.Since(t)), "result": err}
		c.ServeJSON()
	}
}

// 内容：根据条件查询用户
func (c *UserController) UserQuery() {
	t := time.Now()
	if user, _, err := models.UserGetWhereAllInfo(c.GetString("u")); err == nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "查询成功", "time": utils.Millisecond(time.Since(t)), "result": user}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "查询失败", "time": utils.Millisecond(time.Since(t)), "result": err}
		c.ServeJSON()
	}
}

// 内容：查询一个成员信息
func (c *UserController) UserQueryOne() {
	t := time.Now()
	if user, err := models.UserGetInfoOne(c.Ctx.Input.Param(":id")); err == nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "查询成功", "time": utils.Millisecond(time.Since(t)), "result": user}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "查询失败", "time": utils.Millisecond(time.Since(t)), "result": err}
		c.ServeJSON()
	}
}

// 内容：前端查询所有用户  返回 所有用户的数据（不包含密码）
func (c *UserController) MemberAll() {
	t := time.Now()
	//	activate := c.GetString("activate")
	grade := c.GetString("grade") //升降序
	p, _ := c.GetInt("p")         //页码
	//	account := c.GetString("account")

	//按年级查询
	if grade != " " { //两个条件同时达成
		if user, nums, err := models.UserGetAllGrade(grade, p, 5); err == nil {
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "查询成功", "time": utils.Millisecond(time.Since(t)), "result": user, "page": utils.Paginator(p, 5, nums)}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"code": 0, "message": "查询失败", "time": utils.Millisecond(time.Since(t)), "result": user}
			c.ServeJSON()
		}
	}

	/*
		if desc != "" {
			if user, err := models.UserGetAllInfo(1); err == nil {
				c.Data["json"] = map[string]interface{}{"code": 1, "message": "查询成功", "time": utils.Millisecond(time.Since(t)), "result": user}
				c.ServeJSON()
			} else {
				c.Data["json"] = map[string]interface{}{"code": 0, "message": "查询失败", "time": utils.Millisecond(time.Since(t)), "result": user}
				c.ServeJSON()
			}
		} else if activate != "" {  //
			adc := "ASC"
			if desc == "desc" {
				adc = "DESC"
			} else if desc == "asc" {
				adc = "ASC"
			}
			if user, num, err := models.UserGetAllInfo(activate, 0, 0, adc); err == nil {
				if num > 1 {
					c.Data["json"] = map[string]interface{}{"code": 1, "message": "查询成功", "time": utils.Millisecond(time.Since(t)), "result": user}
					c.ServeJSON()
				} else {
					c.Data["json"] = map[string]interface{}{"code": 0, "message": "未查询到", "time": utils.Millisecond(time.Since(t)), "result": user}
					c.ServeJSON()
				}
			} else {
				c.Data["json"] = map[string]interface{}{"code": 0, "message": "查询失败", "time": utils.Millisecond(time.Since(t)), "result": err}
				c.ServeJSON()
			}
		} else {
			c.Data["json"] = map[string]interface{}{"code": 0, "message": "参数错误", "time": utils.Millisecond(time.Since(t)), "result": nil}
			c.ServeJSON()
		}*/
}
