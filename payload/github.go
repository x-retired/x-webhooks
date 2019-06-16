package payload

import (
	"github.com/xiexianbin/webhooks/utils"

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
	logs.Info(event.GetRef())
	logs.Info("--")
	email := event.GetHeadCommit().GetCommitter().GetEmail()
	logs.Info(email)

	//body, err = ioutil.ReadAll(ctx.Request.Body)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer ctx.Request.Body.Close()
	//
	//payload := github.WebHookPayload{}
	//_ = json.Unmarshal(body, &payload)
	//
	//repository := payload.GetRepo()
	//fullName := repository.GetFullName()
	//
	//logs.Info("repository name: " + repository.GetName())
	//logs.Info("repository full_name: " + fullName)
	//logs.Info("user email: " + *payload.GetPusher().Email)

	conf, err := utils.ReadYaml("")
	if err == nil {
		//for _, hook := range conf.WebHooks {
		//	_fullName := hook.Organization + "/" + hook.Repository
		//	if fullName == _fullName {
		//		logs.Info("begin to run : " + hook.Script)
		//		// _, _ = utils.RunBash(beego.AppPath + "/conf/scripts/" + hook.Script)
		//		_, _ = utils.RunBash(beego.AppPath + "/" + beego.AppConfig.String("scripts") + "/" + hook.Script)
		//	}
		//}
		logs.Info(conf)
	} else {
		logs.Error("Read file error!")
	}

}
