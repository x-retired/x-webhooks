package main

import (
	"os"

	"github.com/xiexianbin/webhooks/commands"
	_ "github.com/xiexianbin/webhooks/routers"

	"github.com/astaxie/beego"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] != "run" {
		commands.RunCommand()
	}

	commands.RegisterLogger()
	// utils.LoadArgs()

	// swagger settings
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}
