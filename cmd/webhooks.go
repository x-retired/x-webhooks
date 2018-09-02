package cmd

import (
	"os"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/xiexianbin/webhooks/cmd/commands/help"
	"github.com/xiexianbin/webhooks/cmd/commands/install"
	"github.com/xiexianbin/webhooks/cmd/commands/version"
)

// register logs
func RegisterLogger() {
	logs.SetLogger("console")
	logs.SetLogger("file", `{"filename": "logs/webhooks.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.Async()
}

// command line
func RunCommand() {
	if len(os.Args) > 1 {
		help.Help()
		orm.RunCommand()
		install.Install()
		version.Version()
		os.Exit(0)
	}
}
