package controllers

import (
	"github.com/astaxie/beego"
)

type WebController struct {
	beego.Controller
}

func (c *WebController) Home() {
}
