package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"

	"github.com/xiexianbin/webhooks/controllers"
	"github.com/xiexianbin/webhooks/controllers/payload"
)

func init() {
	api_ns := beego.NewNamespace("/api",
		beego.NSCond(func(ctx *context.Context) bool {
			if ctx.Input.Domain() == "webhooks.xiexianbin.cn" {
				return true
			}
			return false
		}),
		beego.NSRouter("/version", &payload.PayloadController{}, "*:show_payload_version"),
		beego.NSNamespace("v1",
			beego.NSInclude()),
	)
	beego.AddNamespace(api_ns)

	beego.Router("/", &controllers.MainController{})

}
