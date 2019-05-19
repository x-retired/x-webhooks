package controllers

import "github.com/astaxie/beego"

type VersionController struct {
	beego.Controller
}

func (v *VersionController) Get() {
	v.Data["json"] = map[string]string{"version": "0.2.0"}
	v.ServeJSON()
}
