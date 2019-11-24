package notify

import (
	"gopkg.in/gomail.v2"
	"net/smtp"

	"github.com/xiexianbin/webhooks/utils"

	"github.com/astaxie/beego/logs"
)

func SendMail(emailTo []string, subject, body string, conf utils.Smtp) error {
	logs.Info("begin to send to", emailTo)
	msg := gomail.NewMessage()
	msg.SetHeader("From", "XD Game"+"<"+conf.Username+">")
	msg.SetHeader("To", emailTo...)
	//msg.SetAddressHeader("Cc", conf.Username, "webhooks")
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)

	mailer := gomail.Dialer{
		Host: conf.Host,
		Port: conf.Port,
		Auth: smtp.PlainAuth(
			"",
			conf.Username,
			conf.Password,
			conf.Host),
		SSL: conf.SSL}
	if err := mailer.DialAndSend(msg); err != nil {
		logs.Warning("Send mail fail:", err)
		return err
	} else {
		logs.Info("Send mail to", emailTo, "Success")
	}

	return nil
}
