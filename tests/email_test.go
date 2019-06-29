package test

import (
"github.com/xiexianbin/webhooks/utils"
"testing"
)

func TestEmail(t *testing.T) {
	mailTo := []string{"me@xiexianbin.cn"}
	err := utils.SendMail(mailTo, "test", "test")
	if err == nil {
		t.Log("success")
	} else {
		t.Log("run bash error!")
	}
}
