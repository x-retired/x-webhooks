// @APIVersion 1.0.0
// @Title webhooks apis
// @Description webhooks apis
// @Contact me@xiexianbin.cn
// @TermsOfServiceUrl https://xiexianbin.cn/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/xiexianbin/webhooks/controllers"
	"github.com/xiexianbin/webhooks/controllers/apis"
	"github.com/xiexianbin/webhooks/controllers/apis/v1"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	ns := beego.NewNamespace("/api",
		beego.NSRouter("/", &apis.VersionController{}),
		beego.NSNamespace("/v1",
			beego.NSRouter("/", &apis.VersionController{}),
			beego.NSRouter("/payload", &v1.PayloadController{}),
		),
	)
	beego.AddNamespace(ns)
}
