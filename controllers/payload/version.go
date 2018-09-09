package payload

import "github.com/astaxie/beego"

type VersionController struct {
	beego.Controller
}

func (this *VersionController) ShowPayloadVersion() {
	this.Data["content"] = "{version: 1.0.0}"
}
