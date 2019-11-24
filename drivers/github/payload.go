package github

import (
	"bytes"
	"encoding/json"
	"html/template"
	"strings"

	"github.com/xiexianbin/webhooks/notify"
	"github.com/xiexianbin/webhooks/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/google/go-github/github"
)

func GithubPayload(ctx *context.Context) (result string, err error) {
	eventType := github.WebHookType(ctx.Request)
	deliveryID := github.DeliveryID(ctx.Request)
	logs.Info("Webhook X-GitHub-Event: " + eventType)
	logs.Info("Webhook X-GitHub-Delivery: " + deliveryID)

	var secretToken []byte
	payload, err := github.ValidatePayload(ctx.Request, secretToken)
	if err != nil {
		logs.Warn(err)
	}

	event, err := github.ParseWebHook(github.WebHookType(ctx.Request), payload)
	if err != nil {
		logs.Warn(err)
	}

	switch event := event.(type) {
	case *github.PushEvent:
		processPushEvent(event)
	case *github.CommitCommentEvent:
		processCommitCommentEvent(event)
	case *github.CreateEvent:
		processCreateEvent(event)
	case *github.DeleteEvent:
		processDeleteEvent(event)
	default:
		logs.Warn("unkonw event:", event)
	}

	return "", nil
}

func processCommitCommentEvent(event *github.CommitCommentEvent) {
	logs.Info("CommitCommentEvent")

}

func processCreateEvent(event *github.CreateEvent) {
	logs.Info("CreateEvent")

}

func processDeleteEvent(event *github.DeleteEvent) {
	logs.Info("DeleteEvent")
}

func processPushEvent(event *github.PushEvent) {
	repositoryName := event.GetRepo().GetFullName()
	repositoryUrl := event.GetRepo().GetURL()
	branch := strings.Replace(event.GetRef(), "refs/heads/", "", 1)

	logs.Info("repository name:", repositoryName, "branch:", branch)

	conf, err := utils.ReadYaml()
	if err != nil {
		logs.Error("Read file error!")
		return
	}

	for _, provider := range conf.Webhooks {
		for _, action := range provider.Actions {
			if action.Event != "push" {
				continue
			}
			for _, item := range action.Items {
				if branch == item.Branch && repositoryName == item.RepositoryName {
					logs.Info("Matching repo:", item.RepositoryName,
						"branch:", item.Branch, "script:", item.Script)
					scriptPath := beego.AppPath + "/" + beego.AppConfig.String("scripts") +
						"/" + item.Script
					go utils.RunBash(scriptPath)

					mailTo := []string{event.GetHeadCommit().GetCommitter().GetEmail()}
					subject := "[Webhooks] Git Push Event Success"
					_content, _ := json.MarshalIndent(event, "", "  ")
					templateData := struct {
						Title    string
						Result   bool
						GitUrl   string
						GitRepo  string
						Content  string
					}{
						Title:   subject,
						Result:  true,
						GitUrl:  repositoryUrl,
						GitRepo: repositoryName,
						Content: string(_content),
					}

					templateFileName := beego.AppPath + "/views/email/result.html"
					sendEmail(mailTo, subject, templateFileName, templateData)
					return
				}
			}
		}
	}

	logs.Warning("No Matching config. skip.")
	_ = notify.SendMail(utils.GetDefaultNotifyEmail(), "[Webhooks] failed", event.String(), utils.GetSmtp())
}

func sendEmail(emailTo []string, subject, templateFileName string, templateData interface{}) bool {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		logs.Warn(err)
		return false
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, templateData); err != nil {
		logs.Warn(err)
		return false
	}

	notify.SendMail(emailTo, subject, buf.String(), utils.GetSmtp())
	return true
}
