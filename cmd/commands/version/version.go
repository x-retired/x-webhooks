package version

import (
	"fmt"
	"os"
)

// return version
func Version() {
	if os.Args[1] == "version" || os.Args[1] == "-v" {
		fmt.Println("v1.0.0")
	}
}
