package main

import (
	"github.com/xiexianbin/webhooks/cmd"
	_ "github.com/xiexianbin/webhooks/initialize"
	_ "github.com/xiexianbin/webhooks/routers"

	"github.com/astaxie/beego"
)

func main() {
	cmd.RunCommand()

	// swagger settings
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}
