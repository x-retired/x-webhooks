package commands

import (
	"fmt"
	"os"

	"github.com/astaxie/beego/logs"
)

// RegisterLogger bee log
func RegisterLogger() {
	logs.SetLogger("console")
	logs.SetLogger("file", `{"filename":"logs/webhooks.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.Async()
}

// command line
func RunCommand() {
	if len(os.Args) > 1 {
		Help()
		Version()
		os.Exit(0)
	}
}

// help method
func Help() {
	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println("Webhooks is a Web Application for Github, Gitlab and so on.")
		fmt.Println("")
		fmt.Println("USAGE")
		fmt.Println("    webhooks [command] [option]")
		fmt.Println("AVAILABLE OPTIONS")
		fmt.Println("    -h, --help      Prints help information.")
		fmt.Println("    -v, --version   Prints the current version.")
	}
}

// version method
func Version() {
	if os.Args[1] == "-v" || os.Args[1] == "--version" {
		fmt.Println("v1.0.0")
	}
}
