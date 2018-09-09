package v1

import "github.com/astaxie/beego"

type PayloadController struct {
	beego.Controller
}

func (p *PayloadController) Get() {
	p.Data["json"] = map[string]string{"version": "1.0.0"}
	p.ServeJSON()
}
