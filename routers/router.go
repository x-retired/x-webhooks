package routers

import (
	"github.com/xiexianbin/webhooks/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
