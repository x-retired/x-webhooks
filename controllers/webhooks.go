package controllers

import (
	"github.com/xiexianbin/webhooks/payload"
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
		"success": false,
	}

	userAgent := c.Ctx.Request.UserAgent()
	logs.Info("Client User Agent is", userAgent)
	if strings.HasPrefix(userAgent, "GitHub-Hookshot") {
		_, _ = payload.GithubPayload(c.Ctx)
		result["success"] = true
		result["message"] = "success"
	} else {
		logs.Warn("Un-support User Agent:", userAgent)
		result["success"] = false
		result["message"] = "Un-support User Agent"
	}

	c.Data["json"] = result
	c.ServeJSON()
	return
}
