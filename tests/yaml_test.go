package test

import (
	"github.com/xiexianbin/webhooks/utils"
	"testing"
)

func TestYaml(t *testing.T) {
	path := "/Users/xiexianbin/work/code/go/src/github.com/xiexianbin/webhooks/conf/webhooks.yaml"
	t.Log("Begin to read yaml file: " + path)
	conf, err := utils.ReadYaml(path)
	if err == nil {
		t.Log(conf)
		for i, hook := range conf.WebHooks {
			t.Log(i)
			t.Log(hook)
			t.Log(hook.Branch)
			t.Log(hook.Script)
			t.Log(hook.Repository)
			t.Log(hook.Organization)
			t.Log(hook.Secret)
		}
	} else {
		t.Log("Read file error!")
	}
}
