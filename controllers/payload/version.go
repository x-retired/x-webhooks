package payload

import "github.com/astaxie/beego"

type PayloadController struct {
	beego.Controller
}

func (this *PayloadController) show_payload_version() {
	this.Data["content"] = "{version: 1.0.0}"
}
