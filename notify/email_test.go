package notify

import (
	"github.com/xiexianbin/webhooks/utils"
	"testing"
)

func TestEmail(t *testing.T) {
	conf := utils.Smtp{
		Username: "robot@80.xyz",
		Password: "@B94AAwHhe87@Qp",
		Host:     "smtp.qiye.aliyun.com",
		Port:     465,
		SSL:      true,
	}

	mailTo := []string{"me@xiexianbin.cn"}
	err := SendMail(mailTo, "test", "test", conf)
	if err == nil {
		t.Log("send email success!")
	} else {
		t.Log("send email error!")
	}
}
