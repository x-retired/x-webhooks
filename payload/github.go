package payload

import (
	"strings"

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
	branch := strings.Replace(event.GetRef(), "refs/heads/", "", 1)

	logs.Info("repository name:", repositoryName, "branch:", branch)

	conf, err := utils.ReadYaml("")
	if err == nil {
		for _, provider := range conf.Webhooks {
			for _, action := range provider.Actions {
				if action.Event == "push" {
					for _, item := range action.Items {
						if branch == item.Branch && repositoryName == item.RepositoryName {
							logs.Info("Matching repo:", item.RepositoryName, "branch:", item.Branch, "script:", item.Script)
							_, _ = utils.RunBash(beego.AppPath + "/" + beego.AppConfig.String("scripts") + "/" + item.Script)
							email := event.GetHeadCommit().GetCommitter().GetEmail()
							logs.Info("begin to send to", email)
							//mailTo := []string{email}
							//utils.SendMail(mailTo, "go webhooks", "test")
							return
						}
					}
				}
			}
		}
		logs.Warning("No Matching config. skip.")
	} else {
		logs.Error("Read file error!")
	}
}
