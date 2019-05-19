package test

import (
	"github.com/xiexianbin/webhooks/utils"
	"testing"
)

func TestRunBash(t *testing.T) {
	result, err := utils.RunBash("pwd")
	if err == nil {
		t.Log(result)
	} else {
		t.Log("run bash error!")
	}
}
