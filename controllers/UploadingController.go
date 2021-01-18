package controllers

import (
	"XinFeiWebPortal-API/utils"
	"github.com/astaxie/beego"
	"image"
	"os"
	"time"
)

type UploadingController struct {
	beego.Controller
}

var baseUrl = "http://image.ccit.club"

//var baseUrl = "http://127.0.0.1:8080"

// 上传图片
func (c *UploadingController) UploadingImage() {
	f, h, err := c.GetFile("uploadimage")
	if err != nil {
		c.Data["json"] = map[string]interface{}{"errno": 1, "data": ""}
		c.ServeJSON()
		return
	}
	defer f.Close()
	timeUnix := time.Now().Unix() //已知的时间戳
	formatTimeStr := time.Unix(timeUnix, 0).Format("2006-01-02")
	//判断文件夹是否存在，如果不存在就创建
	utils.IsDir("static/upload/" + formatTimeStr + "/")
	c.SaveToFile("uploadimage", "static/upload/"+formatTimeStr+"/"+h.Filename) // 保存位置在 static/upload, 没有文件夹要先创建
	file, _ := os.Open("static/upload/" + formatTimeStr + "/" + h.Filename)
	_, imgtype, _ := image.Decode(file)
	if imgtype == "" {
		var data = []string{}
		msg := baseUrl + "/static/upload/" + formatTimeStr + "/" + h.Filename
		data = append(data, msg)
		c.Data["json"] = map[string]interface{}{"errno": 0, "data": data}
	} else {
		//压缩图片
		config, _, _ := image.DecodeConfig(f)
		// 根据图片宽高来裁剪图片
		// 默认的横图纵横比为 w 4 : h 3 即 640＊480
		// 默认的竖图纵横比为 w 0.618 : h 1 即 475 : 768
		if config.Width/config.Height > 2 {
			utils.Thumbnail("static/upload/"+formatTimeStr+"/"+h.Filename, uint(config.Width), uint(config.Height))
		} else if config.Width < 640 && config.Height < 480 && config.Width > config.Height {

			utils.Thumbnail("static/upload/"+formatTimeStr+"/"+h.Filename, uint(config.Width), uint(config.Height))
		} else if config.Width < 475 && config.Height < 768 && config.Width < config.Height {
			utils.Thumbnail("static/upload/"+formatTimeStr+"/"+h.Filename, uint(config.Width), uint(config.Height))
		} else if config.Width > config.Height {
			// 认为图片为 4 : 3
			utils.Thumbnail("static/upload/"+formatTimeStr+"/"+h.Filename, 640, 480)
		} else if config.Width < config.Height {
			// 认为图片为 0.618 : 1
			utils.Thumbnail("static/upload/"+formatTimeStr+"/"+h.Filename, 475, 768)
		} else {
			utils.Thumbnail("static/upload/"+formatTimeStr+"/"+h.Filename, uint(config.Width), uint(config.Height))
		}
		var data = []string{}
		msg := baseUrl + "/static/upload/" + formatTimeStr + "/" + h.Filename
		data = append(data, msg)
		c.Data["json"] = map[string]interface{}{"errno": 0, "data": data}
	}
	c.ServeJSON()
}

// 上传自拍照
func (c *UploadingController) UploadingSelfie() {
	//接收
	f, h, err := c.GetFile("file")
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "msg": "系统接收不到图片哎!", "data": nil}
		c.ServeJSON()
		return
	}
	defer f.Close()
	timeUnix := time.Now().Unix() //已知的时间戳
	formatTimeStr := time.Unix(timeUnix, 0).Format("2006-01-02")
	fname := utils.Random(4)
	//判断文件夹是否存在，如果不存在就创建
	utils.IsDir("static/selfie/upload/" + formatTimeStr + "/")
	c.SaveToFile("file", "static/selfie/upload/"+formatTimeStr+"/"+fname+h.Filename) // 保存位置在 static/upload, 没有文件夹要先创建
	//压缩图片
	utils.Thumbnail("static/selfie/upload/"+formatTimeStr+"/"+fname+h.Filename, 390, 550)
	msg := baseUrl + "/static/selfie/upload/" + formatTimeStr + "/" + fname + h.Filename
	c.Data["json"] = map[string]interface{}{"code": 0, "msg": msg, "data": nil}
	c.ServeJSON()
}

// 上传荣誉照片
func (c *UploadingController) UploadingHonerImage() {
	//接收
	f, h, err := c.GetFile("file")
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "msg": "系统接收不到图片哎!", "data": nil}
		c.ServeJSON()
		return
	}
	defer f.Close()
	timeUnix := time.Now().Unix() //已知的时间戳
	formatTimeStr := time.Unix(timeUnix, 0).Format("2006-01-02")
	fname := utils.Random(4)
	//判断文件夹是否存在，如果不存在就创建
	utils.IsDir("static/honor/upload/" + formatTimeStr + "/")
	c.SaveToFile("file", "static/honor/upload/"+formatTimeStr+"/"+fname+h.Filename) // 保存位置在 static/upload, 没有文件夹要先创建
	//压缩图片
	config, _, _ := image.DecodeConfig(f) //获取图片信息
	if config.Width > config.Height {
		// 认为图片为 4 : 3
		utils.Thumbnail("static/honor/upload/"+formatTimeStr+"/"+fname+h.Filename, 640, 480)
	} else if config.Width < config.Height {
		// 认为图片为 0.618 : 1
		utils.Thumbnail("static/honor/upload/"+formatTimeStr+"/"+fname+h.Filename, 475, 768)
	}
	msg := baseUrl + "/static/honor/upload/" + formatTimeStr + "/" + fname + h.Filename
	c.Data["json"] = map[string]interface{}{"code": 0, "msg": msg, "data": nil}
	c.ServeJSON()
}

//接收上传的群发邮件模板
func (c *UploadingController) UploadingMailModel() {
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
	err = c.SaveToFile("file", "excel/model/群发模板.xlsx") // 保存位置在 static/upload, 没有文件夹要先创建
	if err == nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "msg": "上传成功", "data": nil}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "msg": "上传失败", "data": nil}
	}
	c.ServeJSON()
}
