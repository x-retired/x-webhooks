package cmds

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"os"
)

// register logs
func RegisterLogger() {
	logs.SetLogger("console")
	logs.SetLogger("file", `{"filename": "logs/webhooks.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.Async()
}

// help
func Help() {
	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println("USAGE")
		fmt.Println("    webhooks command")
		fmt.Println("AVAILABLE COMMANDS")
		fmt.Println("    version     Prints the current webhooks version")
	}
}

// return webhooks version.
func Version() {
	if os.Args[1] == "version" {
		fmt.Println("v1.0.0")
	}
}

// command line
func RunCommand() {
	if len(os.Args) > 1 {
		Help()
		Version()
		os.Exit(0)
	}
}
