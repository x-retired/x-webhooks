package utils

import (
	"os/exec"

	"github.com/astaxie/beego/logs"
)

func RunBash(script string) (result string, err error) {
	out, err := exec.Command("bash", "-c", script).Output()
	result = string(out)
	if err != nil {
		logs.Info("Exec command failed: %s", err)
	}

	logs.Info("Run %s output: %s", script, out)
	return result, err
}
