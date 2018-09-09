package cmd

import (
	"os"

	"github.com/astaxie/beego/orm"
	"github.com/xiexianbin/webhooks/cmd/commands/help"
	"github.com/xiexianbin/webhooks/cmd/commands/install"
	"github.com/xiexianbin/webhooks/cmd/commands/version"
)

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
