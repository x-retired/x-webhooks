package test

import (
	"github.com/xiexianbin/webhooks/notify"
	"testing"
)

func TestEmail(t *testing.T) {
	mailTo := []string{"me@xiexianbin.cn"}
	err := notify.SendMail(mailTo, "test", "test")
	if err == nil {
		t.Log("send email success!")
	} else {
		t.Log("send email error!")
	}
}
