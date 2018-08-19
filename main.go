package main

import (
	_ "github.com/xiexianbin/webhooks/routers"
	"github.com/astaxie/beego"
	"github.com/xiexianbin/webhooks/cmds"
)

func main() {
	cmds.RegisterLogger()

	cmds.RunCommand()
	beego.Run()
}

