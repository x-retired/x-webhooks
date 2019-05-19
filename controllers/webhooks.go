package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/xiexianbin/webhooks/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/google/go-github/github"
)

type WebhooksController struct {
	beego.Controller
}

func (c *WebhooksController) Get() {
	c.Ctx.WriteString("Please call this url use Post method.")
}

func (c *WebhooksController) Post() {
	//payload, err := github.ValidatePayload(c.Ctx.Request, c.webhookSecretKey)
	logs.Info("webhook type: " + github.WebHookType(c.Ctx.Request))
	logs.Info("X-GitHub-Event : " + c.Ctx.Request.Header.Get("X-GitHub-Event"))

	body, err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	payload := github.WebHookPayload{}
	_ = json.Unmarshal(body, &payload)

	repository := payload.GetRepo()
	fullName := repository.GetFullName()

	logs.Info("repository name: " + repository.GetName())
	logs.Info("repository full_name: " + fullName)
	logs.Info("user email: " + *payload.GetPusher().Email)

	path := beego.AppPath + "/" + beego.AppConfig.String("webhooks")  // "/conf/webhooks.yaml"
	logs.Info("Begin to read yaml file: " + path)
	conf, err := utils.ReadYaml(path)
	if err == nil {
		for _, hook := range conf.WebHooks {
			_fullName := hook.Organization + "/" + hook.Repository
			if fullName == _fullName {
				logs.Info("begin to run : " + hook.Script)
				// _, _ = utils.RunBash(beego.AppPath + "/conf/scripts/" + hook.Script)
				_, _ = utils.RunBash(beego.AppPath + "/" + beego.AppConfig.String("scripts") + "/" + hook.Script)
			}
		}
	} else {
		logs.Error("Read file error!")
	}

	c.Ctx.ResponseWriter.WriteHeader(200)
	return
}
