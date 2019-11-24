package commands

import (
	"fmt"
	"os"

	"github.com/astaxie/beego/logs"
)

// RegisterLogger bee log
func RegisterLogger() {
	_ = logs.SetLogger("console")
	_ = logs.SetLogger("file", `{"filename":"logs/webhooks.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.Async()
}

// command line
func RunCommand() {
	if len(os.Args) > 1 {
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			Help()
		}
		if os.Args[1] == "-v" || os.Args[1] == "--version" {
			Version()
		}

		os.Exit(0)
	}
}

// help method
func Help() {
	fmt.Println("Webhooks is a Web Application for Github, Gitlab and so on.")
	fmt.Println("")
	fmt.Println("USAGE")
	fmt.Println("    webhooks [command] [option]")
	fmt.Println("AVAILABLE OPTIONS")
	fmt.Println("    -h, --help      Prints help information.")
	fmt.Println("    -v, --version   Prints the current version.")
}

// version method
func Version() {
	fmt.Println("v1.0.0")
}
