package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "xiexianbin.cn"
	c.Data["Email"] = "me@xiexianbin.cn"
	c.TplName = "index.tpl"
}
