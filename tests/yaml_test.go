package test

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/xiexianbin/webhooks/utils"
)

func TestYaml(t *testing.T) {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	t.Log(path)

	path = "/Users/xiexianbin/work/code/go/src/github.com/xiexianbin/webhooks/tests/webhooks.yaml"
	t.Log("Begin to read yaml file:", path)
	conf, err := utils.ReadYaml(path)
	if err == nil {
		t.Log(conf)
		smtp := conf.Smtp
		t.Log(smtp.Host)
		t.Log(smtp.Password)
		t.Log(smtp.Port)
		t.Log(smtp.Username)
		t.Log(smtp.SSL)

		for i, provider := range conf.Webhooks {
			t.Log(i)
			t.Log("Name: " + provider.Name)
			t.Log(provider.Actions)
			for j, action := range provider.Actions {
				t.Log(j)
				t.Log("Event: " + action.Event)
				t.Log(action.Items)
				for l, item := range action.Items {
					t.Log(l)
					t.Log("Branch: " + item.Branch)
					t.Log("Repository: " + item.RepositoryName)
					t.Log("Script: " + item.Script)
					t.Log("Secret: " + item.Secret)
				}
			}
		}
	} else {
		t.Log("Read file error! err: " + err.Error())
	}
}
