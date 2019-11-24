package gosync

import (
	"fmt"
	"github.com/xiexianbin/webhooks/utils"
)


func GoSync(aliyunOSSConfig *utils.AliyunOSSConfig, sourceDir string) {
	err := SyncLocalToOSS(aliyunOSSConfig, sourceDir, "")
	if err != nil {
		fmt.Println("err", err)
	}
}
