package help

import (
	"fmt"
	"os"
)

// help method
func Help() {
	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println("Webhooks is a webhooks Web Application for Github, Gitlab and so on.")
		fmt.Println("")
		fmt.Println("USAGE")
		fmt.Println("    webhooks command")
		fmt.Println("AVAILABLE COMMANDS")
		fmt.Println("    install     Install webhooks, Use: webhooks install -account=admin -password=123456 -email=me@xiexianbin.cn")
		fmt.Println("    version     Prints the current webhooks version")
	}
}
