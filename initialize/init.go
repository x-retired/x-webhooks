package initialize

import (
	"github.com/xiexianbin/webhooks/utils"
)

func init() {
	// init logs
	utils.InitLogs()
	// init database
	InitDatabase()
}
