package controllers

import (
	"github.com/xiexianbin/webhooks/drivers/github"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type WebhooksController struct {
	beego.Controller
}

func (c *WebhooksController) Get() {
	c.Ctx.WriteString("Please call this url use Post method.")
}

func (c *WebhooksController) Post() {
	result := map[string]interface{}{
		"result": false,
	}

	userAgent := c.Ctx.Request.UserAgent()
	logs.Info("Client User Agent is", userAgent)
	if strings.HasPrefix(userAgent, "GitHub-Hookshot") {
		_, _ = github.GithubPayload(c.Ctx)
		result["result"] = true
		result["message"] = "Success."
	} else {
		logs.Warn("Un-support User Agent:", userAgent)
		result["result"] = false
		result["message"] = "Un-Support User Agent."
	}

	c.Data["json"] = result
	c.ServeJSON()
	return
}
