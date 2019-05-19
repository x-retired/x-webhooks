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

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	// webhooks
	beego.Router("webhooks/", &controllers.WebhooksController{})

	// version
	beego.Router("version/", &controllers.VersionController{})
}
